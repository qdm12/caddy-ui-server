package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileSystem interface {
	ReplaceAllRecursively(directory string, oldToNew map[string]string) (err error)
}

type fileSystem struct {
	fileWalk  func(root string, walkFn filepath.WalkFunc) error
	readFile  func(path string) (content []byte, err error)
	writeFile func(filename string, data []byte, perm os.FileMode) error
}

func NewFileSystem() FileSystem {
	return &fileSystem{
		fileWalk:  filepath.Walk,
		readFile:  ioutil.ReadFile,
		writeFile: ioutil.WriteFile,
	}
}

func (f *fileSystem) ReplaceAllRecursively(directory string, oldToNew map[string]string) (err error) {
	return f.fileWalk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		stat, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("cannot replace strings in directory %s: %w", directory, err)
		} else if stat.IsDir() {
			return nil
		}
		content, err := f.readFile(path)
		if err != nil {
			return fmt.Errorf("cannot replace strings in directory %s: %w", directory, err)
		}
		s := string(content)
		for old, new := range oldToNew {
			s = strings.ReplaceAll(s, old, new)
		}
		if err := f.writeFile(path, []byte(s), 0); err != nil {
			return fmt.Errorf("cannot replace strings in directory %s: %w", directory, err)
		}
		return nil
	})
}

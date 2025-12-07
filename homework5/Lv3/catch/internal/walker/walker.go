package walker

import (
	"path/filepath"
	"io/fs"
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv3/catch/internal/pool"
)

func WalkDir(dir string, taskChan chan<- pool.Task) error {
	return filepath.WalkDir (dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			taskChan <- pool.Task{Filepath: path}
		}
	return nil
})
}

//func filepath.WalkDir(root string, fn fs.WalkDirFunc) error

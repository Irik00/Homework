// package main

// import "log"
	

// func main(){
// 	log.SetFlags(log.Ldate|log.Llongfile|log.Ltime)
// 	log.SetPrefix("[txt]")
// 	log.Println("This is a simple log")
// }

package main

import (
        "io"
        "os"
        "path/filepath"
		//func Dir(path string) string
        "time"
)

type FileState map[string]time.Time

func scanDir(dir string) (FileState, error) {
        state := make(FileState)
        err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			//func Walk(root string, fn WalkFunc) error
                if err != nil {
                        return err
                }
                
                if !info.IsDir() {
					//IsDir() bool        // abbreviation缩写 for Mode().IsDir()
                        relPath, _ := filepath.Rel(dir, path)
						//func Rel(basepath, targpath string) (string, error)
                        state[relPath] = info.ModTime()
                }
                return nil
        })
        return state, err
}

func syncFile(srcPath, dstPath string) error {
        
        if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			//unc MkdirAll(path string, perm FileMode) error
                return err
        }

       
        srcFile, err := os.Open(srcPath)
        if err != nil {
                return err
        }
        defer srcFile.Close()

        
        dstFile, err := os.Create(dstPath)
        if err != nil {
                return err
        }
        defer dstFile.Close()

      
        _, err = io.Copy(dstFile, srcFile)
        return err
}


func monitorAndSync(srcDir, dstDir string, interval time.Duration) error {
   
        prevState, err := scanDir(srcDir)
        if err != nil {
                return err
        }

        
        ticker := time.NewTicker(interval)
                //ticker 心脏 自动收报机
		//func NewTicker(d Duration) *Ticker
        defer ticker.Stop()//停止

        for range ticker.C {
             
            currentState, err := scanDir(srcDir)
            if err != nil {
                        return err
            }

                
            for relPath, modTime := range currentState {
                       
                if prevState[relPath].Before(modTime) {
                    srcPath := filepath.Join(srcDir, relPath)
                    dstPath := filepath.Join(dstDir, relPath)
                    if err := syncFile(srcPath, dstPath); err != nil {
                            return err
                    	}
                    prevState[relPath] = modTime 
                    }
                }
        }
        return nil
}

func main() {
        srcDir := "./test1"
        dstDir := "./test2"
        interval := 2 * time.Second

        if err := monitorAndSync(srcDir, dstDir, interval); err != nil {
                panic(err)
				//func panic(v any)
        }
}
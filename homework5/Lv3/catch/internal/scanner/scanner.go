package scanner

import(
	"bufio"
	"os"
	"strings"
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv3/catch/internal/pool"
)

func ScanFile(task pool.Task, resultChan chan<- pool.Result) error {
	file, err := os.Open(task.Filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(line, task.Keyword) {
			resultChan <- pool.Result{
				Filepath: task.Filepath,
				LineNum:  lineNum,
				Content:  line,
			}
		}
	}
	return scanner.Err()
}
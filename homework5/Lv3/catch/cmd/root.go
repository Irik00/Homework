package cmd

import (
	"fmt"
	"os"
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv3/catch/internal/pool"
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv3/catch/internal/scanner"
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv3/catch/internal/walker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "catch [directory] [keyword]",
	Short: "A simple grep-like tool to search for keywords in files",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		keyword := args[1]
		p := pool.NewPool(10)
		defer p.Close()
		go func() {
			err := walker.WalkDir(dir, p.TaskChan)
			if err != nil {
				fmt.Printf("Error walking directory: %v\n", err)
			}
			close(p.TaskChan)
		}()
		for i := 0; i < 10; i++ {
			p.Wg.Add(1)
			go func() {
				defer p.Wg.Done()
				for task := range p.TaskChan {
					task.Keyword = keyword
					scanner.ScanFile(task, p.ResultChan)
				}
				
			}()
		}
		go func() {
			p.Wg.Wait()
			close(p.ResultChan)
		}()
		fmt.Println("Search Results:")
		for result := range p.ResultChan {
			fmt.Printf("%s:%d: %s\n", result.Filepath, result.LineNum, result.Content)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
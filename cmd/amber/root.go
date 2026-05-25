package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //根命令
	Use:   "amber",
	Short: "一个系统备份软件",
	Long:  `Amber 是一个系统备份软件，支持系统备份、文件备份、磁盘备份`,
}

// Execute 用于启动根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

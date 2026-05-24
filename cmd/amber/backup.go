package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/backup"
	"github.com/hhelibe2006-commits/Amber/internal/cli"
	"github.com/spf13/cobra"
)

var (
	input  cli.Put
	output string
)

var backupCmd = &cobra.Command{
	Use:   "backup [source]",
	Short: "执行备份",
	RunE: func(cmd *cobra.Command, args []string) error {
		typ, _ := cmd.Flags().GetString("type")

		if err := backup.Run(typ, input, output); err != nil {
			fmt.Println("备份出现错误:", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().VarP(&input, "input", "i", "输入路径")
	backupCmd.Flags().StringVarP(&output, "output", "o", "", "输出路径")

	if err := backupCmd.MarkFlagRequired("output"); err != nil {
		fmt.Println("初始化出现错误:", err)
		return
	}

	if err := backupCmd.MarkFlagRequired("input"); err != nil {
		fmt.Println("初始化出现错误:", err)
		return
	}
	backupCmd.Flags().StringP("type", "t", "file", "备份类型")
}

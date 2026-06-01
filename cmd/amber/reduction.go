package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/reduction"
	"github.com/spf13/cobra"
)

func init() {
	var input string

	//还原命令
	var restoreCmd = &cobra.Command{
		Use:   "restore",
		Short: "执行还原",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := reduction.Run(input); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringVarP(&input, "input", "i", "", "输入路径")
	err := restoreCmd.MarkFlagRequired("input")
	if err != nil {
		fmt.Println("初始化错误:", err)
		return
	}
}

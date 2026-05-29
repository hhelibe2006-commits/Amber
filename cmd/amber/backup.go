package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/cli"
	"github.com/spf13/cobra"
)

func init() {
	var input cli.PathList
	var output string

	// 备份命令
	var backupCmd = &cobra.Command{
		Use:   "backup [source]",
		Short: "执行备份",
		//RunE会返回错误，而Run不会，所以使用RunE
		RunE: func(cmd *cobra.Command, args []string) error {
			typ, err := cmd.Flags().GetString("type")
			if err != nil {
				return err
			}
			if err := backup.Run(typ, input, output); err != nil {
				fmt.Println("备份出现错误:", err)
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().VarP(&input, "input", "i", "输入路径")
	backupCmd.Flags().StringVarP(&output, "output", "o", "", "输出路径")

	//确保有输出路径
	if err := backupCmd.MarkFlagRequired("output"); err != nil {
		fmt.Println("初始化出现错误:", err)
		return
	}

	//确保有输入路径
	if err := backupCmd.MarkFlagRequired("input"); err != nil {
		fmt.Println("初始化出现错误:", err)
		return
	}
	//有file、system、disk三种备份方式，默认file
	backupCmd.Flags().StringP("type", "t", "file", "备份类型")
}

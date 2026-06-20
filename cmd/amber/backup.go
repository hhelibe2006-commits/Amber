package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/backup"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/spf13/cobra"
)

func init() {
	clas := value.NewClas()

	// 备份命令
	var backupCmd = &cobra.Command{
		Use:   "backup [source]",
		Short: "执行备份",
		//RunE会返回错误，而Run不会，所以使用RunE
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := backup.Run(clas); err != nil {
				fmt.Println("备份出现错误:", err)
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().VarP(&clas.Input, "input", "i", "输入路径")
	backupCmd.Flags().StringVarP(&clas.Output, "output", "o", "", "输出路径")
	//有fastcdc，rabincdc
	backupCmd.Flags().StringVarP(&clas.Cdc, "cdc", "c", "fastcdc", "CDC算法")
	//有file、system、disk三种备份方式，默认file
	backupCmd.Flags().StringVarP(&clas.Typ, "type", "t", "file", "备份类型")

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

}

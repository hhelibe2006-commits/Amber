package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/reduction"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/spf13/cobra"
)

func init() {
	clas := value.NewReductionClas()
	//还原命令
	var restoreCmd = &cobra.Command{
		Use:   "restore",
		Short: "执行还原",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(clas.Input)
			if err := reduction.Run(clas); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(restoreCmd)

	restoreCmd.Flags().StringVarP(&clas.Input, "input", "i", "", "输入路径")
	if err := restoreCmd.MarkFlagRequired("input"); err != nil {
		fmt.Println("初始化错误:", err)
		return
	}
}

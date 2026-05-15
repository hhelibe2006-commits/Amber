package main

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/src/internal/args"
	"github.com/hhelibe2006-commits/Amber/src/internal/backup"
)

func main() {
	cfg, err := args.Args()
	if err != nil {
		fmt.Println(err)
	}
	backup.Backup(cfg)
}

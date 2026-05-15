package main

import (
	"errors"
	"flag"
	"fmt"
	"slices"
)

type InputPathList []string

func (i *InputPathList) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *InputPathList) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Config struct {
	Mode       string
	OutputPath string
	InputPath  InputPathList
	Target     string
}

func Args() (*Config, error) {
	cfg := &Config{}
	initConfig(cfg)
	flag.Parse()
	err := completenessCheck(cfg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", cfg.InputPath[0])
	return cfg, nil
}

func initConfig(cfg *Config) {
	flag.StringVar(&cfg.Mode, "mode", "full", "备份方式")
	flag.StringVar(&cfg.Mode, "m", "full", "备份方式")

	flag.StringVar(&cfg.OutputPath, "o", "", "输出路径")
	flag.StringVar(&cfg.OutputPath, "out", "", "输出路径")

	flag.Var(&cfg.InputPath, "i", "输入路径或文件")
	flag.Var(&cfg.InputPath, "in", "输入路径或文件")

	flag.StringVar(&cfg.Target, "target", "file", "备份目标")
	flag.StringVar(&cfg.Target, "t", "file", "备份目标")
}

func completenessCheck(cfg *Config) error {
	mode := []string{"full", "incremental", "differential"}
	target := []string{"file", "system", "disk"}
	if cfg.OutputPath == "" {
		return errors.New("未指定输出路径")
	}
	if len(cfg.InputPath) == 0 {
		return errors.New("未指定备份数据")
	}
	if !slices.Contains(mode, cfg.Mode) {
		return errors.New("不存在这种备份方式")
	}
	if !slices.Contains(target, cfg.Target) {
		return errors.New("不存在这种备份目标")
	}
	return nil
}

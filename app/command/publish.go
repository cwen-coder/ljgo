package command

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/cwen-coder/ljgo/app/config"
	"github.com/qiniu/log"
	"github.com/urfave/cli"
)

var CmdPublish = cli.Command{
	Name:   "publish",
	Usage:  "Generate blog to pubilc folder and publish",
	Action: runPublish,
}

func runPublish(c *cli.Context) error {
	cfg, err := config.New(c)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	build(cfg)
	publish(cfg)
	return nil
}

func publish(cfg *config.Config) {
	if cfg.Publish.Cmd == "" {
		return
	}
	var shell, flag string
	shell = "/bin/sh"
	flag = "-c"
	if runtime.GOOS == "windows" {
		shell = "cmd"
		flag = "/C"
	}
	cmd := exec.Command(shell, flag, cfg.Publish.Cmd)
	cmd.Dir = cfg.PublicPath
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("StdoutPipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("StderrPipe: %v", err)
	}
	outStr := bufio.NewScanner(stdout)
	errStr := bufio.NewScanner(stderr)
	go func() {
		for outStr.Scan() {
			fmt.Print(outStr.Text())
		}
	}()
	go func() {
		for errStr.Scan() {
			fmt.Print(errStr.Text())
		}
	}()
	cmd.Run()
}

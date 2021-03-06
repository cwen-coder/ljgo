// Copyright 2016 The Ljgo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"os"

	"github.com/cwen-coder/ljgo/app/command"
	"github.com/urfave/cli"
)

const APP_VER = "v0.1.0-beta"

func main() {
	app := cli.NewApp()
	app.Name = "ligo"
	app.Usage = "An elegant static blog generator"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		command.CmdNew,
		command.CmdBuild,
		command.CmdPublish,
		command.CmdServe,
		command.CmdInfo,
	}
	app.Run(os.Args)
}

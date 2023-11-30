package co

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
)

func init() {
	Z.Conf.SoftInit()
}

var Cmd = &Z.Cmd{

	Name:      `co-author`,
	Summary:   `a unix filter which will output a git co-author line when supplied by a key defined in the config file`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2023 Yongle Xu`,
	License:   `Apache-2.0`,
	Site:      `yonglexu.dev`,
	Source:    `git@github.com:rwx-yxu/co-author.git`,
	Issues:    `github.com/rwx-yxu/co-author/issues`,

	Commands: []*Z.Cmd{
		printCmd,
		// standard external branch imports (see rwxrob/{help,conf,vars})
		help.Cmd, conf.Cmd,
	},

	// Add custom BonzaiMark template extensions (or overwrite existing ones).

	Description: `
		{{cmd .Name}} is a tool that queries the Open weather map API for current weather information.
			`,
}

var printCmd = &Z.Cmd{
	Name:     `print`,
	Summary:  `Print git co-author line when supplied with a key`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(x *Z.Cmd, _ ...string) error {
		fmt.Println("Hello world")
		return nil
	},
}

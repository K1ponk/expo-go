// Copyright 2021 The expo-go Authors.
// This file is part of expo-go.
//
// expo-go is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package expocli

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	// AppHelpTemplate is help template for expo
	AppHelpTemplate = `Name:
	{{.Name}} - {{.Usage}}
	Copyright 2021 The expo-go Authors
Usage:
	{{.HelpName}} [options]{{if .Commands}} [command] [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	{{if .Version}}
Version:
	{{.Version}}
	{{end}}{{if .Commands}}
Commands:
	{{range .Commands}}{{join .Names ", "}}{{"\t"}}{{.Usage}}
	{{end}}{{end}}{{if .Copyright}}
Copyright:
	{{.Copyright}}
	{{end}}
`
)

// NewApp will create new cli app
func NewApp(usage string) *cli.App {
	// configure the cli app
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Usage = usage
	
	return app
}

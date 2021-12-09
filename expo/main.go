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

package main

import (
	"os"

	"github.com/penta-expo/expobar/expocli"
	"github.com/penta-expo/expobar/expoutils"
	"github.com/urfave/cli/v2"
)

var (
	// init the app
	app = expocli.NewApp("EVM-Based project framework.")
)

func init() {
	// configure the app
	app.CustomAppHelpTemplate = expocli.AppHelpTemplate
	app.Commands = []*cli.Command{
		&InitCommand,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		expoutils.Fatalf("cannot start app > %v", err)
	}
}

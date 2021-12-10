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
	"github.com/fatih/color"
	"github.com/penta-expo/expobar/expocli"
	"github.com/penta-expo/expobar/expoutils"
	"github.com/urfave/cli/v2"
)

// initAction is action for InitCommand
func initAction(ctx *cli.Context) error {
	// color plate
	red := color.New(color.FgHiRed).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	// start the spinner
	s := expocli.Spinners()
	s.Message(yellow(" Preparing"))
	s.Start()

	// check config file
	err := expoutils.CheckPath("./expo.json")
	if err == nil {
		s.StopFailMessage(red(" Directory not empty."))
		s.StopFail()

		expoutils.Fatalf("You can't init new project that already initiated")
	}

	// check npm
	err = expoutils.CheckCommand("npm")
	if err != nil {
		s.StopFailMessage(red(" npm not found."))
		s.StopFail()

		expoutils.Fatalf("command 'npm' not found > %v", err)
	}

	// generate and write config file
	var config expocli.Config
	config = expocli.Config{
		Networks: map[string]interface{}{
			"dev": expocli.Network{
				Rpc: "http://127.0.0.1:8545",
				Secret: ".secret",
			},
		},
	}

	// check for name flag
	if ctx.String("name") == "" {
		config.Name = expocli.Prompt("Project name: ")
	} else {
		config.Name = ctx.String("name")
	}

	// get list of available solc-js version
	solcList := expoutils.GetSolcjsList()
	// if flags for compiler version is set, compare with available list
	// and if not ask user to choose one of available version from list
	if ctx.String("compiler.version") == "" || !expoutils.Contains(ctx.String("compiler.version"), solcList) {
		config.Compiler.Version = expocli.PromptSelect("Select solidity version", solcList)
	} else {
		config.Compiler.Version = ctx.String("compiler.version")
	}

	// set other extra option for the compiler
	config.Compiler.Optimize = ctx.Bool("compiler.optimize")
	config.Compiler.Runs = ctx.Uint("compiler.optimize-runs")

	// begin write the config
	err = expoutils.WriteJson("./expo.json", config)
	if err != nil {
		s.StopFailMessage(red(" cannot write configuration file"))
		s.StopFail()

		expoutils.Fatalf("failed writing config file > %v", err)
	}

	s.StopMessage(green(" Happy hacking!!"))
	s.Stop()

	return nil
}

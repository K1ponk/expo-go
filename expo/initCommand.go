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

	s.StopMessage(green(" Happy hacking!!"))
	s.Stop()

	return nil
}

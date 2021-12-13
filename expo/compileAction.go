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
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/penta-expo/expo-go/expocli"
	"github.com/penta-expo/expo-go/expoutils"
	"github.com/urfave/cli/v2"
)

// compileAction is action for CompileCommand
func compileAction(ctx *cli.Context) error {
	// configure spinner and color
	red := color.New(color.FgHiRed).SprintfFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()

	s := expocli.Spinners()

	// check config file
	s.Message(yellow(" checking config file"))
	s.StopFailMessage(red(" failed getting config file."))
	s.StopMessage(green(" config checked."))

	s.Start()

	err := expoutils.CheckPath("./expo.json")
	if err != nil {
		s.StopFail()

		expoutils.Fatalf(" error > %v", err)
	}

	s.Stop()

	// check the solidity file
	s.Message(yellow(" checking file to compile"))
	s.StopFailMessage(red(" failed checking file"))
	s.StopMessage(green(" file ready"))

	s.Start()

	contractsDir := "contracts/"

	// check the directory
	err = expoutils.CheckPath(contractsDir)
	if err != nil {
		s.StopFail()

		expoutils.Fatalf("error > %v", err)
	}

	// get all file
	dirFile, err := ioutil.ReadDir(contractsDir)
	if err != nil {
		s.StopFail()

		expoutils.Fatalf("error > %v", err)
	}

	var solfile []string
	for _, df := range dirFile {
		if !df.IsDir() {
			solfile = append(solfile, "contracts/" + df.Name())
		}
	}

	s.Stop()

	// read the config
	s.Message(yellow(" read config file"))
	s.StopFailMessage(red(" failed reading config file"))
	s.StopMessage(green(" config file readed."))

	s.Start()

	var cfg expocli.Config
	err = expoutils.ReadJson("./expo.json", &cfg)
	if err != nil {
		s.StopFail()

		expoutils.Fatalf("error > %v", err)
	}

	s.Stop()

	// checking solidity compiler
	s.Message(yellow(" checking solidity compiler"))
	s.StopMessage(green(" solidity compiler checked."))
	s.StopFailMessage(red(" failed checking solidity"))

	s.Start()

	err = expoutils.CheckCommand("solcjs")
	if err != nil {
		// install the compiler
		err = installSolc(cfg.Compiler.Version)
		if err != nil {
			s.StopFail()

			expoutils.Fatalf("error > %v", err)
		}
	} else {
		// check if current version match with config file
		fullVer, err := exec.Command("solcjs", "--version").Output()

		curVer := strings.Split(string(fullVer), "+")[0]

		if curVer != cfg.Compiler.Version {
			err = installSolc(cfg.Compiler.Version)
			if err != nil {
				s.StopFail()

				expoutils.Fatalf("error > %v", err)
			}
		}
	}

	s.Stop()

	// configure compiler command with config file
	s.Message(yellow(" compiling contracts"))
	s.StopMessage(green(" contracts compiled"))
	s.StopFailMessage(red(" failed compiling contracts"))

	s.Start()

	args := []string{"--bin", "--abi", "--base-path", "contracts/", "--output-dir", "build/"}

	if cfg.Compiler.Optimize {
		opt := []string{"--optimize", "--optimize-runs", string(cfg.Compiler.Runs)}

		args = append(args, opt...)
	}

	cmd := exec.Command("solcjs", append(args, solfile...)...)
	cmd.Stdout = os.Stdout
	
	if err = cmd.Run(); err != nil {
		s.StopFail()

		expoutils.Fatalf("error > %v", err)
	}

	s.Stop()

	return nil
}

// installSolc will install solidity compiler
func installSolc(version string) error {
	cmd := exec.Command("npm", "install", "solc@" + version, "--global", "--silent")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

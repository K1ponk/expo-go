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

package expoutils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

// WriteJson is a function for writing json file
func WriteJson(target string, data interface{}) error {
	// convert the file into json
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// write json data into a file
	return ioutil.WriteFile(target, jsonData, 0666)
}

// ReadFile is a function for read a file
func ReadFile(target string) ([]byte, error) {
	// open the file for read
	f, err := os.Open(target)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// get the content
	return ioutil.ReadAll(f)
}

// MaptoStruct is a function to convert map into struct
func MaptoStruct(data interface{}, result interface{}) error {
	// marshal the data into json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// unnarshal data into struct
	return json.Unmarshal(jsonData, &result)
}

// CheckCommand is a function for checking command/programs wether exist or not
func CheckCommand(command string) error {
	// check the command and return the error
	_, err := exec.LookPath(command)

	return err
}

// CheckPath is a function for checking a file or directory wether exist or not
func CheckPath(target string) error {
	// check the file or directory
	_, err := os.Stat(target)

	return err
}

// Clear is a function to clear terminal screen
func Clear() {
	// specify the right command for clearing terminal screen
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "android":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	// run the command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fatalf will format message and exit the app
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}

	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

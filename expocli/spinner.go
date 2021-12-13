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
	"time"

	"github.com/penta-expo/expo-go/expoutils"
	"github.com/theckman/yacspin"
)

// Spinners will create new spinner
func Spinners() *yacspin.Spinner {
	// config for spinner
	cfg := yacspin.Config{
		Frequency: 100 * time.Millisecond,
		CharSet: yacspin.CharSets[41],
		SuffixAutoColon: true,
		Colors: []string{"fgYellow"},
		StopColors: []string{"fgGreen"},
		StopCharacter: "√",
		StopFailColors: []string{"fgRed"},
		StopFailCharacter: "✗",
	}

	// create the app
	spinner, err := yacspin.New(cfg)
	if err != nil {
		expoutils.Fatalf("failed initiating spinner %v", err)
	}

	return spinner
}

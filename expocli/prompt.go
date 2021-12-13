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
	"github.com/manifoldco/promptui"
	"github.com/penta-expo/expo-go/expoutils"
)

// Prompt will send user a prompt
func Prompt(q string) string {
	// initiate promptui
	templ := promptui.PromptTemplates{
		Prompt: "{{. | yellow}}",
	}
	prompt := promptui.Prompt{
		Label: q,
		Templates: &templ,
	}

	res, err := prompt.Run()
	if err != nil {
		expoutils.Fatalf("failed run prompt > %v", err)
	}

	return res
}

// PromptSelect will send user select option
func PromptSelect(q string, data []string) string {
	// send user selection input
	templ := promptui.SelectTemplates{
		Label: "{{.}}",
		Inactive: "  {{. | yellow}}",
		Active: "> {{. | green}}",
	}
	prompt := promptui.Select{
		Label: q,
		Items: data,
		Templates: &templ,
		HideSelected: true,
	}

	_, res, err := prompt.Run()
	if err != nil {
		expoutils.Fatalf("error sending input > %v", err)
	}

	return res
}

// Copyright 2021 The expo-go Authors
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

// ContractsTokenTemplate is map for template token result
type ContractsTokenTemplate map[string][]ContractDetails

// ContractDetails details of ContractsTokenTemplate
type ContractDetails struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Raw string `json:"raw"`
}

var (
	// TokenTemplateUrl is url for token template
	TokenTemplateUrl = "https://raw.githubusercontent.com/penta-expo/expo-go/master/templates/contracts.json"
)

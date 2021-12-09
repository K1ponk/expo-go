//	Copyright (C) 2021 The penta-expo Authors
//
//	This program is free software: you can redistribute it and/or modify
//	it under the terms of the GNU General Public License as published by
//	the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//
//	This program is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU General Public License for more details.
//
//	You should have received a copy of the GNU General Public License
//	along with this program.  If not, see <https://www.gnu.org/licenses/>.

package expolibs

// Config structure
type Config struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Extras map[string]interface{} `json:"extras"`
	Networks map[string]interface{} `json:"networks"`
}

// Network structure
type Network struct {
	Rpc string `json:"rpc"`
	PrivateKey string `json:"privateKey"`
}

// Compiler structure
type Compiler struct {
	Version string `json:"version"`
	Optimize bool `json:"optimize"`
	Runs uint `json:"runs"`
}

// SolcCommand is command for solcjs
const SolcCommand = `solcjs{{if .Optimize}} --optimize --optimize-runs {{.Runs}}{{end}} --bin --output-dir ./build --base-path ./contracts --include-path ./node_modules --verbose `

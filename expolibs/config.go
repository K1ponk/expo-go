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

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// WriteConfig for write expo.json confg file
func WriteConfig(data Config) error {
	// Convert struct to json
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Write into expo.json
	err = ioutil.WriteFile("./expo.json", jsonData, 0666)
	if err != nil {
		return err
	}

	return nil
}

// ReadConfig for read expo.json config file
func ReadConfig() ([]byte, error) {
	// Open the file for read
	f, err := os.Open("./expo.json")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// Read the file and return the result
	return ioutil.ReadAll(f)
}

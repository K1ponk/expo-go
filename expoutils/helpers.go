package expoutils

import (
	"encoding/json"
	"net/http"
	"strings"
)

type solcFetch struct{
	Refs string `json:"ref"`
}
// GetSolcjsList will fetch available solc-js version
func GetSolcjsList() []string {
	// fetch data from github
	r, err := http.Get("https://api.github.com/repos/ethereum/solc-js/git/refs/tags")
	if err != nil {
		Fatalf("failed fetching list of solc-js version > %v", err)
	}
	defer r.Body.Close()

	// decode it to struct
	var body []solcFetch
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		Fatalf("error decoding json response > %v", err)
	}

	// return the result in []string type
	var res []string
	for i := range body {
		res = append(res, strings.ReplaceAll(body[i].Refs, "refs/tags/v", ""))
	}

	// re arange the list
	var resSort []string
	for i := len(res) -1; i > 0; i-- {
		resSort = append(resSort, res[i])
	}

	return resSort
}

// Contains will check if x was in []string
func Contains(x string, data []string) bool {
	// check the x
	for _, y := range data {
		if y == x {
			return true
		}
	}

	return false
}

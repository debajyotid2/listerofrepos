// MIT License
// 
// Copyright (c) 2024 Debajyoti Debnath
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"errors"
	"fmt"
	"github.com/pterm/pterm"
	"io"
	"net/http"
)

// Sets up a GET request to the GitHub API based on the user input programming language
func SetupRequest(lang *string) (*http.Request, error) {
	if lang == nil {
		return nil, errors.New("No language provided for making request.")
	}
	requestURL := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars&order=desc", *lang)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// Processes the response received after making the GET request
func ProcessResponse(res *http.Response, nRepos *int) (pterm.TableData, error) {
	if res.StatusCode != 200 {
		return nil, errors.New("The response received is not valid.")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result, err := ReadJson(resBody)
	if err != nil {
		return nil, err
	}

	// Iterate through all returned repositories
	allRepoData := result["items"].([]interface{})

	// Define the data for the table to be rendered
	tableData := pterm.TableData{
		{"Name", "Stars", "Forks", "Watchers", "Issues"},
	}

	for idx, repo := range allRepoData {
		if idx >= *nRepos {
			break
		}
		repoMap := repo.(map[string]interface{})
		repoData := RepoData{
			repoMap["full_name"].(string),
			int(repoMap["stargazers_count"].(float64)),
			int(repoMap["forks_count"].(float64)),
			int(repoMap["watchers_count"].(float64)),
			int(repoMap["open_issues_count"].(float64)),
		}
		tableData = append(tableData, repoData.ToStringSlice())
	}
	return tableData, nil
}

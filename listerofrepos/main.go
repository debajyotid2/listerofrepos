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
	"flag"
	"github.com/pterm/pterm"
	"log"
	"net/http"
	"slices"
	"time"
)

// Timeout for the http client
const TIMEOUT_SECS = 30

// Languages that are supported by the GitHub API
var SUPPORTED_LANGUAGES = []string{
	"c", "cpp", "csharp", "go",
	"java", "javascript", "php", "python", "ruby",
	"rust", "scala", "swift", "typescript",
}

func main() {
	nRepos := flag.Int("n", 100, "Number of repositories to list (cannot exceed 500). Default = 100.")
	lang := flag.String("lang", "c", "Language to fetch repositories for. Default = c.")

	flag.Parse()

	// Validate command line arguments
	if !slices.Contains(SUPPORTED_LANGUAGES, *lang) {
		log.Fatalf("Language %s is not supported by the GitHub API.\n", *lang)
	}
	if *nRepos < 0 || *nRepos > 1000 {
		log.Fatalf("Number of repos requested must be between 1 and 1000.")
	}

	// Make http client
	client := http.Client{
		Timeout: TIMEOUT_SECS * time.Second,
	}

	// Set up and make request
	req, err := SetupRequest(lang)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Process response
	tableData, err := ProcessResponse(res, nRepos)
	if err != nil {
		log.Fatal(err)
	}

	// Create a table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()

	pterm.Println() // Blank line
}

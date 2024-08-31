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
	"net/http"
	"testing"
	"time"
)

// Test for request setup with language as nil
func TestSetupRequestLangNil(t *testing.T) {
	_, err := SetupRequest(nil)
	if err == nil {
		t.Fatalf("Providing a nil language should raise an error.")
	}
}

// Test for correct request setup
func TestSetupRequestCorrectArgs(t *testing.T) {
	lang := "cpp"
	req, err := SetupRequest(&lang)
	if err != nil {
		t.Fatalf("Http request set up should not raise an error")
	}
	if req.Method != http.MethodGet {
		t.Fatalf("Http request method must be GET.")
	}
}

// Test if the data returned from ProcessResponse has the correct dimensions
func TestProcessResponseDataDimensions(t *testing.T) {
	nRepos := 18
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/search/repositories?q=language:cpp&sort=stars&order=desc", nil)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, _ := client.Do(req)
	tableData, err := ProcessResponse(res, &nRepos)
	if err != nil {
		t.Fatalf("Could not process response correctly.")
	}
	if len(tableData) != nRepos+1 {
		t.Fatalf("Could not fetch the right number of repos.")
	}
}

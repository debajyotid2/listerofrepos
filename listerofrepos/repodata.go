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
	"fmt"
	"github.com/pterm/pterm"
)

// Struct to hold information on a single GitHub repository
type RepoData struct {
	name     string
	stars    int
	forks    int
	watchers int
	issues   int
}

const WIDTH = 6 // Number of places to format integers within

// Prints a RepoData instance
func (r *RepoData) Print() {
	pterm.DefaultBasicText.Printf("Name: " + pterm.Yellow(r.name))
	pterm.DefaultBasicText.Printf("  Stars: " + pterm.LightGreen(r.stars))
	pterm.DefaultBasicText.Printf("  Forks: " + pterm.LightCyan(r.forks))
	pterm.DefaultBasicText.Printf("  Watchers: " + pterm.LightWhite(r.watchers))
	pterm.DefaultBasicText.Printf("  Open issues: " + pterm.LightRed(r.issues))
	pterm.DefaultBasicText.Println()
}

// Converts a RepoData instance to a string slice
func (r *RepoData) ToStringSlice() []string {
	return []string{
		r.name,
		fmt.Sprintf("%*d", WIDTH, r.stars),
		fmt.Sprintf("%*d", WIDTH, r.forks),
		fmt.Sprintf("%*d", WIDTH, r.watchers),
		fmt.Sprintf("%*d", WIDTH, r.issues),
	}
}

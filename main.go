package main

import "github.com/chenxifun/ruanzhu-code-merge/core"

func main() {

	c := core.CodeArrange{
		PathList: []string{
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/main.go",
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/core",
		},
		OutPath: "./code1.txt",
		Ext:     []string{".go", ".proto"},
	}
	c.Start()
}

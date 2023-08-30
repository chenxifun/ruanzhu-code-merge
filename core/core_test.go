package core

import (
	"encoding/json"
	"github.com/chenxifun/ruanzhu-code-merge/comm"
	"testing"
)

func TestReadDir(t *testing.T) {

	c := CodeArrange{
		PathList: []string{
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/main.go",
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/core",
		},
		OutPath: "./code1.txt",
		Ext:     []string{".go", ".proto"},
	}
	c.Start()
}

func TestReadFile(t *testing.T) {
	c := CodeArrange{
		PathList: []string{
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/main.go",
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/core",
		},
		Ext:      nil,
		fileList: nil,
	}

	bs := c.readFile("/Users/gaochenxi/Work/distributed-network/dnet-p2p/cmd/dnet/main.go")
	bs = c.convCode(bs)
}

func TestImportJson(t *testing.T) {
	c := CodeArrange{
		PathList: []string{
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/main.go",
			"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/core",
		},
		Ext:     []string{".go", ".proto"},
		OutPath: "./code.txt",
	}

	jb, _ := json.MarshalIndent(&c, "", "\t")
	comm.WriteFile(jb, "../config.json", true)
}

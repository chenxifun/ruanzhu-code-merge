package core

import (
	"github.com/chenxifun/ruanzhu-code-merge/comm"
	"os"
	"regexp"
	"strings"
)

type CodeArrange struct {
	// PathList 代码文件路径，已去重
	PathList []string `json:"pathList,omitempty" yaml:"pathList"`

	// Ext 读取后缀
	Ext []string `json:"ext,omitempty" yaml:"ext"`

	// OutPath 输出的文件路径
	OutPath  string `json:"outPath,omitempty" yaml:"outPath"`
	fileList []string

	codeContent []byte
}

func (a *CodeArrange) Start() error {
	if len(strings.TrimSpace(a.OutPath)) == 0 {
		a.OutPath = "code.txt"
	}
	a.readAllFilePath()
	a.readCode()
	return a.writeFile()
}

func (a *CodeArrange) readCode() {

	for _, s := range a.fileList {
		bs := a.readFile(s)
		bs = a.convCode(bs)
		a.codeContent = append(a.codeContent, bs...)
	}
}

func (a *CodeArrange) readFile(p string) []byte {

	bytes, err := os.ReadFile(p)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func (a *CodeArrange) convCode(c []byte) []byte {
	s := string(c)
	re := regexp.MustCompile(`(?m)^[ \t]*//[^\n]*\n?|/\*.*?\*/`)
	// 将注释替换为空
	result := re.ReplaceAllString(s, "")

	result = strings.ReplaceAll(result, "\n\n", "\n")
	return []byte(result)
}

func (a *CodeArrange) writeFile() error {

	return comm.WriteFile(a.codeContent, a.OutPath, true)
}

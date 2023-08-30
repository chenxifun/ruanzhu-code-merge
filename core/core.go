package core

import (
	"errors"
	"os"
	"path"
	"regexp"
	"strings"
)

type CodeArrange struct {
	// PathList 代码文件路径，已去重
	PathList []string

	// Ext 读取后缀
	Ext []string

	// OutPath 输出的文件路径
	OutPath  string
	fileList []string

	codeContent []byte
}

func (a *CodeArrange) Start() error {
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

	return WriteFile(a.codeContent, a.OutPath, true)
}

func WriteFile(data []byte, filePath string, cover bool) error {

	if !cover {
		if _, err1 := os.Stat(filePath); !os.IsNotExist(err1) {
			return errors.New("file is exist")
		}
	}

	err := os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0600)

}

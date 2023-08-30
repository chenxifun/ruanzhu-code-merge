package core

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (a *CodeArrange) hasFilePath(p string) bool {
	for _, s := range a.fileList {
		if s == p {
			return true
		}
	}
	return false
}

func (a *CodeArrange) setFilePath(filePaths []string) {
	for _, filePath := range filePaths {
		if !a.hasFilePath(filePath) && a.isExt(filePath) {
			a.fileList = append(a.fileList, filePath)
		}
	}
}

func (a *CodeArrange) isExt(p string) bool {
	ext := filepath.Ext(p)
	for _, s := range a.Ext {
		if strings.ToLower(s) == strings.ToLower(ext) {
			return true
		}
	}
	return false
}

func (a *CodeArrange) readAllFilePath() {
	for _, s := range a.PathList {
		a.setFilePath(readFilePath(s))
	}
}

func readFilePath(pathStr string) []string {
	var list []string

	s, err := os.Stat(pathStr)
	if err != nil {
		return list
	}

	if s.IsDir() {
		fs, err := os.ReadDir(pathStr)
		if err != nil {
			return list
		}

		for _, f := range fs {
			dir := path.Join(pathStr, f.Name())
			list = append(list, readFilePath(dir)...)
		}
		return list
	} else {
		list = append(list, pathStr)
		return list
	}
}

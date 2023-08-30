# ruanzhu-code-merge

合并代码到一个文件

## Install
``` shell
go install github.com/chenxifun/ruanzhu-code-merge@latest
```

## Use

```shell
ruanzhu-code-merge -c config.json
```

## Config
``` json
{
	"pathList": [
		"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/main.go",
		"/Users/gaochenxi/Code/github/chenxifun/ruanzhu-code-merge/core"
	],
	"ext": [
		".go",
		".proto"
	],
	"outPath": "./code.txt"
}
```

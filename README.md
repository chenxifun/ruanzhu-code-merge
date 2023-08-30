# ruanzhu-code-merge
合并代码到一个文件
## 参考
```
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

```

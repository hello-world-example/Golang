# VsCode 配置



## 插件配置

- 扩展商店搜索： "Go" 或 "Go for Visual Studio Code" 进行安装
- `Command + Shift + P` 搜索 `go:install/update tools` 全部勾选
- `Command + Shift + P` 搜索 `Go`，在 `settings.json` 文件中编辑，配置参考

```json
{
  "go.inferGopath": true,
  "go.autocompleteUnimportedPackages": true,
  "go.gocodePackageLookupMode": "go",
  "go.gotoSymbol.includeImports": true,
  "go.useCodeSnippetsOnFunctionSuggest": true,
  "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
  "go.docsTool": "gogetdoc",
  "go.formatTool": "goimports",
}
```



## 已安装的 Go 插件

```bash
go @installed 
```



## Read More

- [VS Code 中的代码自动补全和自动导入包](https://maiyang.me/post/2018-09-14-tips-vscode/)
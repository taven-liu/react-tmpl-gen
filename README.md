# React Tmpl Gen
## 生成可执行文件

- `go mod tidy`
- `go build -o rc `

## 配置
1. 新建目录`~/.react-tmpl-gen`
2. 拷贝可执行文件`rc`及`templates`至上一步骤目录
3. 添加步骤1到`PATH`
## 使用
1. 进入到需要生成的模版代码的父级目录
2. 运行`rc g page demo`生产`demo`页面

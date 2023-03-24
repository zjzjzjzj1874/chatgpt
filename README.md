# chatgpt
使用命令行发送chatgpt的api

## 准备

### 二进制文件构建
```Bash
make build
```
cmd目录下会生成gptx的文件

### 获取gpt key

### 设置运行前的环境变量
```Bash
export GPT_KEY="GPT_KEY"
```

```Bash
cd ./cmd && ./gptx -h

gptx is a command-line tool for call openai api

Usage:
  gptx [command]

Available Commands:
  chat        send chat content to chatgpt
  help        Help about any command
  model       list all models that gpt support

Flags:
  -h, --help      help for gptx
  -v, --version   version for gptx

Use "gptx [command] --help" for more information about a command.
```

## 用法

### 查看gpt模型

```Bash
./gptx model

Total Model: 68
ID: babbage
ID: davinci
ID: gpt-3.5-turbo
...
```
### 进行chat对话

```Bash
./gptx chat -c 你是谁

我是一名AI语言模型，服务于OpenAI，可执行各种任务和回答各种问题。
```

## TODO list
- [ ] 选择一个彩色的标准输出

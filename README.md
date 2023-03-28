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
  audio       turn audio into text.
  chat        creates a completion for the chat message
  edit        Given a prompt and an instruction, the model will return an edited version of the prompt.
  help        Help about any command
  img         creates an image given a prompt.
  model       lists the currently available models,

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
./gptx chat -p 你是谁

我是一名AI语言模型，服务于OpenAI，可执行各种任务和回答各种问题。
```

### 图片生成

```Bash
-p: 生成的图片要求
-n: 生成的图片数量(1-10)
-s:生成的图片大小(256x256,512x512,1024x1024)
./gptx img -p "可爱的小猫咪" -n 2 -s 256x256

Total Image: 2
Url: https://oaidalleapiprodscus.blob.core.windows.net/private/org-FszeU94XqTOxWst1f2mp5LpO/user-qcjpFAv1q7NKNH42MHry25KB/img-r3lAOCz0DSmypxl3X5w3ZWyE.png?st=2023-03-24T05%3A27%3A14Z&se=2023-03-24T07%3A27%3A14Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-23T22%3A08%3A23Z&ske=2023-03-24T22%3A08%3A23Z&sks=b&skv=2021-08-06&sig=%2BaFB5nW23BeT6XGdrcSS1M2wvWeWbywJnebdp9wdza8%3D
Url: https://oaidalleapiprodscus.blob.core.windows.net/private/org-FszeU94XqTOxWst1f2mp5LpO/user-qcjpFAv1q7NKNH42MHry25KB/img-r3XgIswuunVwZ6NlwP0NnUAG.png?st=2023-03-24T05%3A27%3A14Z&se=2023-03-24T07%3A27%3A14Z&sp=r&sv=2021-08-06&sr=b&rscd=inline&rsct=image/png&skoid=6aaadede-4fb3-4698-a8f6-684d7786b067&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2023-03-23T22%3A08%3A23Z&ske=2023-03-24T22%3A08%3A23Z&sks=b&skv=2021-08-06&sig=nvVZDD3hsaxPtaS9sxyfvwr2x7u0mF4/9cbts8t60I0%3D

```

### 音频转文字

```Bash
-f: 待转文件
-m: gpt模型,默认使用whisper-1
-l: 语言,参考 https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
./gptx audio trans -f 5.6.mp3 -l en

翻译结果:John, John, you are so dumb. John, John, you are so dumb. John, John, you are so dumb. John, John, you are so dumb.
```

### 编辑提示

```Bash
-i: 待编辑内容
-p: 提示,指示
-m: gpt模型,默认:text-davinci-edit-001
./gptx edit -i "What day of the wek is it?" -p "语法修复"

您的输入:What day of the wek is it?
输出:0.what day of the week is it?
```

## TODO list
- [x] 选择一个彩色的标准输出 [color](github.com/fatih/color)

## 参考资料
- [OpenAI官网](https://openai.com/)
- [Chatgpt客户端](https://chat.openai.com/chat)
- [OpenAI开发者中心](https://platform.openai.com/)
- [OpenAI文档中心](https://platform.openai.com/docs/introduction)
- [OpenAI Api参考](https://platform.openai.com/docs/api-reference)
- [OpenAI计费详情](https://openai.com/pricing)
- [OpenAI计费查询](https://platform.openai.com/account/usage)
- [OpenAI Key生成](https://platform.openai.com/account/api-keys)
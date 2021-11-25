# Bilibili-freestyle

## 前情提要

- 创作背景：做视频，拿来freestyle用
- 注意点：哥们毫无工程化可言，写的比较随意
- 期望：有兴趣的哥们给点建议，[视频](https://www.bilibili.com/video/BV1Qf4y1K7ir) or 代码都ok

## 功能

- 爬取对应`MID`的b站所有视频的评论、弹幕（`MID`在`main.go`中）
- 使用jieba分词，分词后写入本地文件
- 读入分词文件，做过滤操作
    - 去重
    - 去部分标点
    - 去部分语助词

- 开始后，每过5秒开始弹出一个上面处理好的词库里的词语

最后，enjoy yourself～😄

## 使用

- 开始：输入任意字符，回车
- 结束：直接回车

## 编译 && 运行

### 编译

在工程目录下

```shell
go build
```

### 运行

在工程目录下

```shell
./bilibili-freestyle
```


# generateQRCode
> 实现生成带logo的二维码，并在底部添加自己想要的文字
## 简单使用
```bash
go run main.go -text="https://www.baidu.com" -logo="logo.jpg" -out="out.jpg" -bottom="这里是底部文字"
```

## 参数介绍
```text
-text    二维码扫码后内容
-logo    logo文件图标
-percent 二维码Logo的显示比例(默认15%)
-size    二维码的大小(默认256)
-out     输出文件
-bottom  底部文字
```

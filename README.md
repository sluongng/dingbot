# DingBot

## Description

- Massively reworked to make the library easier to use
- Struct comes with builder and opinionated Simple Message that can be used immediately
- Intuitive methods

## How to use

``` go
// If your webhook is https://oapi.dingtalk.com/robot/send?access_token=123456ABCD
// your accessToken is 123456ABCD
const accessToken = "YOUR_ACCESS_TOKEN"

func main() {
	_ = dingbot.SimpleTextMessage("Xin Chao 123").Send(accessToken)
}
```

- See [examples](./example/main.go)

## TODO

- [ ]  Proper testing
- [ ]  Improve code style

## Others

- Official [Java SDK](https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.karFPe&treeId=257&articleId=105735&docType=1#s4)
- [Nodejs SDK](https://github.com/x-cold/dingtalk-robot/)

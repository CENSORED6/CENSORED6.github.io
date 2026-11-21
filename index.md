---
layout: default
---

# 你好，我是 Bo

计算机专业在读，喜欢写代码、记笔记。

## 最近在做什么

- 学Go和C
- 整理课程笔记
- 折腾这个网站
- 研读go圣经 [go圣经](https://www.kancloud.cn/hartnett/gopl-zh/126045)

# 学习路线

## Week 1
主线任务
- [X] 1. 注册GitHub、ChatGPT账号并成功访问
- [X] 2. 学习Markdown的写法，重点关注代码块、多级标题、数字/圆点列表、加粗、删除线......
- [X] 3. 利用Markdown编写简单的个人博客，并利用github pages部署
- [X] 4. 学习go的数组、切片，并完成以下任务：
    - 给定s := make([]byte, 5)，那么len(s) 和 cap(s) 分别是多少？令s = s[2:4]，len(s) 和 cap(s) 又分别是多少？
    - 比较字符串"hello，世界"的长度和for range该字符串的循环次数
- [X] 5. 将学习的笔记/代码上传到GitHub，建议新建一个学习仓库，把每周的任务分文件夹上传

学习成果：
[第一周](https://github.com/CENSORED6/CENSORED6.github.io/tree/main/muxi/week%201) 按照学习的时间

## Week 2
主线任务
- [X] 1. 理解map和slice内存扩容的区别
- [X] 2. 利用闭包实现一个计数器
- [X] 3. 基于下面这个结构体实现四个方法：
    - Average：返回平均分；Scores 为空时返回 0。
    - AddScore：追加一个成绩，调用方能看见这个改动。
    - Top：返回最高的 n 个成绩，从高到低。n 超过长度时返回全部，n <= 0 时返回 nil。不能修改 s.Scores 本来的顺序，也不能让返回的切片存在影响原数据的可能。
    - Merge：合并两个人的成绩，返回一个新的 Student（Name 用接收者的）。合并结果和两个原对象必须互不影响。
- [X] 4. 修改代码
<details>
<summary>代码</summary>

```go
package main

type message struct {
    Topic     string
    Partition int32
    Offset    int64
}

type FeedEventDM struct {
    Type    string
    UserID  int
    Title   string
    Content string
}

type MSG struct {
    ms        message
    feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
    var consumeMSG []MSG
    var lastConsumeTime time.Time // 记录上次消费的时间
    msgs := make(chan MSG)

    //这里源源不断的生产信息
    go func() {
       for i := 0; ; i++ {
          msgs <- MSG{
             ms: message{
                Topic:     "消费主题",
                Partition: 0,
                Offset:    0,
             },
             feedEvent: FeedEventDM{
                Type:    "grade",
                UserID:  i,
                Title:   "成绩提醒",
                Content: "您的成绩是xxx",
             },
          }
          //每次发送信息会停止0.01秒以模拟真实的场景
          time.Sleep(100 * time.Millisecond)
       }
    }()

    //不断接受消息进行消费
    for msg := range msgs {
       // 添加新的值到events中
       consumeMSG = append(consumeMSG, msg)
       // 如果数量达到额定值就批量消费
       if len(consumeMSG) >= ConsumeNum {
          //进行异步消费
          go func() {
             m := consumeMSG[:ConsumeNum]
             fn(m)
          }()
          // 更新上次消费时间
          lastConsumeTime = time.Now()
          // 清除插入的数据
          consumeMSG = consumeMSG[ConsumeNum:]
       } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
          // 如果距离上次消费已经超过5分钟且有未处理的消息
          if len(consumeMSG) > 0 {
             //进行异步消费 
             go func() {
                m := consumeMSG[:ConsumeNum]
                fn(m)
             }()
             // 更新上次消费时间
             lastConsumeTime = time.Now()
             // 清空插入的数据
             consumeMSG = consumeMSG[ConsumeNum:]
          }
       }
    }
}

func fn(m []MSG) {
    fmt.Printf("本次消费了%d条消息\n", len(m))
}
```
</details>

- [X] 5. 使用for循环生成20个goroutine,每个goroutine随机休眠0~1000ms，并向一个channel传入随机数和goroutine编号(从1-20)，等待这些goroutine都生成完后，想办法给这些随机数按照编号进行排序(输出排序前和排序后的结果,要求不使用额外的空间存储这20个数据)

学习成果：
[第二周](https://github.com/CENSORED6/CENSORED6.github.io/tree/main/muxi/week%201) 

## Week 3
主线任务
- [X] 1. 完成小游戏[木犀骇客](https://github.com/weiji6/http-theft-bank-start-template)
- [X] 2. 利用上周学习的并发基础内容（goroutine、channel、sync包），完成招新机试第四题

学习成果：
[第三周](https://github.com/CENSORED6/CENSORED6.github.io/tree/main/muxi/week3) 


## 找我

- GitHub：Bo(https://github.com/censored6)

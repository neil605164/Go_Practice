# Go_Practice

### Direct 模式

最簡單的模式，只會有一個 Producer 負責發送 message 到 Queue 裡、而也只有一個 Consumer 去 Queue 裡消費 message

![](https://kucw.github.io/images/blog/rabbitmq_direct.png)

### Worker 模式

跟 Direct 模式很像，但是差別是 Worker 模式中會 同時 有多個 Consumer 會去消費 Queue 裡的 message，增加 message 消化的速率（一個人做很慢，大家一起做就很快）

![](https://kucw.github.io/images/blog/rabbitmq_worker.png)
# RabbitMQ

## Queue

進入Queue的資料在有人收走以前是不會消失的

- name: string => queue 名稱
- durable: bool => 系統關機後會保存
- delete: bool => 執行完 queue 會自動刪除
- exclusive: bool
- no-wait: bool
- arguments: bool => 提供 header 值，可在 header exchange 使用

## Direct 模式

最簡單的模式，只會有一個 Producer 負責發送 message 到 Queue 裡、而也只有一個 Consumer 去 Queue 裡消費 message

![](https://kucw.github.io/images/blog/rabbitmq_direct.png)

---

## Worker 模式

跟 Direct 模式很像，但是差別是 Worker 模式中會 同時 有多個 Consumer 會去消費 Queue 裡的 message，增加 message 消化的速率（一個人做很慢，大家一起做就很快）

![](https://kucw.github.io/images/blog/rabbitmq_worker.png)

---

## Publish/Subscribe 模式
在先前 ```Direct``` 與 ```Worker``` 模式底下，我們已經認知的角色有：
- ```producer```: a user that sends messages
- ```queue```: a buffer or a space to stores messages
- ```consumer```: a user that reveives messages

```Publish/Subscribe``` 模式的 ```exchange``` 在這扮演了轉遞的角色，目的就是希望 ```producer``` 不要直接將訊息送到 ```queue``` , 事實上使用情境 ```producer``` 也不應該要知道訊息會被送到哪個 ```queue``` 或哪些 ```queue```，那麼 ```Publish/Subscribe``` 模式又可以實現哪些事情:

- 可以過濾掉不需要進入 ```queue``` 的訊息
- 可以清楚了解接收多少資料
- 判斷當前資料該丟向哪個 ```queue``` 或者哪些不同的 ```queue```

### exchange type:
- direct: 照Routing Key將資料倒給對應的Queue，如果沒有綁定Queue會自動丟棄進來的資料
- topic: 同direct，差別: topic支援 # 或 * 的模糊搜尋，如果沒有綁定Queue會自動丟棄進來的資料
    - `#` : 多字串
    - `*`: 單一字串
- headers: 根據資料的header做導向(不實用)
- fanout: 將資料倒給所有已綁定在的 queue，如果沒有綁定Queue會自動丟棄進來的資料

### 注意:
- 如果 ```exchange``` 沒有綁定 ```queue```，訊息將會流失或被丟棄。

![](https://kucw.github.io/images/blog/rabbitmq_subscribe.png)

---

## Router 模式
```router``` 是強化 ```exchange``` 的功能而存在。
當 ```Producer``` 把 ```message``` 丟給 ```Exchange``` 時，須同時附上一個 ```routing key``` ， 而 ```Exchange``` 就會根據這個 ```routing key```，將 ```message``` 丟到指定的 ```Queue``` 上

![](https://kucw.github.io/images/blog/rabbitmq_routing1.png)

## Topic 模式
```Topic``` 與 ```router``` 相同，只是 ```routing key``` 可以進化成使用「**模糊**」綁定。
- ```*``` : 可以取代一個文字
- ```#``` : 可以取代 0 ~ N 個文字

### example: 
- ```quick.orange.rabbit``` : 該 ```router key``` 會被送往 ```Q1```,```Q2``` ```borker```
- ```lazy.orange.elephant``` : 該 ```router key``` 會被送往 ```Q1```,```Q2``` ```borker```
- ```quick.orange.fox``` : 該 ```router key``` 會被送往 ```Q1``` ```borker```
- ```lazy.brown.fox``` : 該 ```router key``` 會被送往 ```Q2``` ```borker```
- ```lazy.pink.rabbit``` : 該 ```router key``` 會被送往 ```Q2``` ```borker```, 但是只會被推送一次
- ```quick.brown.fox``` : 該 ```router key``` 不會被送到 ```broker``` ，且會被忽遺棄掉
- ```orange``` : 該 ```router key``` 不會被送到 ```broker``` ，且會被忽遺棄掉
- ```quick.orange.male.rabbit``` : 該 ```router key``` 不會被送到 ```broker``` ，且會被忽遺棄掉
- ```lazy.orange.male.rabbit``` : 該 ```router key``` 會被送往 ```Q2``` ```borker```

![](https://kucw.github.io/images/blog/rabbitmq_topics.png)
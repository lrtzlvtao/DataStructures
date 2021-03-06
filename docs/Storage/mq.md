# MQ

## MQ 选择

- RabbitMQ

  - 特色
    - 支持非常灵活的路由配置（Exchange）
    - 支持的客户端语言多
  - 缺点
    - 对消息堆积的支持不好
    - 性能较差
    - 使用 Erlang 开发，二次开发及维护性差

- RocketMQ

  - 特色
    - 有着不错的性能，稳定性和可靠性，具备一个现代的消息队列应该有的几乎全部功能和特性
    - 使用 Java 开发，容易扩展和二次开发
  - 缺点
    - 国产消息队列，周边生态较差

- Kafka

  - 特色
    - 有着不错的性能，稳定性和可靠性，具备一个现代的消息队列应该有的几乎全部功能和特性
    - 与周边生态系统的兼容性是最好的没有之一，优势在于专为超高吞吐量的实时日志采集、实时数据同步、实时数据计算等场景来设计
    - 使用 Scala 和 Java 语言开发，设计上大量使用了批量和异步的思想，这种设计使得 Kafka 能做到超高的性能
    - 异步批量的设计性能极高
  - 缺点
    - 同步收发消息的响应时延比较高，不太适合在线业务场景

## MQ 的常见问题

- 消息的顺序问题
- 消息的重复问题

## rabbitmq 广播类型

- fanout: 所有 bind 到此 exchange 的 queue 都可以接收消息（纯广播，绑定到 RabbitMQ 的接受者都能收到消息）
- direct: 通过 routingKey 和 exchange 决定的那个唯一的 queue 可以接收消息
- topic: 所有符合 routingKey(此时可以是一个表达式)的 routingKey 所 bind 的 queue 可以接收消息

## 如何保证 RabbitMQ 消息的顺序性

- 一个 queue 对应一个 consumer
- 一个 queue 但是对应一个 consumer，然后这个 consumer 内部用内存队列做排队，然后分发给底层不同的 worker 来处理

## 消息如何分发

- 消息将以循环（round-robin）的方式发送给订阅的消费者

## 消息怎么路由

- fanout：如果交换器收到消息，将会广播到所有绑定的队列上
- direct：如果路由键完全匹配，消息就被投递到相应的队列
- topic：可以使来自不同源头的消息能够到达同一个队列。 使用 topic 交换器时，可以使用通配符

## 消息基于什么传输

- RabbitMQ 使用信道的方式来传输数据。信道是建立在真实的 TCP 连接内的虚拟连接，且每条 TCP 连接上的信道数量没有限制

## 如何确保消息正确地发送至 RabbitMQ？ 如何确保消息接收方消费了消息

- 发送方确认模式
  - 将信道设置成 confirm 模式（发送方确认模式），则所有在信道上发布的消息都会被指派一个唯一的 ID
  - 一旦消息被投递到目的队列后，或者消息被写入磁盘后（可持久化的消息），信道会发送一个确认给生产者
  - 发送方确认模式是异步的，生产者应用程序在等待确认的同时，可以继续发送消息。当确认消息到达生产者应用程序，生产者应用程序的回调方法就会被触发来处理确认消息
- 接收方确认机制
  - 消费者接收每一条消息后都必须进行确认（消息接收和消息确认是两个不同操作）。只有消费者确认了消息，RabbitMQ 才能安全地把消息从队列中删除

## 如何保证 RabbitMQ 消息的可靠传输

- 生产者丢失消息

  - RabbitMQ 提供 transaction 和 confirm 模式来确保生产者不丢消息
    - transaction 机制
      - 发送消息前，开启事务，然后发送消息，如果发送过程中出现什么异常，事务就会回滚（channel.txRollback()），如果发送成功则提交事务。然而，这种方式有个缺点：吞吐量下降；
    - confirm
      - 一旦 channel 进入 confirm 模式，所有在该信道上发布的消息都将会被指派一个唯一的 ID（从 1 开始），一旦消息被投递到所有匹配的队列之后；abbitMQ 就会发送一个 ACK 给生产者

- 消息队列丢数据

  - 消息持久化
    - 将 queue 的持久化标识 durable 设置为 true, 则代表是一个持久的队列
    - 消息推送投递模式必须设置持久化，deliveryMode=2

- 消费者丢失消息
  - 自动确认消息模式，改为手动确认消息即可

## 为什么不应该对所有的 message 都使用持久化机制

- 然导致性能的下降，因为写磁盘比写 RAM 慢的多

## RabbitMQ 的集群

- 单机模式
- 普通集群模式
  - 在多台机器上启动多个 RabbitMQ 实例，每个机器启动一个
  - 创建的 queue，只会放在一个 RabbitMQ 实例上，但是每个实例都同步 queue 的元数据
  - 消费的时候，实际上如果连接到了另外一个实例，那么那个实例会从 queue 所在实例上拉取数据过来
- 镜像集群模式
  - 创建的 queue，无论元数据还是 queue 里的消息都会存在于多个实例上
  - 每次你写消息到 queue 的时候，都会自动把消息同步到多个实例的 queue 上
  - 性能开销太大

## 延时以及过期失效

- 消息积压

  - 紧急扩容
  - 提高消费效率

- 消息失效

  - RabbtiMQ 是可以设置过期时间的，超过一定的时间就会被 RabbitMQ 给清理掉

- mq 消息队列满了
  - 同上

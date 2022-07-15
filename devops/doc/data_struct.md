# 二叉数
## 四种遍历方式分别为：先序遍历、中序遍历、后序遍历、层序遍历。
>  https://www.cnblogs.com/du001011/p/11229170.html

## 二叉树遍历
> https://baike.baidu.com/item/%E4%BA%8C%E5%8F%89%E6%A0%91%E9%81%8D%E5%8E%86/9796049?fr=aladdin

线索二叉树， 如mysql 索引。
完全二叉数，  深度完全一样。

## https://www.zhihu.com/question/27542473/answer/736401344

## 横向遍历二叉树
https://blog.csdn.net/lbh199466/article/details/104944227

使用一个队列：
把root节点入列，
开始循环
pop队列，获得当下节点
如果有左节点入列，
如果有右节点入列，
打印当下节点。
结束循环

# 原子性和一致性的区别是什么？
一致性是业务意义上的
原子性、隔离性和持久性 是技术特征
https://www.cnblogs.com/stone94/p/10409669.html

事务的四大特性主要是：原子性（Atomicity）、一致性（Consistency）、隔离性（Isolation）、持久性（Durability）。

分布式事务
CAP是Consistency、Avaliability、Partitiontolerance三个词语的缩写，分别表示一致性、可用性、分区容忍性。

# 分布式事务的CAP理论
> https://zhuanlan.zhihu.com/p/102734515

- Two-phase commit protocol  有一个事件管理者，分别告知参与者 提交 或者  怎么感觉和TCC说的是一回事  强一致性 会升级为3阶段提交
- TCC  try commit cancle 补偿模式——最终一致性
 - 基于消息的最大努力通知 （可用性， 分区性。    最终一致，而非强一致）

# 分布式事务中常见的三种解决方案
https://www.cnblogs.com/bluemiaomiao/p/11216380.html

# RabbitMQ消息队列怎样做到服务宕机或重启消息不丢失
消息持久化
>https://www.cnblogs.com/mr-wuxiansheng/p/12809550.html

# 分布式事务2PC和TCC有啥不同
> https://blog.csdn.net/weixin_44062339/article/details/99711087
# TinyGoFrame
轻量级的go开发框架

学Go语言顺便敲的框架，基本模仿Aceld的zinx服务器框架，等后续学习再把这个框架写成分布式的

通过线程池+消息队列设计服务器模型，并进行了简单负载均衡

通过对报文的TLV序列化，解决tcp粘包问题，提供拆包封包方法

提供多路由服务，框架为开发者可以提供了注册任意业务服务

通过策略模式高效执行任务，统一以模板模式进行服务

# 相关需求
一个简单商店示例. 在这家商店中可以:
1. 列出产品
2. 下订单
3. 在远程支付提供商中初始化付款
4. 接收付款的通知, 并将付款的订单标记为已付款

# 业务分析
1. 相关实体: 产品、订单、第三方支付
1. 业务描述的功能:
    1. 产品列表
    1. 创建订单(未支付)
    1. 第三方支付, 通知修改订单
    1. 付款订单(原未支付) 标记为 已支付
1. 隐含业务行为:
    1. 产品库存与订单问题:
        1. 不考虑库存
            1. 不会有库存问题, 直接下单(业务中未描述, 当前场景的需求最简单)
            1. ~~可能会有库存问题, 但可以取消订单(业务中未描述, 不管)~~
        1. ~~有库存(业务中未描述, 不管)~~
            1. ~~一个订单数量对应一个库存~~
    1. 支付成功通知
        1. ~~第三方支付保障支付到位(这边不需要管)~~
        1. 收到通知后, 需要去验证订单是否被支付

## 四色原型进行聚合设计

> For the purposes of the article let’s use an example of a simple shop. This shop will allow us to: list products, place order, initialize payment in remote payments provider, receive notifications about payments and mark order as paid. Let’s take a look at the architecture diagram.
# 相关链接
- [monolith-microservice-shop](https://github.com/ThreeDotsLabs/monolith-microservice-shop)
- [microservices-or-monolith-its-detail](https://threedots.tech/post/microservices-or-monolith-its-detail/)
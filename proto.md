
## websocket                                                                   
**请求URL**

ws://DOMAIN/sub

**HTTP请求方式**

Websocket（JSON Frame），请求和返回协议一致

**请求和返回json**

```json
{
    "ver": 102,
    "op": 10,
    "seq": 10,
    "body": {"data": "xxx"}
}
```

**请求和返回参数说明**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
| ver        | true  | int | 协议版本号 |
| op         | true  | int    | 指令 |
| seq        | true  | int    | 序列号（服务端返回和客户端发送一一对应） |
| body          | true | string | 授权令牌，用于检验获取用户真实用户Id |

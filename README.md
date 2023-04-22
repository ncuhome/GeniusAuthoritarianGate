# GeniusAuthoritarianGate

GeniusAuthoritarian 单体门控程序

## 使用

镜像地址：

`harbor.ncuos.com/genius-auth/gate`

环境变量：

| 键名            | 必须 | 默认值 | 说明                                  |
|---------------|:--:|-----|-------------------------------------|
| Addr          | √  |     | 代理目标地址，不带协议                         |
| Timeout       | x  | 30  | 代理请求超时时间，秒                          |
| Groups        | x  |     | 允许访问的组，详见 GeniusAuth Readme，用英文逗号分隔 |
| JwtKey        | x  | 随机值 |                                     |
| LoginValidate | x  | 7   | 登录身份保持时间，天                          |
| WhiteListPath | x  |     | 免鉴权路径，完全匹配，英文逗号分隔                   |
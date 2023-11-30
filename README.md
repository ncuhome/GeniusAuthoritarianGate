# GeniusAuthoritarianGate

GeniusAuthoritarian 单体门控程序

这个简单的门控并不能规范地处理 token，所以安全度并不高。且在出错时暂时没有对用户进行友好的提示，请谨慎使用

## 使用

镜像地址：

`harbor.ncuos.com/genius-auth/gate`

环境变量：

| 键名            | 必须 | 默认值 | 说明                |
|---------------|:--:|-----|-------------------|
| Addr          | √  |     | 代理目标地址，不带协议       |
| AppCode       | √  |     | 在 GeniusAuth 后台申请 |
| AppSecret     | √  |     | 同上                |
| Timeout       | x  | 30  | 代理请求超时时间，秒        |
| LoginValidate | x  | 7   | 登录身份保持时间，天，7-30    |
| WhiteListPath | x  |     | 免鉴权路径，完全匹配，英文逗号分隔 |

# GeniusAuthoritarianGate

GeniusAuthoritarian 单体门控程序

将会占用 `/login/` 路由段，并设置 `refreshToken`、`accessToken` cookie

创建应用时回调请指向 `/login/`

## 使用

镜像地址：

`harbor.ncuos.com/genius-auth/gate`

环境变量：

| 键名                     | 必须 | 默认值                   | 说明                     |
|------------------------|:--:|-----------------------|------------------------|
| `Addr`                 | √  |                       | 代理目标地址，不带协议            |
| `AesKey`               | √  |                       | 长度必须为 32 位，用于加密 cookie |
| `AppCode`              | √  |                       | 在 GeniusAuth 后台申请      |
| `AppSecret`            | √  |                       | 同上                     |
| `Timeout`              | x  | `30`                  | 代理请求超时时间，秒             |
| `LoginValidate`        | x  | `7`                   | 登录身份保持时间，天，7-30        |
| `WhiteListPath`        | x  |                       | 免鉴权路径，完全匹配，英文逗号分隔      |
| `GeniusAuthHost`       | x  | `v.ncuos.com`         | GeniusAuth Host        |
| `GeniusAuthAppRpcAddr` | x  | `v-app.ncuos.com:443` |                        |

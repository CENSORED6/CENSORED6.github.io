# Cookie / Session / Token / JWT 学习笔记

> 来源：[还分不清 Cookie、Session、Token、JWT？看这一篇就够了](https://developer.aliyun.com/article/1358589)（阿里云开发者社区，2023-10-25，原出处掘金）
> 整理日期：2026-11-21　｜　文中标注「勘误」处为原文表述不严谨、已修正的地方

---

## 目录

1. [三个基础概念](#一三个基础概念)
2. [Cookie](#二cookie存在客户端的小纸条)
3. [Session](#三session把状态挪到服务端)
4. [Cookie vs Session](#四cookie-vs-session)
5. [Token](#五token服务端不发状态只发签名)
6. [JWT](#六jwt-token-的标准格式)
7. [Token vs JWT](#七token-vs-jwt)
8. [分布式 Session 共享](#八分布式下-session-共享的四种方案)
9. [安全清单](#九安全清单)
10. [与实战的对应](#十和你之前折腾的东西对上号)
11. [一页速记](#十一页速记)
12. [原文勘误](#十二原文勘误)

---

## 一、三个基础概念

| 概念 | 回答的问题 | 例子 |
| --- | --- | --- |
| **认证** Authentication | 你是谁 | 指纹打卡、账号密码、短信验证码 |
| **授权** Authorization | 你能干什么 | App 申请相册权限、小程序要昵称头像 |
| **凭证** Credentials | 凭什么证明 | 身份证、sessionId、token |

> 一句话：**认证是验身份，授权是发权限，凭证是那张"通行证"**。
> 后面的 Cookie / Session / Token 全都是凭证的不同实现形态。

---

## 二、Cookie：存在客户端的小纸条

HTTP 是**无状态**协议——服务端记不住上一次请求是谁发的。Cookie 就是浏览器帮忙记的那点状态。

- **位置**：客户端（浏览器），随请求自动携带
- **大小**：单个 ≤ 4KB；单站点约 20 个，浏览器总量约 300 个
- **作用域**：**不可跨域**，靠 `domain` / `path` 绑定；一级域名与二级域名可共享
- **常用属性**

| 属性 | 作用 |
| --- | --- |
| `HttpOnly` | JS 读不到，缓解 XSS 窃取 |
| `Secure` | 只通过 HTTPS 发送 |
| `SameSite` | 限制跨站携带，缓解 CSRF |
| `Max-Age` / `Expires` | 过期控制 |
| `Domain` / `Path` | 作用域控制 |

---

## 三、Session：把状态挪到服务端

Session 是另一种会话状态记录机制，**数据在服务端，只有 ID 在客户端**。

### 认证流程

1. 用户首次请求 → 服务端根据提交信息创建 Session
2. 响应把 `SessionID` 返回给浏览器
3. 浏览器把 `SessionID` 存进 Cookie，并记录所属域名
4. 下次请求自动带上 Cookie → 服务端取 `SessionID` 查 Session
   - 查到 → 已登录，继续业务
   - 查不到 → 未登录 / 已过期

> **`SessionID` 是连接 Cookie 和 Session 的桥梁**——Session 数据在服务端，只有那把"钥匙"在客户端。

---

## 四、Cookie vs Session

| 维度 | Cookie | Session |
| --- | --- | --- |
| 存储位置 | 客户端 | 服务端 |
| 安全性 | 低（可被客户端篡改） | 高（客户端只有一个不可预测的 ID） |
| 数据类型 | 只能字符串 | 任意类型 |
| 大小 | ≤ 4KB | 受服务器内存限制 |
| 生命周期 | 可长期保留 | 较短，超时/会话结束即失效 |
| 集群友好度 | 天然无状态 | 需要共享方案（见第八节） |
| 典型问题 | 篡改、泄露、体积限制 | 内存占用、跨服共享、跨域 |

---

## 五、Token：服务端不发状态，只发签名

### 组成

常见的简单结构：`uid（用户唯一标识） + time（时间戳） + sign（签名，哈希后的定长字符串）`

### 认证流程

1. 客户端用账号密码请求登录
2. 服务端校验通过 → 签发 token
3. 客户端保存 token（localStorage 或 Cookie）
4. 每次请求通过 HTTP Header 携带 token（`Authorization: Bearer <token>`）
5. 服务端验签 → 通过则返回数据

### 特点

- **服务端无状态**：不存 token，用"算签名的时间"换"存 Session 的空间"
- **可扩展**：多台服务器都能独立验签，无需共享存储
- **支持移动端 / 跨程序调用**
- **避开同源策略**：不依赖 Cookie，也就不怕 CSRF（前提是别把 token 塞回 Cookie）

### Access Token 与 Refresh Token

| 类型 | 有效期 | 用途 | 存储 |
| --- | --- | --- | --- |
| Access Token | 短（分钟~小时） | 每次业务请求携带 | 客户端内存/localStorage |
| Refresh Token | 长（天~月） | **只用来换新的 Access Token** | 服务端数据库，仅刷新时校验 |

好处：Access Token 泄露的影响窗口小；用户不必反复输入密码。Refresh Token 也失效时才需要重新登录。

### Token vs Session

- Session 让服务端**有状态**，能记录会话信息；Token 让服务端**无状态**
- 二者不矛盾：Token 做认证，Session 存业务状态，可以叠加使用
- Session 只是"有 ID 就拥有该用户全部权限"的简单认证；Token（尤其 OAuth 场景）区分了认证（用户）与授权（App）
- **选型**：只服务自己的网站/App → Session 够用；要给第三方调 API、多端、水平扩展 → Token/JWT

---

## 六、JWT：Token 的标准格式

JSON Web Token（RFC 7519），目前最流行的跨域认证方案，本质是**带签名的自包含令牌**。

### 结构

```
Header.Payload.Signature
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMifQ.<签名>
```

| 段 | 内容 | 说明 |
| --- | --- | --- |
| Header | `{"alg":"HS256","typ":"JWT"}` | 声明算法与类型 |
| Payload | claims：`sub` / `exp` / `iat` / `nbf` / 自定义字段 | **Base64URL 编码，明文可读** |
| Signature | `HMAC-SHA256(base64(header) + "." + base64(payload), secret)` | HS256 对称密钥；RS256 用私钥签、公钥验 |

> 验签保证**内容没被篡改**，但不保证**内容不可读**。

### 认证流程

1. 用户登录，服务端认证成功 → 返回 JWT
2. 客户端保存（通常 localStorage，也可 Cookie）
3. 访问受保护资源时放在请求头：

```http
GET /calendar/v1/events HTTP/1.1
Host: api.example.com
Authorization: Bearer <token>
```

4. 服务端校验签名与有效期 → 通过则放行

### 三种携带方式

| 方式 | 写法 | 评价 |
| --- | --- | --- |
| **Header**（推荐） | `Authorization: Bearer <token>` | 标准做法，跨域友好 |
| Body | POST body 里带 `token` 字段 | 跨域场景可用 |
| URL 参数 | `?token=xxx` | 会进日志、Referer，**不安全**，尽量避免 |

### 优点

- 自包含，减少查库次数
- 无状态，天然适合水平扩展
- 不依赖 Cookie，不受 CORS 限制

### 缺点

- **签发后无法吊销**：到期前始终有效，除非服务端额外维护黑名单
- payload 明文，**绝不能放机密数据**
- 体积较大，每次请求都要带
- 必须走 HTTPS，否则等同明文传令牌

### 工程实践

- 有效期尽量短（分钟级）+ Refresh Token 续期
- 需要"立即失效"时，用 Redis 维护吊销黑名单（jti 维度）
- 校验 `exp` / `nbf` / `iat`，并留 30~60 秒 clock skew 容忍时钟偏差

---

## 七、Token vs JWT

| 维度 | 普通 Token | JWT |
| --- | --- | --- |
| 用户信息 | 存在服务端（DB/Redis） | 自带在 payload 里 |
| 校验方式 | 查库比对 | 本地验签，**不查库** |
| 能否吊销 | 能（删记录即可） | **难**（签出去就有效到过期） |
| 体积 | 小 | 大 |
| 适用 | 内部系统 | 跨域、多服务、第三方 |

> 相同点：都是访问资源的令牌、都能记录用户信息、都让服务端无状态、都要验证通过才能访问受保护资源。

---

## 八、分布式下 Session 共享的四种方案

| 方案 | 做法 | 优点 | 缺点 |
| --- | --- | --- | --- |
| **Session 复制** | 任一节点 Session 变动就序列化广播给所有节点 | 容错、各节点实时同步 | 网络压力大，量大易堵塞 |
| **粘性 Session / IP 绑定** | Nginx `ip_hash`，同一 IP 固定打到同一台 | 简单，无需改造 | 机器故障即丢失会话 |
| **Session 共享**（主流） | 统一存 Redis / Memcached 集群 | 可水平扩展、重启不丢、可跨平台 | 多一次网络访问，架构变复杂 |
| **Session 持久化** | 落数据库 | 服务重启不丢 | 高并发下数据库压力大 |

---

## 九、安全清单

### Cookie

- 存在客户端，**可能被篡改**，使用前必须校验合法性
- 不要存密码、余额等敏感数据
- 开启 `HttpOnly`，减小 XSS 影响面
- 控制体积（≤ 4KB）与 `domain` / `path`，减少无谓传输
- 无法跨域；单站 20 个、浏览器总量约 300 个的上限
- 移动端支持差

### Session

- 在线用户多时内存占用高，必须有过期清理机制
- 集群部署要解决共享问题
- 多应用共享还会遇到跨域 Cookie 问题
- 浏览器禁用 Cookie 时可用 URL 重写传递 sessionId
- 移动端常用 Token 而非 Session

### Token

- 查询频繁可放 Redis，比查库快
- 不依赖 Cookie，天然免 CSRF
- 完全由应用管理，可避开同源策略
- 移动端首选

### JWT 专项

- 默认**不加密**（可额外加密），**不要写入秘密数据**
- 泄露即拥有全部权限，有效期务必短，重要操作二次认证
- 适合一次性命令认证：每次操作签发一个极短有效期 JWT
- 必须 HTTPS 传输，禁止明文 HTTP

### 密码处理（最容易出错）

- **绝不明文存储密码**
- **绝不用 Base64 / 加密方式存密码**——编码和加密都是**可逆**的，等同于明文
- 只用**单向哈希**，且用强算法（bcrypt / scrypt / argon2）；MD5、SHA1 已不安全
- 从来没有"解哈希"这回事
- 不要以明文形式显示或发送密码；"忘记密码"应生成一次性新密码

---

## 十、和你之前折腾的东西对上号

| 之前的场景 | 对应本文概念 |
| --- | --- |
| 解析 `eyJhbGci...` 得到 `{"alg":"HS256"}` + `{"code":"博","iat":...,"nbf":...}` | 标准 JWT：header + payload + HMAC 签名 |
| 该 token **没有 `exp`** | 即"签发即长期有效"，正是 JWT 最难撤销的坑 |
| 往请求头塞 `Authorization: Bearer <token>` | 本文"方式一"，最标准的携带方式 |
| 校验 `iat` / `nbf` / `exp` | 对应"有效期要短"，并需留 clock skew |
| 网关按路由下发 `map-fragments` 响应头 | 响应头同样是 HTTP 元数据通道，与认证头同理 |

---

## 十一、一页速记

```text
Cookie  = 客户端存的纸条（4KB，不可跨域，可被篡改）
Session = 服务端存的档案，Cookie 里只放档案号（sessionId）
Token   = 服务端签的通行证，不存状态，验签即可
JWT     = Token 的标准格式：header.payload.signature
          能看不能改；能验不能撤；要短命 + HTTPS
密码    = 单向哈希，绝不明文 / 编码 / 可逆加密
```

**选型口诀**

- 自己网站够用就好 → Session-Cookie
- 多端 / 跨域 / 微服务 → Token（JWT + Refresh）
- 接第三方 → OAuth 2.0

---

## 十二、原文勘误

| 原文表述 | 问题 | 正确说法 |
| --- | --- | --- |
| "JWT 将 Token 和 Payload **加密**后存储于客户端" | 与后文"JWT 默认不加密"自相矛盾 | JWT 默认只做 Base64URL **编码 + 签名**，payload 可被任何人解码查看 |
| "Token 可以避免 CSRF 攻击（因为不需要 cookie 了）" | 只在 token 存于 localStorage/内存时成立 | 若把 token 存进 Cookie 并自动携带，仍有 CSRF 风险 |
| "无法废弃某个 Token" 被当作绝对缺点 | 可通过短有效期 + 黑名单缓解 | 需说明工程补丁：Refresh Token、jti 黑名单 |
| 未区分「编码 / 加密 / 哈希」 | 三者常被混用 | 编码可逆（Base64）、加密可逆（需密钥）、哈希单向不可逆 |

---

*本笔记仅为学习整理，原文版权归原作者所有。*

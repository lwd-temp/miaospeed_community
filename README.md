# MiaoSpeed Community Build

---

这里是由Paimon Hub社区维护的MiaoSpeed社区构建版本，我们会在这里发布社区构建版本的更新。

## 社区版本功能
* Clash Meta内核支持
* Cli自定义除根证书外所有预设变量

## 使用方式

基本使用方式与官方版本无异，您可以参考以下指导进行使用。

### 子命令
* **server** 启动miaospeed作为后端服务器。
* **script** 运行临时脚本测试来测试脚本的正确性。
* **misc** 提供miaospeed额外功能。

### server

#### 官方参数
- **bind** 绑定一个套接字，可以是0.0.0.0:8080或/tmp/unix_socket
- **connthread** 并行线程处理正常连接任务(默认64)
- **mtls** 启用miaoko certs TLS验证
- **nospeed** 禁止测速
- **pausesecond** 在每个速度作业之后暂停该时间(秒)
- **speedlimit** speed ratelimit(以字节每秒为单位)，默认没有限制
- **token** 后端请求验证用的令牌
- **verbose** 打印系统日志
- **whitelist** bot白名单，启用后只允许名单内的bot链接后端，格式:1111,2222,3333
- **mmdb** 指定mmdb文件路径，多个文件可以使用逗号分隔，默认情况下不使用。

#### 自定义参数
以下参数可以设置大部分在官方版本中于构建时锁定的预设变量。如果留空，将设置为默认的官方预设。
- **scriptpredefined** 自定义预处理脚本，需填入路径
- **scriptgeo** 自定义geo脚本，需填入路径
- **scriptip** 自定义ip脚本，需填入路径
> **警告**<br>下面的预设变量修改可能具有风险，改动任意一条会导致Miaoko官方主端无法连接至Miaospeed，如果您需要使用第三方软件快速对接Miaospspeed，可以自由修改
- **buildtoken** 自定义构建令牌，详见官方文档
- **serverpublickey** 自定义Miaospeed服务器刚公钥，需填入路径
- **serverprivatekey** 自定义Miaospeed服务器私钥，需填入路径

### script
- **file** 指定测试的脚本文件路径

### misc
- **maxmind-update-license** 更新Maxmind的geoip数据库，需填入Maxmind license key
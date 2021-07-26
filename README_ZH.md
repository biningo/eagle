## eagle

- 自动发现并注册Docker服务到etcd
- etcd配置中心

​

## eagle是如何注册服务的

eagle进程启动之后会启动两个服务

- 监控docker容器，并解析label注册到etcd
- RESTful接口服务，提供服务发现和配置中心

服务注册的范围分为三个等级

- `namespace` 一个命名空间下可以保存多个服务
- `serviceName` 一个服务下可以有多个实例
- `serviceID` 就是docker容器的ID

要被eagle监控的容器需要打上指定的label，用于服务注册的lable不需要赋值，此label可以在配置文件中配置多个，只有包含了这些label才可以被发现，默认是`eagle`

```bash
docker run --name pingip -d -p 9090 --label eagle   biningo/pingip
```

如果要指定namespace则需要在每个容器里面打一个namespace的label，默认的namespace是`default`

```bash
docker run --name pingip -d -p 9090 --label eagle --namespace=ping   biningo/pingip
```

serviceName的值默认是docker镜像的名字，我觉得这是非常合理的，如果需要指定也可以进行手动指定，打一个label即可

```bash
docker run --name pingip -d -p 9090 --label eagle --label namespace=icepan --label serviceName=pingIP   biningo/pingip
```

eagle的唯一缺点就是容器的运维成本略有些加大，不过也没有很大，就是需要手动指定一些label，如果有多个服务则写个shell脚本运维难度也不会很大。当然和他的优点 **代码非侵入性注册** 比起来这点缺点已经不算什么了

​

## etcd中保存了哪些内容

etcd中保存的内容就是`ServiceInstance`的json序列化数据，ServiceInstance中主要就是保存了如下几个内容

- namespace、serviceName、serviceID
- 容器IP和端口   `PrivateIP`和`PrivatePort`
- 宿主机IP和容器映射的端口  `PublicIP`和`PublicPort`
- Labels

> 注意宿主机的公网IP需要在配置文件中指定

​

## eagle是如何进行健康检查的

- 成功注册容器之后会开一个`goroutine`根据配置的`health`
  策略验证服务是否存活，目前实现了两种方式分别是HTTP接口验证和TCP端口检测，默认是TCP端口验证因为这样可以减轻服务的压力，但是HTTP接口验证则更能保证服务是否正常

- 当监控到docker容器挂了则会立即进行服务下线也就是删除etcd中保存的数据

​

## 有哪些可以配置

下面是默认的配置，如果不指定配置文件则使用如下配置

```yaml
labels:
  - "eagle"
server:
  host: "127.0.0.1"
  port: "9999"

etcd:
  endpoints:
    - "127.0.0.1:2380"
  prefix: "eagle"

docker:
  network: "bridge"

health:
  timeout: 5
  interval: 3
  checker:
    type: tcp
```

​

## Usage

启动并且使用默认配置

```bash
eagle
```

启动并指定配置文件

```bash
eagle --config /etc/config.yaml
```

查看帮助

```bash
eagle --help
```

查看所有可以被注册的服务

```bash
eagle service list
eagle service get serviceName
```

​

## RESTful接口

查询服务

```bash
GET /registry/:namespace/services
GET /registry/:namespace/services/serviceName
GET /registry/:namespace/services/serviceName/serviceID
```
配置中心
```bash
GET /config/:namespace/configurations
GET /config/:namespace/configurations/:filename
PUT /config/:namespace/configurations/:filename
DELETE /config/:namespace/configurations/:filename
```
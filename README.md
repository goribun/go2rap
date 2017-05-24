# go2rap

在接口方与调用方并行开发时，对于未发布的接口，mock平台有助于提高开发效率；比如 [RAP](https://github.com/thx/RAP)提供的简单强大的API管理及mock功能。

但是在实际使用中并不会将所有的调用都要切到RAP上，而是当前并行开发的部分需要请求RAP，其他的还要请求实际接口；此时，必然会造成不必要的代码入侵，比如在调用mock接口时单独指定地址，或者为mock请求增加注解标记。

go2rap通过简单的反向代理功能解决了mock时的代码入侵：

![flow-char](https://raw.githubusercontent.com/goribun/go2rap/master/doc/flow-char.png)

如上图所示，go2rap监听80端口，发起接口请求首先经过go2rap（需要将接口的host指到本地），然后通过配置信息决定直接请求接口或是请求RAP，然后响应请求结果。

其实，也可以把它当作反向代理使用，把域名指向本地，方便本地调试时取得登陆信息（cookie）。

## 特点 

1. 功能单一，占用本机资源低，大概占用3M左右内存

2. 配置简单，仅需配置go2rap.json

3. 实时生效，修改配置无需重启（处理每次请求都会读取配置文件，简单粗暴，因为本地调试使用无需考虑性能问题，哈哈哈。。。

## 配置

1. 将接口host指向本地

2. 修改配置文件


  servers部分配置需要代理的服务，name为名称，host为域名，proxy为服务的ip；如果只作为反向代理，只配置该部分即可


  ```json
  ""servers": [
      {
        "name": "test-user-api",
        "host": "test.user.api.lq.wangxs.cn",
        "proxy": "192.168.100.1"
      },
      {
        "name": "rap-user-api",
        "host": "rap.user.api.lq.wangxs.cn",
        "proxy": "192.168.128.6"
      },
      {
        "name": "local-user-web",
        "host": "local.user.wangxs.cn",
        "proxy": "127.0.0.1:8080"
      }
    ]": [
      {
        "name": "test-user-api",
        "host": "test.user.api.lq.wangxs.cn",
        "proxy": "192.168.100.1"
      },
      {
        "name": "rap-user-api",
        "host": "rap.user.api.lq.wangxs.cn",
        "proxy": "192.168.128.6"
      },
      {
        "name": "local-user-web",
        "host": "local.user.wangxs.cn",
        "proxy": "127.0.0.1:8080"
      }
    ]
  ```
  
 condition部分，用来配置路径条件，其中serverA为实际api的server name（server部分配置的name），serverB为mock平台的server name，prefixPath用来指定路径前缀，因为可能多个接口项目都使用同一个mock平台，使用一个前缀路径区分；path就是配置路径条件了，可以指定多个，当调用的路径包含在path时就会请求serverB也就是RAP平台。
 
 ```json
  "conditions": [
    {
      "serverA": "test-user-api",
      "serverB": "rap-user-api",
      "prefixPath": "/mockjsdata/1",
      "path": [
        "/user/vipList"
      ]
    },
    {
      "serverA": "yy-api",
      "serverB": "rap-yy-api",
      "path": [
        "/activity/getAllList",
        "/activity/getDetail"
      ]
    }
  ]
 ```




# microservice-examples 微服务调用示例
- 三个服务http远程调用 http://service-a:1001/api/a -> http://service-b:1002/api/b -> http://service-c:1003/api/c
``` 
curl -i localhost:1001/api/a 
```

- service-b 两个版本 v1 v2

## 接口说明
- /api/a 服务访问链路 svcA -> svcB -> svcC
- /api/header?label-user=groupA 设置svcA->svcB请求header值("label-user":"groupA")
- /api/timeout?second=5 svcA->svcB过程中，B中sleep 5秒后返回
- /api/test1 访问svcB无端口号，返回连接异常

## 外部访问
- istio-gateway暴露端口查看
``` 
kubectl -n istio-system get svc istio-ingressgateway
```
- 通过istio gateway访问 kiali
``` 
 k apply -f expose/expose-kiali.yaml 
```

## 测试场景
VirtualService中destination在Kubernetes平台中host:service-b.default.svc.cluster.local服务的短名称可简写为service-b
 A/B 测试、金丝雀发布、基于流量百分比切分的概率发布,测试前先设置目标地址
``` 
kubectl apply -f traffic-management/destination-rule-service-b.yaml
```
### 网络弹性

- 按版本流量百分比
``` 
kubectl apply -f traffic-management/vs-service-b-90-10.yaml
```
- 按照header值 

request.header.label-user值为groupA的流量全部路由到svcB.v2 , 其他流量路由到svcB.v1
``` 
kubectl apply -f traffic-management/vs-service-b-headers.yaml   
```
``` 
curl -v  http://120.92.117.115:31744/api/header?label-user=groupA    
```
- 超时/重试
``` 
kubectl apply -f traffic-management/vs-service-b-timeout-retries.yaml
```
测试请求超时时间5秒
``` 
http://pingcap-8:31744/api/timeout?second=5
```
- 熔断/断路器
``` 
kubectl apply -f traffic-management/destination-rule-service-b-circuitBreaker.yaml 
```
监控指标暂未看到

- 故障注入
全部流量返回503状态
``` 
kubectl apply -f traffic-management/vs-service-b-fault-injection.yaml 
```
- 本地限速(Local rate limit)
服务实例限速 使用 envoyfilter 实现访问限速，envoyfilter名称filter-local-ratelimit-svc
``` 
kubectl apply -f traffic-management/rate-limit/rate-limit-service-local.yaml 
```
- 全局限速(Global rate limit)
全局服务限速需先部署Redis与限速服务
``` 
kubectl apply -f traffic-management/rate-limit/global/rate-limit-service.yaml 
```
configmap 中配置 gateway 限速设置
``` 
kubectl apply -f traffic-management/rate-limit/global/microservice-examples-gateway-global-rate-limit.yaml
```
删除全局限速
``` 
k delete -f traffic-management/rate-limit/global/microservice-examples-gateway-global-rate-limit.yaml
```

## Tools
- 部署curl
``` 
k apply -f traffic-management/tools/sleep.yaml 
export SOURCE_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS http://httpbin.org/headers
```
- Pod 内访问 svcB
``` 
kubectl exec "$SOURCE_POD" -c sleep -- curl -sS http://service-b:1002/api/b
```

## build
``` 
docker build -t='yusyscloud/gs-service-a' . 
docker build -t='yusyscloud/gs-service-b:v2' . 
```

## 

``` 
kubectl create namespace ns-gs-istio 
kubectl label namespace ns-gs-istio istio-injection=enabled
kubectl apply -n ns-gs-istio -f microservice-examples-deployment.yaml 


kubectl delete -n ns-gs-istio -f microservice-examples-deployment.yaml 
```
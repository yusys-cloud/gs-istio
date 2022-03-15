# gs-istio
Istio service mesh samples.

- [proxywasm-hello](proxywasm-hello) examples of wasme
- [microservice-examples](microservice-examples)  service-a -> service-b -> service-c
- [logginng](logging) Fluentd / Elasticsearch / Kibana fluentd gather all the logs from the Kubernetes node's environment and append the proper metadata to the logs

## Tinygo

```
docker run -v $GOPATH:/go -e "GOPATH=/go" tinygo/tinygo:0.21.0 tinygo build -o /go/src/github.com/yusys-cloud/gs-istio/wasm.wasm -target wasm --no-debug /go/src/github.com/yusys-cloud/gs-istio/main.go
```

### Requirements
Go 1.17  TinyGo  Envoy

### Concepts
- Istio object/configuration Type
This is the type specified in the Istio Config. This could be any of the following types: Gateway, Virtual Service, DestinationRule, ServiceEntry, Rule, Quota or QuotaSpecBinding.

- arm
```
root@orangepiplus2e:~/test/gs-istio/proxywasm-hello# tinygo build -o hello.wasm -target wasm ./main.go
error: Linking globals named 'malloc': symbol multiply defined!
```

### Links
- [Deploying WASM plugin on Istio](https://sirishagopigiri-692.medium.com/deploying-wasm-plugin-on-istio-2323f276d055)
- [Building WASM go filter plugin for envoy](https://sirishagopigiri-692.medium.com/building-wasm-go-filplugin-for-envoy-21d36c568057) 
- [LocalityLoadBalancer](https://istio.io/latest/zh/docs/reference/config/networking/destination-rule/#LocalityLoadBalancerSetting-Failover)

https://github.com/tetratelabs/proxy-wasm-go-sdk/blob/main/doc/OVERVIEW.md
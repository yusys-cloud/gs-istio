

``` 
curl -v 192.168.100.147:32594/api/message/xml
kubectl logs -f $(kubectl get pod -l app=service-a -o jsonpath={.items..metadata.name}) -c istio-proxy
```
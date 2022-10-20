## gs-args
``` 
k logs -f $(k get pod -l app=gs-args -o jsonpath={.items..metadata.name})
k describe pods $(k get pod -l app=gs-args -o jsonpath={.items..metadata.name})
```

## links
- https://kubernetes.io/zh-cn/docs/concepts/storage/volumes/#emptydir
- https://support.huaweicloud.com/usermanual-cce/cce_01_0008.html
- https://kubernetes.io/zh-cn/docs/concepts/workloads/pods/init-containers/
- https://kubernetes.io/zh-cn/docs/tasks/inject-data-application/define-command-argument-container/
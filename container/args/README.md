## gs-args
``` 
k logs -f $(k get pod -l app=gs-args -o jsonpath={.items..metadata.name})
k describe pods $(k get pod -l app=gs-args -o jsonpath={.items..metadata.name})
```
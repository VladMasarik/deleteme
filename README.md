# deleteme

# Dependencies
Module is a collection of packages
All your programs are basically modules. Initiate them with
```
go mod init "name"
```

Resolve / download / record dependencies with
```
go mod tidy
```

# Connection to k8s
You can have "in cluster config" meaning that the app is running inside the k8s that it wants to connect to, in which case it starts using the service account saved somewhere in the pod. OR you can have "out of cluster" connection, in which case you would most likely somehow pass the kubeconfig to the app, and it would use it.

Kubernetes has go-client examples, which hold a lot of information.
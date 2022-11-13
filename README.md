# kevent

Very simple library for creating Kubernetes events.

```go
clientConfig, _ := rest.InClusterConfig()
kubernetesClient, _ := kubernetes.NewForConfig(clientConfig)
kevent.CreateEvent(kubernetesClient, "kube-system", "pod-name", "SomeReason", "Things went boom", true)
```
Kubernetes Play
===============

We'll deploy and play around with Kubernetes
(on Google Container Engine)

## Pods

### create a Pod

kubectl create -f kube/pod.yaml

### look closer

kubectl describe go-env

### port forward

kubectl port-forward -p go-env 8080

### delete the pod

kubectl delete pod go-env

### run a container

kubectl run-container nginx --image=nginx

### delete it

kubectl delete pod nginx-blah

### actually its a replication controller

kubectl get rc

## Replication Controllers

### Take a look

```
cat kube/controller.json

{
   "kind":"ReplicationController",
   "apiVersion":"v1beta3",
   "metadata":{
      "name":"go-env",
      "labels":{
         "name":"go-env"
      }
   },
   "spec":{
      "replicas":2,
      "selector":{
         "name":"go-env"
      },
      "template":{
         "metadata":{
            "labels":{
               "name":"go-env"
            }
         },
         "spec":{
            "containers":[
               {
                  "image":"matthewrudy/go-env:v1",
                  "name":"go-env",
                  "ports":[
                     {
                        "name":"http-server",
                        "containerPort":8080,
                        "protocol":"TCP"
                     }
                  ]
               }
            ]
         }
      }
   }
}
```

### Create it

```
kubectl create -f kube/controller.json
```

### Scale it up

```
kubectl resize --replicas=3 rc go-env
```

### Scale it down

```
kubectl resize --replicas=2 rc go-env
```

### Rolling update

```
kubectl rolling-update go-env --image=matthewrudy/go-env:latest

Error updating pod (Pod "go-env" is invalid: spec: invalid value ... may not update fields other than container.image), retrying after 3 seconds^C%
```
### Rollback

```
kubectl rolling-update go-env --image=matthewrudy/go-env:v1

# same error
```

## Services

### Take a look

```
cat kube/service.json

{
    "kind": "Service",
    "apiVersion": "v1beta3",
    "metadata": {
        "name": "go-env",
        "labels": {
          "name": "go-env"
        }
    },
    "spec": {
        "selector": {
            "name": "go-env"
        },
        "ports": [
            {
                "protocol": "TCP",
                "port": 80,
                "targetPort":"http-server"
            }
        ],
        "createExternalLoadBalancer": true
    }
}
```

### Create it

```
kubectl create -f kube/service.json
```

### Load balancer

```
gcloud compute forwarding-rules list

NAME                             REGION     IP_ADDRESS     IP_PROTOCOL TARGET
a27b0890b0ebe11e5a64f41234567890 asia-east1 107.167.182.94 TCP         asia-east1/targetPools/a27b0890b0ebe11e5a64f41234567890
```

### Create another service

```
kubectl create -f kube/other/service.json
kubectl create -f kube/other/controller.json
```

## Secrets

### Take a look

``` bash
cat kube/secrets/secret.json

{
  "apiVersion": "v1beta3",
  "kind": "Secret",
  "metadata" : {
    "name": "secret"
  },
  "data": {
    "foo": "Zm9vYmFyYmF6"
  }
}
```

### Create the secret

``` bash
kubectl create -f kube/secrets/secret.json
```

### Mount it

``` bash
cat kube/secrets/pod.json

{
 "apiVersion": "v1beta3",
 "kind": "Pod",
  "metadata": {
    "name": "ruby-secrets"
  },
  "spec": {
    "containers": [{
      "name": "ruby-secrets",
      "image": "matthewrudy/ruby-secrets",
      "volumeMounts": [{
        "name": "secrets",
        "mountPath": "/etc/secrets",
        "readOnly": true
      }],
      "ports":[
         {
            "name":"http-server",
            "containerPort":4567,
            "protocol":"TCP"
         }
      ]
    }],
    "volumes": [{
      "name": "secrets",
      "secret": {
        "secretName": "secret"
      }
    }]
  }
}
```

```
kubectl create -f kube/secrets/pod.json
```

Kubernetes Play
===============

We'll deploy and play around with Kubernetes
(on Google Container Engine)

## Images

For the purposes of Demoing,
we have two images we're using.

### matthewrudy/go-env

A mini app in go, that listens on port 8080
and renders all the `env` as a web page.

### matthewrudy/ruby-secrets

A sinatra app that listens on port 4567

It looks for `/etc/secrets`
and if it's there, renders a list of the files.

## Pods

### Take a look

``` bash
cat kube/pod.yml

apiVersion: v1beta3
kind: Pod
metadata:
  labels:
    name: go-env
  name: go-env
  namespace: default
spec:
  containers:
  - image: matthewrudy/go-env:v1
    imagePullPolicy: IfNotPresent
    name: go-env
    ports:
    - containerPort: 8080
      protocol: TCP
  restartPolicy: Always
```

### Create it

```
kubectl create -f kube/pod.yaml
```

### Look closer

kubectl describe pod go-env

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

### Look at the mount

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

### Create it

```
kubectl create -f kube/secrets/pod.json
```

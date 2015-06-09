Kubernetes Play
===============

We'll deploy and play around with Kubernetes
(on Google Container Engine)

# create a Pod

kubectl create -f pod.yaml

# look closer

kubectl describe go-env

# port forward

kubectl port-forward -p go-env 8080

# delete the pod

kubectl delete pod go-env

# run a container

kubectl run-container nginx --image=nginx

# delete it

kubectl delete pod nginx-blah

# actually its a replication controller

kubectl get rc

# create a replication controller

kubectl create -f controller.json

# k8s日志

## configMap

### 创建

把本机项目下的env创建成configMap     
kubectl create configmap data-model-configmap --from-env-file=.env

### 查看

kubectl get configmaps data-model-configmap -o yaml     
kubectl describe configmap data-model-configmap

### 删除

kubectl delete configmap data-model-configmap

## Job

### 创建

kubectl create -f ./dev.job.yaml

### 查看

kubectl get pods kubectl logs project-data-model-dev-4nvd8

## Deployment

### 创建

kubectl apply -f ./deploy/web.deployment.yaml

### 删除

kubectl delete -f ./deploy/web.deployment.yaml

### 查看

kubectl get svc,pod -o wide kubectl get pods

kubectl describe pod xxx kubectl logs xxx

### 进入pod , 但我的docker中没有curl，所以进入容器也不能确定接口中是否可以访问

kubectl exec -ti <your-pod-name>  -n <your-namespace>  -- /bin/sh

## 本地转发端口访问pod, port-forward

kubectl port-forward pod-name 8001:8000

## Service

### 创建

kubectl apply -f ./deploy/service.deployment.yaml

或者在 kubectl apply -f ./deploy/service.web.deployment.yaml service和pod在一个文件中

### 删除

kubectl delete -f ./deploy/service.deployment.yaml

### 查看

kubectl get svc kubectl get service

### 端口暴露

修改service type

## minikube 访问service

> https://stackoverflow.com/questions/63600378/cant-access-minikube-service-using-nodeport-from-host-on-mac

minikube service {svs-name}

## minikube 暴露端口

> https://kubernetes.io/docs/tutorials/hello-minikube/

kubectl expose

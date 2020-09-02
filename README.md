# Wordpress-Operator

The Objective of this Operator is to demonstrate Wordpress kind of resource using Kuberentes controller pattern- [Operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).Another objective of the repository is to show how to build the custom controller that encapsulates specific domain/application level knowledge. The Operator is built using the [operator-sdk framework](https://github.com/operator-framework/operator-sdk).
If Wordpress and MySQL were to be deployed without operator on Kuberenetes, it can be reffered [here](https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/). Lets understand how can this be achieved via Kubernetes Operator.

## Prerequistites

- golang v1.12+.
- set GO111MODULE="on"
- [Install the operator-sdk (version 15)](https://sdk.operatorframework.io/docs/golang/installation/)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [kubectl client](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Wordpress Resource
The Wordpress Operator using the operator-sdk project deploys wordpress using on sql via a [custom resource](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
- A kubernetes resource of kind: `Wordpress` and apiVersion: `example.com/v1` is created that results in the Operator deploying `Deployments`, `PersistentVolumeClaims`, `Services`, that constitute a simple instance of wordpress on sql. 
- The user can specify plaintext password that can be used as a `MYSQL_ROOT_PASSWORD`.
Once the user applies the Wordpress resource (kubectl aply -f ./deploy/crds/example.com_v1_wordpress_cr.yaml) resource, controller could spin up `mysql and wordpress pods` using the `MYSQL_ROOT_PASSWORD` as specified in `spec.sqlrootpassword`.
e.g., 
``` 
apiVersion: example.com/v1
kind: Wordpress
metadata:
  name: mysite
spec:
  sqlRootPassword: plaintextpassword 
  ```
  
## Trying the Operator

`git clone https://github.com/priyanka19-98/Wordpress-Operator.git`
`cd Wordpress-Operator`
We would be trying out the operator locally. By locally we mean that we want to run the operatot logic binary without actually building an image and pushing it to a container registry. Running the operator locally helps in day to day development. 
You can have a [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) single node local cluster to play with the operator.
Before running the operator do:
```
kubectl apply -f ./deploy/crds/example.com_wordpresses_crd.yaml
kubectl apply -f ./deploy/crds/example.com_v1_wordpress_cr.yaml
```

The CRD would be registerd and you can check that by the following command:
`kubectl get crd`
You should be able to see the output as:
```
[pjiandan@pjiandan crds]$ kubectl get crd
NAME                      CREATED AT
wordpresses.example.com   2020-09-01T12:29:55Z
```
After that run the operator locally with `operator-sdk run --local`, you should see the logs as: 
```
INFO[0000] Running the operator locally in namespace default. 
{"level":"info","ts":1598973876.2819793,"logger":"cmd","msg":"Operator Version: 0.0.1"}
{"level":"info","ts":1598973876.2820053,"logger":"cmd","msg":"Go Version: go1.13.10"}
{"level":"info","ts":1598973876.282011,"logger":"cmd","msg":"Go OS/Arch: linux/amd64"}
{"level":"info","ts":1598973876.2820172,"logger":"cmd","msg":"Version of operator-sdk: v0.15.2"}
{"level":"info","ts":1598973876.285575,"logger":"leader","msg":"Trying to become the leader."}
{"level":"info","ts":1598973876.285611,"logger":"leader","msg":"Skipping leader election; not running in a cluster."}
{"level":"info","ts":1598973876.5921307,"logger":"controller-runtime.metrics","msg":"metrics server is starting to listen","addr":"0.0.0.0:8383"}
{"level":"info","ts":1598973876.596543,"logger":"cmd","msg":"Registering Components."}
{"level":"info","ts":1598973876.5967476,"logger":"cmd","msg":"Skipping CR metrics server creation; not running in a cluster."}
{"level":"info","ts":1598973876.5967603,"logger":"cmd","msg":"Starting the Cmd."}
{"level":"info","ts":1598973876.5973437,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"wordpress-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1598973876.5975914,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"wordpress-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1598973876.5977812,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"wordpress-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1598973876.5979419,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"wordpress-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1598973876.5980544,"logger":"controller-runtime.controller","msg":"Starting Controller","controller":"wordpress-controller"}
{"level":"info","ts":1598973876.598183,"logger":"controller-runtime.manager","msg":"starting metrics server","path":"/metrics"}
{"level":"info","ts":1598973876.6982796,"logger":"controller-runtime.controller","msg":"Starting workers","controller":"wordpress-controller","worker count":1}
{"level":"info","ts":1598973876.6983802,"logger":"controller_wordpress","msg":"Reconciling Wordpress","Request.Namespace":"default","Request.Name":"example-wordpress"}
{"level":"info","ts":1598973876.6984997,"logger":"controller_wordpress","msg":"Creating a new PVC","PVC.Namespace":"default","PVC.Name":"wp-pv-claim"}
{"level":"info","ts":1598973876.7138047,"logger":"controller_wordpress","msg":"Creating a new Deployment","Deployment.Namespace":"default","Deployment.Name":"wordpress"}
{"level":"info","ts":1598973876.736821,"logger":"controller_wordpress","msg":"Creating a new Service","Service.Namespace":"default","Service.Name":"wordpress"}
{"level":"info","ts":1598973876.8298655,"logger":"controller_wordpress","msg":"Reconciling Wordpress","Request.Namespace":"default","Request.Name":"example-wordpress"}
{"level":"info","ts":1598973876.8301716,"logger":"controller_wordpress","msg":"Creating a new Service","Service.Namespace":"default","Service.Name":"wordpress"}
```
This indicates that our operator is up and running. 

See if the pods,deployments,pvcs and services are up and running: 
```
kubectl get pods
kubectl get deploy
kubectl get pvc
kubectl get svc
```
You should be able to see the following output:
```
[pjiandan@pjiandan crds]$ kubectl get po
NAME                               READY   STATUS    RESTARTS   AGE
wordpress-6d5b4988ff-dcxfj         1/1     Running   0          16h
wordpress-mysql-59d5d89ff8-qj92r   1/1     Running   0          17h
[pjiandan@pjiandan crds]$ kubectl get svc
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes        ClusterIP   10.96.0.1       <none>        443/TCP        19h
wordpress         NodePort    10.100.123.86   <none>        80:31881/TCP   16h
wordpress-mysql   ClusterIP   None            <none>        3306/TCP       17h
[pjiandan@pjiandan crds]$ kubectl get deploy
NAME              READY   UP-TO-DATE   AVAILABLE   AGE
wordpress         1/1     1            1           16h
wordpress-mysql   1/1     1            1           17h
[pjiandan@pjiandan crds]$ kubectl get pvc
NAME             STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
mysql-pv-claim   Bound    pvc-9ee52dce-b7b7-433d-8596-22392033e55e   10Gi       RWO            standard       17h
wp-pv-claim      Bound    pvc-8674f3fa-acb3-4cd7-9283-5ecec8305945   10Gi       RWO            standard       16h

```
Run the following command to get the IP Address for the Wordpress Service:

`minikube service wordpress --url`

The response should be like this:

`http://192.168.99.101:31881`

Copy the IP address and load the page in your browser to view your site: 


![alt text](https://raw.githubusercontent.com/kubernetes/examples/master/mysql-wordpress-pd/WordPress.png)


# Questions
Please feel free to open up an issue.



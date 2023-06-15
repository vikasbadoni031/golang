# Golang - Build Kubernetes Application.

- Sample program using client-go to interact with the kubernetes API. 
- Creates nginx Deployment, Service, Ingress which is accessible on the localhost.


## Prerequesites Tools
- Minikube
- Docker


## Inputs
Change vars to update the inpust like hostname, port, name and labels.

## How to run locally
1. Start minikube.
    ```
    minikube start
    ```
2. Start minikube tunnel.
    ```
    minikube tunnel
    ```
3. Run the go program
    ```
    go run main.go
    ```
4. Update /etc/hosts to point host defined to localhost
    ```
    127.0.0.1 example.k8test.com 
    ```
5. curl the url to view the 200 reponse and the nginx webpage
    ```
    curl example.k8test.com
    ```
# Docker
> Docker : project opensource untuk deploy aplikasi menggunakan container.

> Container: Image yang di running / hasil instansiasi dari image

> Container registry : untuk untuk menyimpan image docker

> Docker image: distribution file atau aplikasi (hasil bundle aplikasi) aplikasi yang sudah jadi, siap untuk dijalankan.

> Container : image yang kita running

## Install Docker
Get Docker: https://docs.docker.com/get-docker/

### Docker Pull
docker pull nginx:latest

### Docker run
docker run -d -p 8080:80 --name nginx nginx:latest

## Dockerfile
Directory tree:

```
📦Landingpage
 ┣ 📜Dockerfile
 ┗ 📜index.html
```

HTML Code
```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello Word!</title>
</head>
<body>
    <p>Hello from Docker!</p>
</body>
</html>
```

Dockerfile
```
FROM nginx

COPY . /usr/share/nginx/html/
```

Command:
- docker build -t swadharma/landingpage:latest .
- docker push swadharma/landingpage:latest
- docker pull swadharma/landingpage:latest
- docker run -d -p 8181:80 -n landingpage swadharma/landingpage:latest

# Kubernetes
- Sistem internal Google Borg (menjadi Omega)
- Kubernetes diperkenalkan Google 2014, dibuat dengan Golang

Kubernetes adalah platform open-source berbasis Linux yang dirancang untuk
mengotomatisasi penempatan, penskalaan, dan manajemen aplikasi yang berada dalam kontainer.


## Install Kubernetes

### Install choco
```
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```

### Install Virtualbox
https://www.virtualbox.org/wiki/Downloads

### Install minikube
choco install minikube -y

## Ingress
> choco install openssl

> openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout goapp.key -out goapp.crt

> C:\Windows\System32\Drivers\etc\hosts

> kubectl create secret tls goapp-secret-tls --key goapp.key --cert goapp.crt

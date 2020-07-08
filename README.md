# hello-web

![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sha224/hello-web)
![Docker Pulls](https://img.shields.io/docker/pulls/sha224/hello-web)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/sha224/hello-web)

Simple image to run a web service on multiple ports. The web service shows various details about the host and the request. This image can be used for demo and test purposes.

Usage Example:
```
docker run --net="host" -e NAME="My Computer" -e PORTS="8080,9090" hello-web:latest
```

Response Example:
```
Hello, World! You have reached port 8080 of My Computer from 127.0.0.1:43902.

Host Details:
Name: My Computer
Hostname: liberty
IP Address: 127.0.0.1:8080

Request Details:
Method: GET
URL: /
Protocol: HTTP/1.1
Header: map[Accept:[*/*] User-Agent:[curl/7.68.0]]
Body: {}
Content Length: 0
Transfer Encoding: []
Host: localhost:8080
Form Data: map[]
Remote Address: 127.0.0.1:43902
Request URI: /
```

# hello-web

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sha224/hello-web)](https://hub.docker.com/r/sha224/hello-web)
[![Docker Stars](https://img.shields.io/docker/stars/sha224/hello-web)](https://hub.docker.com/r/sha224/hello-web)
[![Docker Pulls](https://img.shields.io/docker/pulls/sha224/hello-web)](https://hub.docker.com/r/sha224/hello-web)
[![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/sha224/hello-web)](https://hub.docker.com/r/sha224/hello-web)

Simple image to run a web service on multiple ports. The web service shows various details about the host and the request. This image can be used for demo and test purposes.

## Usage Examples

### Host Mode

Docker Command:
```
docker run --net="host" -e NAME="My Computer" -e PORTS="8080,9090" sha224/hello-web:latest
```

Sample Response:
```
Hello, World! You have reached port 8080 of My Computer from 127.0.0.1:60292.

Host Details:
Name: My Computer
Hostname: liberty
Server Address: 127.0.0.1:8080
All Addresses: [127.0.0.1/8 192.168.0.213/24 172.17.0.1/16]

Request Details:
Method: GET
URL: /
Protocol: HTTP/1.1
Host: localhost:8080
Header: map[Accept:[*/*] User-Agent:[curl/7.68.0]]
Body: {}
Content Length: 0
Transfer Encoding: []
Form Data: map[]
Remote Address: 127.0.0.1:60292
```

### Bridge Mode

Docker Command:
```
docker run -e NAME="My Container" -e PORTS="8080,9090" -p 8080:8080 -p 9090:9090 sha224/hello-web:latest
```

Sample Response:
```
Hello, World! You have reached port 8080 of My Container from 172.17.0.1:36952.

Host Details:
Name: My Container
Hostname: a9347c43e277
Server Address: 172.17.0.2:8080
All Addresses: [127.0.0.1/8 172.17.0.2/16]

Request Details:
Method: GET
URL: /
Protocol: HTTP/1.1
Host: localhost:8080
Header: map[Accept:[*/*] User-Agent:[curl/7.68.0]]
Body: {}
Content Length: 0
Transfer Encoding: []
Form Data: map[]
Remote Address: 172.17.0.1:36952
```

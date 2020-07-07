# hello-web
Simple image to run a web service on multiple ports. The web service shows various details about the host and the request. This image can be used for demo and test purposes.

Usage Example:
```
docker run --net="host" -e NAME="My Computer" -e PORTS="8080,9090" hello-web:latest
```

Response Example:
```
Hello, World! You have reached port 8080 of My Computer from 127.0.0.1:41732.

Host Details:
Name: My Computer
Hostname: liberty

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
Remote Address: 127.0.0.1:41732
Request URI: /
```

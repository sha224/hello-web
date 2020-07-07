# hello-web
Simple image to run a web service on multiple ports. The web service shows various information about the host and the request. This image can be used for demo and test purposes.

Usage Example:
```
docker run -e NAME="My Computer" -e PORTS="8080,9090" -p 8080:8080 -p 9090:9090 hello-web:latest
```

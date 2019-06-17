## Overview
- Less than 100 lines of go
- Simple Demo for RESTful Style API using golang

## Command
- How to Run
```
$ go build
```

- Curl
```
$ curl -X GET 'localhost:8080/users'
[{"id":1,"username":"Frank"},{"id":2,"username":"Lisa"},{"id":3,"username":"Penny"}]

$ curl -X POST 'localhost:8080/user' -d '{"username": "Alex"}'
{"id":4,"username":"Alex"}

$ curl -X GET 'localhost:8080/user/2'
{"id":2,"username":"Lisa"}

$ curl -X DELETE 'localhost:8080/user/3'
[{"id":1,"username":"Frank"},{"id":2,"username":"Lisa"}]

$ curl -X POST 'localhost:8080/user/2' -d '{"username": "Sunny"}'
{"id":2,"username":"Sunny"}

```

- Reference
    - [Building a RESTful API with Go (Part 1) – John Eckert – Medium](https://medium.com/@johnteckert/building-a-restful-api-with-go-part-1-9e234774b14d)
    - [Building a RESTful API with Go (Part 2) – Codezillas – Medium](https://medium.com/codezillas/building-a-restful-api-with-go-part-2-22d99285c130)
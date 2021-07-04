```shell
curl -d '{"username":"root","password":"root"}' -H 'Content-Type: application/json' -i -X POST  "http://localhost:8888/login"

curl -i "http://localhost:8888/add?book=go-zero&price=10"

curl -i "http://localhost:8888/add?book=go-zero&price=10" -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjUzOTYwNDYsImlhdCI6MTYyNTM2NDA0NiwidXNlcm5hbWUiOiJyb290In0.hdvVvCAutCJvRVFDNShykGg0R2ZOrDv4nAJC4IBJ6eY'


curl -i "http://localhost:8888/check?book=go-zero"

curl -i "http://localhost:8888/check?book=go-zero" -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjUzOTYwNDYsImlhdCI6MTYyNTM2NDA0NiwidXNlcm5hbWUiOiJyb290In0.hdvVvCAutCJvRVFDNShykGg0R2ZOrDv4nAJC4IBJ6eY'

```

## 疑问 为什么要所有的rpc能够启动网关侧才能启动
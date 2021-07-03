wrk -t10 -c1000 -d30s --latency "http://localhost:8888/check?book=go-zero"

curl -i "http://localhost:8888/add?book=go-zero&price=10"

curl -i "http://localhost:8888/check?book=go-zero"

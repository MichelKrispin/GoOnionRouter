curl http://localhost:8888/send \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"from": "127.0.0.1:8001", "to": "127.0.0.1:8002"}'
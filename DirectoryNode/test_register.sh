curl http://localhost:8888/register \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"address": "127.0.0.1:8001"}'

curl http://localhost:8888/register \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"address": "127.0.0.1:8002"}'

curl http://localhost:8888/register \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"address": "127.0.0.1:8003"}'
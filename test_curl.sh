curl -X POST \
     -H "Content-Type: application/json" \
     -H "X-API-KEY: test me 1234567890"  \
     -d '{"name": "John", "age": 30}' \
     -v http://localhost:8081/sensor

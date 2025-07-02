curl -X POST http://localhost:8080/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "123456"
}'

curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "123456"
}'

curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer TOKEN"

curl -X POST http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "password": "654321"
}'

curl -X GET http://localhost:8080/api/v1/users/<user_id> \
  -H "Authorization: Bearer <token>"

curl -X PUT http://localhost:8080/api/v1/users/<user_id> \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john@example.com",
    "password": "123456"
}'

curl -X DELETE http://localhost:8080/api/v1/users/<user_id> \
  -H "Authorization: Bearer <token>"
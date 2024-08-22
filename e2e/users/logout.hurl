# creating user
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe",
    "email": "john@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201

# creating session
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "jhondoe",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201

[Captures]
token: jsonpath "$.token"

# deleting session
DELETE http://{{HOST}}:{{PORT}}/sessions
Authorization: Bearer {{token}}

HTTP 204

# deleting session without authorization returns 401
DELETE http://{{HOST}}:{{PORT}}/sessions
HTTP 401

# deleting session without a valid token returns 401
DELETE http://{{HOST}}:{{PORT}}/sessions
Authorization: Bearer 123
HTTP 401

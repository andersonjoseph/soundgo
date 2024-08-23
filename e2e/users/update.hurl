# creating user
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe",
    "email": "john@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201
[Captures]
userID: jsonpath "$.id"

# creating session
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "jhondoe",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201
[Captures]
token: jsonpath "$.token"

# updating user
PATCH http://{{HOST}}:{{PORT}}/users/{{userID}}
Authorization: Bearer {{token}}
{
    "username": "doejhon"
}
HTTP 200
[Asserts]
jsonpath "$.id" exists

jsonpath "$.username" == "doejhon"
jsonpath "$.email" == "john@mail.com"

jsonpath "$.password" not exists

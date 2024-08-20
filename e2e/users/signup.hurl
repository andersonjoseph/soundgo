# creating user user
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe",
    "email": "john@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201

[Asserts]
jsonpath "$.id" exists

jsonpath "$.username" == "jhondoe"
jsonpath "$.email" == "john@mail.com"

jsonpath "$.password" not exists
jsonpath "$.createdAt" exists

# creating existing user returns conflict error
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe",
    "email": "john@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 409

# creating user with bad username returns bad request error
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhon doe",
    "email": "john@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 400

# creating user with bad email returns bad request error
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe",
    "email": "not-valid-email.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 400

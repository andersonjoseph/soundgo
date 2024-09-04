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

# upload file
POST http://{{HOST}}:{{PORT}}/audios
Authorization: Bearer {{token}}
[MultipartFormData]
title: my first audio
description: just a test
file: file,test.m4a; audio/mp4
HTTP 201
[Captures]
id: jsonpath "$.id"

# stream file
GET http://{{HOST}}:{{PORT}}/audios/{{id}}/file
Authorization: Bearer {{token}}
HTTP 200

# stream file without auth
GET http://{{HOST}}:{{PORT}}/audios/{{id}}/file
HTTP 200

# stream non existing file
GET http://{{HOST}}:{{PORT}}/audios/0191be94-e59d-7218-9478-52d948db3f30/file
HTTP 404

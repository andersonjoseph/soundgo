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
status: published
file: file,test.m4a;
HTTP 201
[Captures]
audioID: jsonpath "$.id"

# update audio
PATCH http://{{HOST}}:{{PORT}}/audios/{{audioID}}
Authorization: Bearer {{token}}
{
    "status": "hidden",
    "title": "my now hidden audio"
}
HTTP 200

# get audio
GET http://{{HOST}}:{{PORT}}/audios/{{audioID}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.title" == "my now hidden audio"
jsonpath "$.description" == "just a test"
jsonpath "$.status" == "hidden"

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
status: hidden
HTTP 201
[Captures]
audioID: jsonpath "$.id"

# get audio
GET http://{{HOST}}:{{PORT}}/audios/{{audioID}}
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.id" exists
jsonpath "$.createdAt" exists

jsonpath "$.title" == "my first audio"
jsonpath "$.description" == "just a test"
jsonpath "$.playCount" == 0
jsonpath "$.status" == "hidden"

# creating other user
POST http://{{HOST}}:{{PORT}}/users
{
    "username": "jhondoe2",
    "email": "john2@mail.com",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201
[Captures]
userID: jsonpath "$.id"

# creating session
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "jhondoe2",
    "password": "Str0ngp4s$w0rd"
}
HTTP 201
[Captures]
token: jsonpath "$.token"

# getting hidden audio
GET http://{{HOST}}:{{PORT}}/audios/{{audioID}}
Authorization: Bearer {{token}}
HTTP 403

# getting hidden audio without auth
GET http://{{HOST}}:{{PORT}}/audios/{{audioID}}
HTTP 403

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
status: published
HTTP 201
[Captures]
id: jsonpath "$.id"

# delete audio
DELETE http://{{HOST}}:{{PORT}}/audios/{{id}}
Authorization: Bearer {{token}}
HTTP 204

# get deleted audio
GET http://{{HOST}}:{{PORT}}/audios/{{id}}
HTTP 404

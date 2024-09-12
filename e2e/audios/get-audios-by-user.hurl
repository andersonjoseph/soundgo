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

# upload a hidden file
POST http://{{HOST}}:{{PORT}}/audios
Authorization: Bearer {{token}}
[MultipartFormData]
title: my first audio
description: just a test
file: file,test.m4a; audio/mp4
status: hidden
HTTP 201

# get audios by user
GET http://{{HOST}}:{{PORT}}/users/{{userID}}/audios
HTTP 200
[Asserts]
jsonpath "$[0].id" exists
jsonpath "$[0].createdAt" exists

jsonpath "$[0].title" == "my first audio"
jsonpath "$[0].description" == "just a test"
jsonpath "$[0].playCount" == 0
jsonpath "$[0].status" == "published"

jsonpath "$[1].id" not exists

# get audios by user while authenticated as owner
GET http://{{HOST}}:{{PORT}}/users/{{userID}}/audios
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$[0].id" exists
jsonpath "$[0].createdAt" exists

jsonpath "$[0].title" == "my first audio"
jsonpath "$[0].description" == "just a test"
jsonpath "$[0].playCount" == 0
jsonpath "$[0].status" == "published"

jsonpath "$[1].id" exists

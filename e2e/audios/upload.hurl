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
[Asserts]
jsonpath "$.id" exists
jsonpath "$.createdAt" exists

jsonpath "$.title" == "my first audio"
jsonpath "$.description" == "just a test"
jsonpath "$.playCount" == 0
jsonpath "$.status" == "published"

# upload invalid file
POST http://{{HOST}}:{{PORT}}/audios
Authorization: Bearer {{token}}
[MultipartFormData]
title: my invalid audio
description: just a test
status: published
file: file,test.jpg;
HTTP 415

# upload a file without auth
POST http://{{HOST}}:{{PORT}}/audios
[MultipartFormData]
title: my audio
description: just a test
file: file,test.m4a;
HTTP 401

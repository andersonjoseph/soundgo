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

[Asserts]
jsonpath "$.token" exists

# creating session with wrong password returns 401
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "jhondoe",
    "password": "wron"
}
HTTP 401

# creating session with unexisting username returns 401
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "missing",
    "password": "Str0ngp4s$w0rd"
}
HTTP 401

# missing username returns 400
POST http://{{HOST}}:{{PORT}}/sessions
{
    "password": "Str0ngp4s$w0rd"
}
HTTP 400

# missing password returns 400
POST http://{{HOST}}:{{PORT}}/sessions
{
    "username": "missing"
}
HTTP 400

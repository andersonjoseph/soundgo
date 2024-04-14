# create new user

POST http://{{host}}/api/v1/users
{
    "email": "test@mail.com",
    "username": "test",
    "password": "Str0ngP4$$w0rD!"
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
jsonpath "$.email" == "test@mail.com"
jsonpath "$.username" == "test"
jsonpath "$.password" not exists

# create existing user

POST http://{{host}}/api/v1/users
{
    "email": "test@mail.com",
    "username": "test",
    "password": "Str0ngP4$$w0rD!"
}
HTTP 409

# create user with invalid email

POST http://{{host}}/api/v1/users
{
    "email": "invalid-email",
    "username": "test",
    "password": "Str0ngP4$$w0rD!"
}
HTTP 400

# create user with invalid username

POST http://{{host}}/api/v1/users
{
    "email": "test@mail.com",
    "username": "u",
    "password": "Str0ngP4$$w0rD!"
}
HTTP 400

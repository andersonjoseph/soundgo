# create new user

POST http://{{host}}/api/v1/users
{
    "email": "login_user@mail.com",
    "username": "login_user",
    "password": "1234567890"
}
HTTP 201

# login user

POST http://{{host}}/api/v1/sessions
{
    "username": "login_user",
    "password": "1234567890"
}
HTTP 201
[Asserts]
jsonpath "$.token" exists

# login user with wrong password

POST http://{{host}}/api/v1/sessions
{
    "username": "login_user",
    "password": "0987654321"
}
HTTP 401

# login user with wrong username

POST http://{{host}}/api/v1/sessions
{
    "username": "login_user_bad",
    "password": "1234567890"
}
HTTP 401

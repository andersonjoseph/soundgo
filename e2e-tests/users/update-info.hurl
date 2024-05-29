# create new user

POST http://{{host}}/api/v1/users
{
    "email": "user_to_update@mail.com",
    "username": "user_to_update",
    "password": "1234567890"
}
HTTP 201
[Captures] 
id: jsonpath "$.id"

# create another user

POST http://{{host}}/api/v1/users
{
    "email": "user_to_update2@mail.com",
    "username": "user_to_update2",
    "password": "1234567890"
}
HTTP 201

# login user

POST http://{{host}}/api/v1/sessions
{
    "username": "user_to_update",
    "password": "1234567890"
}
HTTP 201
[Captures] 
token: jsonpath "$.token"

# update user

PATCH http://{{host}}/api/v1/users/{{id}}
{
    "username": "updated_user"
}
HTTP 200

# update user to an existing username

PATCH http://{{host}}/api/v1/users/{{id}}
{
    "username": "user_to_update2"
}
HTTP 409

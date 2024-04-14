# create new user

POST http://{{host}}/api/v1/users
{
    "email": "user_resetting_password@mail.com",
    "username": "user_resetting_password",
    "password": "1234567890"
}
HTTP 201

# request a password reset
POST http://{{host}}/api/v1/password-reset
{
    "email": "user_resetting_password@mail.com",
    "username": "user_resetting_password"
}
HTTP 201

# request a password reset for a non existing user
POST http://{{host}}/api/v1/password-reset
{
    "email": "non_existing_user@mail.com",
    "username": "non_existing_user"
}
HTTP 201

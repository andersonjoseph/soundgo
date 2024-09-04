// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CheckHealth implements checkHealth operation.
//
// This operation checks the health status of the API and returns a 200 status code
// if the API is functioning correctly. This can be used as a health check endpoint
// for monitoring purposes.
//
// GET /health
func (UnimplementedHandler) CheckHealth(ctx context.Context) (r CheckHealthRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateAudio implements createAudio operation.
//
// This operation allows the client to upload an audio file. The server stores the file and returns
// the ID of the created resource.
//
// POST /audios
func (UnimplementedHandler) CreateAudio(ctx context.Context, req *AudioInputMultipart) (r CreateAudioRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePasswordResetRequest implements createPasswordResetRequest operation.
//
// This operation initiates a password reset process by creating a password reset request.
// If the provided email is associated with a user account, an email with password reset code is sent.
//
// POST /password-reset
func (UnimplementedHandler) CreatePasswordResetRequest(ctx context.Context, req *PasswordResetRequestInput) (r CreatePasswordResetRequestRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSession implements createSession operation.
//
// This operation creates a new session (user login).
//
// POST /sessions
func (UnimplementedHandler) CreateSession(ctx context.Context, req *SessionInput) (r CreateSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// This operation creates a new user in the system using the provided information.
// The request body must include all the necessary details required for creating a user.
// If the user is successfully created, the server will return a 201 status code along with the user
// details.
// If there are any issues with the data or if the user already exists, appropriate error responses
// will be returned.
//
// POST /users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *UserInput) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSession implements deleteSession operation.
//
// This operation logs the user out by deleting the current session.
//
// DELETE /sessions
func (UnimplementedHandler) DeleteSession(ctx context.Context) (r DeleteSessionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAudio implements getAudio operation.
//
// This operation gets an audio with the given ID.
//
// GET /audios/{id}
func (UnimplementedHandler) GetAudio(ctx context.Context, params GetAudioParams) (r GetAudioRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAudioFile implements getAudioFile operation.
//
// This operation streams an audio file with the given ID. The client can request the entire file or
// a specific byte range to enable partial downloads and streaming.
//
// GET /audios/{id}/file
func (UnimplementedHandler) GetAudioFile(ctx context.Context, params GetAudioFileParams) (r GetAudioFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ResetPassword implements resetPassword operation.
//
// This operation resets a user's password. The request requires a valid password reset code and a
// new password.
// If the reset code is valid and the new password meets the required criteria, the password will be
// updated.
//
// PUT /password-reset
func (UnimplementedHandler) ResetPassword(ctx context.Context, req *PasswordResetInput) (r ResetPasswordRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// This operation updates the details of an existing user in the system using the provided
// information.
// The request body must include all the necessary details required for updating the user.
// If the user is successfully updated, the server will return a 200 status code along with the
// updated user details.
// If there are any issues with the input data or if the user does not exist, appropriate error
// responses will be returned.
//
// PATCH /users/{id}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UpdateUserInput, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

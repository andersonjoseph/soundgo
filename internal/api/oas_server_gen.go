// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CheckHealth implements checkHealth operation.
	//
	// This operation checks the health status of the API and returns a 200 status code
	// if the API is functioning correctly. This can be used as a health check endpoint
	// for monitoring purposes.
	//
	// GET /health
	CheckHealth(ctx context.Context) (CheckHealthRes, error)
	// CreateAudio implements createAudio operation.
	//
	// This operation allows the client to upload an audio file. The server stores the file and returns
	// the ID of the created resource.
	//
	// POST /audios
	CreateAudio(ctx context.Context, req *AudioInputMultipart) (CreateAudioRes, error)
	// CreatePasswordResetRequest implements createPasswordResetRequest operation.
	//
	// This operation initiates a password reset process by creating a password reset request.
	// If the provided email is associated with a user account, an email with password reset code is sent.
	//
	// POST /password-reset
	CreatePasswordResetRequest(ctx context.Context, req *PasswordResetRequestInput) (CreatePasswordResetRequestRes, error)
	// CreateSession implements createSession operation.
	//
	// This operation creates a new session (user login).
	//
	// POST /sessions
	CreateSession(ctx context.Context, req *SessionInput) (CreateSessionRes, error)
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
	CreateUser(ctx context.Context, req *UserInput) (CreateUserRes, error)
	// DeleteSession implements deleteSession operation.
	//
	// This operation logs the user out by deleting the current session.
	//
	// DELETE /sessions
	DeleteSession(ctx context.Context) (DeleteSessionRes, error)
	// GetAudioFile implements getAudioFile operation.
	//
	// This operation streams an audio file with the given ID. The client can request the entire file or
	// a specific byte range to enable partial downloads and streaming.
	//
	// GET /audios/{id}/file
	GetAudioFile(ctx context.Context, params GetAudioFileParams) (GetAudioFileRes, error)
	// ResetPassword implements resetPassword operation.
	//
	// This operation resets a user's password. The request requires a valid password reset code and a
	// new password.
	// If the reset code is valid and the new password meets the required criteria, the password will be
	// updated.
	//
	// PUT /password-reset
	ResetPassword(ctx context.Context, req *PasswordResetInput) (ResetPasswordRes, error)
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
	UpdateUser(ctx context.Context, req *UpdateUserInput, params UpdateUserParams) (UpdateUserRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}

// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// CreatePasswordReset invokes createPasswordReset operation.
	//
	// This operation initiates a password reset process by creating a password reset request.
	// If the provided email is associated with a user account, an email with password reset code is sent.
	//
	// POST /password-reset
	CreatePasswordReset(ctx context.Context, request *PasswordResetRequestInput) (CreatePasswordResetRes, error)
	// CreateSession invokes createSession operation.
	//
	// This operation creates a new session (user login).
	//
	// POST /sessions
	CreateSession(ctx context.Context, request *SessionInput) (CreateSessionRes, error)
	// CreateUser invokes createUser operation.
	//
	// This operation creates a new user in the system using the provided information.
	// The request body must include all the necessary details required for creating a user.
	// If the user is successfully created, the server will return a 201 status code along with the user
	// details.
	// If there are any issues with the data or if the user already exists, appropriate error responses
	// will be returned.
	//
	// POST /users
	CreateUser(ctx context.Context, request *UserInput) (CreateUserRes, error)
	// DeleteSession invokes deleteSession operation.
	//
	// This operation logs the user out by deleting the current session.
	//
	// DELETE /sessions
	DeleteSession(ctx context.Context) (DeleteSessionRes, error)
	// ResetPassword invokes resetPassword operation.
	//
	// This operation resets a user's password. The request requires a valid password reset code and a
	// new password.
	// If the reset code is valid and the new password meets the required criteria, the password will be
	// updated.
	//
	// PUT /password-reset
	ResetPassword(ctx context.Context, request *PasswordResetInput) (ResetPasswordRes, error)
	// UpdateUser invokes updateUser operation.
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
	UpdateUser(ctx context.Context, request *UpdateUserInput, params UpdateUserParams) (UpdateUserRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// CreatePasswordReset invokes createPasswordReset operation.
//
// This operation initiates a password reset process by creating a password reset request.
// If the provided email is associated with a user account, an email with password reset code is sent.
//
// POST /password-reset
func (c *Client) CreatePasswordReset(ctx context.Context, request *PasswordResetRequestInput) (CreatePasswordResetRes, error) {
	res, err := c.sendCreatePasswordReset(ctx, request)
	return res, err
}

func (c *Client) sendCreatePasswordReset(ctx context.Context, request *PasswordResetRequestInput) (res CreatePasswordResetRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPasswordReset"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/password-reset"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreatePasswordReset",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/password-reset"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreatePasswordResetRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreatePasswordResetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSession invokes createSession operation.
//
// This operation creates a new session (user login).
//
// POST /sessions
func (c *Client) CreateSession(ctx context.Context, request *SessionInput) (CreateSessionRes, error) {
	res, err := c.sendCreateSession(ctx, request)
	return res, err
}

func (c *Client) sendCreateSession(ctx context.Context, request *SessionInput) (res CreateSessionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSession"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/sessions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateSession",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/sessions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateSessionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateSessionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateUser invokes createUser operation.
//
// This operation creates a new user in the system using the provided information.
// The request body must include all the necessary details required for creating a user.
// If the user is successfully created, the server will return a 201 status code along with the user
// details.
// If there are any issues with the data or if the user already exists, appropriate error responses
// will be returned.
//
// POST /users
func (c *Client) CreateUser(ctx context.Context, request *UserInput) (CreateUserRes, error) {
	res, err := c.sendCreateUser(ctx, request)
	return res, err
}

func (c *Client) sendCreateUser(ctx context.Context, request *UserInput) (res CreateUserRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createUser"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/users"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateUser",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateUserRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteSession invokes deleteSession operation.
//
// This operation logs the user out by deleting the current session.
//
// DELETE /sessions
func (c *Client) DeleteSession(ctx context.Context) (DeleteSessionRes, error) {
	res, err := c.sendDeleteSession(ctx)
	return res, err
}

func (c *Client) sendDeleteSession(ctx context.Context) (res DeleteSessionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteSession"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/sessions"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "DeleteSession",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/sessions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "DeleteSession", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteSessionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ResetPassword invokes resetPassword operation.
//
// This operation resets a user's password. The request requires a valid password reset code and a
// new password.
// If the reset code is valid and the new password meets the required criteria, the password will be
// updated.
//
// PUT /password-reset
func (c *Client) ResetPassword(ctx context.Context, request *PasswordResetInput) (ResetPasswordRes, error) {
	res, err := c.sendResetPassword(ctx, request)
	return res, err
}

func (c *Client) sendResetPassword(ctx context.Context, request *PasswordResetInput) (res ResetPasswordRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resetPassword"),
		semconv.HTTPMethodKey.String("PUT"),
		semconv.HTTPRouteKey.String("/password-reset"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ResetPassword",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/password-reset"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeResetPasswordRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeResetPasswordResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateUser invokes updateUser operation.
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
func (c *Client) UpdateUser(ctx context.Context, request *UpdateUserInput, params UpdateUserParams) (UpdateUserRes, error) {
	res, err := c.sendUpdateUser(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdateUser(ctx context.Context, request *UpdateUserInput, params UpdateUserParams) (res UpdateUserRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateUser"),
		semconv.HTTPMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/users/{id}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "UpdateUser",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PATCH", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateUserRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "UpdateUser", r); {
			case err == nil: // if NO error
				satisfied[0] |= 1 << 0
			case errors.Is(err, ogenerrors.ErrSkipClientSecurity):
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, ogenerrors.ErrSecurityRequirementIsNotSatisfied
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

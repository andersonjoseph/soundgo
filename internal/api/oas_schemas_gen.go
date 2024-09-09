// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"time"

	"github.com/go-faster/errors"

	ht "github.com/ogen-go/ogen/http"
)

// Ref: #/components/schemas/audio
type Audio struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description OptString   `json:"description"`
	CreatedAt   time.Time   `json:"createdAt"`
	User        string      `json:"user"`
	Status      AudioStatus `json:"status"`
	PlayCount   int         `json:"playCount"`
}

// GetID returns the value of ID.
func (s *Audio) GetID() string {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Audio) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *Audio) GetDescription() OptString {
	return s.Description
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Audio) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUser returns the value of User.
func (s *Audio) GetUser() string {
	return s.User
}

// GetStatus returns the value of Status.
func (s *Audio) GetStatus() AudioStatus {
	return s.Status
}

// GetPlayCount returns the value of PlayCount.
func (s *Audio) GetPlayCount() int {
	return s.PlayCount
}

// SetID sets the value of ID.
func (s *Audio) SetID(val string) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Audio) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *Audio) SetDescription(val OptString) {
	s.Description = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Audio) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUser sets the value of User.
func (s *Audio) SetUser(val string) {
	s.User = val
}

// SetStatus sets the value of Status.
func (s *Audio) SetStatus(val AudioStatus) {
	s.Status = val
}

// SetPlayCount sets the value of PlayCount.
func (s *Audio) SetPlayCount(val int) {
	s.PlayCount = val
}

func (*Audio) createAudioRes() {}
func (*Audio) getAudioRes()    {}
func (*Audio) updateAudioRes() {}

// Ref: #/components/schemas/audio-input
type AudioInputMultipart struct {
	// The audio file to upload.
	File        ht.MultipartFile          `json:"file"`
	Title       string                    `json:"title"`
	Description OptString                 `json:"description"`
	Status      AudioInputMultipartStatus `json:"status"`
}

// GetFile returns the value of File.
func (s *AudioInputMultipart) GetFile() ht.MultipartFile {
	return s.File
}

// GetTitle returns the value of Title.
func (s *AudioInputMultipart) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *AudioInputMultipart) GetDescription() OptString {
	return s.Description
}

// GetStatus returns the value of Status.
func (s *AudioInputMultipart) GetStatus() AudioInputMultipartStatus {
	return s.Status
}

// SetFile sets the value of File.
func (s *AudioInputMultipart) SetFile(val ht.MultipartFile) {
	s.File = val
}

// SetTitle sets the value of Title.
func (s *AudioInputMultipart) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *AudioInputMultipart) SetDescription(val OptString) {
	s.Description = val
}

// SetStatus sets the value of Status.
func (s *AudioInputMultipart) SetStatus(val AudioInputMultipartStatus) {
	s.Status = val
}

type AudioInputMultipartStatus string

const (
	AudioInputMultipartStatusPublished AudioInputMultipartStatus = "published"
	AudioInputMultipartStatusHidden    AudioInputMultipartStatus = "hidden"
)

// AllValues returns all AudioInputMultipartStatus values.
func (AudioInputMultipartStatus) AllValues() []AudioInputMultipartStatus {
	return []AudioInputMultipartStatus{
		AudioInputMultipartStatusPublished,
		AudioInputMultipartStatusHidden,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s AudioInputMultipartStatus) MarshalText() ([]byte, error) {
	switch s {
	case AudioInputMultipartStatusPublished:
		return []byte(s), nil
	case AudioInputMultipartStatusHidden:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *AudioInputMultipartStatus) UnmarshalText(data []byte) error {
	switch AudioInputMultipartStatus(data) {
	case AudioInputMultipartStatusPublished:
		*s = AudioInputMultipartStatusPublished
		return nil
	case AudioInputMultipartStatusHidden:
		*s = AudioInputMultipartStatusHidden
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type AudioStatus string

const (
	AudioStatusPublished AudioStatus = "published"
	AudioStatusPending   AudioStatus = "pending"
	AudioStatusHidden    AudioStatus = "hidden"
)

// AllValues returns all AudioStatus values.
func (AudioStatus) AllValues() []AudioStatus {
	return []AudioStatus{
		AudioStatusPublished,
		AudioStatusPending,
		AudioStatusHidden,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s AudioStatus) MarshalText() ([]byte, error) {
	switch s {
	case AudioStatusPublished:
		return []byte(s), nil
	case AudioStatusPending:
		return []byte(s), nil
	case AudioStatusHidden:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *AudioStatus) UnmarshalText(data []byte) error {
	switch AudioStatus(data) {
	case AudioStatusPublished:
		*s = AudioStatusPublished
		return nil
	case AudioStatusPending:
		*s = AudioStatusPending
		return nil
	case AudioStatusHidden:
		*s = AudioStatusHidden
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

type CheckHealthOK struct {
	// The current health status of the API. It will return "healthy" if
	// the API is functioning properly, and "unhealthy" otherwise.
	Status CheckHealthOKStatus `json:"status"`
}

// GetStatus returns the value of Status.
func (s *CheckHealthOK) GetStatus() CheckHealthOKStatus {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *CheckHealthOK) SetStatus(val CheckHealthOKStatus) {
	s.Status = val
}

func (*CheckHealthOK) checkHealthRes() {}

// The current health status of the API. It will return "healthy" if
// the API is functioning properly, and "unhealthy" otherwise.
type CheckHealthOKStatus string

const (
	CheckHealthOKStatusHealthy   CheckHealthOKStatus = "healthy"
	CheckHealthOKStatusUnhealthy CheckHealthOKStatus = "unhealthy"
)

// AllValues returns all CheckHealthOKStatus values.
func (CheckHealthOKStatus) AllValues() []CheckHealthOKStatus {
	return []CheckHealthOKStatus{
		CheckHealthOKStatusHealthy,
		CheckHealthOKStatusUnhealthy,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CheckHealthOKStatus) MarshalText() ([]byte, error) {
	switch s {
	case CheckHealthOKStatusHealthy:
		return []byte(s), nil
	case CheckHealthOKStatusUnhealthy:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CheckHealthOKStatus) UnmarshalText(data []byte) error {
	switch CheckHealthOKStatus(data) {
	case CheckHealthOKStatusHealthy:
		*s = CheckHealthOKStatusHealthy
		return nil
	case CheckHealthOKStatusUnhealthy:
		*s = CheckHealthOKStatusUnhealthy
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type CheckHealthServiceUnavailable struct {
	// The current health status of the API.
	Status CheckHealthServiceUnavailableStatus `json:"status"`
}

// GetStatus returns the value of Status.
func (s *CheckHealthServiceUnavailable) GetStatus() CheckHealthServiceUnavailableStatus {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *CheckHealthServiceUnavailable) SetStatus(val CheckHealthServiceUnavailableStatus) {
	s.Status = val
}

func (*CheckHealthServiceUnavailable) checkHealthRes() {}

// The current health status of the API.
type CheckHealthServiceUnavailableStatus string

const (
	CheckHealthServiceUnavailableStatusHealthy   CheckHealthServiceUnavailableStatus = "healthy"
	CheckHealthServiceUnavailableStatusUnhealthy CheckHealthServiceUnavailableStatus = "unhealthy"
)

// AllValues returns all CheckHealthServiceUnavailableStatus values.
func (CheckHealthServiceUnavailableStatus) AllValues() []CheckHealthServiceUnavailableStatus {
	return []CheckHealthServiceUnavailableStatus{
		CheckHealthServiceUnavailableStatusHealthy,
		CheckHealthServiceUnavailableStatusUnhealthy,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CheckHealthServiceUnavailableStatus) MarshalText() ([]byte, error) {
	switch s {
	case CheckHealthServiceUnavailableStatusHealthy:
		return []byte(s), nil
	case CheckHealthServiceUnavailableStatusUnhealthy:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CheckHealthServiceUnavailableStatus) UnmarshalText(data []byte) error {
	switch CheckHealthServiceUnavailableStatus(data) {
	case CheckHealthServiceUnavailableStatusHealthy:
		*s = CheckHealthServiceUnavailableStatusHealthy
		return nil
	case CheckHealthServiceUnavailableStatusUnhealthy:
		*s = CheckHealthServiceUnavailableStatusUnhealthy
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type Conflict struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *Conflict) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *Conflict) SetError(val string) {
	s.Error = val
}

func (*Conflict) updateUserRes() {}

type CreateAudioBadRequest struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *CreateAudioBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *CreateAudioBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*CreateAudioBadRequest) createAudioRes() {}

type CreateAudioUnsupportedMediaType struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *CreateAudioUnsupportedMediaType) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *CreateAudioUnsupportedMediaType) SetError(val OptString) {
	s.Error = val
}

func (*CreateAudioUnsupportedMediaType) createAudioRes() {}

type CreatePasswordResetRequestBadRequest struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *CreatePasswordResetRequestBadRequest) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *CreatePasswordResetRequestBadRequest) SetError(val string) {
	s.Error = val
}

func (*CreatePasswordResetRequestBadRequest) createPasswordResetRequestRes() {}

// CreatePasswordResetRequestNoContent is response for CreatePasswordResetRequest operation.
type CreatePasswordResetRequestNoContent struct{}

func (*CreatePasswordResetRequestNoContent) createPasswordResetRequestRes() {}

type CreateUserBadRequest struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *CreateUserBadRequest) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *CreateUserBadRequest) SetError(val string) {
	s.Error = val
}

func (*CreateUserBadRequest) createUserRes() {}

type CreateUserConflict struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *CreateUserConflict) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *CreateUserConflict) SetError(val string) {
	s.Error = val
}

func (*CreateUserConflict) createUserRes() {}

type DeleteAudioForbidden struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *DeleteAudioForbidden) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *DeleteAudioForbidden) SetError(val OptString) {
	s.Error = val
}

func (*DeleteAudioForbidden) deleteAudioRes() {}

// DeleteAudioNoContent is response for DeleteAudio operation.
type DeleteAudioNoContent struct{}

func (*DeleteAudioNoContent) deleteAudioRes() {}

type DeleteAudioNotFound struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *DeleteAudioNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *DeleteAudioNotFound) SetError(val OptString) {
	s.Error = val
}

func (*DeleteAudioNotFound) deleteAudioRes() {}

// DeleteSessionNoContent is response for DeleteSession operation.
type DeleteSessionNoContent struct{}

func (*DeleteSessionNoContent) deleteSessionRes() {}

type Forbidden struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *Forbidden) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *Forbidden) SetError(val string) {
	s.Error = val
}

func (*Forbidden) updateUserRes() {}

type GetAudioFileForbidden struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GetAudioFileForbidden) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GetAudioFileForbidden) SetError(val OptString) {
	s.Error = val
}

func (*GetAudioFileForbidden) getAudioFileRes() {}

type GetAudioFileNotFound struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GetAudioFileNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GetAudioFileNotFound) SetError(val OptString) {
	s.Error = val
}

func (*GetAudioFileNotFound) getAudioFileRes() {}

type GetAudioFileOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetAudioFileOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// GetAudioFileOKHeaders wraps GetAudioFileOK with response headers.
type GetAudioFileOKHeaders struct {
	AcceptRanges OptString
	Response     GetAudioFileOK
}

// GetAcceptRanges returns the value of AcceptRanges.
func (s *GetAudioFileOKHeaders) GetAcceptRanges() OptString {
	return s.AcceptRanges
}

// GetResponse returns the value of Response.
func (s *GetAudioFileOKHeaders) GetResponse() GetAudioFileOK {
	return s.Response
}

// SetAcceptRanges sets the value of AcceptRanges.
func (s *GetAudioFileOKHeaders) SetAcceptRanges(val OptString) {
	s.AcceptRanges = val
}

// SetResponse sets the value of Response.
func (s *GetAudioFileOKHeaders) SetResponse(val GetAudioFileOK) {
	s.Response = val
}

func (*GetAudioFileOKHeaders) getAudioFileRes() {}

type GetAudioFilePartialContent struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetAudioFilePartialContent) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

// GetAudioFilePartialContentHeaders wraps GetAudioFilePartialContent with response headers.
type GetAudioFilePartialContentHeaders struct {
	AcceptRanges OptString
	ContentRange OptString
	Response     GetAudioFilePartialContent
}

// GetAcceptRanges returns the value of AcceptRanges.
func (s *GetAudioFilePartialContentHeaders) GetAcceptRanges() OptString {
	return s.AcceptRanges
}

// GetContentRange returns the value of ContentRange.
func (s *GetAudioFilePartialContentHeaders) GetContentRange() OptString {
	return s.ContentRange
}

// GetResponse returns the value of Response.
func (s *GetAudioFilePartialContentHeaders) GetResponse() GetAudioFilePartialContent {
	return s.Response
}

// SetAcceptRanges sets the value of AcceptRanges.
func (s *GetAudioFilePartialContentHeaders) SetAcceptRanges(val OptString) {
	s.AcceptRanges = val
}

// SetContentRange sets the value of ContentRange.
func (s *GetAudioFilePartialContentHeaders) SetContentRange(val OptString) {
	s.ContentRange = val
}

// SetResponse sets the value of Response.
func (s *GetAudioFilePartialContentHeaders) SetResponse(val GetAudioFilePartialContent) {
	s.Response = val
}

func (*GetAudioFilePartialContentHeaders) getAudioFileRes() {}

type GetAudioFileRequestedRangeNotSatisfiable struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GetAudioFileRequestedRangeNotSatisfiable) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GetAudioFileRequestedRangeNotSatisfiable) SetError(val OptString) {
	s.Error = val
}

func (*GetAudioFileRequestedRangeNotSatisfiable) getAudioFileRes() {}

type GetAudioForbidden struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GetAudioForbidden) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GetAudioForbidden) SetError(val OptString) {
	s.Error = val
}

func (*GetAudioForbidden) getAudioRes() {}

type GetAudioNotFound struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *GetAudioNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *GetAudioNotFound) SetError(val OptString) {
	s.Error = val
}

func (*GetAudioNotFound) getAudioRes() {}

type NotFound struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *NotFound) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *NotFound) SetError(val string) {
	s.Error = val
}

func (*NotFound) updateUserRes() {}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUpdateAudioInputStatus returns new OptUpdateAudioInputStatus with value set to v.
func NewOptUpdateAudioInputStatus(v UpdateAudioInputStatus) OptUpdateAudioInputStatus {
	return OptUpdateAudioInputStatus{
		Value: v,
		Set:   true,
	}
}

// OptUpdateAudioInputStatus is optional UpdateAudioInputStatus.
type OptUpdateAudioInputStatus struct {
	Value UpdateAudioInputStatus
	Set   bool
}

// IsSet returns true if OptUpdateAudioInputStatus was set.
func (o OptUpdateAudioInputStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUpdateAudioInputStatus) Reset() {
	var v UpdateAudioInputStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUpdateAudioInputStatus) SetTo(v UpdateAudioInputStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUpdateAudioInputStatus) Get() (v UpdateAudioInputStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUpdateAudioInputStatus) Or(d UpdateAudioInputStatus) UpdateAudioInputStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/password-reset-input
type PasswordResetInput struct {
	Code     string `json:"code"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetCode returns the value of Code.
func (s *PasswordResetInput) GetCode() string {
	return s.Code
}

// GetEmail returns the value of Email.
func (s *PasswordResetInput) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *PasswordResetInput) GetPassword() string {
	return s.Password
}

// SetCode sets the value of Code.
func (s *PasswordResetInput) SetCode(val string) {
	s.Code = val
}

// SetEmail sets the value of Email.
func (s *PasswordResetInput) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *PasswordResetInput) SetPassword(val string) {
	s.Password = val
}

// Ref: #/components/schemas/password-reset-request-input
type PasswordResetRequestInput struct {
	Email string `json:"email"`
}

// GetEmail returns the value of Email.
func (s *PasswordResetRequestInput) GetEmail() string {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *PasswordResetRequestInput) SetEmail(val string) {
	s.Email = val
}

type ResetPasswordBadRequest struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *ResetPasswordBadRequest) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *ResetPasswordBadRequest) SetError(val string) {
	s.Error = val
}

func (*ResetPasswordBadRequest) resetPasswordRes() {}

// ResetPasswordNoContent is response for ResetPassword operation.
type ResetPasswordNoContent struct{}

func (*ResetPasswordNoContent) resetPasswordRes() {}

// Ref: #/components/schemas/session
type Session struct {
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *Session) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Session) SetToken(val string) {
	s.Token = val
}

func (*Session) createSessionRes() {}

// Ref: #/components/schemas/session-input
type SessionInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *SessionInput) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *SessionInput) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *SessionInput) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *SessionInput) SetPassword(val string) {
	s.Password = val
}

type Unauthorized struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *Unauthorized) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *Unauthorized) SetError(val string) {
	s.Error = val
}

func (*Unauthorized) createSessionRes() {}
func (*Unauthorized) deleteSessionRes() {}
func (*Unauthorized) updateUserRes()    {}

type UpdateAudioBadRequest struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *UpdateAudioBadRequest) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *UpdateAudioBadRequest) SetError(val OptString) {
	s.Error = val
}

func (*UpdateAudioBadRequest) updateAudioRes() {}

type UpdateAudioForbidden struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *UpdateAudioForbidden) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *UpdateAudioForbidden) SetError(val OptString) {
	s.Error = val
}

func (*UpdateAudioForbidden) updateAudioRes() {}

// Ref: #/components/schemas/update-audio-input
type UpdateAudioInput struct {
	Title       OptString                 `json:"title"`
	Description OptString                 `json:"description"`
	Status      OptUpdateAudioInputStatus `json:"status"`
}

// GetTitle returns the value of Title.
func (s *UpdateAudioInput) GetTitle() OptString {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *UpdateAudioInput) GetDescription() OptString {
	return s.Description
}

// GetStatus returns the value of Status.
func (s *UpdateAudioInput) GetStatus() OptUpdateAudioInputStatus {
	return s.Status
}

// SetTitle sets the value of Title.
func (s *UpdateAudioInput) SetTitle(val OptString) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *UpdateAudioInput) SetDescription(val OptString) {
	s.Description = val
}

// SetStatus sets the value of Status.
func (s *UpdateAudioInput) SetStatus(val OptUpdateAudioInputStatus) {
	s.Status = val
}

type UpdateAudioInputStatus string

const (
	UpdateAudioInputStatusPublished UpdateAudioInputStatus = "published"
	UpdateAudioInputStatusHidden    UpdateAudioInputStatus = "hidden"
)

// AllValues returns all UpdateAudioInputStatus values.
func (UpdateAudioInputStatus) AllValues() []UpdateAudioInputStatus {
	return []UpdateAudioInputStatus{
		UpdateAudioInputStatusPublished,
		UpdateAudioInputStatusHidden,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UpdateAudioInputStatus) MarshalText() ([]byte, error) {
	switch s {
	case UpdateAudioInputStatusPublished:
		return []byte(s), nil
	case UpdateAudioInputStatusHidden:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UpdateAudioInputStatus) UnmarshalText(data []byte) error {
	switch UpdateAudioInputStatus(data) {
	case UpdateAudioInputStatusPublished:
		*s = UpdateAudioInputStatusPublished
		return nil
	case UpdateAudioInputStatusHidden:
		*s = UpdateAudioInputStatusHidden
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type UpdateAudioNotFound struct {
	Error OptString `json:"error"`
}

// GetError returns the value of Error.
func (s *UpdateAudioNotFound) GetError() OptString {
	return s.Error
}

// SetError sets the value of Error.
func (s *UpdateAudioNotFound) SetError(val OptString) {
	s.Error = val
}

func (*UpdateAudioNotFound) updateAudioRes() {}

type UpdateUserBadRequest struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *UpdateUserBadRequest) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *UpdateUserBadRequest) SetError(val string) {
	s.Error = val
}

func (*UpdateUserBadRequest) updateUserRes() {}

// Ref: #/components/schemas/update-user-input
type UpdateUserInput struct {
	Username string `json:"username"`
}

// GetUsername returns the value of Username.
func (s *UpdateUserInput) GetUsername() string {
	return s.Username
}

// SetUsername sets the value of Username.
func (s *UpdateUserInput) SetUsername(val string) {
	s.Username = val
}

// Ref: #/components/schemas/user
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetID returns the value of ID.
func (s *User) GetID() string {
	return s.ID
}

// GetUsername returns the value of Username.
func (s *User) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val string) {
	s.ID = val
}

// SetUsername sets the value of Username.
func (s *User) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

func (*User) createUserRes() {}
func (*User) updateUserRes() {}

// Ref: #/components/schemas/user-input
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUsername returns the value of Username.
func (s *UserInput) GetUsername() string {
	return s.Username
}

// GetEmail returns the value of Email.
func (s *UserInput) GetEmail() string {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *UserInput) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *UserInput) SetUsername(val string) {
	s.Username = val
}

// SetEmail sets the value of Email.
func (s *UserInput) SetEmail(val string) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *UserInput) SetPassword(val string) {
	s.Password = val
}

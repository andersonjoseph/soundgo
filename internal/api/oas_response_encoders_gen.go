// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func encodeCheckHealthResponse(response CheckHealthRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CheckHealthOK:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CheckHealthServiceUnavailable:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateAudioResponse(response CreateAudioRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Audio:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateAudioBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateAudioUnsupportedMediaType:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(415)
		span.SetStatus(codes.Error, http.StatusText(415))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreatePasswordResetRequestResponse(response CreatePasswordResetRequestRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreatePasswordResetRequestNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *CreatePasswordResetRequestBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateSessionResponse(response CreateSessionRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Session:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Unauthorized:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeCreateUserResponse(response CreateUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *User:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateUserBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *CreateUserConflict:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteAudioResponse(response DeleteAudioRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteAudioNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *DeleteAudioForbidden:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *DeleteAudioNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteSessionResponse(response DeleteSessionRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteSessionNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *Unauthorized:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetAudioResponse(response GetAudioRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Audio:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetAudioForbidden:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetAudioNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetAudioFileResponse(response GetAudioFileRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetAudioFileOKHeaders:
		w.Header().Set("Content-Type", "audio/mp4")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Accept-Ranges" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Accept-Ranges",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AcceptRanges.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Accept-Ranges header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response.Response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetAudioFilePartialContentHeaders:
		w.Header().Set("Content-Type", "audio/mp4")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Accept-Ranges" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Accept-Ranges",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.AcceptRanges.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Accept-Ranges header")
				}
			}
			// Encode "Content-Range" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Content-Range",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.ContentRange.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Content-Range header")
				}
			}
		}
		w.WriteHeader(206)
		span.SetStatus(codes.Ok, http.StatusText(206))

		writer := w
		if _, err := io.Copy(writer, response.Response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetAudioFileNotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *GetAudioFileRequestedRangeNotSatisfiable:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(416)
		span.SetStatus(codes.Error, http.StatusText(416))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeResetPasswordResponse(response ResetPasswordRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ResetPasswordNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *ResetPasswordBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateUserResponse(response UpdateUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *User:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *UpdateUserBadRequest:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Unauthorized:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Forbidden:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *NotFound:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *Conflict:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

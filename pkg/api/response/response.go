// Package response contains typed structs for various responses to api requests.
// These can be used comfortably to produce swagger api documentation.
package response

import "time"

// HTTPError is the response type for any error that occur for a HTTP request.
type HTTPError struct {
	Error string `json:"error" example:"JWT is expired"`
}

// HTTPLogin is the response for login url request using OAuth service.
type HTTPLogin struct {
	LoginURL string `json:"login_url" example:"https://accounts.google.com/auth/..."`
}

// HTTPAuthorization is the response for login / register request.
type HTTPAuthorization struct {
	Token     string        `json:"token" example:"abcdefghijklmnopqrstuvwxyz1234567890"`
	ExpiresIn time.Duration `json:"expires_in" example:"3600"` // in seconds
	UserID    uint          `json:"user_id" example:"10"`
	UserEmail string        `json:"user_email" example:"myname@example.com"`
	UserName  string        `json:"user_name" example:"Go Pher"`
}

// HTTPRefreshToken is the response when client asks for refreshing the old token.
type HTTPRefreshToken struct {
	Token     string        `json:"token" example:"abcdefghijklmnopqrstuvwxyz1234567890"`
	ExpiresIn time.Duration `json:"expires_in" example:"3600"` // in seconds
}
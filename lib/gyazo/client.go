package gyazo

import (
	"context"
	"errors"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const (
	// APIEndpoint is Gyazo API Endpoint
	APIEndpoint = "https://api.gyazo.com"
	// UploadEndpoint is Gyazo Upload Endpoint
	UploadEndpoint = "https://upload.gyazo.com"
	// UserPath is Gyazo user API Path
	UserPath = "/api/users/me"
	// UploadPath is Gyazo Upload API Path
	UploadPath = "/api/upload"
	// ListPath is Gyazo list API Path
	ListPath = "/api/images"
	// AuthorizePath is Gyazo authorization API Path
	AuthorizePath = "/oauth/authorize"
	// TokenPath is Gyazo Access Token API Path
	TokenPath = "/oauth/token"
	// DeletePathPrefix is Gyazo delete API Path Prefix
	DeletePathPrefix = ListPath + "/"
	// DIDUploadPath is Gyazo upload with device id Path
	DidUploadPath = "/upload.cgi"
)

// func userEndpoint() string                 { return APIEndpoint + UserPath }
// func uploadEndpoint() string               { return UploadEndpoint + UploadPath }
// func listEndpoint() string                 { return APIEndpoint + ListPath }
func authorizeEndpoint() string { return APIEndpoint + AuthorizePath }
func TokenEndpoint() string     { return APIEndpoint + TokenPath }

// func imageEndpoint(imageID string) string  { return APIEndpoint + DeletePathPrefix + imageID }
// func didUploadEndpoint() string            { return UploadEndpoint + DidUploadPath }
// func deleteEndpoint(imageID string) string { return APIEndpoint + DeletePathPrefix + imageID }

func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint(),
			TokenURL: TokenEndpoint(),
		},
		RedirectURL: os.Getenv("CALLBACK_URL"),
	}
	return config
}

func NewClient(token string) (*http.Client, error) {
	if token == "" {
		return nil, errors.New("access token is empty")
	}
	config := GetConnect()
	context := context.Background()
	c := config.Client(context, &oauth2.Token{AccessToken: token})
	return c, nil
}

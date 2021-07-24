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

func UserEndpoint() string      { return APIEndpoint + UserPath }
func ListEndpoint() string      { return APIEndpoint + ListPath }
func AuthorizeEndpoint() string { return APIEndpoint + AuthorizePath }
func TokenEndpoint() string     { return APIEndpoint + TokenPath }

// func UploadEndpoint() string               { return UploadEndpoint + UploadPath }
// func ImageEndpoint(imageID string) string  { return APIEndpoint + DeletePathPrefix + imageID }
// func DidUploadEndpoint() string            { return UploadEndpoint + DidUploadPath }
// func DeleteEndpoint(imageID string) string { return APIEndpoint + DeletePathPrefix + imageID }

func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthorizeEndpoint(),
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

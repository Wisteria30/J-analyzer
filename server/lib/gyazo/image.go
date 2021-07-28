package gyazo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

type Image struct {
	ID           string `json:"image_id"`
	PermalinkURL string `json:"permalink_url"`
	ThumbURL     string `json:"thumb_url"`
	URL          string `json:"url"`
	Type         string `json:"type"`
	Star         bool   `json:"star"`
	CreatedAt    string `json:"created_at"`
}

type ImageDetails struct {
	ID           string      `json:"image_id"`
	PermalinkURL string      `json:"permalink_url"`
	ThumbURL     string      `json:"thumb_url"`
	URL          string      `json:"url"`
	Type         string      `json:"type"`
	CreatedAt    string      `json:"created_at"`
	Metadata     struct {
		App   string      `json:"app"`
		Title string      `json:"title"`
		URL   string      `json:"url"`
		Desc  string      `json:"desc"`
	} `json:"metadata"`
	Ocr struct {
		Locale      string `json:"locale"`
		Description string `json:"description"`
	} `json:"ocr"`
}

// ErrorResponse reports error caused by API request.
type ErrorResponse struct {
	Status  string
	Message string `json:"message"`
}

// Error returns the error response status and message.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v", r.Status, r.Message)
}

// List represents the returned images and http headers from an API request.
type List struct {
	Meta   Meta
	Images *[]Image
}

// Meta represents the returned http headers from an API request.
type Meta struct {
	TotalCount  int
	CurrentPage int
	PerPage     int
	UserType    string
}

// ListOptions specifies the optional parameters to an API request.
type ListOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// List lists the images the specified user.
func GetList(c http.Client, opts *ListOptions) (*List, error) {
	url := ListEndpoint()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new request: %w", err)
	}
	// Build and set query parameters
	if opts != nil {
		params, err := query.Values(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to build query parameters: %w", err)
		}
		req.URL.RawQuery = params.Encode()
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to a get request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, buildErrorResponse(res)
	}

	list := &List{
		Images: new([]Image),
		Meta:   createMeta(res.Header),
	}

	if err = json.NewDecoder(res.Body).Decode(&list.Images); err != nil {
		return nil, fmt.Errorf("failed to decode a responsed JSON: %w", err)
	}

	return list, nil
}

func GetImage(c http.Client, image_id string) (*ImageDetails, error) {
	url := ImageEndpoint(image_id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new request: %w", err)
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to a get request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, buildErrorResponse(res)
	}
	image := &ImageDetails{}

	if err = json.NewDecoder(res.Body).Decode(&image); err != nil {
		return nil, fmt.Errorf("failed to decode a responsed JSON: %w", err)
	}

	return image, nil
}

// buildErrorResponse builds an error information from a HTTP response.
func buildErrorResponse(res *http.Response) error {
	er := &ErrorResponse{Status: res.Status}
	if err := json.NewDecoder(res.Body).Decode(er); err != nil {
		er.Message = err.Error()
	}
	return er
}

// createMeta creates a meta data from a HTTP response headers.
func createMeta(h http.Header) Meta {
	return Meta{
		TotalCount:  atoi(h["X-Total-Count"][0]),
		CurrentPage: atoi(h["X-Current-Page"][0]),
		PerPage:     atoi(h["X-Per-Page"][0]),
		UserType:    h["X-User-Type"][0],
	}
}

// atoi is shorthand for strconv.Atoi(s).
// Golang package docs: https://golang.org/pkg/strconv/#Atoi
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}


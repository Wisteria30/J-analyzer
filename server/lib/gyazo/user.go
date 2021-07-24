package gyazo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	User struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		UID          string `json:"uid"`
		ProfileImage string `json:"profile_image"`
	} `json:"user"`
}

func GetUser(c http.Client) (*Users, error) {
	url := UserEndpoint()
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
	var users *Users
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("failed to decode a responsed JSON: %w", err)
	}

	return users, nil
}
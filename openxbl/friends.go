package openxbl

import (
	"fmt"
	"net/url"
)

// Generated structs from https://mholt.github.io/json-to-go/

// FriendsService provides access to the Friends related functions in the OpenXBL API.
//
// OpenXBL API docs: https://xbl.io/console
type FriendsService apiService

// FriendSeachResp contains the response from the /friends/search operation.
type FriendSeachResp struct {
	ProfileUsers []User `json:"profileUsers,omitempty"`
}

// Search will perform a search of an XBox Gamertag.
func (s *FriendsService) Search(gt string) (User, error) {
	u := fmt.Sprintf("friends/search/?gt=%v", url.QueryEscape(gt))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return User{}, err
	}

	var friendSeachResp *FriendSeachResp
	resp, err := s.client.Do(req, &friendSeachResp)
	if err != nil {
		return User{}, err
	}
	if friendSeachResp == nil {
		return User{}, fmt.Errorf("something went wrong, try again: %s", resp.Status)
	}

	if len(friendSeachResp.ProfileUsers) < 1 {
		return User{}, fmt.Errorf("no user returned")
	}

	return friendSeachResp.ProfileUsers[0], nil
}

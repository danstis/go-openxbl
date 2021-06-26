package openxbl

// User contains the details of a single user.
type User struct {
	ID       string `json:"id"`
	HostID   string `json:"hostId"`
	Settings []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"settings"`
	IsSponsoredUser bool `json:"isSponsoredUser"`
}

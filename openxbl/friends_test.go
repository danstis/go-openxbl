package openxbl

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// JSON minified using https://codebeautify.org/jsonminifier

// TestFriendsService_Search tests the FriendsService Search() method.
func TestFriendsService_Search(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/friends/search/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		switch r.URL.RawQuery {
		case "gt=john":
			fmt.Fprint(w, `{"profileUsers":[{"id":"555123456789","hostId":"555123456789","settings":[{"id":"GameDisplayPicRaw","value":"https://images-eds-ssl.xboxlive.com/image?url=123&mode=Padding&format=png"},{"id":"Gamerscore","value":"987"},{"id":"Gamertag","value":"John"},{"id":"AccountTier","value":"Silver"},{"id":"XboxOneRep","value":"GoodPlayer"},{"id":"PreferredColor","value":"https://dlassets-ssl.xboxlive.com/public/content/ppl/colors/00000.json"},{"id":"RealName","value":""},{"id":"Bio","value":""},{"id":"TenureLevel","value":"0"},{"id":"Watermarks","value":""},{"id":"Location","value":""},{"id":"ShowUserAsAvatar","value":"1"}],"isSponsoredUser":false}]}`)
		case "gt=nonexisting":
			fmt.Fprint(w, `{"code":28,"source":"Profile","description":"The server found no data for the requested entity.","traceInformation":null}`)
		}
	})

	type args struct {
		gt string
	}
	tests := []struct {
		name    string
		s       *FriendsService
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "Known good user",
			s:    client.FriendsService,
			args: args{
				gt: "john",
			},
			want: User{
				ID:     "555123456789",
				HostID: "555123456789",
				Settings: []struct {
					ID    string "json:\"id\""
					Value string "json:\"value\""
				}{
					{ID: "GameDisplayPicRaw", Value: "https://images-eds-ssl.xboxlive.com/image?url=123&mode=Padding&format=png"},
					{ID: "Gamerscore", Value: "987"},
					{ID: "Gamertag", Value: "John"},
					{ID: "AccountTier", Value: "Silver"},
					{ID: "XboxOneRep", Value: "GoodPlayer"},
					{ID: "PreferredColor", Value: "https://dlassets-ssl.xboxlive.com/public/content/ppl/colors/00000.json"},
					{ID: "RealName", Value: ""},
					{ID: "Bio", Value: ""},
					{ID: "TenureLevel", Value: "0"},
					{ID: "Watermarks", Value: ""},
					{ID: "Location", Value: ""},
					{ID: "ShowUserAsAvatar", Value: "1"},
				},
				IsSponsoredUser: false,
			},
			wantErr: false,
		},
		{
			name: "Missing user error",
			s:    client.FriendsService,
			args: args{
				gt: "nonexisting",
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Search(tt.args.gt)
			if (err != nil) != tt.wantErr {
				t.Errorf("FriendsService.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FriendsService.Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

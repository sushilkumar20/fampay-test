package client

import (
	"net/http"
	"testing"
	"time"
)

func TestVideoFetcher_FetchLatestVideo(t *testing.T) {
	type fields struct {
		client  *http.Client
		baseUrl string
		token   string
	}
	type args struct {
		fetchVideoAfter string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "latest video fetcher",
			fields: fields{
				client:  &http.Client{},
				baseUrl: "https://www.googleapis.com/youtube/v3",
				token:   "ya29.a0AeXRPp5_OLnlWxALUWTLgqxdrQ6AXTie5f7QXxD3UWvg_PqTnd3MO4Dc9aE8G7Z7D3LzxqQBb6oTjXdTM9zqxmrcebM9UeXVtDqacpiZvo9o81OPQbO7Ot_PLCLj8OtSRQWft1y-1OHhw7E2JX2nNeq7kT2b1HhJhF0q6rA5aCgYKAaESARESFQHGX2MiWown5MhDMy92B_sTwMPUbQ0175",
			},
			args: args{
				fetchVideoAfter: time.Now().UTC().Format(time.RFC3339),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &YouTubeClient{
				client:  tt.fields.client,
				baseUrl: tt.fields.baseUrl,
				token:   tt.fields.token,
			}
			v.FetchLatestVideo(tt.args.fetchVideoAfter)
		})
	}
}

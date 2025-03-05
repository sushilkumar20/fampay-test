package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type YouTubeClient struct {
	client  *http.Client
	baseUrl string
	token   string
}

func NewYouTubeClient(client *http.Client, baseUrl, token string) *YouTubeClient {
	return &YouTubeClient{client: client, baseUrl: baseUrl, token: token}
}

// TODO:error handling
func (v *YouTubeClient) FetchLatestVideo(fetchVideoAfter string) ([]VideoItem, error) {

	parsedUrl, err := url.Parse(v.baseUrl + "/search")
	if err != nil {
		fmt.Println(err)
	}

	queryParams := parsedUrl.Query()
	queryParams.Set("publishedAfter", fetchVideoAfter)
	queryParams.Set("part", "snippet")
	queryParams.Set("order", "date")
	queryParams.Set("q", "cricket")
	queryParams.Set("maxResults", "5")
	parsedUrl.RawQuery = queryParams.Encode()

	request, err := http.NewRequest("GET", parsedUrl.String(), nil)

	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("Authorization", "Bearer "+v.token)

	response, err := v.client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != http.StatusOK {
		//TODO: pass custom error
		fmt.Println(response.StatusCode)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	videos, err := v.parseYoutubeResponse(body)

	if err != nil {
		fmt.Println("Error parsing response:", err)
		return nil, err
	}
	return videos, nil
}

func (v *YouTubeClient) parseYoutubeResponse(responseBody []byte) ([]VideoItem, error) {
	var ytResponse YouTubeResponse

	err := json.Unmarshal(responseBody, &ytResponse)
	if err != nil {
		return nil, err
	}

	return ytResponse.Items, nil
}

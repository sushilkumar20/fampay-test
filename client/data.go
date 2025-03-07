package client

type YouTubeResponse struct {
	Items []VideoItem `json:"items"`
}

type VideoItem struct {
	ID      VideoID `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type VideoID struct {
	VideoID string `json:"videoId"`
}

type Snippet struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishedAt string     `json:"publishedAt"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

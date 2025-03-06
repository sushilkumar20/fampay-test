package service

type VideoMetaDataListResponse struct {
	Videos []VideoResponseData `json:"videos"`
}

type VideoResponseData struct {
	VideoId          string `json:"videoId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	PublishedAt      string `json:"publishedAt"`
	DefaultThumbnail string `json:"defaultThumbnail"`
}

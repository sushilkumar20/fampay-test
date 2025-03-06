package db

type VideoMetaData struct {
	Id               string `gorm:"primary_key"`
	VideoId          string
	Title            string
	Description      string
	PublishedAt      int64 //no of millisecond
	DefaultThumbnail string
}

type Page struct {
	LastFetchedTime string
	Size            int
}

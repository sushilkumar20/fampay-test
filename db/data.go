package db

type VideoMetaData struct {
	Id               string `gorm:"primary_key"`
	VideoId          string
	Title            string `gorm:"index"`
	Description      string `gorm:"index"`
	PublishedAt      int64  `gorm:"index:idx_published_at,sort:desc"` //no of millisecond
	DefaultThumbnail string
}

type Page struct {
	LastFetchedTime string
	Size            int
}

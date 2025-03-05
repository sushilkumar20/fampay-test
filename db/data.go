package db

type VideoMetaData struct {
	Id               string `gorm:"primary_key"`
	VideoId          string
	Title            string
	Description      string
	PublishedAt      int64
	DefaultThumbnail string
}

type Page struct {
	Limit  int32
	Offset int32
}

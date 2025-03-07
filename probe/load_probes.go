package probe

import (
	"fam/config"
	"fam/db"
	"fam/manager"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func LoadProbe(dataBase *gorm.DB) {
	duration := config.Configuration.ProbeDuration
	youtubeVideoSyncer := YouTubeVideoSyncer{
		DBClinet: db.NewPostgresStore(dataBase),
		YouTubeClient: *manager.NewYouTubeClient(&http.Client{},
			config.Configuration.YoutubeConfig.BaseUrl,
			config.Configuration.YoutubeConfig.ApiKey)}
	runner := NewProbeRunner(duration, nil, &youtubeVideoSyncer)

	err := runner.Start()
	if err != nil {
		log.Fatalf("Failed to start probe, reason: %s ", err)
	}

	log.Println("Probe started")
}

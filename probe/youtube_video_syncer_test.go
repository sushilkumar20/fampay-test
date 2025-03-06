package probe

import (
	"fam/config"
	"fam/db"
	"fam/manager"
	"net/http"
	"testing"
)

var postgresConfig *config.PostgresConfig = &config.PostgresConfig{
	Host:         "localhost",
	UserName:     "postgres",
	Password:     "root",
	DatabaseName: "fam",
	Port:         "5432",
}
var dataBase = db.InitDb(postgresConfig)
var baseUrl = "https://www.googleapis.com/youtube/v3"
var token = "ya29.a0AeXRPp7JGl-Iwg2h6rCCe0xsIUYoVqO0CZXWJ-it0xLMIvPWJgbceRbWw2WsveBZHFqV62yWJLwy_kvLCEkBhLm2VqdaYoOQTZPuc2dgTVADOE5KjMWaYHUCTfd0OyJD4TRza-Tc40eEqYtB_k6kb-WJ-HuJx_KJ_vGKr2F2aCgYKAZQSARESFQHGX2MiDKAR_ldADXDj_ZUSm4__oQ0175"
var youtubeClient = manager.NewYouTubeClient(&http.Client{}, baseUrl, token)

func TestYouTubeVideoSyncer_Execute(t *testing.T) {
	db.LoadSchema(dataBase)
	type fields struct {
		DBClinet      db.DBClinet
		YouTubeClient manager.YouTubeClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "youtube video syncer",
			fields: fields{
				DBClinet:      db.NewPostgresStore(dataBase),
				YouTubeClient: *youtubeClient,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &YouTubeVideoSyncer{
				DBClinet:      tt.fields.DBClinet,
				YouTubeClient: tt.fields.YouTubeClient,
			}
			if err := y.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

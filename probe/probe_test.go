package probe

import (
	"fam/db"
	"testing"
	"time"
)

func TestProbeRunnerImpl_Start(t *testing.T) {
	type fields struct {
		runEvery time.Duration
		quit     chan struct{}
		probe    Probe
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "probe testing",
			fields: fields{
				runEvery: 10 * time.Second,
				quit:     nil,
				probe: &YouTubeVideoSyncer{
					db.NewPostgresStore(dataBase),
					*youtubeClient,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProbeRunnerImpl{
				runEvery: tt.fields.runEvery,
				quit:     tt.fields.quit,
				probe:    tt.fields.probe,
			}
			err := p.Start()
			time.Sleep(time.Duration(120 * time.Second))
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

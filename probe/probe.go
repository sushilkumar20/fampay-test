package probe

import (
	"errors"
	"log"
	"time"
)

type Probe interface {
	Execute() error
}

type ProbeRunner interface {
	Start() error
	Stop()
}

type ProbeRunnerImpl struct {
	runEvery time.Duration
	quit     chan struct{}
	probe    Probe
}

func NewProbeRunner(runEvery time.Duration, quit chan struct{}, probe Probe) *ProbeRunnerImpl {
	return &ProbeRunnerImpl{
		runEvery: runEvery,
		quit:     quit,
		probe:    probe,
	}
}

func (p *ProbeRunnerImpl) Start() error {
	if p.quit != nil {
		return errors.New("probe already started")
	}
	ticker := time.NewTicker(p.runEvery)
	p.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err := p.probe.Execute()
				if err != nil {
					log.Println("Probe call failed", err)
				}
			case <-p.quit:
				ticker.Stop()
				p.quit = nil
				return
			}
		}
	}()

	return nil
}

func (p *ProbeRunnerImpl) Stop() {
	if p.quit == nil {
		close(p.quit)
		p.quit = nil
	}
}

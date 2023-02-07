package app

import (
	"context"
	"log"
	"time"
	"yandex-devops/config"
	"yandex-devops/internal/server/service"
)

type MyFile struct {
	cfg      *config.Server
	services *service.Services
}

func NewMyFile(config *config.Server, services *service.Services) *MyFile {
	return &MyFile{
		cfg:      config,
		services: services,
	}
}

func (f *MyFile) Restore() {
	if !f.cfg.FileRestore {
		return
	}
	if f.services.Fss == nil {
		return
	}

	metrics, err := f.services.Fss.GetAll()
	if err != nil {
		return
	}

	_, err = f.services.Mss.SaveOrUpdateAll(*metrics)
	if err != nil {
		return
	}
}

func (f *MyFile) Start(ctx context.Context) {
	ticker := time.NewTicker(f.cfg.FileInterval)

	for {
		select {
		case <-ticker.C:
			if err := f.recordFile(); err != nil {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (f *MyFile) Finish() {
	if err := f.recordFile(); err != nil {
		log.Println(err)
	}
}

func (f *MyFile) recordFile() error {
	metrics, err := f.services.Mss.GetAll()

	if err != nil {
		return err
	} else {
		_, err := f.services.Fss.SetAll(*metrics)
		if err != nil {
			return err
		}
	}
	return nil
}

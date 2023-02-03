package app

import (
	"context"
	"time"
	"yandex-devops/config"
	"yandex-devops/storage/file"
	"yandex-devops/storage/memory"
)

type MyFile struct {
	ctx         context.Context
	cfg         *config.Server
	fileStorage *file.FileStorage
	memStorage  *memory.MemStorage
}

func NewMyFile(ctx context.Context, config *config.Server, memStorage *memory.MemStorage, fileStorage *file.FileStorage) *MyFile {
	return &MyFile{
		ctx:         ctx,
		cfg:         config,
		fileStorage: fileStorage,
		memStorage:  memStorage,
	}
}

func (f *MyFile) Restore() {
	if !f.cfg.FileRestore {
		return
	}
	if f.fileStorage == nil {
		return
	}

	metrics, err := f.fileStorage.GetAll()
	if err != nil {
		return
	}

	_, err = f.memStorage.SetAll(*metrics)
	if err != nil {
		return
	}
}

func (f *MyFile) Start() {
	ticker := time.NewTicker(f.cfg.FileInterval)

	for {
		select {
		case <-ticker.C:
			metrics, err := f.memStorage.GetAll()
			if err != nil {
				continue
			} else {
				_, err := f.fileStorage.SetAll(*metrics)
				if err != nil {
					continue
				}
			}
		case <-f.ctx.Done():
			return
		}
	}
}

func (f *MyFile) Finish() {

}

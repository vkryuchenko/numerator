package main

import (
	"flag"
	"github.com/labstack/gommon/log"
	"handlers"
	"storage/filesystem"
	"time"
)

func main() {
	masterToken := flag.String("token", "436a5a1c810f4edbe8a4728621efd7e8", "MASTER_TOKEN for delete and manual update update values")
	dumpPath := flag.String("dump", "/data/dump.json", "data dump path")
	flag.Parse()

	storage, err := filesystem.New(*dumpPath)
	if err != nil {
		log.Warn(err)
	}
	go persistFilesystem(storage)
	provider := handlers.Provider{MasterToken: *masterToken, Storage: storage}
	log.Info(*masterToken)
	log.Fatal(provider.Serve())
}

func persistFilesystem(fs *filesystem.FSStorage) {
	throttling := time.Tick(5 * time.Second)
	for {
		<-throttling
		err := fs.SaveData()
		if err != nil {
			log.Warn(err)
		}
	}
}

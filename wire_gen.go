// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/cloudsonic/sonic-server/engine"
	"github.com/cloudsonic/sonic-server/itunesbridge"
	"github.com/cloudsonic/sonic-server/persistence/ledis"
	"github.com/cloudsonic/sonic-server/scanner"
	"github.com/deluan/gomate"
	ledis2 "github.com/deluan/gomate/ledis"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func initImporter(musicFolder string) *scanner.Importer {
	checkSumRepository := ledis.NewCheckSumRepository()
	itunesScanner := scanner.NewItunesScanner(checkSumRepository)
	mediaFileRepository := ledis.NewMediaFileRepository()
	albumRepository := ledis.NewAlbumRepository()
	artistRepository := ledis.NewArtistRepository()
	artistIndexRepository := ledis.NewArtistIndexRepository()
	playlistRepository := ledis.NewPlaylistRepository()
	propertyRepository := ledis.NewPropertyRepository()
	db := newDB()
	search := engine.NewSearch(artistRepository, albumRepository, mediaFileRepository, db)
	importer := scanner.NewImporter(musicFolder, itunesScanner, mediaFileRepository, albumRepository, artistRepository, artistIndexRepository, playlistRepository, propertyRepository, search)
	return importer
}

// wire_injectors.go:

var allProviders = wire.NewSet(itunesbridge.NewItunesControl, ledis.Set, engine.Set, scanner.Set, newDB)

func newDB() gomate.DB {
	return ledis2.NewEmbeddedDB(ledis.Db())
}
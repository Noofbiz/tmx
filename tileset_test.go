package tmx

import (
	"os"
	"testing"
)

func TestTMXURL(t *testing.T) {
	TMXURL = "testData"
	if TMXURL != "testData" {
		t.Errorf("TMXURL was not updated to 'testData' after changing it.")
	}
}

func TestTilesetLoading(t *testing.T) {
	TMXURL = "testData/tilesheetTest.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse %v. Error was: %v", TMXURL, err)
		return
	}
	if m.Tilesets[0].Image[0].Source != "roguelikeIndoor_transparent.png" {
		t.Errorf("Image not properly parsed from embedded tileset")
		return
	}
	if m.Tilesets[1].Image[0].Source != "roguelikeHoliday_transparent.png" {
		t.Errorf("Image not properly parsed from external tileset")
		return
	}
}

func TestTilesetPreLoading(t *testing.T) {
	preloaded := "testData/preloaded.tsx"
	PreloadedTilesets[preloaded] = preloadedTsx

	TMXURL = "testData/tilesheetTestPreloaded.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse %v. Error was: %v", TMXURL, err)
		return
	}
	if m.Tilesets[0].Image[0].Source != "roguelikeIndoor_transparent.png" {
		t.Errorf("Image not properly parsed from embedded tileset")
		return
	}
	if m.Tilesets[1].Image[0].Source != "roguelikeHoliday_transparent.png" {
		t.Errorf("Image not properly parsed from external tileset")
		return
	}
}

func TestTilesetTSXNotExist(t *testing.T) {
	TMXURL = "testData/tsxNotExist.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the tsx does not exist", TMXURL)
		return
	}
}

func TestTilesetTSXMalformed(t *testing.T) {
	TMXURL = "testData/tsxMalformed.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the tsx file is not a valid tsx file", TMXURL)
	}
}

func TestMalformedTileset(t *testing.T) {
	TMXURL = "testData/malformedTilesheet.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the tileset was not valid", TMXURL)
	}
}

var preloadedTsx = []byte(`
	<?xml version="1.0" encoding="UTF-8"?>
		<tileset name="external" tilewidth="16" tileheight="16" spacing="1" tilecount="48" columns="12">
		<image source="roguelikeHoliday_transparent.png" width="203" height="67"/>
		</tileset>
`)

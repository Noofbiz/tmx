package tmx

import (
	"os"
	"testing"
)

func TestDataMalformed(t *testing.T) {
	TMXURL = "testData/malformedData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the data was not valid", TMXURL)
	}
}

var testDataExpected = []uint32{
	235, 236, 237,
	247, 356, 282,
	323, 324, 273,
}

func TestDataCSV(t *testing.T) {
	TMXURL = "testData/csvData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse CSV encoded data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Errorf("Decoded CSV data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
			return
		}
	}
}

func TestDataMalformedCSV(t *testing.T) {
	TMXURL = "testData/malformedCSVData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the data was not valid", TMXURL)
	}
}

func TestDataTiles(t *testing.T) {
	TMXURL = "testData/tileData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse CSV encoded data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Errorf("Tile data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
			return
		}
	}
}

func TestDataEmptyCSV(t *testing.T) {
	TMXURL = "testData/csvEmptyData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v when the data was empty", TMXURL)
	}
}

func TestDataUnknownEncoding(t *testing.T) {
	TMXURL = "testData/unknownEncodingData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v data without proper encoding", TMXURL)
	}
	if err.Error() != "Unknown Encoding" {
		t.Errorf("Error recieved trying to parse %v was incorrect. \n Wanted: %v\nGot: %v\n", TMXURL, "Unknown Encoding", err.Error())
	}
}

func TestDataBase64(t *testing.T) {
	TMXURL = "testData/base64Data.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse base64 encoded data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Errorf("Tile data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
			return
		}
	}
}

func TestDataMalformedBase64(t *testing.T) {
	TMXURL = "testData/malformedBase64Data.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v data with bad base64 encoding", TMXURL)
	}
}

func TestDataUnknownCompression(t *testing.T) {
	TMXURL = "testData/unknownCompressionData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v data without proper encoding", TMXURL)
	}
	if err.Error() != "Unknown Compression" {
		t.Errorf("Error recieved trying to parse %v was incorrect. \n Wanted: %v\nGot: %v\n", TMXURL, "Unknown Encoding", err.Error())
	}
}

func TestDataZlib(t *testing.T) {
	TMXURL = "testData/zlibData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse zilb compressed data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Errorf("Decoded ZLib data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
			return
		}
	}
}

func TestDataMalformedZlib(t *testing.T) {
	TMXURL = "testData/malformedZlibData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v data with bad zlib compression", TMXURL)
		return
	}
}

func TestDataGZip(t *testing.T) {
	TMXURL = "testData/gzipData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Errorf("Unable to parse gzip compressed data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Errorf("Decoded GZip data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
			return
		}
	}
}

func TestDataMalformedGZip(t *testing.T) {
	TMXURL = "testData/malformedGZipData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Errorf("Able to parse %v data with bad gzip compression", TMXURL)
		return
	}
}

func TestDataFlipped(t *testing.T) {
	TMXURL = "testData/flipData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Error("unable to parse flip data")
		return
	}
	exp := []struct {
		GID      uint32
		Flipping uint32
	}{
		{235, VerticalFlipFlag | DiagonalFlipFlag},
		{236, HorizontalFlipFlag | DiagonalFlipFlag},
		{237, HorizontalFlipFlag},
		{247, VerticalFlipFlag},
	}
	for i, e := range exp {
		if m.Layers[0].Data[0].Tiles[i].GID != e.GID {
			t.Errorf("Flipped tile data does not match GIDs\nWanted: %v\nGot: %v", e.GID, m.Layers[0].Data[0].Tiles[i].GID)
			return
		}
		if m.Layers[0].Data[0].Tiles[i].Flipping != e.Flipping {
			t.Errorf("Flipped tile data does not match expected flip state\nWanted: %v\nGot: %v", e.Flipping, m.Layers[0].Data[0].Tiles[i].Flipping)
			return
		}
	}
}

func TestDataChunks(t *testing.T) {
	TMXURL = "testData/chunkData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	m, err := Parse(f)
	if err != nil {
		t.Error("Unable to parse chunk data")
		return
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Chunks[0].Tiles[i].GID != e {
			t.Errorf("Test data did not match expected data\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Chunks[0].Tiles[i].GID)
			return
		}
	}
}

func TestDataMalformedChunks(t *testing.T) {
	TMXURL = "testData/malformedChunkData.tmx"
	f, err := os.Open(TMXURL)
	if err != nil {
		t.Errorf("Unable to open %v. Error was: %v", TMXURL, err)
		return
	}
	defer f.Close()
	_, err = Parse(f)
	if err == nil {
		t.Error("Able to parse malformed chunk data")
		return
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"gopkg.in/yaml.v3"
)

type Track struct {
	Name      string
	AlbumName string
	Artist    string
}

var emptyTrack = Track{"", "", ""}

var tracks = []Track{
	{"foo", "album1", "artist1"},
	{"bar", "album1", "artist1"},
	{"foo", "album2", "artist1"},
	{"baz", "album2", "artist2"},
	{"baz", "album3", "artist2"},
}

func rejectEmpty(tracks []Track) []Track {
	filteredTracks := make([]Track, 0, len(tracks))
	for _, track := range tracks {
		if track.Name != "" {
			filteredTracks = append(filteredTracks, track)
		}
	}
	return filteredTracks
}

func writeTracksToFile(tracks []Track) error {
	yaml, err := yaml.Marshal(rejectEmpty(tracks))
	if err != nil {
		return err
	}
	if err := os.WriteFile("albums.yaml", yaml, 0644); err != nil {
		log.Fatalf("Error writing tracks to file: %v", err)
		return err
	}
	return nil
}

func main() {

	yamlData, err := os.ReadFile("albums.yaml")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error reading albums.yaml: %v", err)
	}

	var existingTracks []Track
	if len(yamlData) > 0 {
		if err := yaml.Unmarshal(yamlData, &existingTracks); err != nil {
			log.Fatalf("Error parsing albums.yaml: %v", err)
		}
	}

	existingTracks = append(existingTracks, emptyTrack)
	existingTracksLength := len(existingTracks)

	idx, err := fuzzyfinder.FindMulti(
		tracks,
		func(i int) string {
			return tracks[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				existingTracks[existingTracksLength-1] = emptyTrack
				writeTracksToFile(existingTracks)
				return ""
			}
			track := tracks[i]

			existingTracks[existingTracksLength-1] = track
			writeTracksToFile(existingTracks)

			return fmt.Sprintf("Track: %s (%s)\nAlbum: %s",
				tracks[i].Name,
				tracks[i].Artist,
				tracks[i].AlbumName)
		}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("selected: %v\n", idx)
}

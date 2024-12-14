package main

import (
	"bytes"
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

func writeTracksToFile(tracks []Track, content string) error {
	yaml, err := yaml.Marshal(rejectEmpty(tracks))
	if err != nil {
		return err
	}

	content = fmt.Sprintf("---\n%s\n---\n%s", yaml, content)
	if err := os.WriteFile("albums.md", []byte(content), 0644); err != nil {
		log.Fatalf("Error writing tracks to file: %v", err)
		return err
	}
	return nil
}

func readTracksFromMarkdown(mdData []byte) ([]Track, error) {
	// Split frontmatter from content
	parts := bytes.Split(mdData, []byte("---\n"))
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown file format - missing frontmatter")
	}

	// Parse YAML frontmatter
	var mdTracks []Track
	if err := yaml.Unmarshal(parts[1], &mdTracks); err != nil {
		return nil, fmt.Errorf("error parsing frontmatter: %v", err)
	}

	return mdTracks, nil
}

func readContentFromMarkdown(mdData []byte) (string, error) {
	parts := bytes.Split(mdData, []byte("---\n"))
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid markdown file format - missing frontmatter")
	}
	return string(parts[2]), nil
}

func main() {
	mdData, err := os.ReadFile("albums.md")
	if err != nil {
		log.Fatal(err)
	}

	existingTracks, err := readTracksFromMarkdown(mdData)
	if err != nil {
		log.Fatal(err)
	}

	content, err := readContentFromMarkdown(mdData)
	if err != nil {
		log.Fatal(err)
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
				writeTracksToFile(existingTracks, content)
				return ""
			}
			track := tracks[i]

			existingTracks[existingTracksLength-1] = track
			writeTracksToFile(existingTracks, content)

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

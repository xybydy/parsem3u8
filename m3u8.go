package parsem3u8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var linePattern = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

type MediaType uint

const (
	Stream = iota + 1
	Vod
)

var videoTypes = []string{
	"3g2",
	"3gp",
	"aaf",
	"asf",
	"avchd",
	"avi",
	"drc",
	"flv",
	"m2v",
	"m4p",
	"m4v",
	"mkv",
	"mng",
	"mov",
	"mp2",
	"mp4",
	"mpe",
	"mpeg",
	"mpg",
	"mpv",
	"mxf",
	"nsv",
	"ogg",
	"ogv",
	"qt",
	"rm",
	"rmvb",
	"roq",
	"svi",
	"vob",
	"webm",
	"wmv",
	"yuv",
}

type Segments []Segment
type Segment struct {
	Name     string     `json:"name"`
	Duration float64    `json:"duration"`
	Ctags    customTags `json:"tags"`
	URL      string     `json:"url"`
	Type     MediaType  `json:"-"`
}

type customTags []customTag
type customTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c customTags) Get(key string) string {
	for _, tag := range c {
		if tag.Key == strings.ToLower(key) {
			return tag.Value
		}
	}

	return ""
}

func (c *customTags) Set(key, value string) error {
	for _, tag := range *c {
		if tag.Key == strings.ToLower(key) {
			tag.Value = value
			return nil
		}
	}

	return errors.New("no such key")
}

func ParseFile(filepath string) Segments {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return Parse(f)
}

func Parse(reader io.Reader) Segments {
	var (
		extInf   bool
		segments Segments
		segment  Segment
		lines    []string
	)

	s := bufio.NewScanner(reader)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if i == 0 && line != "#EXTM3U" {
			log.Panic("not a proper m3u file.")
		}

		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "#EXTINF"):
			var duration string

			if extInf {
				log.Fatalf("duplicate EXTINF: %s, line: %d", line, i+1)
			}

			_, err := fmt.Sscanf(line, "#EXTINF:%s", &duration)
			if err != nil {
				log.Fatal(err)
			}

			segment.Duration, err = strconv.ParseFloat(duration, 64)
			if err != nil {
				log.Fatal(err)
			}

			splitName := strings.Split(line, ",")
			segment.Name = splitName[len(splitName)-1]
			segment.Ctags = parseLineParameters(line)

			extInf = true
		case !strings.HasPrefix(line, "#"):
			segment.URL = line
			extInf = false

			lineType := getType(line)
			segment.Type = lineType

			segments = append(segments, segment)
		}
	}

	return segments
}

func parseLineParameters(line string) customTags {
	r := linePattern.FindAllStringSubmatch(line, -1)
	tags := make(customTags, 0)

	for _, arr := range r {
		tag := customTag{
			Key:   strings.ToLower(arr[1]),
			Value: strings.Trim(arr[2], "\""),
		}
		tags = append(tags, tag)
	}

	return tags
}

func getType(text string) MediaType {
	ext := strings.Trim(filepath.Ext(text), ".")

	for i := 0; i < len(videoTypes); i++ {
		if ext == videoTypes[i] {
			return Vod
		}
	}

	return Stream
}

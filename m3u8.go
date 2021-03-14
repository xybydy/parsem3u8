package parsem3u8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var linePattern = regexp.MustCompile(`([a-zA-Z-]+)=("[^"]+"|[^",]+)`)

type Segments []Segment
type Segment struct {
	Name     string     `json:"name"`
	Duration float64    `json:"duration"`
	Ctags    customTags `json:"tags"`
	URL      string     `json:"url"`
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

			segment.Name = strings.Split(line, ",")[1]
			segment.Ctags = parseLineParameters(line)

			extInf = true
		case !strings.HasPrefix(line, "#"):
			segment.URL = line
			extInf = false

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

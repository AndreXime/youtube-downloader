package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var OutputDir = "./downloads"

// Gerencia as 9 linhas do terminal
func MonitorProgress(stream <-chan string) {
	const maxLines = 9
	buffer := []string{}
	lastHeight := 0

	for line := range stream {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}

		buffer = append(buffer, text)
		if len(buffer) > maxLines {
			buffer = buffer[len(buffer)-maxLines:]
		}

		if lastHeight > 0 {
			fmt.Printf("\033[%dA\033[J", lastHeight)
		}

		for _, l := range buffer {
			if len(l) > 100 {
				l = l[:97] + "..."
			}
			fmt.Println(l)
		}
		lastHeight = len(buffer)
	}
}

func CheckDeps() error {
	for _, bin := range []string{"yt-dlp", "node", "ffmpeg"} {
		if _, err := exec.LookPath(bin); err != nil {
			return fmt.Errorf("Dependência ausente: %s", bin)
		}
	}
	return nil
}

func Run(stream chan<- string) error {
	defer close(stream)
	if AppState.Folder != "" {
		OutputDir += "/" + AppState.Folder
	}
	os.MkdirAll(OutputDir, os.ModePerm)

	// Monta argumentos
	args := []string{"--no-warnings"}
	if AppState.Playlist {
		args = append(args, "--yes-playlist")
	} else {
		args = append(args, "--no-playlist")
	}

	if AppState.Format == "mp3" {
		args = append(args, "-x", "--audio-format", "mp3", "--audio-quality", "0")
	} else {
		args = append(args, "-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best")
	}

	args = append(args, "--add-metadata", "--embed-thumbnail", "--js-runtime", "node",
		"--extractor-args", "youtube:player_client=ios,web", "-o", fmt.Sprintf("%s/%%(title)s.%%(ext)s", OutputDir))
	args = append(args, AppState.Link)

	// Executa
	cmd := exec.Command("yt-dlp", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	// Leitura com Split customizado para tratar \r
	scanner := bufio.NewScanner(stdout)
	scanner.Split(func(d []byte, eof bool) (int, []byte, error) {
		if i := bytes.IndexAny(d, "\r\n"); i >= 0 {
			return i + 1, d[0:i], nil
		}
		if eof && len(d) > 0 {
			return len(d), d, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		stream <- scanner.Text()
	}

	return cmd.Wait()
}

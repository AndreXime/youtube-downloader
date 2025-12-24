package main

import (
	"fmt"
	"os"
	"strings"
)

// Helper para formatar o tamanho dos arquivos
func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

// Obtém nomes dos arquivos e o tamanho total da pasta
func getFolderStats() ([]string, string, int) {
	files, err := os.ReadDir(AppState.OutputDir)
	if err != nil {
		return nil, "0 B", 0
	}

	var names []string
	var totalSize int64
	count := 0

	for _, f := range files {
		if !f.IsDir() {
			info, err := f.Info()
			if err == nil {
				count++
				totalSize += info.Size()
				names = append(names, fmt.Sprintf("%-40s (%s)", f.Name(), formatBytes(info.Size())))
			}
		}
	}
	return names, formatBytes(totalSize), count
}

// MonitorProgress agora vive aqui e gerencia as duas visões
func MonitorProgress(stream <-chan string) {
	const maxLines = 9
	buffer := []string{}
	lastHeight := 0

	fileList, totalSize, fileCount := getFolderStats()

	fmt.Println(Inf.Render("--- Logs do yt-dlp --- "))

	for line := range stream {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}

		// Gatilhos extras para atualizar a lista de arquivos
		lowerText := strings.ToLower(text)
		isUpdate := strings.Contains(lowerText, "sleeping") ||
			strings.Contains(lowerText, "metadata") ||
			strings.Contains(lowerText, "destination:")

		if isUpdate {
			fileList, totalSize, fileCount = getFolderStats()
		}

		// Gerencia buffer de logs
		buffer = append(buffer, text)
		if len(buffer) > maxLines {
			buffer = buffer[len(buffer)-maxLines:]
		}

		// Limpa apenas as linhas impressas anteriormente
		if lastHeight > 0 {
			fmt.Printf("\033[%dA\033[J", lastHeight)
		}

		// 1. Renderiza Logs
		for _, l := range buffer {
			if len(l) > 100 {
				l = l[:97] + "..."
			}
			fmt.Println(l)
		}

		// 2. Renderiza Status da Pasta (se houver arquivos)
		currentHeight := len(buffer)
		if len(fileList) > 0 {
			header := fmt.Sprintf("\n--- Arquivos na Pasta (%d) --- ", fileCount)
			fmt.Println(Inf.Render(header) + Suc.Render("Total: "+totalSize))

			for _, f := range fileList {
				fmt.Printf("✅ %s\n", f)
			}
			currentHeight += len(fileList) + 2
		}

		lastHeight = currentHeight
	}
}

package main

import (
	"fmt"
)

var AppState struct {
	Link      string
	Format    string
	Playlist  bool
	OutputDir string
}

func main() {
	fmt.Println(Title.Render("YouTube Downloader CLI"))

	if err := CheckDeps(); err != nil {
		fmt.Println(Err.Render(err.Error()))
		return
	}

	if err := GetInputs(); err != nil {
		fmt.Println(Err.Render(err.Error()))
		return
	}

	progressCh := make(chan string)
	errCh := make(chan error)

	// Roda download em background
	go func() {
		errCh <- Run(progressCh)
	}()

	// Trava a main thread mostrando a UI até acabar
	MonitorProgress(progressCh)

	if err := <-errCh; err != nil {
		fmt.Println(Err.Render("Erro: " + err.Error()))
	} else {
		fmt.Println(Suc.Render("Sucesso!"))
	}
}

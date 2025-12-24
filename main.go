package main

import (
	"fmt"
)

var AppState struct {
	Link     string
	Format   string
	Playlist bool
	Folder   string
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

	fmt.Println(Inf.Render("Iniciando... (ctrl+z para encerrar)"))

	// Prepara canais
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

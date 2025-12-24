package main

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title = lipgloss.NewStyle().Background(lipgloss.Color("0")).Foreground(lipgloss.Color("15")).Bold(true)
	Err   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	Suc   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
	Inf   = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
)

func GetInputs() error {
	customTheme := huh.ThemeCharm()

	customTheme.Focused.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true)
	customTheme.Blurred.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	customTheme.Blurred.SelectedOption = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))

	var folder string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Link do Vídeo ou Playlist").
				Value(&AppState.Link).
				Validate(func(str string) error {
					if !strings.HasPrefix(str, "http") {
						return errors.New("link inválido!")
					}
					return nil
				}),

			huh.NewSelect[string]().
				Title("Formato desejado").
				Options(
					huh.NewOption("MP3 (Áudio)", "mp3"),
					huh.NewOption("MP4 (Vídeo)", "mp4"),
				).
				Value(&AppState.Format),

			huh.NewConfirm().
				Title("Baixar playlist completa se o link pertencer a uma?").
				Affirmative("Sim").
				Negative("Não").
				Value(&AppState.Playlist),

			huh.NewInput().
				Title("Digite o nome da subpasta para esse download, deixe vazio para ser na raiz").
				Value(&folder),
		),
	).WithTheme(customTheme)

	err := form.Run()
	if err != nil {
		return err
	}

	AppState.OutputDir = "./downloads"
	if folder != "" {
		AppState.OutputDir += "/" + folder
	}

	return nil
}

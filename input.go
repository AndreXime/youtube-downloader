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

func GetInputs() (string, string, bool, bool) {
	var link string
	var format string
	var playlist bool

	customTheme := huh.ThemeCharm() // Baseado no tema padrão, mas vamos tunar:
	
	// Cor da pergunta quando você está nela (Focused)
	customTheme.Focused.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true)
	
	// Cor da pergunta após você já ter respondido (Blurred/Compelted)
	// Isso dá o feedback de "aceito" que você quer
	customTheme.Blurred.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Cinza escuro
	customTheme.Blurred.SelectedOption = lipgloss.NewStyle().Foreground(lipgloss.Color("42")) // Verde para a opção escolhida

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Link do Vídeo ou Playlist").
				Value(&link).
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
				Value(&format),

			huh.NewConfirm().
				Title("Baixar playlist completa se o link pertencer a uma?").
				Affirmative("Sim").
				Negative("Não").
				Value(&playlist),
		),
	).WithTheme(customTheme) // Inline mantém o histórico

	err := form.Run()

	if err != nil {
		return "", "", false, false
	}

	return link, format, playlist, true
}


# YouTube Downloader CLI

CLI interativa em Go para baixar vídeos e músicas do YouTube. Interface no terminal com [Bubble Tea](https://github.com/charmbracelet/bubbletea) e [huh](https://github.com/charmbracelet/huh); downloads e conversão via **yt-dlp** e **FFmpeg**.

## Funcionalidades

- Escolha entre **áudio (MP3)** ou **vídeo (MP4)**
- Download de **vídeo individual** ou **playlist completa**
- Fluxo guiado por menus no terminal

## Stack

| Camada | Tecnologia |
|--------|------------|
| Linguagem | Go |
| Runtime | Go 1.24+ |
| TUI | Bubble Tea, huh, lipgloss |
| Download | yt-dlp |
| Conversão | FFmpeg |

## Pré-requisitos

Antes de usar a ferramenta, instale em sua máquina:

1. **yt-dlp** — motor de download
2. **FFmpeg** — conversão de mídia
3. **Node.js** — usado pelo extrator do yt-dlp em alguns fluxos

Verifique se os binários estão no `PATH`:

```bash
yt-dlp --version
ffmpeg -version
```

## Instalação

### Opção 1: Binário pré-compilado (recomendado)

1. Acesse a [página de Releases](https://github.com/AndreXime/youtube-downloader/releases) e baixe o binário para o seu sistema operacional.
2. Dê permissão de execução (Linux/macOS):

```bash
chmod +x youtube-downloader
```

3. Execute:

```bash
./youtube-downloader
```

### Opção 2: Compilar do código-fonte

```bash
git clone https://github.com/AndreXime/youtube-downloader.git
cd youtube-downloader
go build .
./youtube-downloader
```

## Estrutura do projeto

```text
youtube-downloader/
├── main.go      # entrada da aplicação
├── engine.go    # integração com yt-dlp
├── input.go     # formulários interativos (huh)
├── view.go      # renderização TUI
└── go.mod
```

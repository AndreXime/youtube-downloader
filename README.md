# YouTube Downloader CLI

Uma ferramenta de linha de comando moderna, rápida para baixar vídeos e músicas do YouTube.

## Funcionalidades

* Escolha entre Áudio (MP3) ou Vídeo (MP4).
* Opção para baixar vídeos individuais ou playlists completas.

## Pré-requisitos

Antes de rodar o script, você precisa ter instalado em sua máquina:

1. **yt-dlp** (Motor de download)
2. **FFmpeg** (Necessário para conversão de mídia)
3. **Node.js** (Utilizado pelo extrator do yt-dlp)

## Como usar

1. Clone o repositório:
```bash
git clone https://github.com/AndreXime/youtube-downloader.git
cd youtube-downloader
```

2. Faça build:
```bash
go build .
```

3. Executa o binario:

```bash
./youtube-downloader
```
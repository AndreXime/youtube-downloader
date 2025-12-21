# YouTube Downloader CLI

Uma ferramenta de linha de comando moderna, rápida para baixar vídeos e músicas do YouTube.

## Funcionalidades

* Adicione quantos links quiser antes de iniciar o processamento.
* Escolha entre Áudio (MP3) ou Vídeo (MP4).
* Opção para baixar vídeos individuais ou playlists completas.

## Pré-requisitos

Antes de rodar o script, você precisa ter instalado em sua máquina:

1. **[Bun](https://bun.sh/)** (Runtime principal)
2. **[yt-dlp](https://github.com/yt-dlp/yt-dlp)** (Motor de download)
3. **[FFmpeg](https://ffmpeg.org/)** (Necessário para conversão de mídia)
4. **Node.js** (Exigido pelo extrator do yt-dlp)

## Instalação

1. Clone o repositório:
```bash
git clone https://github.com/AndreXime/youtube-downloader.git
cd youtube-downloader
```


2. Instale as dependências:
```bash
bun install
```


## Como usar

Para iniciar basta executar:

```bash
bun start
```
import { existsSync, mkdirSync } from "node:fs";
import logUpdate from "log-update";
import { intro, text, isCancel, confirm, log, select } from "@clack/prompts";

const outputDir = "./downloads";

async function download(
	link: string[],
	downloadPlaylist: boolean,
	format: string,
) {
	if (!existsSync(outputDir)) {
		mkdirSync(outputDir, { recursive: true });
	}

	const args = [
		"yt-dlp",
		"--no-warnings",
		downloadPlaylist ? "--yes-playlist" : "--no-playlist",
		...(format === "mp3"
			? ["-x", "--audio-format", "mp3", "--audio-quality", "0"]
			: ["-f", "bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best"]),
		"--add-metadata",
		"--embed-thumbnail",
		"--js-runtime",
		"node",
		"--extractor-args",
		"youtube:player_client=ios,web",
		"-o",
		`${outputDir}/%(title)s.%(ext)s`,
		...link,
	];

	const proc = Bun.spawn(args, { stdout: "pipe" });
	const decoder = new TextDecoder();
	const logBuffer: string[] = [];
	const MAX_LINES = 9;

	logUpdate("\n".repeat(MAX_LINES - 1));

	for await (const chunk of proc.stdout) {
		const lines = decoder.decode(chunk).split("\n").filter(Boolean);

		for (const line of lines) {
			logBuffer.push(line.replace(/\r/g, ""));
			if (logBuffer.length > MAX_LINES) logBuffer.shift();
			logUpdate(logBuffer.join("\n"));
		}
	}

	logUpdate.done();
	await proc.exited;
}

async function checkDependencies() {
	const deps = [
		{ name: "yt-dlp", command: ["yt-dlp", "--version"] },
		{ name: "Node.js", command: ["node", "--version"] },
		{ name: "ffmpeg", command: ["ffmpeg", "-version"] },
	];

	for (const dep of deps) {
		try {
			const proc = Bun.spawnSync(dep.command);
			if (proc.exitCode !== 0) throw new Error();
		} catch {
			log.error(`Dependência ausente: ${dep.name}`);
			log.info(`Certifique-se de que o ${dep.name} está instalado.`);
			process.exit(1);
		}
	}
}

async function main() {
	intro("YouTube Downloader CLI");

	const links: string[] = [];

	while (true) {
		const input = await text({
			message: `Cole o link #${links.length + 1} (ou deixe vazio para finalizar)`,
			placeholder: "https://youtube.com/...",
			validate(value) {
				if (value.length > 0 && !value.startsWith("http"))
					return "Link inválido!";
			},
		});

		if (isCancel(input) || !input) break;
		links.push(input);
	}

	if (links.length === 0) {
		log.error("Nenhum link adicionado.");
		return;
	}

	const format = await select({
		message: "O que você deseja baixar?",
		options: [
			{
				value: "mp3",
				label: "Apenas Áudio (MP3)",
			},
			{
				value: "mp4",
				label: "Vídeo (MP4)",
			},
		],
	});

	if (isCancel(format)) return;

	const isPlaylist = await confirm({
		message: "Se houver playlists nos links, baixar todas as músicas/videos?",
		initialValue: false,
	});

	if (isCancel(isPlaylist)) return;

	checkDependencies();

	log.info(`Fazendo download na pasta ${outputDir}...`);

	await download(links, isPlaylist, format);

	log.success("Processo concluído");
}

await main();

import type { OS } from "./ostype"

export function getOS() {
	const q = new URLSearchParams(window.location.search)
	let currentos = (q.get("os") || "").toLowerCase() as OS
	if (currentos == ("" as OS)) {
		if (navigator.userAgent.includes("Mac")) currentos = "macos"
		else if (navigator.userAgent.includes("Windows")) currentos = "windows"
		else if (navigator.userAgent.includes("Linux")) currentos = "linux"
	}
	return currentos
}

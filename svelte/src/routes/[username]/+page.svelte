<script lang="ts">
	import type { Dotfile } from "$lib/server/db"
	import type { OS } from "$lib/ostype"
	import type { PageServerData } from "./$types"
	import OSSelector from "$lib/components/OSSelector.svelte"
	import { onMount } from "svelte"
	import Navbar from "$lib/components/Navbar.svelte"

	function getFilename(df: Dotfile) {
		const fn = (path: string) => path.split("/").at(-1) as string
		let fileNames: string[] = []
		if (df.paths.linux) fileNames.push(fn(df.paths.linux))
		if (df.paths.windows) fileNames.push(fn(df.paths.windows))
		if (df.paths.macos) fileNames.push(fn(df.paths.macos))

		let counts: { [key: string]: number } = {}
		for (let i of fileNames) {
			if (!counts[i]) counts[i] = 0
			counts[i]++
		}
		let max = Math.max(...Object.values(counts))
		return Object.keys(counts).filter((key) => counts[key] == max)
	}

	export let data: PageServerData

	let currentos: OS
	onMount(() => {
		const q = new URLSearchParams(window.location.search)
		currentos = (q.get("os") || "").toLowerCase() as OS
		if (currentos == ("" as OS)) {
			if (navigator.userAgent.includes("Mac")) currentos = "macos"
			else if (navigator.userAgent.includes("Windows")) currentos = "windows"
			else if (navigator.userAgent.includes("Linux")) currentos = "linux"
		}
	})
</script>

<Navbar depotname={data.depotname}>
	<OSSelector {currentos} slot="right-items" />
</Navbar>

{#if currentos}
	{#each data.docs as doc}
		{getFilename(doc)}<br />
	{/each}
{/if}

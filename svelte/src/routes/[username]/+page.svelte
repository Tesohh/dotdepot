<script lang="ts">
	import type { Dotfile } from "$lib/server/db"
	import type { OS } from "$lib/ostype"
	import type { PageServerData } from "./$types"
	import OSSelector from "$lib/components/OSSelector.svelte"
	import { onMount } from "svelte"
	import Navbar from "$lib/components/Navbar.svelte"
	import { treeifyPaths, type PathTree, type PathContexts } from "treeify-paths"
	import type { WithId } from "mongodb"
	import FileTree from "$lib/components/FileTree.svelte"

	export let data: PageServerData

	let tree: PathTree<WithId<Dotfile>>

	let currentos: OS
	onMount(() => {
		const q = new URLSearchParams(window.location.search)
		currentos = (q.get("os") || "").toLowerCase() as OS
		if (currentos == ("" as OS)) {
			if (navigator.userAgent.includes("Mac")) currentos = "macos"
			else if (navigator.userAgent.includes("Windows")) currentos = "windows"
			else if (navigator.userAgent.includes("Linux")) currentos = "linux"
		}
		if (currentos != ("" as OS) && currentos != undefined) {
			let paths = data.docs.map((v) => [v.paths[currentos], v])
			console.log(paths)
			tree = treeifyPaths(paths as PathContexts)
			console.log(tree)
		}
	})
</script>

<Navbar depotname={data.depotname}>
	<OSSelector {currentos} slot="right-items" />
</Navbar>

{#if currentos}
	<div class="flex flex-col items-center">
		<FileTree {tree} />
	</div>
{/if}

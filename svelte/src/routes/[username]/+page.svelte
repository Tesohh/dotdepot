<script lang="ts">
	import type { Dotfile } from "$lib/server/db"
	import type { OS } from "$lib/ostype"
	import type { PageServerData } from "./$types"
	import OSSelector from "$lib/components/OSSelector.svelte"
	import { onMount } from "svelte"
	import Navbar from "$lib/components/Navbar.svelte"
	import { treeifyPaths, type PathTree, type PathContexts } from "treeify-paths"
	import FileTree from "$lib/components/FileTree.svelte"
	import SvelteMarkdown from "svelte-markdown"
	import renderers from "$lib/mdrenderers"

	function getReadme(p: Dotfile[]) {
		return p.find((v) => v.paths[currentos] == "~/.config/dotdepot/readme.md")
	}

	export let data: PageServerData

	let tree: PathTree<Dotfile>
	let readme: Dotfile | undefined
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
			tree = treeifyPaths(paths as PathContexts)

			readme = getReadme(data.docs)
		}
		console.log(renderers)
	})
</script>

<Navbar depotname={data.depotname}>
	<OSSelector {currentos} slot="right-items" />
</Navbar>

{#if currentos}
	<div class="flex flex-col items-center">
		<div class="ml-[-4rem]">
			<FileTree {tree} />
		</div>

		<div class=" mt-10 w-80 border p-4 rounded-md">
			{#if readme}
				readme.md
				<hr />
				<div class="mkd">
					<SvelteMarkdown source={readme.content} {renderers} />
				</div>
			{/if}
		</div>
	</div>
{/if}

<style>
	.mkd {
		font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu,
			Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
		font-size: 8pt;
		line-height: 10pt;
	}
	.mkd :global(h1) {
		font-size: 16pt;
		line-height: 20pt;
	}
	.mkd :global(strong) {
		font: bold;
	}
	.mkd :global(em) {
		font-style: italic;
	}
	.mkd :global(a) {
		text-decoration: underline;
	}
	.mkd :global(a):hover {
		color: #93c5fd;
	}
	.mkd :global(li) {
		display: list-item;
		margin-left: 15px;
	}
	.mkd :global(ul) {
		list-style: initial;
	}
	.mkd :global(ol) {
		list-style-type: decimal;
		list-style-position: inside;
		margin-left: -15px;
	}
	.mkd :global(code) {
		font-size: 8pt;
	}
</style>

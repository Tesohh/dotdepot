<script lang="ts">
	import Navbar from "$lib/components/Navbar.svelte"
	import renderers from "$lib/mdrenderers"
	import { onMount } from "svelte"
	import SvelteMarkdown from "svelte-markdown"

	let readme: string
	onMount(() =>
		fetch("https://raw.githubusercontent.com/Tesohh/dotdepot/main/README.md")
			.then((res) => res.text())
			.then((res) => (readme = res))
	)
</script>

<Navbar depotname={""}>
	<div slot="right-items" class="text-2xl">
		<a class="fa-brands fa-github" href="https://github.com/Tesohh/dotdepot" />
	</div>
</Navbar>

<div class="flex flex-col items-center">
	{#if readme}
		<div class="mt-4 w-9/12 border p-4 rounded-md">
			readme.md
			<hr />
			<div class="mkd">
				<SvelteMarkdown source={readme} {renderers} />
			</div>
		</div>
	{/if}
</div>

<style>
	.mkd {
		font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu,
			Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
		font-size: 8pt;
		line-height: 10pt;
	}
	.mkd :global(h1) {
		font-size: 16pt;
		line-height: 26pt;
	}
	.mkd :global(h2) {
		font-size: 14pt;
		line-height: 24pt;
	}
	.mkd :global(h3) {
		font-size: 12pt;
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

<script lang="ts">
	import Navbar from "$lib/components/Navbar.svelte"
	import type { OS } from "$lib/ostype"
	import { getOS } from "$lib/getos"
	import { onMount } from "svelte"
	import OSSelector from "$lib/components/OSSelector.svelte"
	import CodeBlock from "$lib/components/CodeBlock.svelte"

	let currentos: OS
	onMount(() => {
		currentos = getOS()
	})
	export let data
</script>

<Navbar depotname={data.depotname}>
	<OSSelector {currentos} slot="right-items" />
</Navbar>

<div class="flex flex-col items-center">
	<div class="mt-5 w-80 border p-4 rounded-md">
		{data.df?.paths[currentos]?.split("/").at(-1) ?? "unknown"}
		<hr class="mb-2" />
		<CodeBlock markdownbackticks={false}>{data.df?.content}</CodeBlock>
	</div>
</div>

<style>
	:global(code) {
		font-size: 8pt;
	}
</style>

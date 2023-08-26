<script lang="ts">
	let data: any // oops
	import { HighlightAuto } from "svelte-highlight"
	import "svelte-highlight/styles/ros-pine.css"
	export let markdownbackticks = true
	$: code = data?.innerText
	$: (() => {
		if (code && markdownbackticks) {
			code = code.replaceAll("```", "")
			let codeSegments = code.split("\n")
			codeSegments.shift()

			code = codeSegments.join("\n")
			code.trim()
		}
	})()
</script>

<span bind:this={data} class="hidden"><slot /></span>
{#if code}
	<HighlightAuto {code} />
{/if}

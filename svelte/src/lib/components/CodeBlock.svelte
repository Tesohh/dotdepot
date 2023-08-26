<script lang="ts">
	let data: any // oops
	import Highlight, { HighlightAuto } from "svelte-highlight"
	import go from "svelte-highlight/languages/go"
	import "svelte-highlight/styles/ros-pine.css"
	$: code = data?.innerText
	$: (() => {
		if (code) {
			code = code.replaceAll("```", "")
			let codeSegments = code.split("\n")
			codeSegments.shift()

			code = codeSegments.join("\n")
			code.trim()
		}
	})()
	// console.log(code.split(" ")[0].replace("```", ""))
</script>

<span bind:this={data} class="hidden"><slot /></span>
{#if code}
	<HighlightAuto {code} />
{/if}

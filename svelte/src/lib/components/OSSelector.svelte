<script lang="ts">
	import type { OS } from "$lib/ostype"
	import Osbutton from "./osbutton.svelte"
	export let currentos: OS
</script>

<div
	class="ml-10 flex flex-row max-w-fit
    rounded-xl border-gray-300 border
    [&>*:not(:first-child)]:border-l [&>*:first-child]:rounded-l-xl [&>*:last-child]:rounded-r-xl"
>
	<Osbutton os="windows" selected={currentos == "windows"} />
	<Osbutton os="macos" selected={currentos == "macos"} />
	<Osbutton os="linux" selected={currentos == "linux"} />
</div>

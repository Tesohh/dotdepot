<script lang="ts">
	import type { OS } from "$lib/ostype"

	const os2color = { windows: "bg-blue-500", macos: "bg-gray-500", linux: "bg-yellow-500" }

	export let os: OS
	export let selected: boolean

	const select = () => {
		const urlParams = new URLSearchParams(window.location.search)
		urlParams.set("os", os)
		window.location.search = urlParams.toString()
	}
</script>

<button
	class="w-12 h-8 md:w-16 md:h-10 text-lg md:text-xl {selected ? os2color[os] : ''}"
	on:click={select}
>
	{#if os == "windows"}
		<i class="fa-brands fa-windows fa-fw" />
	{:else if os == "macos"}
		<i class="fa-brands fa-apple fa-fw" />
	{:else if os == "linux"}
		<i class="fa-brands fa-linux fa-fw" />
	{/if}
</button>

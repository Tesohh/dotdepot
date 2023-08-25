<script lang="ts">
	import type { Dotfile } from "$lib/server/db"
	import type { WithId } from "mongodb"
	import { onMount } from "svelte"
	import type { PathTree } from "treeify-paths"

	export let tree: PathTree<WithId<Dotfile>>
</script>

<ul>
	{#each tree.children as child}
		<li>
			{#if child.children}
				{#if child.name == ""}
					📁 {child.path.split("/").at(-1)}
				{:else}
					<a href={child.ctx.safeID}>📄 {child.name}</a>
				{/if}
				<br />
				{#if child.name == ""}
					<svelte:self tree={child} />
				{/if}
			{/if}
		</li>
	{/each}
</ul>

<style>
	ul {
		padding-left: 1.1rem;
	}
</style>

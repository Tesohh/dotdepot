<script lang="ts">
	import type { Dotfile } from "$lib/server/db"
	import type { WithId } from "mongodb"
	import { onMount } from "svelte"
	import type { PathTree } from "treeify-paths"

	export let tree: PathTree<Dotfile>
</script>

{#if tree}
	<ul>
		{#each tree.children as child}
			<li>
				{#if child.children}
					{#if child.name == ""}
						📁 {child.path.split("/").at(-1)}
					{:else}
						<a class="hover:text-blue-300" href={child.ctx.safeID ?? ""}>📄 {child.name}</a>
					{/if}
					<br />
					{#if child.name == ""}
						<svelte:self tree={child} />
					{/if}
				{/if}
			</li>
		{/each}
	</ul>
{:else}
	hmmm... it seems there are no files here.
{/if}

<style>
	ul {
		padding-left: 1.1rem;
	}
</style>

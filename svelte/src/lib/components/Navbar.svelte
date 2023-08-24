<script lang="ts">
	export let depotname: string
</script>

<div class="flex flex-row items-center p-2">
	<span class="font-bold text-lg md:text-xl">📦 {depotname ? "/" + depotname : "dotdepot"}</span>
	<slot name="left-items" />

	<div class="grow" id="navbar-divider" />

	<div class="flex flex-row-reverse">
		<slot name="right-items" />
	</div>
</div>

<!-- Logic -->
<script>
	import * as Icons from "lucide-svelte";

	export let icon = null
	const IconComponent = icon && Icons[icon] ? Icons[icon] : null
</script>


<!-- Component body -->
<div class="snackbar">

	{#if IconComponent}
		<div class="icon">
			<IconComponent style="width: 100%; height: 100%;" />
		</div>
	{/if}

	{#if $$slots.default && IconComponent}
		<div class="spacer"></div>
	{/if}

	<div class="text">
		<slot></slot>
	</div>

</div>


<!-- Styling -->
<style>
	.snackbar {
		color: #236;
		background: #68F;
		box-shadow: 0 0 32px #000;
		max-width: 720px;
		padding: 12px 18px;
		border-radius: 24px;
		display: inline-flex;
		align-items: center;
		position: fixed;
		left: 50%;
		bottom: 32px;
		transform: translateX(-50%);
		transition: all .2s;
	}

	.icon {
		position: relative;
		width: 20px;
		height: 20px;
		margin: 0 -4px 0 -4px;
	}

	.spacer {
		width: 12px;
	}
</style>

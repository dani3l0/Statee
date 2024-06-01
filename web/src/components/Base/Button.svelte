<!-- Logic -->
<script>
	import * as Icons from "lucide-svelte";

	export let highlighted = false
	export let warn = false
	export let danger = false
	export let icon = null

	const IconComponent = icon && Icons[icon] ? Icons[icon] : null
</script>


<!-- Component body -->
<div class="button
	{highlighted ? "highlighted" : ""}
	{warn ? "warn" : ""}
	{danger ? "danger" : ""}
	">

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
	.button {
		color: #DDD;
		background: #333;
		padding: 12px 16px;
		border-radius: 24px;
		display: inline-block;
		cursor: pointer;
		margin: 4px;
		display: inline-flex;
		align-items: center;
		transition: all .2s;
	}
	.button:hover {
		opacity: .8;
	}

	.highlighted {
		color: #236;
		background: #68F;
	}

	.warn {
		color: #620;
		background: #F80;
	}

	.danger {
		color: #511;
		background: #F33;
	}

	.icon {
		position: relative;
		width: 20px;
		height: 20px;
		margin: 0 -4px 0 -4px;
	}

	.spacer {
		width: 8px;
	}
</style>

<script>
	export let onClick = () => {}
	export let highlighted = false
	export let icon = null
</script>


<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div on:click={onClick} class="button {highlighted ? "highlighted" : ""}">
	{icon ? icon : ""}
	<slot></slot>
</div>


<style>
	.button {
		color: #DDD;
		background: #333;
		padding: 12px 16px;
		border-radius: 24px;
		display: inline-block;
		cursor: pointer;
		margin: 4px;
		transition: all .2s;
	}
	.button:hover {
		opacity: .8;
	}

	.button.highlighted {
		color: #222;
		background: #68F;
	}
</style>

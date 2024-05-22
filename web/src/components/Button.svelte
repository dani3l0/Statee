<script>
	export let onClick = () => {}
</script>


<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div on:click={onClick} class="button">
	<slot></slot>
</div>


<style>
	.button {
		color: #222;
		background: #68F;
		padding: 12px 16px;
		border-radius: 24px;
		display: inline-block;
		cursor: pointer;
		transition: all .2s;
	}
	.button:hover {
		opacity: .8;
	}
</style>

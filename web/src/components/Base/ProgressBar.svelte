<!-- Logic -->
<script>
	export let min = 0
	export let max = 100
	export let value = 0

	$: fill = 100 * (value - min) / (max - min)
</script>


<!-- Component body -->
<div class="progress-bar">
	<div class="fill" style="width: {fill}%;"></div>
</div>


<!-- Styling -->
<style>
	.progress-bar {
		position: relative;
		height: 16px;
		width: 100%;
		background: #282828;
		border-radius: 16px;
		overflow: hidden;
	}

	.fill {
		width: 0;
		height: 100%;
		background: #68F;
		transition: all .8s ease;
	}
</style>

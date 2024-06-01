<!-- Logic -->
<script>
	export let hidden = false
</script>


<!-- Component body -->
<div class="spinner {hidden ? "hidden" : ""}">
	<div class="dot"></div>
</div>


<!-- Styling -->
<style>
	@keyframes spin {
		from { transform: rotateZ(0) }
		to { transform: rotateZ(360deg) }
	}

	.spinner {
		position: relative;
		width: 64px;
		height: 64px;
		background: #282828;
		border-radius: 32px;
		overflow: hidden;
		animation: spin 1s linear infinite;
		transition: all .5s;
	}

	.spinner.hidden {
		opacity: 0;
		height: 0;
		visibility: hidden;
		animation: none;
	}

	.dot {
		width: 36px;
		height: 36px;
		margin: 10px;
		background: #666;
		border-radius: 32px;
	}
</style>

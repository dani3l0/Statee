<!-- Logic -->
<script>
	export let value = 0

	setInterval(() => {
		value = Math.round(Math.random() * 100)
	}, 3000);
</script>


<!-- Component body -->
<div style="--num: {value}"></div>


<!-- Styling -->
<style>
	@property --num {
		syntax: '<integer>';
		initial-value: 0;
		inherits: false;
	}

	div {
		transition: --num 1s;
		counter-reset: num var(--num);
	}
	div::after {
		content: counter(num);
	}
</style>

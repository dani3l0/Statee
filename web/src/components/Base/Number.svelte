<!-- Logic -->
<script>
	// Read parameters
	export let value = 0
	export let precision = 2
	
	// Runtime
	let displayValue = 0

	// easeInOutCubic functionm
	// aka cubic-bezier(0.65, 0, 0.35, 1)
	function ease(x) {
		return x < 0.5 ? 4 * x * x * x : 1 - Math.pow(-2 * x + 2, 3) / 2;
	}
	
	// Function to animate the incrementing value with easing
	function animateValue(targetValue, duration) {
		const startValue = displayValue
		const startTime = performance.now()

		function updateValue(currentTime) {
			const elapsedTime = currentTime - startTime
			const progress = Math.min(elapsedTime / duration, 1)
			const easedProgress = ease(progress)
			displayValue = startValue + (targetValue - startValue) * easedProgress

			if (progress < 1) requestAnimationFrame(updateValue)
			else displayValue = targetValue
		}

		requestAnimationFrame(updateValue)
	}

	// Animate number when 'value' changes
	$: {
		animateValue(value, 1000)
	}
</script>


<!-- Component body -->
<div>{displayValue.toFixed(precision)}</div>


<!-- Styling -->
<style>
	div {
		display: inline;
	}
</style>

<script>
    import ProgressBar from "./ProgressBar.svelte"

    export let value = 0
    let displayValue = 0

    // Easing function
    function ease(x) {
        return -(Math.cos(Math.PI * x) - 1) / 2
    }

    // Function to animate the incrementing value with easing
    function animateValue(targetValue, duration) {
        const startValue = displayValue
        const startTime = performance.now()

        function updateValue(currentTime) {
            const elapsedTime = currentTime - startTime
            const progress = Math.min(elapsedTime / duration, 1)
            const easedProgress = ease(progress) // Apply easing function
            displayValue = startValue + (targetValue - startValue) * easedProgress

            if (progress < 1) {
                requestAnimationFrame(updateValue)
            } else {
                displayValue = targetValue // Ensure final value is exact
            }
        }

        requestAnimationFrame(updateValue)
    }

    // Update value every 3 seconds
    setInterval(() => {
        value = Math.random() * .5 + 3.2
        animateValue(value, 800) // Animate the new value
    }, 3000)
</script>

<!-- Component body -->
<div>{displayValue.toFixed(2)}</div>
<ProgressBar value={value} min={3} max={6.00} />

<!-- Styling -->
<style>
    div {
        font-size: 24px;
        padding: 20px;
        background-color: #333;
        color: #ddd;
        border-radius: 8px;
        display: inline-block;
    }
</style>

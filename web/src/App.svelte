<!-- Imports -->
<script>
	import "@fontsource/quicksand/600.css"
    import ProgressBar from "./components/Base/ProgressBar.svelte";
    import Incremental from "./components/Base/Float.svelte";
</script>


<!-- Main wrapper -->
<main>
	<!-- component testing for now -->
	<Incremental value={3} />
</main>


<!-- Global, system-wide styling -->
<style>
	:global(body) {
		background: #000;
		color: #DDD;
		font-family: 'Quicksand', sans-serif;
		user-select: none;
	}
</style>

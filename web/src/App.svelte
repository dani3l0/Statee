<!-- Imports -->
<script>
	import "@fontsource/quicksand/600.css"
	import * as Icons from 'lucide-svelte'
    import Button from "./components/Button.svelte";
</script>


<!-- Main wrapper -->
<main>
	<Icons.ArrowDown />
	<h1>Statee</h1>
	<p>Yeah, it works</p>
	<Button>Click xd</Button>
	<Button highlighted>Click xd</Button>

	<Button highlighted>Click xd</Button>
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

<script>
	import { marked } from 'marked';
    import { Button } from 'flowbite-svelte';

	let value = $state(`Some words are *italic*, some are **bold**\n\n- lists\n- are\n- cool`);
	let foo = 'baz'
	let bar = 'qux'
	let result = $state("");

	async function doPost () {
		const res = await fetch('/api', {
			method: 'POST',
			body: JSON.stringify({
				foo,
				bar
			})
		})
		
		const json = await res.json();
		result = JSON.stringify(json);
	}
</script>

<div class="grid">
	input
	<textarea bind:value={value}></textarea>

	output
	<div>{@html marked(value)}</div>

	<Button on:click={doPost}>
		Post it.
	</Button>

	<p>
	Result:
	</p>
	<pre>
	{result}
	</pre>
</div>

<style>
	.grid {
		display: grid;
		grid-template-columns: 5em 1fr;
		grid-template-rows: 1fr 1fr;
		grid-gap: 1em;
		height: 100%;
	}

	textarea {
		flex: 1;
		resize: none;
	}
</style>

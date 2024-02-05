<script lang="ts">
	import { onMount } from 'svelte';

	let categories: Category[] = [];

	const fetchCategories = async () => {
		const response = await fetch('http://localhost:9090/categories');

		if (response.ok) {
			categories = await response.json();
		} else {
			console.error('Failed to fetch categories');
		}
	};

	onMount(() => {
		fetchCategories();
	});
</script>

<ul role="list" class="marker:text-red-500 list-disc pl-5 my-2 space-y-1 text-slate-600">
	{#each categories as category}
		<li><a href={`/categories/${category.ID}/posts`}>{category.Title}</a></li>
	{/each}
</ul>

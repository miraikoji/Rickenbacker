<script lang="ts">
	import { onMount } from 'svelte';

	let posts: Post[] = [];

	const fetchPosts = async () => {
		const response = await fetch('http://localhost:9090/posts');

		if (response.ok) {
			posts = await response.json();
		} else {
			console.error('Failed to fetch posts');
		}
	};

	onMount(() => {
		fetchPosts();
	});
</script>

<h1>Posts</h1>

{#if posts.length > 0}
	<ul>
		{#each posts as post}
			<li>{post.Title}</li>
		{/each}
	</ul>
{:else}
	<p>Loading...</p>
{/if}

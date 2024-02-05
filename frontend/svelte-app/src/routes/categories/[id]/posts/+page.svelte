<script lang="ts">
	export let data;
</script>

<h1>Posts</h1>

{#if data.posts.length > 0}
	<ul>
		{#each data.posts as post}
			<li>{post.Title}</li>
		{/each}
	</ul>
{:else}
	<p>Loading...</p>
{/if}

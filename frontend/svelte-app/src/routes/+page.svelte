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

<div class="bg-gradient-to-bl flex justify-center lg:h-screen mt-12">
	<div class="container mx-auto">
		<div class="grid grid-cols-12 gap-12">
			{#each posts as post (post.ID)}
				<div class="bg-white rounded-sm border col-span-4">
					<img
						src={post.ImageURL || '/posts/default_hero.jpg'}
						alt=""
						class="w-full h-48 rounded-sm object-cover"
					/>
					<div class="px-1 py-4">
						<div class="font-bold text-xl mb-2">{post.Title}</div>
						<p class="text-gray-700 text-base">
							{post.Body}
						</p>
					</div>
					<div class="px-1 py-4">
						<a href="/posts/1" class="text-blue-500 hover:underline">Read More</a>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

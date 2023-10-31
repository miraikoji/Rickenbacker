<script>
	export let data;
	import Badge from '@/components/Badge.svelte';
	import PostMetaData from '@/components/PostMetaData.svelte';
</script>

<div class="mt-16">
	<div class="text-center my-4">
		<Badge name="GitHub" />
		<Badge name="GitHub Actions" />
		<Badge name="Fargate" />
		<Badge name="Fargate Spot" />
		<Badge name="Terraform" />
		<Badge name="ECS" />
	</div>
	<h1 class="text-4xl font-light text-center">
		{data.title}
	</h1>
	<div class="my-6">
		<PostMetaData />
	</div>
	<img src="/posts/default_hero.jpg" alt="" />
	<main class="mt-8">
		{data.content}
	</main>
</div>

<style>
	main p {
		font-size: 1.1rem;
		margin: 1.5rem 0;
	}
</style>

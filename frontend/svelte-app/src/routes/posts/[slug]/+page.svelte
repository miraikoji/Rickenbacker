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
		CIの料金を気にせずじゃぶじゃぶ使いたいので、GitHub ActionsのSelf
		hosted-runnerをFargateSpot+Terraform+ECSで動かした
	</h1>
	<div class="my-6">
		<PostMetaData />
	</div>
	<img src="/posts/default_hero.jpg" alt="" />
	<main class="mt-8">
		<p>
			CIの料金、元々はCircleCIを使ったいたのですが、料金形態が微妙に分かりづらい上に少々お高め。わかりやすさと管理のしやすさなどを求めてGitHub
			Actionsにしてみたもののこちらも高額になってきた。
		</p>
		<p>
			開発組織として、生産性やスピードアップのために、小さめなPRをたくさん出したいという事を考えると、このCIの事で頭を悩ませるのは嫌だなと。
		</p>
	</main>
</div>

<!-- <h1 class="text-4xl font-black">{data.title}</h1>
<div>{data.content}</div> -->

<style>
	main p {
		font-size: 1.1rem;
		margin: 1.5rem 0;
	}
</style>

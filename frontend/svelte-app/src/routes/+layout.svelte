<script>
	import '../app.css';
	import Header from '@/components/Header.svelte';
	import SideHeader from '@/components/SideHeader.svelte';
	import Badge from '@/components/Badge.svelte';
</script>

<Header />

<div class="mx-auto max-w-7xl p-6 lg:px-8">
	<div class="grid grid-cols-12 grid-rows-2 gap-12">
		<div class="col-span-9">
			<slot />
		</div>
		<aside class="col-span-3 my-24">
			<div class="mb-6">
				<SideHeader title="カテゴリー" />
				<ul role="list" class="marker:text-red-500 list-disc pl-5 my-2 space-y-1 text-slate-600">
					<li>エンジニアブログ</li>
					<li>ポッドキャスト</li>
				</ul>
			</div>
			<div class="mb-6">
				<SideHeader title="タグ" />
				<div class="my-2">
					<Badge name="GitHub" />
					<Badge name="Terraform" />
					<Badge name="Lambda" />
					<Badge name="AWS" />
					<Badge name="Go" />
					<Badge name="Echo" />
					<Badge name="Ruby" />
					<Badge name="Ruby on Rails" />
					<Badge name="RSpec" />
				</div>
			</div>
			<div>
				<div class="isolate flex -space-x-2 overflow-hidden">
					<img
						class="relative z-30 inline-block h-8 w-8 rounded-full ring-2 ring-white"
						src="https://avatars.githubusercontent.com/u/60030224?v=4"
						alt=""
					/>
					<img
						class="relative z-20 inline-block h-8 w-8 rounded-full ring-2 ring-white"
						src="https://avatars.githubusercontent.com/u/6788936?v=4"
						alt=""
					/>
					<img
						class="relative z-10 inline-block h-8 w-8 rounded-full ring-2 ring-white"
						src="https://avatars.githubusercontent.com/u/36214433?v=4"
						alt=""
					/>
				</div>
			</div>
		</aside>
	</div>
</div>

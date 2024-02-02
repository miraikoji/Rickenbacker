import { error } from '@sveltejs/kit';

export const load = async ({ fetch }): Promise<{ posts: Post[] }> => {
	const url = 'http://localhost:9090/posts';
	const response = await fetch(url);
	if (response.ok) {
		const posts = await response.json();
		return {
			posts
		};
	} else {
		console.error('Failed to fetch posts');
		throw error(404, 'Posts not found');
	}
};

import { error } from '@sveltejs/kit';

export const load = async ({ params, fetch }): Promise<{ posts: Post[] }> => {
	const fetchPost = async (url: string): Promise<Post[] | null> => {
		const response = await fetch(url);

		if (response.ok) {
			return await response.json();
		} else {
			console.error('Failed to fetch posts');
			return null;
		}
	};

	if (!isNaN(Number(params.id))) {
		const response = await fetchPost(`http://localhost:9090/categories/${params.id}/posts`);

		if (response !== null) {
			return {
				posts: response
			};
		} else {
			throw error(404, 'Post not found');
		}
	}

	throw error(404, 'Not found');
};

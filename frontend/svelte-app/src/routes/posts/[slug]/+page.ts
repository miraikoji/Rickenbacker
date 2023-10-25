import { error } from '@sveltejs/kit';

/** @type {import('./$types').PageLoad} */

export async function load({ params, fetch }) {
	const fetchPost = async (url: string): Promise<Post | null> => {
		const response = await fetch(url);

		if (response.ok) {
			return await response.json();
		} else {
			console.error('Failed to fetch posts');
			return null;
		}
	};

	if (!isNaN(Number(params.slug))) {
		const response = await fetchPost(`http://backend:9090/posts/${params.slug}`);

		if (response !== null) {
			return {
				title: response.Title,
				content: response.Body
			};
		} else {
			throw error(404, 'Post not found');
		}
	}

	throw error(404, 'Not found');
}

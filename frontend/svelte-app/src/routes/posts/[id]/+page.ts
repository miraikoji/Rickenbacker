/** @type {import('./$types').PageLoad} */

import { error } from '@sveltejs/kit';

export const load = async ({ params, fetch }): Promise<{ post: Post }> => {
	const fetchPost = async (url: string): Promise<Post | null> => {
		const response = await fetch(url);

		if (response.ok) {
			return await response.json();
		} else {
			console.error('Failed to fetch posts');
			return null;
		}
	};

	if (!isNaN(Number(params.id))) {
		const response = await fetchPost(`http://localhost:9090/posts/${params.id}`);

		if (response !== null) {
			return {
				post: {
					ID: response.ID,
					Title: response.Title,
					Body: response.Body,
					CreatedAt: response.CreatedAt,
					UpdatedAt: response.UpdatedAt,
					UserID: response.UserID,
					ImageURL: response.ImageURL,
					DeletedAt: null
				}
			};
		} else {
			throw error(404, 'Post not found');
		}
	}

	throw error(404, 'Not found');
};

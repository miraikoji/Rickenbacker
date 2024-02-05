// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}
}

declare type Post = {
	ID: number;
	CreatedAt: string;
	UpdatedAt: string;
	DeletedAt: string | null;
	Title: string;
	Body: string;
	UserID: number;
	ImageURL: string | null;
};

declare type Category = {
	ID: number;
	Title: string;
};

declare type Breadcrumb = {
	Name: string;
	Path: string;
};

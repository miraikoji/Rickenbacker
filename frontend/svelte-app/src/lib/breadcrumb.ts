// src/lib/breadcrumb.ts
import { writable } from 'svelte/store';

// Breadcrumbアイテムの型定義
interface Breadcrumb {
	Name: string;
	Path: string;
}

// パンくずリストのためのストアを作成
const breadcrumbStore = writable<Breadcrumb[]>([]);

export default breadcrumbStore;

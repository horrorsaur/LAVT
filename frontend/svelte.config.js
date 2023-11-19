import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: vitePreprocess(),
  kit: {
    adapter: adapter({
      // Not using Svelte server-side logic. Add fallback for adapter-static.
      fallback: "index.html",
    }),
  }
};

export default config;

import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

const config = async () => {
  if (process.env.NODE_ENV !== 'production') {
    const dotenv = await import('dotenv');
    dotenv.config();
  }
};

config();

export default {
  preprocess: [vitePreprocess({})],
};
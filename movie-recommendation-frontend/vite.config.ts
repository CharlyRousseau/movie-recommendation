/// <reference types="vitest" />
import { defineConfig } from 'vite'
import { configDefaults } from 'vitest/config'

import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from 'path'

export default defineConfig({
  plugins: [svelte()],
  test: {
    environment: 'jsdom',
    globals: true,
    alias: [{ find: /^svelte$/, replacement: 'svelte/internal' }],
    exclude: [...configDefaults.exclude, './postcss.config.cjs'],
  },

  resolve: {
    alias: {
      $lib: path.resolve("./src/lib"),
      $components: path.resolve("./src/components"),
      $shadcn: path.resolve("./src/lib/components"),
    },
  },
})

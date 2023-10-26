import {
  fileURLToPath,
  URL
} from 'node:url'

import {
  defineConfig
} from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx()],
  server: {
    host: '0.0.0.0',
    port: 3008,
    strictPort: false,
    open: false,
    proxy: {
      '/v1': {
        target: ' http://localhost:3089',
        changeOrigin: true,
      }
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src',
        import.meta.url)),
      "comps": fileURLToPath(new URL('src/components',
        import.meta.url)),
      "views": fileURLToPath(new URL('src/views',
        import.meta.url)),
      "store": fileURLToPath(new URL('src/store',
        import.meta.url)),
      "utils": fileURLToPath(new URL('src/utils',
        import.meta.url)),
      "libs": fileURLToPath(new URL('src/libs',
        import.meta.url)),
      "api": fileURLToPath(new URL('src/ApiDevice',
        import.meta.url)),
      "styles": fileURLToPath(new URL('src/styles',
        import.meta.url)),
    }
  },
})
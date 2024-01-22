import vue from "@vitejs/plugin-vue";
import { fileURLToPath, URL } from "node:url";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig, loadEnv } from "vite";
import vuetify from "vite-plugin-vuetify";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env: Record<string, string> = loadEnv(mode, process.cwd(), "");

  return {
    plugins: [
      vue(),
      vuetify()
    ],
    resolve: {
      alias: {
        "@": fileURLToPath(new URL("./src", import.meta.url))
      }
    },
    server: {
      port: 5175,
      host: "0.0.0.0",
      open: false,
      hmr: true,
      cors: true,
    },
    build: {
      // sourcemap: env.VITE_PROD_SOURCE_MAPS, // 生产环境是否生成 source map 文件
      chunkSizeWarningLimit: 2000, // 忽略chunk大小警告的限制
      rollupOptions: {
        plugins: [
          visualizer() // 查看打包体积
        ]
      }
    }
  };
});

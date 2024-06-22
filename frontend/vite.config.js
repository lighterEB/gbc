import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

const rootPath = new URL('.', import.meta.url).pathname
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      imports: [
        {
          'naive-ui': ['useDialog', 'useMessage', 'useNotification', 'useLoadingBar'],
        },
      ],
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
    Icons(),
  ],
  resolve: {
    alias: {
      '@': rootPath + 'src',
      stores: rootPath + 'src/stores',
      wailsjs: rootPath + 'wailsjs',
    },
  },
})

import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import federation from '@originjs/vite-plugin-federation';

// https://vitejs.dev/config/
export default defineConfig(({ command }) => ({
  plugins: [
    react(),
    federation({
      name: 'taskManagement',
      filename: 'remoteEntry.js',
      exposes: {
        './App': './src/App.tsx',
      },
      shared: ['react', 'react-dom'],
    }),
  ],
  server: {
    port: 5001,
    origin: 'http://localhost:5001',
    cors: true,
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
  },
  preview: {
    port: 5001,
  },
  build: {
    target: 'esnext',
    minify: false,
    cssCodeSplit: false,
    modulePreload: false,
    rollupOptions: {
      output: {
        format: 'esm',
      },
    },
  },
}));

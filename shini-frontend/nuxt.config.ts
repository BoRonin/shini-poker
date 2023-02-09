// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  runtimeConfig: {
    BASE_URL: process.env.BASE_URL
  },
  app: {

    head: {
      meta:
        [{
          'http-equiv': 'Content-Security-Policy',
          content: 'upgrade-insecure-requests'
        }],
      script: [],
      noscript: []
    }

  },
  css: [
    '@/assets/css/main.sass'
  ],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },


})

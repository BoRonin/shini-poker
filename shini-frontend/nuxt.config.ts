// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    pageTransition: { name: 'page', mode: 'out-in' },
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
  modules:[
    '@vueuse/nuxt',
  ]


})

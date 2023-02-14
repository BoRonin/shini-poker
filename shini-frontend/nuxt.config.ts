// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devServer:{
    port: 3001,
    url: "http://localhost:3001"
  },
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
  ],

})

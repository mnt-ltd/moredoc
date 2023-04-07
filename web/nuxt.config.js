export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    meta: [
      { charset: 'utf-8' },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1.0, maximum-scale=1.0',
      },
      { name: 'keywords', content: '' },
      { name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  generate: {
    dir: '../dist',
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    'element-ui/lib/theme-chalk/index.css',
    'element-ui/lib/theme-chalk/display.css',
    '~assets/font-awesome-4.7.0/css/font-awesome.min.css',
    '~assets/css/var.scss',
    '~assets/css/app.scss',
    '~assets/css/github-markdown-light.css',
  ],
  styleResources: {
    scss: ['~/assets/css/*.scss'],
  },

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ['@/plugins/element-ui'],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    // '@nuxtjs/eslint-module',
    '@nuxtjs/style-resources',
    '@nuxtjs/composition-api/module',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    // https://go.nuxtjs.dev/pwa
    '@nuxtjs/pwa',
  ],

  // 路由切换了，回到页面顶部
  router: {
    scrollBehavior(to, from, savedPosition) {
      return { x: 0, y: 0 }
    },
  },

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
    proxy: true,
  },
  server: {
    port: 3000, // default: 3000
    host: '0.0.0.0', // default: localhost
  },
  proxy: {
    '/api': {
      target: 'http://127.0.0.1:8880', // 目标服务器
      changeOrigin: true,
    },
    '/view': {
      target: 'http://127.0.0.1:8880', // 目标服务器
      changeOrigin: true,
    },
    '/uploads': {
      target: 'http://127.0.0.1:8880', // 目标服务器
      changeOrigin: true,
    },
    '/download': {
      target: 'http://127.0.0.1:8880', // 目标服务器
      changeOrigin: true,
    },
  },
  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: 'en',
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    postcss: null,
    transpile: [/^element-ui/],
  },
}

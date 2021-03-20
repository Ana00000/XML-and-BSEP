module.exports = {
  devServer: {
      proxy: {
        '^/api': {
          target: 'https://localhost:8080',
          changeOrigin: true,
          secure: false,
          pathRewrite: {'^/api': '/api'},
          logLevel: 'debug'
        },
      }
    },
  transpileDependencies: [
    'vuetify'
  ]
}
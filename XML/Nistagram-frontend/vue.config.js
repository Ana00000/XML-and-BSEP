module.exports = {
  devServer: {
    port: 8081,
    watchOptions: {
      poll: true
    }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
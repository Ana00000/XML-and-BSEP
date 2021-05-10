module.exports = {
  devServer: { proxy: 'http://localhost:8082/' },
  transpileDependencies: [
    'vuetify'
  ]
}
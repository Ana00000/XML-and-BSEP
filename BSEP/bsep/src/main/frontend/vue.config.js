const fs = require('fs');
module.exports = {
  devServer: {
    disableHostCheck: true,
    https: true
  },

  transpileDependencies: [
    'vuetify'
  ]
}
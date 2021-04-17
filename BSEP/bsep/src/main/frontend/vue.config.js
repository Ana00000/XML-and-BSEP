const fs = require('fs');
module.exports = {
  devServer: {
    host: 'localhost',
    disableHostCheck: true,
    https: true,
    key: fs.readFileSync('../../../localhost.pem'),
    cert: fs.readFileSync('../../../localhost.cer'),
    public: 'https://localhost:8080/'
  },

  transpileDependencies: [
    'vuetify'
  ]
}
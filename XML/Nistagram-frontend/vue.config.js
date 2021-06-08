const fs = require('fs');
module.exports = {
  devServer: {
    port: 8081,
    https:true,
    key: fs.readFileSync('/etc/ssl/localhost.pem'),
    cert: fs.readFileSync('/etc/ssl/localhost.cer'),
    ca: fs.readFileSync('/etc/ssl/localhost.pem'),
    watchOptions: {
      poll: true
    }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
const fs = require('fs');
module.exports = {
  devServer:{
    https:true,
    key: fs.readFileSync('../../../localhost.pem'),
    cert: fs.readFileSync('../../../localhost.cer'),
    ca: fs.readFileSync('../../../localhost.pem')
  },
  transpileDependencies: [
    'vuetify'
  ]
}
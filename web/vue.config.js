module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  pwa: {
    name: 'Web',
    iconPaths: {
      favicon32: 'img/icons/favicon.png',
      favicon16: 'img/icons/favicon.png',
      appleTouchIcon: 'img/icons/favicon.png',
      maskIcon: 'img/icons/favicon.png',
      msTileImage: 'img/icons/favicon.png'
    },
    manifestOptions:{
      icons:[
        {
          'src': './img/icons/android-chrome-192x192.png',
          'sizes': '192x192',
          'type': 'image/png'
        },
        {
          'src': './img/icons/android-chrome-512x512.png',
          'sizes': '512x512',
          'type': 'image/png'
        }
      ]
    }
}
}
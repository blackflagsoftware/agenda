import { createApp } from 'vue'
// import './style.css'
import App from './App.vue'

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as labsComponents from 'vuetify/labs/components'
import * as directives from 'vuetify/directives'
import * as icons from 'vuetify/iconsets/mdi-svg'

const vuetify = createVuetify({
  components: {
    ...components,
    ...labsComponents,
  },
  directives,
  icons,
})

createApp(App).use(vuetify).mount('#app')

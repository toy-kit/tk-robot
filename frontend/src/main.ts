import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import { createI18n } from 'vue-i18n'
import locales from './locales/index'
import "./assets/fonts/iconfont.css"
import { createPinia } from 'pinia'
const i18n = createI18n({
  locale: 'zh',
  fallbackLocale: 'en',
  messages: locales,
})
const pinia = createPinia()
createApp(App).use(i18n).use(pinia).mount('#app')

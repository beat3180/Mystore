import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'






//import VueGoogleMaps from '@fawmi/vue-google-maps'
// const VueGeolocation = require('vue3-geolocation')
// const GMaps = require('vuejs3-google-maps')
//const Geocoder = require('@pderas/vue2-geocoder')


axios.defaults.baseURL = 'http://localhost:80/api/'
axios.defaults.withCredentials = true
createApp(App).use(store).use(router).mount('#app')

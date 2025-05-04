/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import { registerPlugins } from '@/plugins'

// Components
import App from './App.vue'


// Composables
import { createApp } from 'vue'
import VueSweetalert2 from 'vue-sweetalert2';
import 'sweetalert2/dist/sweetalert2.min.css';
import axios from 'axios';

axios.defaults.baseURL =  'http://localhost:6969';
axios.defaults.withCredentials = true;
axios.defaults.withXSRFToken = true;
axios.defaults.headers.common['Content-Type'] = 'application/json';

const app = createApp(App)

registerPlugins(app)

app.use(VueSweetalert2);

app.mount('#app')

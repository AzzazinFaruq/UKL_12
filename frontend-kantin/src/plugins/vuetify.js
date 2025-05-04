/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'
import { VDateInput } from "vuetify/labs/VDateInput";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { VNumberInput } from 'vuetify/labs/VNumberInput'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    components,
    directives,
    defaultTheme: 'myCustomTheme',
    themes: {
      myCustomTheme: {
        dark: false,
        colors: {
          background: '#F9F2EE',
          primary: '#B06F4F',
          alert: '#C37B58',
          success: '#B06F4F',
          info: '#B06F4F',
          warning: '#B06F4F',
          error: '#B06F4F',
        },
        variables: {
          'font-family': 'Helvetica, Arial, sans-serif',
        },
      },
    },
  },
  components: {
    VDateInput,
    VNumberInput,
  },
})

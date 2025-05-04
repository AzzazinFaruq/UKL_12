<template>
  <v-app>
    <Navbar v-if="showNavbar" :checkEnabled="!isExcludedRoute" />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import Navbar from './components/navbar.vue'

export default {
  name: 'App',
  components: {
    Navbar
  },
  data() {
    return {
      excludedRoutes: ['/login', '/register', '/siswa/register', '/stan/register']
    }
  },
  computed: {
    showNavbar() {
      return !this.isExcludedRoute
    },
    isExcludedRoute() {
      return this.excludedRoutes.includes(this.$route.path)
    }
  }
}
</script>

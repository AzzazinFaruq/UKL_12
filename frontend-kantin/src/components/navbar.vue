<template>
  <div>
    <v-navigation-drawer v-model="drawer" >
        <v-list>
          <v-list-item-title class="">
            <div class="logo-wrap">
              <v-img src="@/assets/register/logo-kantin.png" alt="" max-width="200" class="mb-3 ml-5"></v-img>
            </div>
          </v-list-item-title>
          <v-divider></v-divider>
          <v-list-item
            v-for="item in menuItems"
            :key="item.title"
            :to="item.path"
            :prepend-icon="item.icon"
          :title="item.title"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <!-- Navbar -->
    <v-app-bar color="white" elevation="0">
      <!-- <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon> -->
      <!-- <v-toolbar-title>Kantin Digital</v-toolbar-title> -->
      <v-spacer></v-spacer>

      <!-- Profile Menu -->
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn icon v-bind="props">
            <v-avatar size="40">
              <v-icon>mdi-account-circle</v-icon>
            </v-avatar>
          </v-btn>
        </template>
        <v-list>
          <v-list-item @click="logout">
            <v-list-item-title>Keluar</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <!-- Sidebar -->
  </div>
</template>

<script>
import axios from 'axios'
import Swal from 'sweetalert2'

export default {
  props: {
    checkEnabled: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      drawer: true,
      status: false,
      role: '',
      menuItems:[],
      menuItemsSiswa: [
        {
          title: 'Dashboard',
          icon: 'mdi-view-dashboard',
          path: '/siswa'
        },
        {
          title: 'Menu',
          icon: 'mdi-food',
          path: '/products'
        },
        {
          title: 'Transaksi',
          icon: 'mdi-receipt',
          path: '/transactions'
        },
        {
          title: 'History',
          icon: 'mdi-file-document',
          path: '/reports'
        },
      ],
      MenuItemStand:[
        {
          title: 'Dashboard',
          icon: 'mdi-file-document',
          path: '/stan'
        },
        {
          title: 'Transaksi',
          icon: 'mdi-file-document',
          path: '/stan/transaksi'
        },
        {
          title: 'Manajemen Menu',
          icon: 'mdi-file-document',
          path: '/stan/menu'
        },
        {
          title: 'Manajemen Diskon',
          icon: 'mdi-file-document',
          path: '/stan/diskon'
        },
        {
          title: 'History',
          icon: 'mdi-file-document',
          path: '/stan/histori'
        },
      ]
    }
  },
  watch: {
    checkEnabled(newValue) {
      if (newValue) {
        this.checkUserData()
      }
    }
  },
  mounted() {
    if (this.checkEnabled) {
      this.checkUserData()
    }
  },
  methods: {
    checkUserData() {

      axios.get('/api/user').then(res => {
        const userId = res.data.data.Id
        var role = res.data.data.Role.role

        if (role === 'Siswa') {
          this.menuItems = this.menuItemsSiswa;
          axios.get('/api/siswa-by-user-id/' + userId, {
            validateStatus: function (status) {
              return status <= 500
            }
          }).then(siswaRes => {
            if (siswaRes.status === 500) {
              this.showAlert('siswa')
            }
          })
        } else if (role === 'Admin Stan') {
          this.menuItems = this.MenuItemStand;
          axios.get('/api/stan-by-user-id/' + userId, {
            validateStatus: function (status) {
              return status <= 500
            }
          }).then(stanRes => {
            if (stanRes.status === 500) {
              this.showAlert('stan')
            }
          })
        }
      })
    },
    showAlert(type) {
      const config = {
        title: type === 'siswa' ? 'Data Siswa Belum Lengkap' : 'Data Stan Belum Lengkap',
        text: type === 'siswa' ?
          'Anda perlu melengkapi data siswa untuk melanjutkan' :
          'Anda perlu melengkapi data stan untuk melanjutkan',
        icon: 'warning',
        showConfirmButton: true,
        confirmButtonText: 'Lengkapi Data',
        allowOutsideClick: false,
        allowEscapeKey: false,
        allowEnterKey: false
      }

      Swal.fire(config).then((result) => {
        if (result.isConfirmed) {
          this.$router.push(type === 'siswa' ? '/siswa/register' : '/stan/register')
        }
      })
    },
    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('role')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.v-app-bar {
  border-bottom: 1px solid #e0e0e0;
}

.v-list-item {
  cursor: pointer;
}

.v-list-item:hover {
  background-color: #f5f5f5;
}

/* Custom style untuk tombol di SweetAlert */
:deep(.swal2-styled) {
  margin: 5px !important;
  padding: 12px 25px !important;
}

:deep(#adminBtn) {
  background-color: #1976d2 !important;
}

:deep(#siswaBtn) {
  background-color: #1976d2 !important;
}

:deep(.swal2-popup) {
  padding: 2rem;
}
</style>

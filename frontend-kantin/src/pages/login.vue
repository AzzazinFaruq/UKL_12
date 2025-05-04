<template>
  <div class="container-login">
    <v-card elevation="0" class="card-login">
      <div class="d-flex justify-center">
        <v-img src="@/assets/register/logo-kantin.png" alt="logo" max-width="150"></v-img>
      </div>
      <div class="">
        <v-form>

          <label for="">Username</label>

          <v-text-field
            rounded="lg"
            color="primary"
            variant="outlined"
            placeholder="Masukkan Username"
            v-model="form.username"
            persistent-placeholder
          ></v-text-field>

          <label for="">Kata Sandi</label>
          <v-text-field
            rounded="lg"
            variant="outlined"
            color="primary"
            :type="showPassword ? 'text' : 'password'"
            placeholder="Masukkan Kata Sandi"
            v-model="form.password"
            persistent-placeholder
            :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append-inner="showPassword = !showPassword"
            :error-messages="passwordError"
          ></v-text-field>

          <!-- <v-label for="">Mendaftar Sebagai</v-label>

          <v-radio-group inline v-model="form.role_id" row>
            <v-radio label="Admin" :value="1"></v-radio>
            <v-radio label="Siswa" :value="2"></v-radio>
          </v-radio-group> -->

          <v-btn width="100%" color="#C37B58" @click="login">Daftar</v-btn>

        </v-form>
        <v-dialog v-model="dialog" width="400">
          <v-card class="pa-4">
            <v-card-title class="text-center">
              Pilih Role Anda
            </v-card-title>
            <v-card-text>
              <div class="d-flex flex-column gap-2">
                <v-btn
                  block
                  color="primary"
                  variant="elevated"
                  class="mb-2"
                  @click="confirmRole('admin')"
                >
                  Admin Stan
                </v-btn>
                <v-btn
                  block
                  color="primary"
                  variant="elevated"
                  @click="confirmRole('siswa')"
                >
                  Siswa
                </v-btn>
              </div>
            </v-card-text>
          </v-card>
        </v-dialog>
      </div>
      <div class="text-center mt-3">
        <p style="font-size: 14px; color: #595959;">Belum punya akun? mari kita <a href="/register"><b>daftar</b></a></p>
      </div>
    </v-card>
  </div>
</template>
<script>
import axios from "axios";
import Swal from 'sweetalert2'
export var success=false;
export default {
  data() {
    return {
      dialog: false,
      success:false,
      showPassword: false,
      form: {
        id: "",
        username: "",
        password: "",
        role_id: "",
      },
    };
  },
  computed: {
    passwordError() {
      if (this.form.password && this.form.password.length < 6) {
        return 'Username atau Password Salah'
      }
      return ''
    },
  },
  mounted(){
  },
  methods: {
    showErrorToast(message) {
      const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener('mouseenter', Swal.stopTimer)
          toast.addEventListener('mouseleave', Swal.resumeTimer)
        }
      })

      Toast.fire({
        icon: 'error',
        title: message,
        background: '#FFF0F0',
        color: '#DE3545'
      })
    },
    showSuccessToast(message) {
      const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener('mouseenter', Swal.stopTimer)
          toast.addEventListener('mouseleave', Swal.resumeTimer)
        }
      })

      Toast.fire({
        icon: 'success',
        title: message,
        background: '#F0FFF0',
        color: '#28A745'
      })
    },
    confirmRole(role) {
      const roleText = role === 'admin' ? 'Admin Stan' : 'Siswa'

      // Tutup dialog sebelum menampilkan SweetAlert
      this.dialog = false

      // Delay sedikit untuk animasi dialog menutup
      setTimeout(() => {
        Swal.fire({
          title: 'Konfirmasi Role',
          text: `Anda yakin ingin mendaftar sebagai ${roleText}?`,
          icon: 'question',
          showCancelButton: true,
          confirmButtonColor: '#C37B58',
          cancelButtonColor: '#d33',
          confirmButtonText: 'Ya, Saya Yakin',
          cancelButtonText: 'Batal'
        }).then((result) => {
          if (result.isConfirmed) {
            this.selectRole(role)
          } else {
            // Jika user membatalkan, tampilkan dialog pemilihan role kembali
            this.dialog = true
          }
        })
      }, 300) // Delay 300ms untuk animasi
    },
    selectRole(role) {
      if (role === 'admin') {
        axios.put(`/api/user-role/${this.form.id}`, { role_id: 1 }).then((res) => {
          console.log(res.data)
          this.showSuccessToast('Role berhasil diupdate')
          this.$router.push('/stan/register')
        })
      } else if (role === 'siswa') {
        axios.put(`/api/user-role/${this.form.id}`, { role_id: 2 }).then((res) => {
          console.log(res.data)
          this.showSuccessToast('Role berhasil diupdate')
          this.$router.push('/siswa/register')
        })
      }
    },
    login() {
      if (this.passwordError) {
        this.showErrorToast('Mohon periksa kembali kata sandi Anda')
        return
      }

      try {
        const loginData = {
          username: this.form.username,
          password: this.form.password,
        }

        axios.post("/login", loginData).then((res) => {
          axios.get(`api/user`).then((res) => {
            this.form.id = res.data.data.Id
          })
          this.showSuccessToast('Login Berhasil')
          localStorage.setItem("role", res.data.role)
          localStorage.setItem("authenticated", res.data.authenticated)

          if (res.data.role == "Admin Stan") {
            this.$router.push('/stan')
          } else if (res.data.role == "default") {
            this.dialog = true // Tampilkan dialog pilih role
          } else if (res.data.role == "Siswa") {
            this.$router.push('/siswa')
          }
        })
      } catch (error) {
        this.showErrorToast('Username atau Password Salah')
      }
    },
  },
};
</script>
<style lang="scss">

.v-btn{
  text-transform: none !important;
}

a{
  text-decoration: none;
  color: black;
}

label{
  font-weight: 600;
  font-size: 16px;
}

.container-login {
  width: 600px;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    @media screen and (max-width:600px) {
    width: 95%;
    }
    .card-login{
    padding: 32px;
    padding-right: 64px;
    padding-left: 64px;
    @media screen and (max-width:600px) {
      padding: 32px;
    }
    .image-login{
      margin-bottom: 32px;
    }
  }
}
.v-container{
  padding: 0px !important;
}
.no-underline {
  text-decoration: none !important;
  color: grey;
}
.no-underline:hover {
  color: black; /* Ganti dengan warna yang diinginkan */
}

/* Untuk border normal */
:deep(.v-field.v-field--outlined .v-field__outline) {
  border-width: 2px !important; /* sesuaikan ketebalan yang diinginkan */
}

/* Untuk border saat hover */
:deep(.v-field.v-field--outlined:hover .v-field__outline) {
  border-width: 2px !important;
}

/* Untuk border saat focus */
:deep(.v-field--focused .v-field__outline) {
  border-width: 2px !important;
}

/* Untuk border bottom */
:deep(.v-field.v-field--outlined .v-field__outline-bottom) {
  border-bottom-width: 2px !important;
}

/* Kustomisasi tambahan untuk toast */
.swal2-toast {
  padding: 10px !important;
  margin: 15px !important;
  width: auto !important;
  max-width: 300px !important;
  box-shadow: 0 3px 10px rgba(0,0,0,0.1) !important;
}

.swal2-toast .swal2-title {
  font-size: 14px !important;
  margin: 0 10px !important;
}

.swal2-toast .swal2-icon {
  margin: 0 !important;
  font-size: 12px !important;
}

/* Tambahkan style untuk success toast */
.swal2-toast.swal2-icon-success {
  background-color: #F0FFF0 !important;
}

.swal2-toast.swal2-icon-success .swal2-title {
  color: #28A745 !important;
}

.swal2-toast.swal2-icon-success .swal2-icon {
  color: #28A745 !important;
}

.v-card-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 1rem;
}

.v-btn {
  text-transform: none;
  font-weight: 500;
  height: 48px;
}

:deep(.swal2-popup) {
  padding: 2rem;
}

:deep(.swal2-confirm) {
  text-transform: none !important;
}

:deep(.swal2-cancel) {
  text-transform: none !important;
}
</style>

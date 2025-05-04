<template>
  <div class="wrapper">
    <div class="container-login">
      <v-card elevation="0" class="card-login">

        <!-- Foto Profile Preview -->
        <div class="text-center mt-5 mb-8">
          <v-avatar size="150" color="grey-lighten-2">
            <v-img
              v-if="imagePreview"
              :src="imagePreview"
              alt="Preview"
              cover
            ></v-img>
            <v-icon
              v-else
              size="100"
              color="grey-darken-2"
              icon="mdi-account-circle"
            ></v-icon>
          </v-avatar>

          <!-- Upload Button -->
          <div class="mt-3">
            <v-btn
              color="primary"
              size="small"
              @click="$refs.fileInput.click()"
            >
              Pilih Foto
            </v-btn>
            <input
              ref="fileInput"
              type="file"
              accept="image/*"
              class="d-none"
              @change="handleFileUpload"
            >
          </div>
        </div>

        <div class="">
          <v-form @submit.prevent="register" enctype="multipart/form-data">
            <label>Nama Stand</label>
            <v-text-field
              rounded="lg"
              color="primary"
              variant="outlined"
              placeholder="Masukkan Nama Stand"
              v-model="form.nama_stan"
              persistent-placeholder
            ></v-text-field>

            <label>No. Telepon</label>
            <v-text-field
              rounded="lg"
              color="primary"
              variant="outlined"
              placeholder="Masukkan No. Telepon"
              v-model="form.telp"
              persistent-placeholder
            ></v-text-field>

            <label>Alamat</label>
            <v-textarea
              rounded="lg"
              color="primary"
              variant="outlined"
              placeholder="Masukkan Alamat"
              v-model="form.nama_pemilik"
              persistent-placeholder
            ></v-textarea>

            <v-btn width="100%" color="#C37B58" @click="register" :loading="loading">
              Daftar
            </v-btn>
          </v-form>
        </div>
      </v-card>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Swal from 'sweetalert2'

export default {
  data() {
    return {
      loading: false,
      imagePreview: null,
      form: {
        nama_stan: "",
        telp: "",
        nama_pemilik:"",
        foto: null,
        user_id: "",
      },
    };
  },
  mounted(){
    axios.get('/api/user').then(response => {
      this.form.user_id = response.data.data.Id
      console.log(this.form.user_id)
    }).catch(error => {
      console.log(error)
    })
  },
  methods: {
    handleFileUpload(event) {
      const file = event.target.files[0]
      if (!file) return

      // Validasi ukuran file (2MB)
      if (file.size > 2000000) {
        this.showErrorToast('Ukuran foto harus kurang dari 2MB')
        return
      }

      // Validasi tipe file
      if (!file.type.includes('image/')) {
        this.showErrorToast('File harus berupa gambar')
        return
      }

      // Set file untuk form
      this.form.foto = file

      // Buat preview
      const reader = new FileReader()
      reader.onload = e => {
        this.imagePreview = e.target.result
      }
      reader.readAsDataURL(file)
    },

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

    async register() {
      try {
        this.loading = true

        // Validasi input
        if (!this.form.nama_pemilik || !this.form.telp || !this.form.nama_stan|| !this.form.foto) {
          this.showErrorToast('Semua field harus diisi')
          return
        }



        // Buat FormData object
        const formData = new FormData()
        formData.append('nama_siswa', this.form.nama_siswa)
        formData.append('telp', this.form.telp)
        formData.append('alamat', this.form.alamat)
        formData.append('foto', this.form.foto)
        formData.append('user_id', this.form.user_id)

        // Kirim request
        const response = await axios.post("/api/stan", formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })

        this.showSuccessToast('Pendaftaran berhasil')
        this.$router.push('/dashboard')

      } catch (error) {
        console.error('Error:', error)
        this.showErrorToast(error.response?.data?.error || 'Terjadi kesalahan')
      } finally {
        this.loading = false
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.wrapper {
  min-height: 100vh;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.container-login {
  width: 600px;
  @media screen and (max-width:600px) {
    width: 95%;
  }
}

.card-login {
  padding: 32px 64px;
  @media screen and (max-width:600px) {
    padding: 32px;
  }
}

label {
  font-weight: 600;
  font-size: 16px;
  display: block;
  margin-bottom: 8px;
}

.v-btn {
  margin-top: 16px;
  text-transform: none !important;
}

.v-avatar {
  border: 2px solid #C37B58;
}
</style>

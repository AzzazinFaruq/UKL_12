<template>
  <div>
    <v-container>
      <div class="my-5">
      </div>
      <div>
        <v-row>
          <v-col  v-for="(item,index) in menu.slice(0,3)" :key="item.id" cols="12" sm="6" md="3">
          <v-card v-if="index <3" rounded="lg" class="card-program" :elevation="elevation" @click="enter()">
            <div class="list-container">
              <img  :src="`http://localhost:6969/${item.foto}`" alt="" width="100%" class="list-image">
            </div>
            <div class="pa-3">
              <h3 class="">{{ item.nama_makanan }}</h3>
              <p class="description-text mb-3">{{ item.deskripsi }}</p>
              <p>21.000</p>
            </div>
          </v-card>
          <v-card v-else class="border-lg pa-4 card-program d-flex justify-center align-center">
            <a href="/program">
              <h2 style="color: #BF3232;">Lihat <br>Lebih <br>Banyak</h2>
            </a>
          </v-card>
        </v-col>
      </v-row>
      </div>

    </v-container>
  </div>
</template>
<script>
import axios from "axios";
export default {

  data() {
    return {
      status: false,
      menu:[],
      siswa:[],
      elevation:2
    }
  },
  mounted() {
    this.CheckSiswa()
    this.getMenu()
  },
  methods: {
    enter(){
    },
    getMenu() {
      axios.get('/api/menu').then((res) => {
        this.menu = res.data.data
        console.log(this.menu)
      })
    },
    CheckSiswa() {
      axios.get('/api/user').then((res) => {
        axios.get('/api/siswa-by-user-id/' + res.data.data.Id).then((res) => {
          this.status = true
          this.siswa = res.data.data
          console.log(this.siswa)
        }).catch((err) => {
          this.status = false
        })
      })
    }
  }

}
</script>
<style lang="scss">
.list-container{
  height: 230px;
  .list-image{
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
}
</style>

<template>
  <div class="home">
    <v-parallax
    height="300"
      dark
      src="https://cdn.vuetifyjs.com/images/parallax/material.jpg"
    >
      <v-row
        align="center"
        justify="center"
      >
        <v-col class="text-center" cols="12">
          <h1 class="display-1 font-weight-thin mb-4">Netflix Info</h1>
        </v-col>
      </v-row>
    </v-parallax>
    <v-container>
      <v-row justify-center>
        <v-col xs="6" sm="6" md="6" lg="10" xl="10">
          <Search v-model="searchText"/>
        </v-col>
        <v-col xs="6" sm="6" md="6" lg="2" xl="2">
          <v-btn
            class="ma-2"
            tile color="primary"
            @click="searchNetflix"
            large
            :disabled="(!searchText || searchText.length < 3) ? true : false">
              Search
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <program-list-skeleton v-if="isLoading"></program-list-skeleton>
    <netflix-info-list :list="result" v-else ></netflix-info-list>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios';
import Search from '@/components/Search.vue';
import NetflixInfoList from '@/components/NetflixInfoList.vue';
import ProgramListSkeleton from '@/components/ProgramListSkeleton.vue';

// import { link } from '@/helpers/http-common';

export default {
  name: 'netflixinfo',
  data() {
    return {
      isLoading: false,
      apiURL: 'localhost:8080/api/unogs/search',
      searchText: '',
      result: [],
    };
  },
  components: {
    Search,
    NetflixInfoList,
    ProgramListSkeleton,
  },
  mounted() {
  },
  methods: {
    async searchNetflix() {
      try {
        const link = axios.create({
          baseURL: `${process.env.VUE_APP_API_URL}`,
        });
        this.isLoading = true;
        const response = await link.get(`/api/unogs/search?title=${this.searchText}&skip=0&limit=5`);
        console.log(response);
        this.result = response.data.data.results;
        this.isLoading = false;
      } catch (error) {
        this.isLoading = false;
        console.log(error);
      }
    },
  },
  watch: {},
};
</script>

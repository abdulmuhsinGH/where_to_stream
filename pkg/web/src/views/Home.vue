<template>
  <div class="home">
    <v-parallax
      height="300"
      dark
      src="https://cdn.vuetifyjs.com/images/parallax/material2.jpg"
    >
      <v-row
        align="center"
        justify="center"
      >
        <v-col class="text-center" cols="12">
          <h1 class="display-1 font-weight-thin mb-4">
            Amazon Prime and Netflix Info(US and UK locations)</h1>
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
            @click="searchUtelly"
            large
            :disabled="(!searchText || searchText.length < 3) ? true : false">
              Search
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <program-list-skeleton v-if="isLoading"></program-list-skeleton>
    <program-list :list="result" v-else ></program-list>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios';
import Search from '@/components/Search.vue';
import ProgramList from '@/components/ProgramList.vue';
import ProgramListSkeleton from '@/components/ProgramListSkeleton.vue';

// import { link } from '@/helpers/http-common';

export default {
  name: 'home',
  data() {
    return {
      isLoading: false,
      apiURL: 'localhost:8080/api/utelly/search',
      searchText: '',
      result: [],
    };
  },
  components: {
    Search,
    ProgramList,
    ProgramListSkeleton,
  },
  mounted() {
  },
  methods: {
    async searchUtelly() {
      try {
        const link = axios.create({
          baseURL: `${process.env.VUE_APP_API_URL}`,
        });
        this.isLoading = true;
        const response = await link.get(`/api/utelly/search?program=${this.searchText}`);
        // console.log(response);
        this.result = response.data.data.results;
        this.isLoading = false;
      } catch (error) {
        this.isLoading = false;
        // console.log(error);
      }
    },
  },
  watch: {},
};
</script>

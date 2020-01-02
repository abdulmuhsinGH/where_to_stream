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
          <v-alert
            v-model="showAlert"
            dense
            dismissible
            type=alertType
          >
            {{alertMessage}}
          </v-alert>
        </v-col>
        <v-col xs="6" sm="6" md="6" lg="2" xl="2">
          <v-btn
            class="ma-2"
            tile color="primary"
            @click="searchNetflix"
            v-on:keyup.enter="searchNetflix"
            large
            :disabled="(!searchText || searchText.length < 3) ? true : false">
              Search
          </v-btn>
        </v-col>
      </v-row>
      <v-row justify-center>
        <v-col v-if="result.length > 0" align-self="center" xs="12" sm="12" md="12" lg="12" xl="12">
          <v-btn color="success" @click="searchNetflix">See More...</v-btn>
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
      showAlert: false,
      alertMessage: '',
      alertType: '',
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
        if (this.searchText.length <= 0) {
          this.setAlertInfo('Please enter a tv show or movie', 'warning');
          return;
        }
        const link = axios.create({
          baseURL: `${process.env.VUE_APP_API_URL}`,
        });
        this.isLoading = true;
        const response = await link.get(`/api/unogs/search?title=${this.searchText}&skip=${this.result.length}&limit=4`);
        this.result = response.data.data;
        if (this.result.length <= 0) {
          this.setAlertInfo(`No results found for "${this.searchText}"; try another movie or tv show`, 'warning');
        }
        this.isLoading = false;
      } catch (error) {
        this.isLoading = false;
        this.setAlertInfo('Oops Something Went Wrong :(', 'error');
      }
    },
    setAlertInfo(alertMessage, alertType) {
      console.log('hello');
      this.alertType = alertType;
      this.showAlert = true;
      this.alertMessage = alertMessage;
    },
  },
  watch: {},
};
</script>

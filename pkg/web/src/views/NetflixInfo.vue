<template>
  <div class="home">
    <v-parallax
    height="300"
      dark
      :src="require('@/assets/material.webp')"
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
            @click="newSearchNetflix"
            large
            :disabled="(!searchText || searchText.length < 3) ? true : false">
              Search
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <netflix-info-list :list="result"></netflix-info-list>
    <program-list-skeleton v-if="isLoading"></program-list-skeleton>
    <v-container grid-list-xs>
      <v-row justify-center>
        <v-col v-if="result.length > 0" align-self="center" xs="12" sm="12" md="12" lg="12" xl="12">
          <v-btn color="success" @click="searchNetflix">See More...</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
// @ is an alias to /src
import Search from '@/components/Search.vue';
import NetflixInfoList from '@/components/NetflixInfoList.vue';
import ProgramListSkeleton from '@/components/ProgramListSkeleton.vue';

import link from '@/helpers/http-common';

export default {
  name: 'netflixinfo',
  data() {
    return {
      isLoading: false,
      searchText: '',
      result: [],
      showAlert: false,
      alertMessage: '',
      alertType: '',
      appUrl: link,
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
        this.isLoading = true;
        const response = await this.appUrl.get(`/api/unogs/search?title=${this.searchText}&skip=${this.result.length}&limit=4`);
        await this.populateResults(response.data.data);
        this.hasResults();
        this.isLoading = false;
      } catch (error) {
        console.log(error);
        this.isLoading = false;
        this.setAlertInfo('Oops Something Went Wrong :(', 'error');
      }
    },
    async newSearchNetflix() {
      this.result = [];
      await this.searchNetflix();
    },
    setAlertInfo(alertMessage, alertType) {
      this.alertType = alertType;
      this.showAlert = true;
      this.alertMessage = alertMessage;
    },
    hasResults() {
      if (this.result.length <= 0) {
        this.setAlertInfo(`No results found for "${this.searchText}"; try another movie or tv show`, 'warning');
      }
    },
    async populateResults(moviesAndOrTvShows) {
      const vm = this;
      if (moviesAndOrTvShows) {
        await moviesAndOrTvShows.forEach(movieAndOrTvShow => vm.result.push(movieAndOrTvShow));
      }
    },
  },
};
</script>

<template>
  <div class="home">
    <v-parallax
      height="300"
      dark
      :src="require('@/assets/material2.jpg')"
    >
      <v-row
        align="center"
        justify="center"
      >
        <v-col class="text-center" cols="12">
          <h1 class="display-1 font-weight-thin mb-4">
            Prime Video and Netflix Info(US and UK only)</h1>
        </v-col>
      </v-row>
    </v-parallax>
    <v-container>
      <v-row justify-center>
        <v-col xs="6" sm="6" md="6" lg="10" xl="10">
          <SearchMedia v-model="searchText"/>
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
    <prime-video-and-netflix-list :list="result" v-else ></prime-video-and-netflix-list>
  </div>
</template>

<script>
// @ is an alias to /src
import SearchMedia from '@/components/Search.vue';
import PrimeVideoAndNetflixList from '@/components/PrimeVideoAndNetflixList.vue';
import ProgramListSkeleton from '@/components/ProgramListSkeleton.vue';

import link from '@/helpers/http-common';

export default {
  name: 'prime-video-and-netflix',
  data() {
    return {
      isLoading: false,
      searchText: '',
      result: [],
    };
  },
  components: {
    SearchMedia,
    PrimeVideoAndNetflixList,
    ProgramListSkeleton,
  },
  mounted() {
  },
  methods: {
    async searchUtelly() {
      try {
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

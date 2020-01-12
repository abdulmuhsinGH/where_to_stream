<template>
  <v-container>
      <v-row >
        <v-col v-bind:key="item.weight" v-for="item in list" xs="12" sm="12" md="4" lg="3" xl="3">
          <PrimeVideoAndNetflixCard :data="item" />
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
import PrimeVideoAndNetflixCard from '@/components/PrimeVideoAndNetflixCard.vue';

export default {
  name: 'ProgramList',
  props: {
    list: Array,
  },
  components: {
    PrimeVideoAndNetflixCard,
  },
};
</script>

<style>

</style>

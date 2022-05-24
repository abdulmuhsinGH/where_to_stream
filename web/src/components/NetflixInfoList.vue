<template>
  <v-container>
      <v-row >
        <v-col v-bind:key="item.netflix_id"
          v-for="item in list" xs="12" sm="12" md="4" lg="4" xl="4">
          <NetflixInfoCard :data="item" />
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
import NetflixInfoCard from '@/components/NetflixInfoCard.vue';

export default {
  name: 'NetflixInfoList',
  props: {
    list: Array,
  },
  components: {
    NetflixInfoCard,
  },
  data() {
    return {
      count: 0,
    };
  },
};
</script>

<style>

</style>

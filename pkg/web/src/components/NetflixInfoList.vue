<template>
  <v-container>
      <v-row >
        <v-col v-bind:key="item.RESULT.nfinfo.netflixid"
          v-for="item in list" xs="12" sm="12" md="6" lg="6" xl="6">
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

<template>
  <v-container>
      <v-row >
        <v-col v-bind:key="item.weight" v-for="item in list" xs="12" sm="12" md="4" lg="3" xl="3">
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
};
</script>

<style>

</style>

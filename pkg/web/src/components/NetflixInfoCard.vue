<template>
  <v-card
    class="mx-auto"
    max-width="400"
  >
    <v-img
      :alt="data.RESULT.nfinfo.title"
      class="white--text align-end"
      height="200px"
      :src="data.RESULT.nfinfo.image1"
    >
      <v-card-title>{{data.RESULT.nfinfo.title}}</v-card-title>
    </v-img>
    <v-list class="transparent">
      <v-list-item
        v-for="item in data.RESULT.country"
        :key="item.cid"
      >
        <v-list-item-subtitle class="text-right">
          {{ item.country }}
        </v-list-item-subtitle>
      </v-list-item>
    </v-list>
    <v-divider></v-divider>
    <v-card-actions>
      <v-btn
        color="orange"
        text
      >
        See More Details
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'NetflixCard',
  props: {
    data: Object,
  },
};
</script>

<style>

</style>

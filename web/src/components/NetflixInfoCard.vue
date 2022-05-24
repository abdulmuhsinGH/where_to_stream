<template>
  <v-card
    class="mx-auto"
    outlined
  >
    <v-img
      :alt="data.title"
      :aspect-ratio="16/9"
      class="white--text align-end"
      :src="data.img"
    >
      <v-card-title>{{data.title}}</v-card-title>
    </v-img>
    <v-divider></v-divider>
    <!-- <v-card-text>
      <span class="subheading">Available In:</span>
      <v-chip-group
        v-model="selection"
        active-class="deep-purple--text text--accent-4"
      >
        <v-chip v-for="item in data"
          :key="item.cid">
          {{ item.country }}
          <flag :iso="item.ccode" />
        </v-chip>
      </v-chip-group>
    </v-card-text> -->
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

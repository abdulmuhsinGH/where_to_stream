<template>
  <v-card
    class="mx-auto"
    max-width="400"
  >
    <v-img
      :alt="data.name"
      class="white--text align-end"
      height="200px"
      :src="data.picture"
    >
      <v-card-title>{{data.name}}</v-card-title>
    </v-img>
    <v-list class="transparent">
      <v-list-item
        v-for="item in data.Locations"
        :key="item.id"
      >
        <v-list-item-title><v-img :src="item.icon"></v-img></v-list-item-title>
        <v-list-item-subtitle class="text-right">
          {{ item.display_name }}
        </v-list-item-subtitle>
        <v-list-item-icon>
          <v-btn text icon color :href="item.url">
              <v-icon small>fa-external-link-alt</v-icon>
          </v-btn>
        </v-list-item-icon>
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
  name: 'ProgramCard',
  props: {
    data: Object,
  },
};
</script>

<style>

</style>

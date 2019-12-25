<template>
  <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Where To Stream</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn text small to="/netflixinfo">Netflix Info</v-btn>
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn text small v-on="on" to="/">*Amazon Prime and Netflix</v-btn>
        </template>
        <span>Only movies and TV shows in US and UK</span>
      </v-tooltip>
    </v-app-bar>
</template>

<script>
export default {
  name: 'Header',

};
</script>

<style></style>

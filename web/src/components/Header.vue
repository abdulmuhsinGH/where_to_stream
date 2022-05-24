<template>
  <v-app-bar app color="indigo" dark>
      <v-toolbar-title>Where To Stream</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-tooltip bottom>
        <template v-slot:activator="{ on }">
          <v-btn text small v-on="on" to="/">*Prime Video and Netflix</v-btn>
        </template>
        <span>Only TV Shows and movies showing in US and/or UK</span>
      </v-tooltip>
      <v-btn text small to="/netflixinfo">Netflix Info</v-btn>
    </v-app-bar>
</template>

<script>
export default {
  name: 'AppHeader',

};
</script>

<style></style>

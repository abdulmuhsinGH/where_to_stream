<template>
  <v-container>
      <v-row>
        <v-col xs="12" sm="12" md="4" lg="4" xl="4">
          <v-sheet
          :color="`grey ${theme.isDark ? 'darken-2' : 'lighten-4'}`"
          class="px-3 pt-3 pb-3"
        >
          <v-skeleton-loader
            class="mx-auto"
            max-width="300"
            type="card"
          ></v-skeleton-loader>
        </v-sheet>
        </v-col>
        <v-col xs="12" sm="12" md="4" lg="4" xl="4">
          <v-sheet
          :color="`grey ${theme.isDark ? 'darken-2' : 'lighten-4'}`"
          class="px-3 pt-3 pb-3"
        >
          <v-skeleton-loader
            class="mx-auto"
            max-width="300"
            type="card"
          ></v-skeleton-loader>
        </v-sheet>
        </v-col>
        <v-col xs="12" sm="12" md="4" lg="4" xl="4">
          <v-sheet
          :color="`grey ${theme.isDark ? 'darken-2' : 'lighten-4'}`"
          class="px-3 pt-3 pb-3"
        >
          <v-skeleton-loader
            class="mx-auto"
            max-width="300"
            type="card"
          ></v-skeleton-loader>
        </v-sheet>
        </v-col>
      </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'ProgramListSkeleton',
  inject: ['theme'],
};
</script>

<style>

</style>

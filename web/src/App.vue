<template>
  <v-app>
    <AppHeader/>
    <v-main>
      <router-view></router-view>
    </v-main>
    <AppFooter/>
  </v-app>
</template>

<script>
import AppHeader from './components/Header.vue';
import AppFooter from './components/Footer.vue';

export default {
  name: 'App',

  components: {
    AppHeader,
    AppFooter,
  },

  data: () => ({
    //
  }),
};
</script>

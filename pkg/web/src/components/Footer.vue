<template>
  <v-footer
    dark
    padless
  >
    <v-card
      tile
      class="flex indigo lighten-1 white--text text-center"
    >
      <v-card-text>
        <v-btn
          v-for="contact in contacts"
          :key="contact.icon"
          class="mx-4 white--text"
          icon
          :href="contact.url"
        >
          <v-icon size="24px">{{ contact.icon }}</v-icon>
        </v-btn>
      </v-card-text>
      <v-divider></v-divider>

      <v-card-text class="white--text">
        &copy; {{ new Date().getFullYear() }}<strong><v-btn text small href="https://abdulmuhsin.dev">Abdul Muhsin Iddris</v-btn></strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  name: 'Footer',
  data() {
    return {
      contacts: [
        {
          icon: 'fab fa-twitter',
          url: 'https://twitter.com/abdul_muhsin_',
        },
        {
          icon: 'fab fa-linkedin',
          url: 'https://www.linkedin.com/in/abdulmuhsin/',
        },
        {
          icon: 'fab fa-github',
          url: 'https://github.com/abdulmuhsinGH',
        },
        /* {
          icon: 'fab fa-user',
          url: 'https://abdulmuhsin.dev'
        }, */
      ],
    };
  },
};
</script>

<style></style>

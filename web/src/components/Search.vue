<template>
  <v-container>
    <v-layout
      text-center
      wrap
    >
      <v-flex xs12>
        <v-text-field
          dense
          outlined
          clearable
          label="Search TV Show/Movie"
          v-bind:value="value"
          v-on:input="$emit('input', $event)"
        >
        </v-text-field>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'SearchMedia',
  props: ['value'],
  data: () => ({}),
};
</script>

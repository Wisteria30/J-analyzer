<template>
  <v-row>
    <v-col v-for="n in 9" :key="n" class="d-flex child-flex" cols="4">
      <v-img
        :src="`https://i.gyazo.com/08f8fda49373fcada6f28e4cd0c3dbb7.jpg`"
        aspect-ratio="1"
        class="grey lighten-2"
      >
        <template #placeholder>
          <v-row class="fill-height ma-0" align="center" justify="center">
            <v-progress-circular
              indeterminate
              color="grey lighten-5"
            ></v-progress-circular>
          </v-row>
        </template>
      </v-img>
    </v-col>
  </v-row>
</template>

<script>
export default {
  computed: {
    images() { return this.$store.getters.getImages }
  }
}
</script>

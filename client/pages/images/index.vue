<template>
  <v-row>
    <v-col
      v-for="n in images"
      :key="n.image_id"
      class="d-flex child-flex"
      cols="4"
    >
      <nuxt-link :to="`/images/${n.image_id}`">
        <v-img :src="`${n.url}`" aspect-ratio="1" class="grey lighten-2">
          <template #placeholder>
            <v-row class="fill-height ma-0" align="center" justify="center">
              <v-progress-circular
                indeterminate
                color="grey lighten-5"
              ></v-progress-circular>
            </v-row>
          </template>
        </v-img>
      </nuxt-link>
    </v-col>
  </v-row>
</template>

<script>
import ROUTES from '~/routes/api'
export default {
  middleware: 'authenticated',
  async fetch() {
    const payload = {
      uri: ROUTES.GET.LIST,
    }
    await this.$store.dispatch('getList', payload)
  },
  computed: {
    images() {
      return this.$store.getters.list
    },
  },
}
</script>

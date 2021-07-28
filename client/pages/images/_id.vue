<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-img :src="`${item.url}`" class="grey lighten-2">
        <template #placeholder>
          <v-row class="fill-height ma-0" align="center" justify="center">
            <v-progress-circular
              indeterminate
              color="grey lighten-5"
            ></v-progress-circular>
          </v-row>
        </template>
      </v-img>
      <v-card>
        <v-card-title class="headline">
          {{ item.metadata.title }}
        </v-card-title>
        <v-card-text>
          <hr class="my-3" />
          <p>{{ new Date(item.created_at) }}</p>
          <p>App: {{ item.metadata.app }}</p>
          <p>From: {{ item.metadata.title }}</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="info" nuxt @click="Analyze"> 画像の解析 </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import ROUTES from '~/routes/api'

export default {
  computed: {
    item() {
      return this.$store.getters.image
    },
  },
  async fetch({ store, route }) {
    await store.dispatch('findImage', {
      uri: ROUTES.GET.IMAGE.replace(':id', route.params.id),
    })
  },
  methods: {
    Analyze() {
      console.log('Analyzeする')
      // const payload = {
      //   uri: ROUTES.POST.ANALYZE,
      // }
      // this.$store.dispatch('imageAnalyze', payload)
    },
  },
}
</script>

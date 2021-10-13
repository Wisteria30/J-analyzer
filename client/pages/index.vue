<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card class="logo py-4 d-flex justify-center">
        <Itihime />
        <GyazoLogo />
      </v-card>
      <v-card>
        <v-card-title class="headline"> Jantama Analyzer </v-card-title>
        <v-card-text>
          <p>雀魂のプレイ画面を読み込んで、盤面から何切るを解析する</p>
          <p>Gyazoと連携します</p>
          <hr class="my-3" />
          <a
            href="https://game.mahjongsoul.com/"
            target="_blank"
            rel="noopener noreferrer"
          >
            雀魂
          </a>
          <br />
          <a
            href="https://gyazo.com/"
            target="_blank"
            rel="noopener noreferrer"
          >
            Gyazo
          </a>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="accent" to="/howto"> 使い方 </v-btn>
          <v-btn color="primary" nuxt @click="Authorize">
            Gyazoにログイン
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import ROUTES from '~/routes/api'
export default {
  middleware: 'authenticated',
  methods: {
    Authorize() {
      const payload = {
        uri: ROUTES.GET.AUTH,
      }
      this.$store.dispatch('userAuthorization', payload)
    },
  },
}
</script>

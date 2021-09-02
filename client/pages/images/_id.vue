<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="10">
      <v-img :src="`${image.url}`" class="grey lighten-2">
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
          {{ image.metadata.title }}
        </v-card-title>
        <v-card-text>
          <hr class="my-3" />
          <p>{{ new Date(image.created_at) }}</p>
          <p>App: {{ image.metadata.app }}</p>
          <p>From: {{ image.metadata.title }}</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn :disabled="isJuntama" color="info" nuxt @click="Analyze">
            画像の解析
          </v-btn>
        </v-card-actions>
      </v-card>
      <template v-if="success">
        <hr class="my-3" />
        <hr class="my-3" />
        <div class="request">
          <v-card class="mx-auto">
            <v-card-text>
              <p class="text-h4 text--primary">識別結果</p>
              <div class="text--primary">
                <h3>場風: {{ bakaze }}</h3>
                <br />
                <h3>自風: {{ zikaze }}</h3>
                <br />
                <h3>巡目: {{ request.turn }}</h3>
                <br />
                <h3>ドラ表示牌:</h3>
                <br />
                <v-row no-gutters>
                  <div
                    v-for="(n, index) in request.dora_indicators"
                    :key="index"
                  >
                    <v-img
                      :src="require(`@/assets/pai/${n}.png`)"
                      :aspect-ratio="32 / 45"
                      width="45"
                    ></v-img>
                  </div>
                </v-row>
                <br />
                <h3>手牌:</h3>
                <br />
                <v-row no-gutters>
                  <div v-for="(n, index) in request.hand_tiles" :key="index">
                    <v-img
                      :src="require(`@/assets/pai/${n}.png`)"
                      :aspect-ratio="32 / 45"
                      width="45"
                    ></v-img>
                  </div>
                </v-row>
              </div>
            </v-card-text>
          </v-card>
        </div>
        <hr class="my-3" />
        <hr class="my-3" />
        <div class="request">
          <v-card class="mx-auto" max-width="1200">
            <v-card-text>
              <p class="text-h4 text--primary">解析結果</p>
              <div class="text--primary">
                <h3>向聴(シャンテン)数: {{ response.syanten }}</h3>
                <br />
              </div>
              <v-data-table
                :headers="headers"
                :items="items"
                :items-per-page="15"
                hide-default-footer
                item-key="tile"
                :sort-by="['exp_value']"
                :sort-desc="[true]"
                class="elevation-1"
              >
                <template #[`item.tile`]="{ item }">
                  <v-img
                    :src="require(`@/assets/pai/${item.tile}.png`)"
                    :aspect-ratio="32 / 45"
                    width="45"
                  ></v-img>
                </template>
                <template #[`item.yuko`]="{ item }">
                  <v-row no-gutters>
                    <div v-for="(n, index) in item.yuko" :key="index">
                      <v-img
                        :src="require(`@/assets/pai/${n}.png`)"
                        :aspect-ratio="32 / 45"
                        width="30"
                      ></v-img>
                    </div>
                  </v-row>
                </template>
                <template #[`item.exp_value`]="{ item }">
                  {{ `${item.exp_value}点` }}
                </template>
                <template #[`item.win_prob`]="{ item }">
                  {{ `${item.win_prob}%` }}
                </template>
                <template #[`item.tenpai_prob`]="{ item }">
                  {{ `${item.tenpai_prob}%` }}
                </template>
              </v-data-table>
            </v-card-text>
          </v-card>
        </div>
      </template>
      <template v-if="loading">
        <v-data-table
          class="elevation-1"
          hide-default-footer
          loading
          loading-text="Loading... Please wait"
        ></v-data-table>
      </template>
    </v-col>
  </v-row>
</template>

<script>
import ROUTES from '~/routes/api'

export default {
  async fetch({ store, route }) {
    await store.dispatch('findImage', {
      uri: ROUTES.GET.IMAGE.replace(':id', route.params.id),
    })
    store.dispatch('resetResponse', {})
  },
  computed: {
    image() {
      return this.$store.getters.image
    },
    success() {
      return this.$store.getters.analysisSuccess
    },
    bakaze() {
      if (this.$store.getters.analysisRequest.bakaze === 27) {
        return '東'
      } else if (this.$store.getters.analysisRequest.bakaze === 28) {
        return '南'
      } else if (this.$store.getters.analysisRequest.bakaze === 29) {
        return '西'
      } else if (this.$store.getters.analysisRequest.bakaze === 30) {
        return '北'
      } else {
        return ''
      }
    },
    zikaze() {
      if (this.$store.getters.analysisRequest.zikaze === 27) {
        return '東'
      } else if (this.$store.getters.analysisRequest.zikaze === 28) {
        return '南'
      } else if (this.$store.getters.analysisRequest.zikaze === 29) {
        return '西'
      } else if (this.$store.getters.analysisRequest.zikaze === 30) {
        return '北'
      } else {
        return ''
      }
    },
    headers() {
      return [
        {
          text: '打牌',
          value: 'tile',
        },
        {
          text: '受入枚数',
          value: 'ukeire',
        },
        {
          text: '有効牌',
          value: 'yuko',
        },
        {
          text: '期待値',
          value: 'exp_value',
        },
        {
          text: '和了確率',
          value: 'win_prob',
        },
        {
          text: '聴牌確率',
          value: 'tenpai_prob',
        },
      ]
    },
    items() {
      if (this.$store.getters.analysisResponse.candidates === null) {
        return []
      }
      const ret = []
      this.$store.getters.analysisResponse.candidates.forEach((candidate) => {
        let sum = 0
        const yuko = []
        candidate.required_tiles.forEach((required) => {
          sum += required.count
          yuko.push(required.tile)
        })
        ret.push({
          tile: candidate.tile,
          ukeire: `${candidate.required_tiles.length}種${sum}枚`,
          yuko: yuko.sort((a, b) => {
            return a - b
          }),
          exp_value: candidate.exp_values
            ? Math.round(candidate.exp_values[0])
            : 0,
          win_prob: candidate.win_probs
            ? Math.round(candidate.win_probs[0] * 10000) / 100
            : 0,
          tenpai_prob: candidate.tenpai_probs
            ? Math.round(candidate.tenpai_probs[0] * 10000) / 100
            : 0,
        })
      })
      return ret
    },
    request() {
      return this.$store.getters.analysisRequest
    },
    response() {
      return this.$store.getters.analysisResponse
    },
    isJuntama() {
      return (
        this.$store.getters.image.metadata.title !==
        '雀魂 -じゃんたま-| 麻雀を無料で気軽に'
      )
    },
    loading() {
      return this.$store.getters.loading
    },
  },
  methods: {
    Analyze() {
      const payload = {
        uri: ROUTES.GET.ANALYZE,
        params: {
          img_url: this.$store.getters.image.url,
        },
      }
      this.$store.dispatch('imageAnalyze', payload)
    },
  },
}
</script>

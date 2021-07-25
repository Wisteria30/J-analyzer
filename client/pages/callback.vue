<template>
  <div>
    <h1>コールバック</h1>
    <p>しばらくお待ちください・・・</p>
  </div>
</template>

<script>
import ROUTES from '~/routes/api'
export default {
  async mounted() {
    const payload = {
      uri: ROUTES.GET.TOKEN,
      state: this.$route.query.state,
      code: this.$route.query.code,
    }
    await this.$store.dispatch('getToken', payload)
  },
}
</script>

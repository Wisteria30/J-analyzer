import { createRequestClient } from './request-client'

export const state = () => ({
  images: [],
  meta: {},
  token: '',
})

export const actions = {
  async hello({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateHello', res)
    this.app.router.push('/list')
  },
  async getImages({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateGetImages', res)
  },
  async setToken({ commit }, payload) {
    await this.$cookie.set('access_token', payload)
    commit('mutateToken', payload)
  },
}

export const mutations = {
  mutateGetImages(state, payload) {
    state.images = payload.Images
    state.meta = payload.Meta
  },
  mutateHello(state, payload) {
    console.log(payload)
  },
  mutateToken(state, payload) {
    state.token = payload
  },
}

export const getters = {
  getImages(state) {
    return state.images
  },
  getMeta(state) {
    return state.meta
  },
  IsAuth(state) {
    return !!state.images.length
  },
}

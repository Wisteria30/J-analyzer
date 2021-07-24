import { createRequestClient } from './request-client'

export const state = () => ({
  images: [],
  meta: {},
  token: '',
})

export const actions = {
  async userAuthorization({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    this.$cookies.set('state', res.state)
    window.location.replace(res.url)
  },
  async getToken({ commit }, payload) {
    if (
      payload.state !== undefined &&
      payload.code !== undefined &&
      payload.state === this.$cookies.get('state')
    ) {
      const client = createRequestClient(this.$axios)
      const res = await client.get(payload.uri, { code: payload.code })
      this.$cookies.set('jwt_token', res)
      commit('mutateToken', res)
      this.app.router.push('/')
    }
  },
  async getImages({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateGetImages', res)
  },
  setToken({ commit }, payload) {
    commit('mutateToken', payload)
  },
}

export const mutations = {
  mutateToken(state, payload) {
    state.token = payload
  },
  mutateGetImages(state, payload) {
    state.images = payload.Images
    state.meta = payload.Meta
  },
}

export const getters = {
  getToken(state) {
    return state.token
  },
  getImages(state) {
    return state.images
  },
  getMeta(state) {
    return state.meta
  },
  isLoggedIn(state) {
    return !!state.token
  },
}

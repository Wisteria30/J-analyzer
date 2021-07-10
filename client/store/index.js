import { createRequestClient } from './request-client'

export const state = () => ({
  images: [],
  meta: {},
  isAuth: false,
  token: '',
})

export const actions = {
  async userAuthorization({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateAuthorization', res)
  },
  async hello({ commit }, payload) {
    const client = createRequestClient(this.$axios)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateHello', res)
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
  mutateAuthorization(state, payload) {
    state.isAuth = Boolean(payload)
  },
  mutateGetImages(state, payload) {
    state.images = payload.Images
    state.meta = payload.Meta
    console.log(state.images)
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
  getIsAuth(state) {
    return state.isAuth
  },
}

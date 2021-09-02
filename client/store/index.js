import firebase from '@/plugins/firebase'
import { createRequestClient } from './request-client'

export const state = () => ({
  images: [],
  image: {},
  meta: {},
  token: '',
  analysisRequest: {},
  analysisResponse: {},
  analysisSuccess: false,
  loading: false,
})

export const actions = {
  async userAuthorization({ commit }, payload) {
    const client = createRequestClient(this.$axios, this.$cookies, this)
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
      const client = createRequestClient(this.$axios, this.$cookies, this)
      const customToken = await client.get(payload.uri, { code: payload.code })

      const res = await firebase.auth().signInWithCustomToken(customToken)
      const token = await res.user.getIdToken()
      this.$cookies.set('jwt_token', token)
      const refreshToken = res.user.refreshToken
      this.$cookies.set('refresh_token', refreshToken)
      commit('mutateToken', token)
      this.app.router.push('/')
    }
  },
  async getList({ commit }, payload) {
    const client = createRequestClient(this.$axios, this.$cookies, this)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateList', res)
  },
  setToken({ commit }, payload) {
    this.$cookies.set('jwt_token', payload)
    commit('mutateToken', payload)
  },
  async findImage({ commit }, payload) {
    const client = createRequestClient(this.$axios, this.$cookies, this)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateImage', res)
  },
  async imageAnalyze({ commit }, payload) {
    commit('mutateLoading', { loading: true })
    const client = createRequestClient(this.$axios, this.$cookies, this)
    const res = await client.get(payload.uri, payload.params)
    commit('mutateAnalysis', res)
    commit('mutateLoading', { loading: false })
  },
  resetResponse({ commit }, payload) {
    commit('mutateLoading', { loading: false })
    commit('mutateAnalysis', {
      success: null,
      request: {},
      response: {},
    })
  },
}

export const mutations = {
  mutateLoading(state, payload) {
    state.loading = payload.loading
  },
  mutateToken(state, payload) {
    state.token = payload
  },
  mutateList(state, payload) {
    state.images = payload.Images
    state.meta = payload.Meta
  },
  mutateImage(state, payload) {
    state.image = payload
  },
  mutateAnalysis(state, payload) {
    state.analysisSuccess = payload.success
    state.analysisRequest = payload.request
    state.analysisResponse = payload.response
  },
}

export const getters = {
  image(state) {
    return state.image
  },
  analysisRequest(state) {
    return state.analysisRequest
  },
  analysisResponse(state) {
    return state.analysisResponse
  },
  analysisSuccess(state) {
    return state.analysisSuccess
  },
  token(state) {
    return state.token
  },
  list(state) {
    return state.images
  },
  getMeta(state) {
    return state.meta
  },
  isLoggedIn(state) {
    return !!state.token
  },
  loading(state) {
    return state.loading
  },
}

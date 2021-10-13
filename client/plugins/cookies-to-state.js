export default ({app, store}) => {
    const token = app.$cookies.get('jwt_token')
    if (token) {
        store.dispatch('setToken', token)
    }
    const refreshToken = app.$cookies.get('refresh_token')
    if (refreshToken) {
        store.dispatch('setRefreshToken', refreshToken)
    }
}

export default ({app, store}) => {
    const token = app.$cookies.get('access_token')
    if (token) {
        store.dispatch('setToken', token)
    }
}
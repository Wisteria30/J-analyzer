export default ({ $axios, app }) => {
  $axios.onRequest((config) => {
    const token = app.$cookies.get('jwt_token')
    if (token) {
      console.log('load plugins axios!!')
      config.headers.common['X-Requested-With'] = 'XMLHttpRequest'
      config.headers.common.Authorization = `Bearer ${token}`
    }
  })
}

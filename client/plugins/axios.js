export default ({ $axios }) => {
    $axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'
}
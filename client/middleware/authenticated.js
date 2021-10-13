export default function ({ store, route, redirect }) {
  if (route.name === 'index' && store.getters.isLoggedIn) {
    return redirect('/images')
  }
  if (
    (route.name === 'images' || route.name === 'images-id') &&
    !store.getters.isLoggedIn &&
    !store.getters.hasRefreshToken
  ) {
    return redirect('/')
  }
}

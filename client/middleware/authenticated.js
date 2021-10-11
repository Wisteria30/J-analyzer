export default function ({ store, route, redirect }) {
  if (
    route.name !== 'index' &&
    !store.getters.isLoggedIn &&
    !store.getters.hasRefreshToken
  ) {
    return redirect('/')
  }
  if (route.name === 'index' && store.getters.isLoggedIn) {
    return redirect('/images')
  }
}

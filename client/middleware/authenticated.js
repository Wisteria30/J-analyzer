export default function ({ store, route, redirect }) {
  if (route.name === 'list' && !store.getters.isLoggedIn) {
    return redirect('/')
  }
  if (route.name === 'index' && store.getters.isLoggedIn) {
    return redirect('/list')
  }
}

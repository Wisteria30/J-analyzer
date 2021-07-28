export default function ({ store, route, redirect }) {
  if (route.name === 'images' && !store.getters.isLoggedIn) {
    return redirect('/')
  }
  if (route.name === 'index' && store.getters.isLoggedIn) {
    return redirect('/images')
  }
}

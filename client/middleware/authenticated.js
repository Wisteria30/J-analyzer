export default function ({ store, route, redirect }) {
  if (route.name === 'index' && !store.getters.IsAuth) {
    return redirect('/')
  }
  if (!store.getters.IsAuth) {
    return redirect('/')
  }
  if (route.name === 'index' && store.getters.IsAuth) {
    return redirect('/list')
  }
}

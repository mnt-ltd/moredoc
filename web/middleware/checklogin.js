export default function ({ store, route, redirect }) {
  const user = store.getters['user/user'] || { id: 0 }
  if (user.id) {
    redirect(`/user/${user.id}`)
  }
}

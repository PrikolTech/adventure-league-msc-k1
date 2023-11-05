import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
// import { useUser } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        requiresUnAuth: true
      }
    }
  ]
})

// router.beforeEach(async (to, from, next) => {
//   const userStore = useUser()
//   console.log(userStore.user)
//   if (to.matched.some((record) => record.meta.requiresAuth)) {
//     if (userStore.user && userStore.user.id) {
//       next()
//     } else {
//       next('/login')
//     }
//   } 
//   if (to.matched.some((record) => record.meta.requiresUnAuth)) {
//     if (userStore.user && userStore.user.id) {
//       next(from)
//     } else {
//       next()
//     }
//   } else {
//     next()
//   }
// })

export default router

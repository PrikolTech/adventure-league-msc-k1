import { createRouter, createWebHistory } from 'vue-router'
import { useUser } from '@/stores/user'
import NotFoundView from '@/views/NotFoundView.vue'
import HomeView from '@/views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'

import ProfileView from '@/views/ProfileView.vue'
import ProfileCourses from '@/components/profile/courses/ProfileCourses.vue'
import ProfileTasks from '@/components/profile/ProfileTasks.vue'
import ProfileMe from '@/components/profile/ProfileMe.vue'
import ProfileApplications from '@/components/profile/ProfileApplications.vue'
import ProfileSafety from '@/components/profile/ProfileSafety.vue'
import ProfileSettings from '@/components/profile/ProfileSettings.vue'

import MeCourses from '@/components/profile/me/MeCourses.vue'
import MeMarks from '@/components/profile/me/MeMarks.vue'
import MeComments from '@/components/profile/me/MeComments.vue'

import CourseView from '@/views/CourseView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: false
      }
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'notFound',
      component: NotFoundView,
      // meta: {
      //   requiresUnAuth: true
      // }
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
      meta: {
        requiresUnAuth: true
      }
    },
    {
      path: '/profile/',
      name: 'profile',
      component: ProfileView,
      children: [
        {
          path: 'me',
          name: 'profile/me/',
          component: ProfileMe,
          children: [
            {
              path: 'meCourses',
              name: 'profile/me/meCourses',
              component: MeCourses
            },
            {
              path: 'meMarks',
              name: 'profile/me/meMarks',
              component: MeMarks
            },
            {
              path: 'meComments',
              name: 'profile/me/meComments',
              component: MeComments
            },
          ]
        },
        {
          path: 'tasks',
          name: 'profile/tasks',
          component: ProfileTasks
        },
        {
          path: 'courses',
          name: 'profile/courses',
          component: ProfileCourses
        },
        {
          path: 'applications',
          name: 'profile/applications',
          component: ProfileApplications
        },
        {
          path: 'safety',
          name: 'profile/safety',
          component: ProfileSafety
        },
        {
          path: 'settings',
          name: 'profile/settings',
          component: ProfileSettings
        },
      ],
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/courses/:id',
      name: 'courses',
      component: CourseView
    },
  ]
})

// router.beforeEach(async (to, from, next) => {
//   const userStore = useUser()
//   if (to.matched.some((record) => record.meta.requiresAuth)) {
//     if (userStore.user) {
//       next()
//     } else {
//       next('/login')
//     }
//   } 
//   if (to.matched.some((record) => record.meta.requiresUnAuth)) {
//     if (userStore.user) {
//       next(from)
//     } else {
//       next()
//     }
//   } else {
//     next()
//   }
// })

export default router

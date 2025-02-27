import { useUserStore } from '@/stores'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: () => import('@/views/login/LoginNewPage.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/user',
      component: () => import('@/views/user/UserPage.vue'),
      children: [
        {
          path: '/user/avatar',
          component: () => import('@/views/user/UserAvatar.vue'),
        },
        {
          path: '/user/Password',
          component: () => import('@/views/user/UserPassword.vue'),
        },
        {
          path: '/user/Profile',
          component: () => import('@/views/user/UserProfile.vue'),
        },
      ],
      meta: { requiresAuth: true },
    },
    {
      path: '/',
      component: () => import('@/views/layout/HomePage.vue'),
      redirect: 'article/myday',
      children: [
        {
          path: '/article/myday',
          component: () => import('@/views/article/MyDay.vue'),
        },
        {
          path: '/article/ImportantToDo',
          component: () => import('@/views/article/ImportantToDo.vue'),
        },
        {
          path: '/article/TaskInPlan',
          component: () => import('@/views/article/TaskInPlan.vue'),
        },
        {
          path: '/article/AssignedToMe',
          component: () => import('@/views/article/AssignedToMe.vue'),
        },
        {
          path: '/article/TaskList',
          component: () => import('@/views/article/TaskList.vue'),
        },
      ],
      meta: { requiresAuth: true },
    },
  ],
})
router.beforeEach((to, from, next) => {
  const useStore = useUserStore()
  if (to.meta.requiresAuth && !useStore.token) {
    next('/login') // 如果需要授权且没有token，则跳转到登录页
  } else {
    next() // 否则继续正常路由跳转
  }
})

export default router

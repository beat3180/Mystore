import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '../pages/index.vue'
import Reset from '../pages/join/reset.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  { path: '/reset/:token', component: Reset},


]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

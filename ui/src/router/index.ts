import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '../pages/index.vue'
import Reset from '../pages/join/reset.vue'
import Post from '../pages/post.vue'
import Content from '../pages/content.vue'
import MyStore from '../pages/mystore.vue'
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  { path: '/reset/:token', component: Reset },
  {
    path: '/post',
    name: 'Post',
    component: Post
  },
  {
    path: '/content/:id',
    name: 'Content',
    component: Content
  },
  {
    path: '/mystore',
    name: 'MyStore',
    component: MyStore
  },


]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

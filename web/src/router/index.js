import { createRouter, createWebHashHistory } from 'vue-router'
const HomeView = ()  => import('../views/HomeView.vue')
const MessageView = () => import('../views/MessageView.vue')
const ResumeView = () => import('../views/Resume.vue')
const NotFound = () => import('../views/NotFound.vue')

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
    children:[
      {
        path: 'message',
        name: 'message',
        component: MessageView
      },
      {
        path: 'resume',
        name: 'resume',
        component: ResumeView
      }
    ]
  },
  {
    path: '/404',
    name: '404',
    component: NotFound
  },
  {
    path: '/:pathMatch(.*)',
    redirect: '/404'
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

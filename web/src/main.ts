import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Home from './views/Home.vue'
import Diaries from './views/Diaries.vue'
import LandingPage from './views/LandingPage.vue'
import { createWebHistory, createRouter } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'Landing Page',
      path: '/',
      component: LandingPage
    },
    {
      name: 'Home',
      path: '/home',
      component: Home
    },
    {
      name: 'Diaries',
      path: '/diaries',
      component: Diaries
    },
  ]
})

createApp(App)
  .use(router)
  .mount('#app')

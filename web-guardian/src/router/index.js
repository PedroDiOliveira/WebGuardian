import { createRouter, createWebHistory } from 'vue-router'
import Guardian from "@/view/Guardian.vue"
import Home from "../view/Home.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path:"/",
      component:Home
    },
    {
      path: "/guardian",
      component: Guardian
    }
  ]
})

export default router

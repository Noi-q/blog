import {createRouter,createWebHistory,RouteRecordRaw} from "vue-router";
import Layout from "../layout/index.vue"
const defaultRoutes:Array<RouteRecordRaw> = [
  {
    path:"/",
    redirect:"/home"
  },
  {
    path:"/home",
    component:Layout,
    children:[
      {
        path:"/home",
        component:()=>import("../views/home/index.vue")
      }
    ]
  }
]

const router = createRouter({
  history:createWebHistory(),
  routes:[...defaultRoutes]
})

export default router
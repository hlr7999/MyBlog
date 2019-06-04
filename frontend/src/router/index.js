import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/Home'
import Login from '@/views/Login'
import Register from '@/views/Register'
import Class from '@/views/Class'
import Article from '@/views/Article'
import AboutMe from '@/views/AboutMe'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/Login',
      name: 'Login',
      component: Login
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    },
    {
      path: '/Class/:class_id',
      name: 'Class',
      component: Class
    },
    {
      path: '/Class/:class_id/:son_class_id',
      name: 'Class',
      component: Class
    },
    {
      path: '/Article/:article_id',
      name: 'Article',
      component: Article
    },
    {
      path: '/AboutMe',
      name: 'AboutMe',
      component: AboutMe
    }
  ]
})

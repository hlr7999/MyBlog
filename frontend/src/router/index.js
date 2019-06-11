import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/Home'
import Login from '@/views/Login'
import Register from '@/views/Register'
import Class from '@/views/Class'
import Article from '@/views/Article'
import AboutMe from '@/views/AboutMe'
import ArticleEdit from '@/views/ArticleEdit'
import Forbidden from '@/views/Forbidden'
import LikeList from '@/views/LikeList'
import UserInfo from '@/views/UserInfo'
import AdminHome from '@/views/AdminHome'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/Forbidden',
      name: 'Forbidden',
      component: Forbidden
    },
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/page/:page',
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
    },
    {
      path: '/NewArticle',
      name: 'ArticleEdit',
      component: ArticleEdit
    },
    {
      path: '/ModifyArticle/:article_id',
      name: 'ArticleEdit',
      component: ArticleEdit
    },
    {
      path: '/Like',
      name: 'LikeList',
      component: LikeList
    },
    {
      path: '/Like/page/:page',
      name: 'LikeList',
      component: LikeList
    },
    {
      path: '/Collect',
      name: 'CollectList',
      component: LikeList
    },
    {
      path: '/Collect/page/:page',
      name: 'CollectList',
      component: LikeList
    },
    {
      path: '/UserInfo',
      name: 'UserInfo',
      component: UserInfo
    },
    {
      path: '/Admin',
      name: 'AdminHome',
      component: AdminHome
    },
    {
      path: '/Admin/page/:page',
      name: 'AdminHome',
      component: AdminHome
    }
  ]
})

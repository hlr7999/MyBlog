<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
  <div class="container">
    <div class="articleItem" v-for="article in articles" :key="article.id">
      <article-card :articleInfo="article"></article-card>
    </div>
    <div class="viewmore">
      <a v-show="hasMore" class="viewMoreBtn" href="javascript:void(0);" @click="viewMoreFun">点击加载更多</a>
      <a v-show="!hasMore" class="viewMoreBtn" href="javascript:void(0);">暂无更多数据</a>
    </div>
  </div>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleCard from "../components/articleCard.vue" 

export default {
  name: 'home',
  components: {
    'blog-header': header,
    'article-card': articleCard
  },
  data() {
    return {
      articles: [{
        id: "0",
        title: "Vue.js搭建博客",
        year: 2019,
        month: 5,
        day: 25,
        browse_count: 77,
        comment_count: 77,
        like_count: 77,
        collect_count: 77,
        class_id: 0,
        class_name: "Vue.js",
        description: "Vue.js搭建博客",
        image: "/static/img/vuelogo.jpg"
      },{
        id: "1",
        title: "Vue.js搭建博客",
        year: 2019,
        month: 5,
        day: 25,
        browse_count: 77,
        comment_count: 77,
        like_count: 77,
        collect_count: 77,
        class_id: 0,
        class_name: "Vue.js",
        description: "Vue.js搭建博客",
        image: "/static/img/vuelogo.jpg"
      }],
      hasMore: true
    }
  },
  methods: {
    viewMoreFun() {

    }
  }
}
</script>

<style>

</style>

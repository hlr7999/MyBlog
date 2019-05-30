<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
  <div class="container">
    
    <div class="articleItem" v-for="article in articles" :key="article.id">
      <article-card :articleInfo="article"></article-card>
    </div>
    <div class="viewmore">
      <a v-show="hasMore" class="tcolors-bg" href="javascript:void(0);" @click="viewMoreFun">点击加载更多</a>
      <a v-show="!hasMore" class="tcolors-bg" href="javascript:void(0);">暂无更多数据</a>
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
        id: "fuck1",
        title: "Vue.js搭建博客",
        year: 2019,
        month: 5,
        day: 25,
        browse_count: 77,
        comment_count: 77,
        like_count: 77,
        collect_count: 77,
        class_id: 0,
        class_name: "vue",
        description: "Vue.js搭建博客",
        image: "/static/img/vuelogo.jpg"
      },{
        id: "fuck2",
        title: "Vue.js搭建博客",
        year: 2019,
        month: 5,
        day: 25,
        browse_count: 77,
        comment_count: 77,
        like_count: 77,
        collect_count: 77,
        class_id: 0,
        class_name: "vue",
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
a {
  text-decoration: none;
  color: black;
}

.container {
  max-width: 50%;
  margin:100px auto;
  padding:0 10px;
  transition: all 0.5s ease-out;
  font-size: 15px;
}

.articleItem {
  margin: 10px auto;
}

.viewmore{
  text-align: center;
  width:100%;
}

.viewmore a{
  background-color: #97dffd;
  text-decoration: none;
  border-radius: 5px;
  text-align: center;
  font-size: 16px;
  color:#fff;
  height:30px;
  line-height: 30px;
  display: block;
}

.tcolors-bg{
  background: #97dffd;
  transition: all .3s ease-in-out;
}

.tcolors-bg:hover{
  background: #48456D;
}

@media screen and (max-width:1400px){
  .container {
    max-width: 60%!important;
    margin:100px auto;
  }
}

@media screen and (max-width:1000px){
  .container{
    max-width: 80%!important;
    margin:100px auto;
  }
}

@media screen and (max-width:800px){
  .container{
    max-width: 100%!important;
    margin:100px auto;
  }
}

</style>

<template>
<div>
  <blog-header :activeIndex="2"></blog-header>
  <div class="container">
    <div class="classTitle">
      <div class="classLabel" >
        <a :href="'#/Class/'+class_id">{{class_name}}</a>
      </div>
      <ul v-if="sonclassList" class="sonclassTitle" >
        <li v-for="(item,index) in sonclassList" :key="index">
          <a :href="'#/Class/'+class_id+'/'+item.id" 
            :class="item.id == son_class_id ? 'active' : ''">{{item.name}}</a>
        </li>
      </ul>
    </div>

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
      sonclassList: [{
          id: "00",
          name: "Vue.js"
        },{
          id: "11",
          name: "Golang"
        }
      ],
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
        class_name: "Vue.js",
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
        class_name: "Vue.js",
        description: "Vue.js搭建博客",
        image: "/static/img/vuelogo.jpg"
      }],
      hasMore: true,
      class_id: "0",
      son_class_id: "0",
      class_name: ""
    }
  },

  methods: {
    getData() {
      this.class_id = this.$route.params.class_id
      this.son_class_id = this.$route.params.son_class_id
      this.class_name = "技术分享"
    },

    viewMoreFun() {

    }
  },

  mounted () {
    this.getData()
  },

  watch: {
    '$route': 'getData'
  }
}
</script>

<style>
.classTitle {
  margin-bottom: 40px;
  position: relative;
  border-radius: 5px;
  background: #fff;
  padding: 15px;
  text-align: left;
}

.sonclassTitle {
  width: 100%;
  margin: 0;
  padding: 0;
}

.sonclassTitle li {
  display: inline-block;
}

.sonclassTitle li a {
  display: inline-block;
  padding: 3px 7px;
  margin: 5px 10px;
  color:#fff;
  border-radius: 4px;
  background: #64609E;
  border: 1px solid #64609E;
  transition: transform 0.2s linear;
  -webkit-transition: transform 0.2s linear;
}

.sonclassTitle li a:hover{
  transform: translate(0,-3px);
  -webkit-transform: translate(0,-3px);
}

.sonclassTitle li a.active{
  background-color: #fff !important;
  color:#64609E !important;
}
</style>

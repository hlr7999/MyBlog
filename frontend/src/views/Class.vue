<template>
<div>
  <blog-header :activeIndex="2"></blog-header>

  <div class="container">
    <div class="classTitle">
      <div class="classLabel" >
        <a :href="'#/Class/'+class_id">{{className}}</a>
      </div>
      <ul v-if="sonclassList" class="sonclassTitle" >
        <li v-for="(item,index) in sonclassList" :key="index">
          <a :href="'#/Class/'+class_id+'/'+item._id" 
            :class="item._id == son_class_id ? 'active' : ''">{{item.name}}</a>
        </li>
      </ul>
    </div>

    <div class="articleItem" v-for="article in articles" :key="article.id">
      <article-card :articleInfo="article"></article-card>
    </div>
  </div>

  <blog-footer></blog-footer>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleCard from "../components/articleCard.vue" 
import footer from "../components/footer.vue"

import { GetSecondClass, GetArticlesById } from "../api/api"

export default {
  name: 'home',

  components: {
    'blog-header': header,
    'article-card': articleCard,
    'blog-footer': footer
  },

  data() {
    return {
      sonclassList: [],
      articles: [],
      currentPage: 1,
      totalNum: 2,
      class_id: "0",
      son_class_id: "0",
      className: ""
    }
  },

  methods: {
    getData() {
      this.class_id = this.$route.params.class_id
      this.son_class_id = this.$route.params.son_class_id

      for (let item of this.$store.state.classList) {
        if (item._id == this.class_id) {
          this.className = item.name
          break
        }
      }
      
      GetSecondClass(this.class_id)
      .then(res => {
        this.sonclassList = res.data
      })
      .catch(() => {
        this.$message.error("获取分类失败")
      })
      var classId = this.class_id
      if (this.son_class_id) {
        classId = this.son_class_id
      }
      GetArticlesById(classId)
      .then(res => {
        this.articles = res.data
      })
      .catch(() => {
        this.$message.error("获取文章失败")
      })
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

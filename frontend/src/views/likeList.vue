<template>
<div>
  <blog-header :activeIndex="0"></blog-header>

  <div class="container">
    <div class="likeBoxTitle">
      <h1 v-show="like==1"><i class="fa fa-wa fa-heart"></i>喜欢的文章</h1>
      <h1 v-show="like==2"><i class="fa fa-wa fa-star"></i>收藏的文章</h1>
    </div>

    <div class="articleItem" v-for="article in articles" :key="article._id">
      <article-card :articleInfo="article"></article-card>
    </div>

    <el-pagination
      @current-change="handleCurrentChange"
      :current-page.sync="currentPage"
      :page-size="9"
      layout="total, prev, pager, next"
      :total="totalNum">
    </el-pagination>
  </div>

  <blog-footer></blog-footer>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleCard from "../components/articleCard.vue" 
import footer from "../components/footer.vue"

import { GetLCArticles } from "../api/api"

export default {
  components: {
    'blog-header': header,
    'article-card': articleCard,
    'blog-footer': footer
  },

  data() {
    return {
      totalArticles: [],
      articles: [],
      currentPage: 1,
      totalNum: 0,
      like: 0,
      homePath: ""
    }
  },

  methods: {
    handleCurrentChange() {
      this.articles = this.totalArticles.slice(
        (this.currentPage - 1) * 9,
        this.currentPage * 9
      );
      this.router.push({
        path: this.homePath + "page/" + this.currentPage
      });
    },

    changePage() {
      if (this.$route.params.page) {
        this.currentPage = Number(this.$route.params.page)
        if (this.currentPage == 1) {
          this.$router.push({
            path: this.homePath
          })
        }
      } else {
        this.currentPage = 1
      }
      
      if (this.currentPage > (this.totalNum-1)/6 + 1) {
        this.currentPage = 1
        this.$router.push({
          path: this.homePath
        })
      }
    },

    getData() {
      if (!this.$store.state.hasLogin) {
        this.$message.error("您未登录")
        this.$router.push({
          path: "/Forbidden"
        })
        return
      }
      
      var r = this.$route.path
      r = r.toLowerCase()
      if (r == "/like") {
        this.like = 1
        this.homePath = "/Like"
      } else if (r == "/collect") {
        this.like = 2
        this.homePath = "/Collect"
      }
      this.changePage()

      GetLCArticles(r)
      .then(res => {
        if (res.data) {
          this.totalArticles = res.data
          this.totalNum = this.totalArticles.length
          this.articles = this.totalArticles.slice(
            (this.currentPage - 1) * 9,
            this.currentPage * 9
          )
        }
      })
      .catch(() => {
        this.$message({
          message: "未知错误",
          duration: 1500
        })
      })
    }
  },

  created () {
    this.getData()
  },

  watch: {
    '$route':'getData'
  },
}
</script>

<style>
.likeBoxTitle{
  text-align: center;
  padding:40px 0;
  margin-bottom: 40px;
  background: #fff;
  border-radius: 5px;
}

.likeBoxTitle h1{
  font-size: 25px;
  font-weight: 700;
  color: black;
  margin: 0;
}

.likeBoxTitle h1 i{
  color:#DF2050;
  margin-right: 10px;
  font-size: 30px;
}
</style>

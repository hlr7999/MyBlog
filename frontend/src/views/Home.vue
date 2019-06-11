<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
  <div class="container">
    <div class="articleItem" v-for="article in articles" :key="article.id">
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

import { GetHomeArticles } from "../api/api"

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
      totalNum: 0
    }
  },

  methods: {
    handleCurrentChange() {
      this.articles = this.totalArticles.slice(
        (this.currentPage - 1) * 9,
        this.currentPage * 9
      )
      this.router.push({
        path: "/page/" + this.currentPage
      })
    },

    changePage() {
      this.currentPage = Number(this.$route.params.page)
      if (this.currentPage > (this.totalNum-1)/6 + 1) {
        this.currentPage = 1
        this.$router.push({
          path: "/"
        })
      }
    },

    getData() {
      GetHomeArticles()
      .then(res => {
        if (res.data) {
          this.totalArticles = res.data.articles
          this.totalNum = this.articles.length
          this.articles = this.articles.slice(
            (this.currentPage - 1) * 9,
            this.currentPage * 9
          )
        }
      })
      .catch(() => {
        this.$message({
          message: "未知错误",
          type: "error",
          duration: 1500
        })
      })
    }
  },

  watch: {
    '$route': 'changePage'
  },

  created () {
    this.getData()
  }
}
</script>

<style>

</style>

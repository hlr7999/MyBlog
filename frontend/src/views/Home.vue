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
      currentPage: this.$store.state.currentPage,
      totalNum: 0
    }
  },

  methods: {
    handleCurrentChange() {
      this.articles = this.totalArticles.slice(
        (this.currentPage - 1) * 9,
        this.currentPage * 9
      )
      this.$store.commit("changePage", this.currentPage)
    },

    getData() {
      GetHomeArticles()
      .then(res => {
        if (res.data) {
          this.totalArticles = res.data.articles
          this.totalNum = this.articles.length
          if (this.currentPage > (this.totalNum-1)/6 + 1 ||
              this.$store.state.lastUsePage != "H") {
            this.currentPage = 1
            this.$store.commit("changePage", this.currentPage)
            if (this.$store.state.lastUsePage != "H") {
              this.$store.commit("changeLastUsePage", "H")
            }
          }
          this.articles = this.articles.slice(
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
  }
}
</script>

<style>

</style>

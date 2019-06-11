<template>
  <div>
    <div class="articleItem" v-for="article in articles" :key="article.id">
      <admin-article-card v-if="type == 1" :articleInfo="article"></admin-article-card>
      <article-card v-else :articleInfo="article"></article-card>
    </div>
    <el-pagination
      @current-change="handleCurrentChange"
      :current-page.sync="currentPage"
      :page-size="9"
      layout="total, prev, pager, next"
      :total="totalNum"
    ></el-pagination>
  </div>
</template>

<script>
import articleCard from "../components/articleCard.vue"
import adminArticleCard from "../components/adminArticleCard.vue";

import { GetHomeArticles } from "../api/api";

export default {
  components: {
    "article-card": articleCard,
    "admin-article-card": adminArticleCard
  },

  data() {
    return {
      totalArticles: [],
      articles: [],
      currentPage: 1,
      totalNum: 0,
      homePath: "/"
    };
  },

  props: {
    type: Number
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
      if (this.type == 1) {
        this.homePath = "/Admin"
        if (!this.$store.state.hasLogin || !this.$store.state.isAdmin) {
          this.$router.push({
            path: "/Forbidden"
          })
        }
      }

      GetHomeArticles()
      .then(res => {
        if (res.data) {
          this.totalArticles = res.data.articles;
          this.totalNum = this.articles.length;
          this.articles = this.articles.slice(
            (this.currentPage - 1) * 9,
            this.currentPage * 9
          );
        }
      })
      .catch(() => {
        this.$message({
          message: "未知错误",
          type: "error",
          duration: 1500
        });
      });
    }
  },

  watch: {
    $route: "changePage"
  },

  created() {
    this.getData();
  }
};
</script>

<style>
</style>

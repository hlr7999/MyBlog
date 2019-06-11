<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
  <div class="container">
    <article-list :type="0"></article-list>
  </div>
  <blog-footer></blog-footer>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleList from "../components/allArticleList.vue" 
import footer from "../components/footer.vue"

export default {
  components: {
    'blog-header': header,
    'article-list': articleList,
    'blog-footer': footer
  },

  data() {
    return {

    }
  }
}
</script>

<style>

</style>

<template>
<div>
  <blog-header :activeIndex="0"></blog-header>
  <div class="container">
    <article-detail></article-detail>
  </div>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleDetailCard from "../components/articleDetailCard.vue";

export default {
  data() {
    return {
    }
  },

  components: {
    "article-detail": articleDetailCard,
    "blog-header": header
  }
};
</script>

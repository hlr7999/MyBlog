<template>
<div>
  <blog-header :activeIndex="3"></blog-header>
  <div class="container">
    <about-content></about-content>
    <about-comment></about-comment>
  </div>
</div>
</template>

<script>
import header from "../components/header.vue"
import aboutContent from "../components/aboutContent.vue"
import comment from "../components/comment.vue"

export default {
  data() {
    return {
    }
  },

  components: {
    "about-content": aboutContent,
    "blog-header": header,
    "about-comment": comment
  }
};
</script>

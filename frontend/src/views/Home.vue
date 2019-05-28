<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
  <el-card v-for="article in articles" :key="article.id">
    
  </el-card>
</div>
</template>

<script>
import header from "../components/header.vue"
export default {
  name: 'home',
  components: {
    'blog-header': header
  },
  data() {
    return {
      atricles: [{
        id: "fuck1",
        title: "Vue.js搭建博客",
        image: ""
      }]
    }
  }
}
</script>

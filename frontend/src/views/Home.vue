<template>
<div>
  <blog-header :activeIndex="1"></blog-header>
</div>
</template>

<script>
import header from "../components/header.vue"
export default {
  name: 'home',
  components: {
    'blog-header': header
  }
}
</script>

<template>
<div>
  <blog-header></blog-header>
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

<template>
<div>
  <blog-header :activeIndex="0"></blog-header>
  <div class="container">
    <edit-card></edit-card>
  </div>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleEditCard from "../components/articleEditCard.vue"

export default {
  components: {
    'blog-header': header,
    'edit-card': articleEditCard
  },

  data() {
    return {

    }
  },

  methods: {
    
  }
}
</script>

<style>

</style>

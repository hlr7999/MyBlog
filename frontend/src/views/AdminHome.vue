<template>
<div class="adminHome">
  <blog-header :activeIndex="0"></blog-header>
  <div class="container">
    <el-tabs v-model="activeName" :lazy="true">
      <el-tab-pane label="文章管理" name="first">
        <article-list :type="1"></article-list>
      </el-tab-pane>
      <el-tab-pane label="用户管理" name="second">
        <user-list></user-list>
      </el-tab-pane>
      <el-tab-pane label="分类管理" name="third">
        <class-list></class-list>
      </el-tab-pane>
    </el-tabs>
  </div>
  <blog-footer></blog-footer>
</div>
</template>

<script>
import header from "../components/header.vue"
import articleList from "../components/allArticleList.vue" 
import userList from "../components/userList.vue"
import classList from "../components/classList.vue"
import footer from "../components/footer.vue"

export default {
  components: {
    'blog-header': header,
    'article-list': articleList,
    'user-list': userList,
    'class-list': classList,
    'blog-footer': footer
  },

  data() {
    return {
      activeName: 'first'
    }
  }
}
</script>

<style>
.adminHome .el-tabs__header {
  padding: 10px;
}

.adminHome .el-tabs__item {
  color: white !important;
  font-size: 16px !important;
  font-weight: 600!important;
}

.adminHome .el-tabs__item.is-active {
  color: #409eff !important;
}
</style>

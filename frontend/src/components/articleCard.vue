<template>
  <el-row>
    <el-col
      :span="24"
      class="a-articleInfo"
    >
      <span class="a-round-date">
        <span class="month" v-html="articleInfo.month+'月'"></span>
        <span class="day" v-html="articleInfo.day"></span>
      </span>
      <header>
        <h1>
          <a :href="'#/Article/'+articleInfo._id">{{articleInfo.title}}</a>
        </h1>
        <h2>
          <i class="fa fa-fw fa-user"></i>发表于
          <i class="fa fa-fw fa-clock-o"></i>
          <span
            v-html="articleInfo.year + '年' + articleInfo.month + '月' + articleInfo.day + '日'"></span> •
          <i class="fa fa-fw fa-eye"></i>
          {{articleInfo.browseCount}} 次浏览 •
          <i class="fa fa-fw fa-comments"></i>
          {{articleInfo.commentCount}} 条评论 •
          <span class="rateBox">
            <i class="fa fa-fw fa-heart"></i>
            {{articleInfo.likeCount?articleInfo.likeCount:0}}点赞 •
            <i class="fa fa-fw fa-star"></i>
            {{articleInfo.collectCount?articleInfo.collectCount:0}}收藏
          </span>
        </h2>
        <div class="classLabel">
          <div>{{articleInfo.className}}</div>
        </div>
      </header>
      <div class="article-content">
        <p style="text-indent:2em;">{{articleInfo.description}}</p>
        <div :class="articleInfo.imageCover ? 'imageCover' : 'noCover' ">
          <img :src="articleInfo.imageCover" alt style="max-width:100%">
        </div>
      </div>
      <div class="viewdetail">
        <a class="readWholeArticle" :href="'#/Article/'+articleInfo._id">阅读全文>></a>
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  data() {
    return {
    }
  },

  props: {
    articleInfo: Object
  }
};
</script>

<style>
.a-articleInfo .article-content .imageCover {
  max-height:300px;
  overflow:hidden;
  text-align:center;
}

.a-articleInfo .article-content .noCover {
  display: none;
}
</style>

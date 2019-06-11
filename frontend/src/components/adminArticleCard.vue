<template>
  <el-row>
    <el-col
      :span="24"
      class="a-articleInfo"
    >
      <header>
        <h1>
          <a :href="'#/Article/'+articleInfo.id">{{articleInfo.title}}</a>
        </h1>
        <h2>
          <i class="fa fa-fw fa-user"></i>发表于
          <i class="fa fa-fw fa-clock-o"></i>
          <span
            v-html="articleInfo.year + '年' + articleInfo.month + '月' + articleInfo.day + '日'"></span> •
          <i class="fa fa-fw fa-eye"></i>
          {{articleInfo.browse_count}} 次浏览 •
          <i class="fa fa-fw fa-comments"></i>
          {{articleInfo.comment_count}} 条评论 •
          <span class="rateBox">
            <i class="fa fa-fw fa-heart"></i>
            {{articleInfo.like_count?articleInfo.like_count:0}}点赞 •
            <i class="fa fa-fw fa-star"></i>
            {{articleInfo.collect_count?articleInfo.collect_count:0}}收藏
          </span>
        </h2>
        <div class="classLabel">
          <a :href="'#/Class/'+articleInfo.class_id">{{articleInfo.class_name}}</a>
        </div>
      </header>
      <div class="article-content">
        <p style="text-indent:2em;">{{articleInfo.description}}</p>
      </div>
      <div class="btnBox">
        <el-button size="mini">修改</el-button>
        <el-button type="danger" size="mini">删除</el-button>
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
  },

  methods: {

  }
};
</script>

<style>
.classLabel  {
left:-18px!important;
padding-left:18px!important;
}

.classLabel ::after{
border:none!important;
}

.btnBox {
  text-align: center;
  margin-top: 10px;
  line-height: 20px;
}
</style>

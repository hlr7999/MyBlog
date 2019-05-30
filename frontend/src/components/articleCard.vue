<template>
  <el-row>
    <el-col
      :span="24"
      class="s-articleInfo"
    >
      <span class="s-round-date">
        <span class="month" v-html="articleInfo.month+'月'"></span>
        <span class="day" v-html="articleInfo.day"></span>
      </span>
      <header>
        <h1>
          <a :href="'#/article?id='+articleInfo.id" target="_blank">{{articleInfo.title}}</a>
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
          <a :href="'#/class?id='+articleInfo.class_id">{{articleInfo.class_name}}</a>
        </div>
      </header>
      <div class="article-content">
        <p style="text-indent:2em;">{{articleInfo.description}}</p>
        <p style="max-height:300px;overflow:hidden;text-align:center;">
          <img :src="articleInfo.image" alt style="max-width:100%">
        </p>
      </div>
      <div class="viewdetail">
        <a class="tcolors-bg" :href="'#/article?id='+articleInfo.id" target="_blank">阅读全文>></a>
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  data() {
    return {
    };
  },

  props: {
    articleInfo: Object
  }
};
</script>

<style>
.s-round-date {
  position: absolute;
  top:-20px;
  left:-35px;
  height:60px;
  width:70px;
  padding-top:10px;
  border-radius: 100px;
  color:#fff;
  background: #97dffd;
  box-shadow: 0 2px 2px 0 rgba(0,0,0,.14), 0 3px 1px -2px rgba(0,0,0,.2), 0 1px 5px 0 rgba(0,0,0,.12);
}

.s-round-date span {
  text-align: center;
  display: block;
}

.s-round-date .day {
  font-size: 30px;
  font-weight: 700;
}

.s-articleInfo {
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  position: relative;
  background: #fff;
  padding:15px;
  border-radius: 5px;
  margin-bottom: 40px;
  font-size: 15px;
}

.s-articleInfo header {
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  text-align: left;
}

.s-articleInfo header h1{
  margin: 10px 0;
  font-size: 25px;
  font-weight: 700;
  text-align: center;
  line-height: 30px;
}

.s-articleInfo header h2{
  margin:10px 0;
  line-height: 24px;
  text-align: center;
  color:#555;
  font-size: 14px;
}

.s-articleInfo section p{
  line-height: 24px;
  margin:10px 0;
}

.s-articleInfo header h2 .el-rate,
.rateBox{
  display:inline-block;
}

.s-articleInfo header h2 .el-rate i{
  font-size: 16px;
}

.s-articleInfo .article-coontent{
  font-size: 15px;
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
}

.s-articleInfo .article-content p{
  margin:10px 0;
  line-height:24px;
}

.s-articleInfo .viewdetail{
  margin:10px 0 ;
  line-height: 24px;
  text-align: center;
}

.s-articleInfo .viewdetail a{
  color:#fff;
  font-size: 12px;
  padding:5px 10px;
  border-radius: 5px;
}

.classLabel {
  border-color:#97dffd;
  box-shadow: 0 2px 2px 0 rgba(0,0,0,.14), 0 3px 1px -2px rgba(0,0,0,.2), 0 1px 5px 0 rgba(0,0,0,.12);
  background-color: #97dffd;
  margin:5px 0 15px 2px;
  font-size: 14px;
  text-align: left;
  position: relative;
  left:-32px;
  border-radius: 0 4px 4px 0;
  padding:7px 11.2px 7px 32px;
  display: inline-block;
  color:#fff;
}

.classLabel ::after{
  position: absolute;
  content: '';
  top:100%;
  left:0;
  border-top: 0 solid transparent;
  border-right-width:1em;
  border-right-color:#48456D;
  border-right-style:solid;
  border-bottom: 1em solid transparent;
  border-left: 0 solid transparent;
  width:0;
  height:0;
}

.classLabel  a{
  color:#fff;
  line-height: 100%;
}

.tcolors-bg{
  background: #97dffd;
  transition: all .3s ease-in-out;
}

.tcolors-bg:hover{
  background: #48456D;
}

@media screen and (max-width:800px){
  .s-round-date {
    display: none;
  }
  .classLabel  {
    left:-18px!important;
    padding-left:18px!important;
  }
  .classLabel ::after{
    border:none!important;
  }
}

</style>

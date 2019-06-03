<template>
  <div class="detailArticleBox a-articleInfo">
    <span class="a-round-date">
      <span class="month" v-html="articleInfo.month+'月'"></span>
      <span class="day" v-html="articleInfo.day"></span>
    </span>
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

    <mavon-editor
      :value = "articleInfo.content"
      :boxShadow = "false"
      :subfield = "false"
      defaultOpen = "preview"
      :toolbarsFlag	= "false"
      :ishljs = "true"
      code-style = "github-gist"
    />

    <div class="likeColBox">
      <div class="likeBox" @click="likecollectHandle(1)">
        <i :class="likeArt ? 'fa fa-fw fa-heart' : 'fa fa-fw fa-heart-o'"></i>
        喜欢 | {{articleInfo.like_count}}
      </div>
      <div class="collectBox" @click="likecollectHandle(2)">
        <i :class="collectArt ? 'fa fa-fw fa-star' : 'fa fa-fw fa-star-o'"></i>
        收藏 | {{articleInfo.collect_count}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      aid: "0",
      articleInfo: "",
      likeArt: false,
      collectArt: false,
      haslogin: false,
      userId: "",
    }
  },

  methods: {
    getData: function() {
      if (this.$route.params.article_id) {
        this.aid = this.$route.params.article_id
      }
      
      // if (localStorage.getItem("userInfo")) {
      //   that.haslogin = true;
      //   that.userInfo = JSON.parse(localStorage.getItem("userInfo"));
      //   that.userId = that.userInfo.userId;
      //   // console.log(that.userInfo);
      // } else {
      //   that.haslogin = false;
      // }

      //获取详情接口
      // getArticleInfo(that.aid, that.userId, function(msg) {
      //   // console.log('文章详情',msg);
      //   that.articleInfo = msg;
      //   that.likeCount = msg.like_count ? msg.like_count : 0;
      //   that.collectCount = msg.collect_count ? msg.collect_count : 0;
      //   that.likeArt = msg.user_like_start == 0 ? false : true;
      //   that.collectArt = msg.user_collect_start == 0 ? false : true;
      //   that.create_time = initDate(that.articleInfo.create_time, "all");
      // });

      this.articleInfo = {
        id: "0",
        title: "Vue.js搭建博客",
        year: 2019,
        month: 5,
        day: 25,
        browse_count: 77,
        comment_count: 77,
        like_count: 77,
        collect_count: 77,
        class_id: 0,
        class_name: "Vue.js",
        description: "Vue.js搭建博客",
        image: "/static/img/vuelogo.jpg",
        content: "# Not found"
      }

      const axios = require('axios')
      axios.get('http://localhost:10080/markdown')
        .then(res => {
          this.articleInfo.content = res.data
        })
    },

    // likecollectHandle: function(islike) {
    //   //用户点击喜欢0,用户点击收藏1
    //   var that = this;
    //   if (localStorage.getItem("userInfo")) {
    //     //判断是否登录
    //     var tip = "";
    //     if (islike == 1) {
    //       if (!that.likeArt) {
    //         that.likeCount += 1;
    //         that.likeArt = true;
    //         tip = "已点赞";
    //       } else {
    //         that.likeCount -= 1;
    //         that.likeArt = false;
    //         tip = "已取消点赞";
    //       }
    //     } else {
    //       if (!that.collectArt) {
    //         that.collectCount += 1;
    //         that.collectArt = true;
    //         tip = "已收藏";
    //       } else {
    //         that.collectCount -= 1;
    //         that.collectArt = false;
    //         tip = "已取消收藏";
    //       }
    //     }
    //     getArtLikeCollect(that.userId, that.aid, islike, function(msg) {
    //       // console.log('喜欢收藏成功',msg);
    //       that.$message({
    //         message: tip,
    //         type: "success"
    //       });
    //     });
    //   } else {
    //     //未登录 前去登录。
    //     that
    //       .$confirm("登录后即可点赞和收藏，是否前往登录页面?", "提示", {
    //         confirmButtonText: "确定",
    //         cancelButtonText: "取消",
    //         type: "warning"
    //       })
    //       .then(() => {
    //         //确定，跳转至登录页面
    //         //储存当前页面路径，登录成功后跳回来
    //         localStorage.setItem("logUrl", that.$route.fullPath);
    //         that.$router.push({ path: "/Login?login=1" });
    //       })
    //       .catch(() => {
    //         //取消关闭弹窗
    //       });
    //   }
    // }
  },

  watch: {
    $route: "getData"
  },

  components: {
    
  },

  mounted () {
    this.getData();
  }
};
</script>

<style>
.detailArticleBox .article-content {
  font-size: 15px;
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  overflow-x: hidden;
}

/*点赞 收藏*/
.likeColBox {
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  flex-wrap: wrap;
  margin: 10px;
}

.likeBox,
.collectBox {
  display: inline-block;
  padding: 0 10px;
  margin: 0 10px;
  height: 35px;
  color: #e26d6d;
  line-height: 35px;
  border-radius: 35px;
  border: 1px solid #e26d6d;
  cursor: pointer;
}

.markdown-body {
  min-width: 0 !important;
}

.markdown-body .v-note-panel {
  border: 0 !important;
}

.markdown-body .v-note-panel .v-show-content {
  background-color: #fff !important;
}

.markdown-body .highlight pre, 
.markdown-body pre {
  padding: 3px !important;
}

</style>

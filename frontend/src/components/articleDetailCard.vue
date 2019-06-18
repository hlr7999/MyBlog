<template>
  <div class="detailArticleBox a-articleInfo">
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
        <a :href="'#/Class/'+articleInfo.classId">{{articleInfo.className}}</a>
      </div>
    </header>

    <mavon-editor
      :value = "articleInfo.content"
      :boxShadow = "false"
      :subfield = "false"
      defaultOpen = "preview"
      :editable="false"
      :toolbarsFlag	= "false"
      :ishljs = "true"
      code-style = "github-gist"
      :shortCut = "false"
      :imageClick = "imageClick"
    />

    <div class="likeColBox">
      <div class="likeBox" @click="likecollectHandle(1)">
        <i :class="likeArt ? 'fa fa-fw fa-heart' : 'fa fa-fw fa-heart-o'"></i>
        喜欢 | {{articleInfo.likeCount}}
      </div>
      <div class="collectBox" @click="likecollectHandle(2)">
        <i :class="collectArt ? 'fa fa-fw fa-star' : 'fa fa-fw fa-star-o'"></i>
        收藏 | {{articleInfo.collectCount}}
      </div>
    </div>
  </div>
</template>

<script>
import { GetArticle, IsLikeCollect, DoLikeCollect } from "../api/api"

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
      
      GetArticle(this.aid)
      .then(res => {
        this.articleInfo = res.data
      })
      .catch(() => {
        this.$message.error("错误")
      })

      IsLikeCollect(this.aid)
      .then(res => {
        this.likeArt = res.data.isLike
        this.collectArt = res.data.isCollect
      })
    },

    likecollectHandle: function(islike) {
      if (this.$store.state.hasLogin) {
        var tip = ""
        if (islike == 1) {
          if (!this.likeArt) {
            this.articleInfo.likeCount += 1
            this.likeArt = true
            tip = "已点赞"
          } else {
            this.articleInfo.likeCount -= 1
            this.likeArt = false
            tip = "已取消点赞"
          }
        } else {
          if (!this.collectArt) {
            this.articleInfo.collectCount += 1
            this.collectArt = true
            tip = "已收藏"
          } else {
            this.articleInfo.collectCount -= 1
            this.collectArt = false
            tip = "已取消收藏"
          }
        }
        DoLikeCollect(this.aid, islike, this.likeArt, this.collectArt)
        .then(res => {
          this.$message({
            message: tip,
            type: "success"
          })
        })
      } else {
        this.$confirm("登录才可点赞和收藏，是否前往登录页面?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
        .then(() => {
          this.$router.push("/Login");
        })
        .catch(() => {

        });
      }
    },

    imageClick() {
      
    }
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

.detailArticleBox .markdown-body {
  min-width: 0 !important;
}

.detailArticleBox .markdown-body .v-note-panel {
  border: 0 !important;
}

.detailArticleBox .markdown-body .v-note-panel .v-show-content {
  background-color: #fff !important;
}

.detailArticleBox .markdown-body .highlight pre, 
.detailArticleBox .markdown-body pre {
  padding: 3px !important;
}

.detailArticleBox .markdown-body ul {
  list-style: disc;
}

</style>

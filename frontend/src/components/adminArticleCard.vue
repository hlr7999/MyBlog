<template>
  <el-row>
    <el-col
      :span="24"
      class="a-articleInfo"
      v-if="articleInfo.newArticle"
    >
      <div class="newArticleBox" @click="newArticleFunc">
        <i class="el-icon-plus"></i>
      </div>
    </el-col>

    <el-col
      :span="24"
      class="a-articleInfo"
      v-else
    >
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
            {{articleInfo.like_count?articleInfo.likeCount:0}}点赞 •
            <i class="fa fa-fw fa-star"></i>
            {{articleInfo.collect_count?articleInfo.collectCount:0}}收藏
          </span>
        </h2>
        <div class="classLabel">
          <a :href="'#/Class/'+articleInfo.classId">{{articleInfo.className}}</a>
        </div>
      </header>
      <div class="article-content">
        <p style="text-indent:2em;">{{articleInfo.description}}</p>
      </div>
      <div class="btnBox">
        <el-button type="danger" size="mini" @click="dialogVisible = true">删除</el-button>
      </div>
    </el-col>

    <el-dialog
      title="删除"
      :visible.sync="dialogVisible"
      width="30%">
      <span>确认删除</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="deleteArticle">确 定</el-button>
      </span>
    </el-dialog>
  </el-row>
</template>

<script>
import { DeleteArticle } from "../api/api"

export default {
  data() {
    return {
      dialogVisible: false
    }
  },

  props: {
    articleInfo: Object
  },

  methods: {
    newArticleFunc() {
      this.$router.push({
        path: "/NewArticle"
      })
    },

    deleteArticle() {
      DeleteArticle(this.articleInfo._id)
      .then(res => {
        this.$message.success("删除成功")
        this.$emit("changeData")
      })
      .catch(() => {
        this.$message.error("删除失败")
      })
      this.dialogVisible = false
    }
  }
}
</script>

<style>
.newArticleBox {
  text-align: center;
  height: 70px;
  line-height: 70px;
  font-size: 40px;
  font-weight: 400;
  cursor: pointer;
}

/* .classLabel  {
  left:-18px!important;
  padding-left:18px!important;
} */

/* .classLabel::after{
  border:none!important;
} */

.btnBox {
  text-align: center;
  margin-top: 10px;
  line-height: 20px;
}
</style>

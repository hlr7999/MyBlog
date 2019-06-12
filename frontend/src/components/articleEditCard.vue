<template>
<div class="editBox">
  <div class="editFormItem">
    <div class="editFormTitle">标题:</div>
    <el-input 
      placeholder="输入标题"
      v-model="articleInfo.title">
    </el-input>
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">描述:</div>
    <el-input 
      placeholder="输入描述" 
      type="textarea"
      :rows="2"
      v-model="articleInfo.description">
    </el-input>
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">一级分类:</div>
    <el-select v-model="articleInfo.firstClass" placeholder="请选择一级分类">
      <el-option
        v-for="item in firstClassList"
        :key="item.id"
        :label="item.name"
        :value="item.id">
      </el-option>
    </el-select>
    <el-button type="text" class="newClass" @click="newFirstClass">新建</el-button>
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">二级分类:</div>
    <el-select v-model="articleInfo.secondClass" placeholder="请选择二级分类">
      <el-option
        v-for="item in secondClassList"
        :key="item.id"
        :label="item.name"
        :value="item.id">
      </el-option>
    </el-select>
    <el-button type="text" class="newClass" @click="newSecondClass">新建</el-button>
  </div>

  <div class="editFormItem">
    
  </div>

  <div class="editFormItem">
    <el-button class="submit" @click="submitEdit">完成</el-button>
    <el-button class="cancel" @click="cancelEdit">取消</el-button>
  </div>

  <el-dialog
    :title = "'新建'+newClassLevel+'级分类'"
    :visible.sync = "dialogVisible"
    width = "40%">
    <el-input placeholder="输入分类" v-model="newClassName"></el-input>
    <span slot="footer" class="dialog-footer">
      <el-button class="cancel" @click="dialogVisible = false">取 消</el-button>
      <el-button class="submit" @click="newClass">确 定</el-button>
    </span>
  </el-dialog>

</div>
</template>

<script>
import { GetArticle } from "../api/api"

export default {
    data() {
    return {
      aid: "",
      articleInfo: {
        title: "",
        description: "",
        firstClass: "",
        secondClass: "",
        content: "",
        imageCover: ""
      },
      firstClassList: [],
      secondClassList: [],
      newClassLevel: "",
      newClassName: "",
      dialogVisible: false
    }
  },

  methods: {
    getData() {
      if (!this.$store.state.hasLogin || !this.$store.state.isAdmin) {
        this.$router.push({
          path: "/Forbidden"
        })
      }

      this.aid = this.$route.params.article_id
      if (this.aid) {
        GetArticle(this.aid)
        .then(res => {
          this.articleInfo = res.data 
        })
        .catch(() => {
          this.$message({
            message: "错误",
            type: "error",
            duration: 1500
          })
        })
      }
    },

    newFirstClass() {
      this.newClassLevel = "一"
      this.dialogVisible = true
    },

    newSecondClass() {
      this.newClassLevel = "二"
      this.dialogVisible = true
    },

    newClass() {

    },

    $imgAdd(pos, $file) {
      console.log(pos, $file)
    },

    submitEdit() {

    },

    cancelEdit() {
      this.$router.push({
        path: '/Admin'
      })
    }
  },

  mounted () {
    this.getData()
  }
}
</script>

<style>
.editBox {
  background-color: #fff;
  font-size: 15px;
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  overflow-x: hidden;
  padding:15px;
  border-radius: 5px;
}

.editBox .editFormItem {
  margin: 15px 20px;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: flex-start;
  align-items: center;
}

.editBox .editFormItem .editFormTitle {
  width: 100px;
  text-align: left;
}

.editBox .el-input,
.editBox .el-textarea,
.editBox .el-select {
  width: 60%;
}

.editBox .el-select .el-input {
  width: 100%;
}

.editBox .markdown-body .v-note-panel {
  border: 1px solid #e0e0e0;
}

.editBox .submit {
  background-color: #97dffd;
}

.editBox .cancel {
  background-color: #c0c4cc;
}

.editBox .newClass {
  color: #67c23a;
  margin-left: 15px;
}

.editBox .el-dialog .el-input {
  width: 70%;
}

@media screen and (max-width:800px) {
  .editBox .el-dialog {
    width: 96% !important;
  }
}
</style>

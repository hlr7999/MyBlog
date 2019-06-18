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
    <el-select 
      v-model="articleInfo.firstClass" 
      placeholder="请选择一级分类"
      @change="changeFirstClass">
      <el-option
        v-for="item in firstClassList"
        :key="item._id"
        :label="item.name"
        :value="item._id">
      </el-option>
    </el-select>
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">二级分类:</div>
    <el-select v-model="articleInfo.secondClass" placeholder="请选择二级分类">
      <el-option
        v-for="item in secondClassList"
        :key="item._id"
        :label="item.name"
        :value="item">
      </el-option>
    </el-select>
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">文章内容:</div>
    <el-button type="primary" size="small" @click="contentFile">选择文件</el-button>
    {{ articleInfo.content }}
    <input type="file" class="fileUploadInput" accept=".md"
      id="contentInput" @change="uploadContent($event)">
  </div>

  <div class="editFormItem">
    <div class="editFormTitle">文章封面:</div>
    <el-button type="primary" size="small" @click="coverFile">选择文件</el-button>
    {{ articleInfo.imageCover }}
    <input type="file" class="fileUploadInput"  accept="image/jpg, image/jpeg, image/png"
      id="coverInput" @change="uploadCover($event)">
  </div>

  <div class="editFormItem">
    <el-button class="submit" @click="submitEdit">完成</el-button>
    <el-button class="cancel" @click="cancelEdit">取消</el-button>
  </div>

</div>
</template>

<script>
import { GetArticle, NewArticle } from "../api/api"
import { GetClasses, GetSecondClass } from "../api/api"

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
      formData: new FormData(),
      firstClassList: [],
      secondClassList: [],
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

      GetClasses()
      .then(res => {
        if (res.data) {
          this.firstClassList = res.data
        }
      })
      .catch(() => {
        this.$message.error("错误")
      })
    },

    changeFirstClass(item) {
      this.articleInfo.secondClass = ""
      GetSecondClass(item)
      .then(res => {
        if (res.data) {
          this.secondClassList = res.data
        }
      })
      .catch(() => {
        this.$message.error("错误")
      })
    },

    contentFile() {
      var file = document.getElementById("contentInput")
      file.click()
    },

    coverFile() {
      var file = document.getElementById("coverInput")
      file.click()
    },

    uploadContent(e) {
      var files = e.target.files
      if (files.length > 1) {
        this.$message.error("选择文件过多")
      }

      var file = files[0]
      if (file.size / 1024 / 1024 > 1) {
        this.$message.error("文件大小不能超过 1MB!");
        return
      }
      var type = file.name.split(".").pop()
      if (type != "md" && type != "markdown") {
        this.$message.error("只能是 markdown 文件!");
      }

      this.formData.append("content", file)
      this.articleInfo.content = file.name
    },

    uploadCover(e) {
      var files = e.target.files
      if (files.length > 1) {
        this.$message.error("选择文件过多")
      }

      var file = files[0]
      if (file.size / 1024 / 1024 > 1) {
        this.$message.error("图片大小不能超过 1MB!");
        return
      }
      if (file.type != "image/png" && file.type != "image/jpg" 
          && file.type != "image/jpeg") {
        this.$message.error("图片只能是 JPG/JPEG/PNG 格式!");
      }

      this.formData.append("imageCover", file)
      this.articleInfo.imageCover = file.name
    },

    submitEdit() {
      if (!this.articleInfo.title || !this.articleInfo.description ||
          !this.articleInfo.firstClass || !this.articleInfo.secondClass ||
          !this.articleInfo.content) {
        this.$message.error("信息未填写完整")
        return
      }

      this.formData.append("title", this.articleInfo.title)
      this.formData.append("description", this.articleInfo.description)
      this.formData.append("classId", this.articleInfo.secondClass._id)
      this.formData.append("className", this.articleInfo.secondClass.name)
      NewArticle(this.formData)
      .then(res => {
        this.$message.success("创建成功")
        this.$router.push("/Admin")
      })
      .catch(() => {
        this.$message.error("错误")
      })
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

.editBox .fileUploadInput {
  display: none;
}

@media screen and (max-width:800px) {
  .editBox .el-dialog {
    width: 96% !important;
  }
}
</style>

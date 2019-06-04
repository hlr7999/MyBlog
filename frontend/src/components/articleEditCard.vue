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
    <mavon-editor
      :value = "articleInfo.content"
      :boxShadow = "false"
      :ishljs = "true"
      code-style = "github-gist"
      :toolbars = "toolbars"
    />
  </div>

  <div class="editFormItem">
    <el-button class="submit" @click="submitEdit">完成</el-button>
    <el-button class="cancel" @click="cancelEdit">取消</el-button>
  </div>

</div>
</template>

<script>
export default {
    data() {
    return {
      isNew: true,
      aid: "",
      articleInfo: {
        title: "",
        description: "",
        firstClass: "",
        secondClass: "",
        content: ""
      },
      firstClassList: [{
        id: "0",
        name: "技术分享"
      }, {
        id: "1",
        name: "生活随笔"
      }, {
        id: "2",
        name: "读书分享"
      }],
      secondClassList: [{
        id: "0",
        name: "Vue.js"
      }, {
        id: "1",
        name: "Golang"
      }],
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        strikethrough: true, // 中划线
        mark: true, // 标记
        superscript: true, // 上角标
        subscript: true, // 下角标
        quote: true, // 引用
        ol: true, // 有序列表
        ul: true, // 无序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        code: true, // code
        table: true, // 表格
        fullscreen: false,
        readmodel: false,
        htmlcode: true, // 展示html源码
        help: false,
        undo: false,
        redo: false,
        trash: true,
        save: false,
        navigation: false,
        alignleft: true, // 左对齐
        aligncenter: true, // 居中
        alignright: true, // 右对齐
        subfield: true,
        preview: true,
      }
    }
  },

  methods: {
    getData() {
      this.aid = this.$route.params.article_id
      if (this.aid) {
        this.isNew = false
      }

      if (!this.isNew) {
        // get article info
      }
    },

    newFirstClass() {
      
    },

    newSecondClass() {

    },

    submitEdit() {

    },

    cancelEdit() {
      this.$router.push({
        path: '/'
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

.markdown-body .v-note-panel {
  border: 1px solid #e0e0e0;
}

.submit {
  background-color: #97dffd;
}

.cancel {
  background-color: #c0c4cc;
}

.newClass {
  color: #67c23a;
  margin-left: 15px;
}
</style>

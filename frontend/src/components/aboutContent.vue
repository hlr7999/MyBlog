<template>
  <div class="aboutContentBox a-articleInfo">
    <header>
      <h1>
        <a :href="'#/AboutMe'">{{aboutInfo.title}}</a>
      </h1>
    </header>

    <mavon-editor
      :value = "aboutInfo.content"
      :boxShadow = "false"
      :subfield = "false"
      defaultOpen = "preview"
      :editable="false"
      :toolbarsFlag	= "false"
      :ishljs = "true"
      code-style = "github-gist"
      :shortCut = "false"
    />

  </div>
</template>

<script>
import { AboutMe } from "../api/api"

export default {
  data() {
    return {
      aid: "0",
      aboutInfo: "",
      likeArt: false,
      collectArt: false,
      haslogin: false,
      userId: "",
    }
  },

  methods: {
    getData: function() {
     
      this.aboutInfo = {
        title: "关于我",
        content: "# Not found"
      }

      this.$store.commit("changeLoading", true)
      AboutMe()
      .then(res => {
        this.aboutInfo.content = res.data
      })
      .catch(() => {
        this.$message.error({
          message: '未知错误',
          type: 'error',
          showClose: true,
          duration: 2000
        })
        this.$store.commit("changeLoading", false)
      })
    },

  },

  mounted () {
    this.getData();
  }
};
</script>

<style>
.aboutContentBox .article-content {
  font-size: 15px;
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  overflow-x: hidden;
}

.aboutContentBox .markdown-body {
  min-width: 0 !important;
}

.aboutContentBox .markdown-body .v-note-panel {
  border: 0 !important;
}

.aboutContentBox .markdown-body .v-note-panel .v-show-content {
  background-color: #fff !important;
}

.aboutContentBox .markdown-body .highlight pre, 
.aboutContentBox .markdown-body pre {
  padding: 3px !important;
}

.aboutContentBox .markdown-body ul {
  list-style: disc;
}

</style>

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
        id: "0",
        title: "关于我",
        content: "# Not found"
      }

      const axios = require('axios')
      axios.get('http://localhost:10080/aboutMe')
        .then(res => {
          this.aboutInfo.content = res.data
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

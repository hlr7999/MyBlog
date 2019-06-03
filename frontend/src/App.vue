<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.classLabel {
  border-color:#97dffd;
  box-shadow: 0 2px 2px 0 rgba(0,0,0,.14), 0 3px 1px -2px rgba(0,0,0,.2), 0 1px 5px 0 rgba(0,0,0,.12);
  background-color: #97dffd;
  margin:5px 0 15px 2px;
  font-size: 16px;
  text-align: left;
  position: relative;
  left:-33px;
  border-radius: 0 4px 4px 0;
  padding:7px 11.2px 7px 32px;
  display: inline-block;
  color:#fff;
}

.classLabel ::after {
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

.classLabel a {
  color:#fff;
  line-height: 100%;
}

a {
  text-decoration: none;
  color: black;
}

.container {
  max-width: 50%;
  margin:100px auto;
  padding:0 10px;
  transition: all 0.5s ease-out;
  font-size: 16px;
}

.articleItem {
  margin: 10px auto;
}

.a-round-date {
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

.a-round-date span {
  text-align: center;
  display: block;
}

.a-round-date .day {
  font-size: 30px;
  font-weight: 700;
}

.a-articleInfo {
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

.a-articleInfo header {
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
  text-align: left;
}

.a-articleInfo header h1 {
  margin: 10px 0;
  font-size: 25px;
  font-weight: 700;
  text-align: center;
  line-height: 30px;
}

.a-articleInfo header h2 {
  margin:10px 0;
  line-height: 24px;
  text-align: center;
  color:#555;
  font-size: 14px;
}

.a-articleInfo header h2 .el-rate,
.rateBox{
  display:inline-block;
}

.a-articleInfo header h2 .el-rate i{
  font-size: 16px;
}

.a-articleInfo .article-coontent {
  font-size: 15px;
  white-space: normal;
  word-wrap: break-word;
  word-break: break-all;
}

.a-articleInfo .article-content p {
  margin:10px 0;
  line-height:24px;
  text-align: left;
}

.a-articleInfo .viewdetail {
  margin:10px 0 ;
  line-height: 24px;
  text-align: center;
}

.a-articleInfo .viewdetail a {
  color:#fff;
  font-size: 12px;
  padding:5px 10px;
  border-radius: 5px;
}

.readWholeArticle {
  background: #97dffd;
  transition: all .3s ease-in-out;
}

.readWholeArticle:hover {
  background: #48456D;
}

.viewmore {
  text-align: center;
  width:100%;
}

.viewmore a {
  background-color: #97dffd;
  text-decoration: none;
  border-radius: 5px;
  text-align: center;
  font-size: 16px;
  color:#fff;
  height:30px;
  line-height: 30px;
  display: block;
}

.viewMoreBtn {
  background: #97dffd;
  transition: all .3s ease-in-out;
}

.viewMoreBtn:hover {
  background: #48456D;
}

@media screen and (max-width:1400px) {
  .container {
    max-width: 60%!important;
    margin:100px auto;
  }
}

@media screen and (max-width:1000px) {
  .container{
    max-width: 80%!important;
    margin:100px auto;
  }
}

@media screen and (max-width:800px) {
  .container{
    max-width: 100%!important;
    margin:100px auto;
  }
  .a-round-date {
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

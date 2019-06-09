<template>
  <div class="commenBox">
    <div class="commenInputBox">
      <h3> 发表评论 </h3>
      <div>
        <el-input type="textarea" :rows="2" placeholder="留下宝贵的评论吧~" v-model="commentText"></el-input>
        <div :class="isEmojiOpen ? 'emoji' : 'emoji emoji-open'">
          <div class="emoji-logo" @click="isEmojiOpen = !isEmojiOpen">
            <span>emoji表情</span>
          </div>
          <div class="emoji-body">
            <ul class="emoji-items emoji-items-show">
              <li
                class="emoji-item"
                v-for="(item,index) in emojilist"
                :key="index"
                @click="choseEmoji(item.title)"
              >
                <img :src="'static/img/emoji/' + item.url" alt>
              </li>
            </ul>
          </div>
        </div>
        <el-row class="commentSendBox">
          <el-col :span="24" class="commentSendBtn" :onclick="sendComment">
            {{ sendTip }}
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="commentContent" ref="commentContent">
      <a class="commentContentTitle">共 {{ commentList ? commentList.length : 0}} 条</a>
      <div class="commentContentBody">
        <ul class="commentContentList">
          <li class="commentContentListItem" v-for="(item,index) in commentList" :key="index">
            <div class="commentItem">
              <header>
                <img :src="item.avatar">
                <div class="i-name">{{item.username}}</div>
                <div class="i-time">{{item.time}}</div>
              </header>
              <section>
                <p v-html="analyzeEmoji(item.content)">{{analyzeEmoji(item.content)}}</p>
                <div
                  v-if="haslogin"
                  class="replyBtn"
                  @click="replyComment(item.id)"
                >回复</div>
              </section>
            </div>
            <ul v-show="item.sonComment" class="commentContentList" style="padding-left:60px;">
              <li
                class="commentContentListItem"
                v-for="(item,index) in item.sonComment"
                :key="index"
              >
                <div class="commentItem">
                  <header>
                    <img :src="item.avatar">
                    <div class="i-name">
                      {{item.username}}
                      <span>回复</span>
                      {{item.reply_name}}
                    </div>
                    <div class="i-time">{{ item.time }}</div>
                  </header>
                  <section>
                    <p v-html="analyzeEmoji(item.content)">{{item.content}}</p>
                    <div
                      v-show="haslogin"
                      class="replyBtn"
                      @click="replyComment(item.comment_id)"
                    >回复</div>
                  </section>
                </div>
              </li>
            </ul>
          </li>
        </ul>
        <div class="viewmore">
          <h1 v-show="hasMore" class="viewMoreBtn" @click="addMoreFun">查看更多</h1>
          <h1 v-show="!hasMore" class="viewMoreBtn">没有更多</h1>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "comment",

  data() {
    //选项 / 数据
    return {
      commenInputBox: "",
      commentContent: "",
      commentText: "",
      isEmojiOpen: true, 
      commentList: "",
      pageId: 0,
      aid: 0,
      hasMore: true,
      haslogin: false,
      userId: "",
      // leaveId: 0,
      // leavePid: "",
      // pid: "",
      sendTip: "发送",
      emojilist: [
        { title: "微笑", url: "weixiao.gif" },
        { title: "嘻嘻", url: "xixi.gif" },
        { title: "哈哈", url: "haha.gif" },
        { title: "可爱", url: "keai.gif" },
        { title: "可怜", url: "kelian.gif" },
        { title: "挖鼻", url: "wabi.gif" },
        { title: "吃惊", url: "chijing.gif" },
        { title: "害羞", url: "haixiu.gif" },
        { title: "挤眼", url: "jiyan.gif" },
        { title: "闭嘴", url: "bizui.gif" },
        { title: "鄙视", url: "bishi.gif" },
        { title: "爱你", url: "aini.gif" },
        { title: "泪", url: "lei.gif" },
        { title: "偷笑", url: "touxiao.gif" },
        { title: "亲亲", url: "qinqin.gif" },
        { title: "生病", url: "shengbing.gif" },
        { title: "太开心", url: "taikaixin.gif" },
        { title: "白眼", url: "baiyan.gif" },
        { title: "右哼哼", url: "youhengheng.gif" },
        { title: "左哼哼", url: "zuohengheng.gif" },
        { title: "嘘", url: "xu.gif" },
        { title: "衰", url: "shuai.gif" },
        { title: "吐", url: "tu.gif" },
        { title: "哈欠", url: "haqian.gif" },
        { title: "抱抱", url: "baobao.gif" },
        { title: "怒", url: "nu.gif" },
        { title: "疑问", url: "yiwen.gif" },
        { title: "馋嘴", url: "chanzui.gif" },
        { title: "拜拜", url: "baibai.gif" },
        { title: "思考", url: "sikao.gif" },
        { title: "汗", url: "han.gif" },
        { title: "困", url: "kun.gif" },
        { title: "睡", url: "shui.gif" },
        { title: "钱", url: "qian.gif" },
        { title: "失望", url: "shiwang.gif" },
        { title: "酷", url: "ku.gif" },
        { title: "色", url: "se.gif" },
        { title: "哼", url: "heng.gif" },
        { title: "鼓掌", url: "guzhang.gif" },
        { title: "晕", url: "yun.gif" },
        { title: "悲伤", url: "beishang.gif" },
        { title: "抓狂", url: "zhuakuang.gif" },
        { title: "黑线", url: "heixian.gif" },
        { title: "阴险", url: "yinxian.gif" },
        { title: "怒骂", url: "numa.gif" },
        { title: "互粉", url: "hufen.gif" },
        { title: "书呆子", url: "shudaizi.gif" },
        { title: "愤怒", url: "fennu.gif" },
        { title: "感冒", url: "ganmao.gif" },
        { title: "心", url: "xin.gif" },
        { title: "伤心", url: "shangxin.gif" },
        { title: "猪", url: "zhu.gif" },
        { title: "熊猫", url: "xiongmao.gif" },
        { title: "兔子", url: "tuzi.gif" },
        { title: "喔克", url: "ok.gif" },
        { title: "耶", url: "ye.gif" },
        { title: "棒棒", url: "good.gif" },
        { title: "不", url: "no.gif" },
        { title: "赞", url: "zan.gif" },
        { title: "来", url: "lai.gif" },
        { title: "弱", url: "ruo.gif" },
        { title: "草泥马", url: "caonima.gif" },
        { title: "神马", url: "shenma.gif" },
        { title: "囧", url: "jiong.gif" },
        { title: "浮云", url: "fuyun.gif" },
        { title: "给力", url: "geili.gif" },
        { title: "围观", url: "weiguan.gif" },
        { title: "威武", url: "weiwu.gif" },
        { title: "话筒", url: "huatong.gif" },
        { title: "蜡烛", url: "lazhu.gif" },
        { title: "蛋糕", url: "dangao.gif" },
        { title: "发红包", url: "fahongbao.gif" }
      ]
    };
  },
  methods: {
    getData() {
      this.commentList = []
    },

    choseEmoji: function(inner) {
      this.commentText += "[" + inner + "]";
    },

    analyzeEmoji: function(cont) {
      var pattern1 = /\[[\u4e00-\u9fa5]+\]/g;
      var pattern2 = /\[[\u4e00-\u9fa5]+\]/;
      var content = cont.match(pattern1);
      var str = cont;
      if (content) {
        for (var i = 0; i < content.length; i++) {
          for (var j = 0; j < this.emojilist.length; j++) {
            if ("[" + this.emojilist[j].title + "]" == content[i]) {
              var src = this.emojilist[j].url;
              break;
            }
          }
          str = str.replace(
            pattern2,
            '<img src="static/img/emoji/' + src + '"/>'
          );
        }
      }
      return str;
    },

    //发送留言
    sendComment: function() {
      //留言
      var that = this
      console.log(this.commentText)
      if (this.commentText) {
        this.sendTip = "发送中";
        // if (that.leaveId == 0) {
        //   setcommentItemComment(
        //     that.textarea,
        //     that.userId,
        //     that.aid,
        //     that.leavePid,
        //     that.pid,
        //     function(msg) {
        //       //   console.log(msg);
        //       that.textarea = "";
        //       that.routeChange();
        //       that.removeRespond();
        //       var timer02 = setTimeout(function() {
        //         that.sendTip = "发送~";
        //         clearTimeout(timer02);
        //       }, 1000);
        //     }
        //   );
        // } else {
        //   //其他模块留言回复
        //   setOuthComment(
        //     that.textarea,
        //     that.userId,
        //     that.aid,
        //     that.leaveId,
        //     that.leavePid,
        //     that.pid,
        //     function(msg) {
        //       //   console.log(msg);
        //       that.textarea = "";
        //       that.removeRespond();
        //       that.routeChange();
        //     }
        //   );
        // }
        this.commentText = ""
        this.getData()
        var timer = setTimeout(function() {
          that.sendTip = "发送"
          clearTimeout(timer)
        }, 1000)
      } else {
        this.sendTip = "内容不能为空"
        this.commentText = ""
        this.getData()
        var timer = setTimeout(function() {
          that.sendTip = "发送"
          clearTimeout(timer)
        }, 3000)
      }
      console.log(this.sendTip)
    },

    replyComment: function(leavePid, pid) {
      //回复留言
      // console.log(leavePid,pid);
      // var that = this;
      // if (localStorage.getItem("userInfo")) {
      //   var dom = event.currentTarget;
      //   dom = dom.parentNode;
      //   this.leavePid = leavePid;
      //   this.pid = pid;
      //   dom.appendChild(this.$refs.commenInputBox);
      // } else {
      //   that
      //     .$confirm("登录后即可点赞和收藏，是否前往登录页面?", "提示", {
      //       confirmButtonText: "确定",
      //       cancelButtonText: "取消",
      //       type: "warning"
      //     })
      //     .then(() => {
      //       //确定，跳转至登录页面
      //       //储存当前页面路径，登录成功后跳回来
      //       localStorage.setItem("logUrl", that.$route.fullPath);
      //       that.$router.push({ path: "/Login?login=1" });
      //     })
      //     .catch(() => {});
      // }
    },

    addMoreFun: function() {
      
    }
  },
  
  watch: {
    '$route':'getData'
  },

  mounted() {
    this.getData()
  }
};
</script>

<style>
.commenBox {
  position: relative;
  background: #fff;
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 5px;
}

.commenBox ul {
  list-style: none;
}

.commenBox .commenInputBox h3 {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 8px;
  text-align: left;
}

.commenBox .commenInputBox textarea {
  background: #f4f6f7;
  height: 100px;
  margin-bottom: 10px;
}

.commenBox .emoji {
  position: relative;
  z-index: 1;
  text-align: left;
}

.commenBox .emoji .emoji-logo {
  position: relative;
  border-radius: 4px;
  color: #444;
  display: inline-block;
  background: #fff;
  border: 1px solid #ddd;
  font-size: 13px;
  padding: 0 6px;
  cursor: pointer;
  height: 30px;
  box-sizing: border-box;
  z-index: 2;
  line-height: 30px;
}

.commenBox .emoji .emoji-logo:hover {
  animation: a 5s infinite ease-in-out;
  -webkit-animation: a 5s infinite ease-in-out;
}

.commenBox .emoji .emoji-body {
  position: absolute;
  background: #fff;
  border: 1px solid #ddd;
  z-index: 1;
  top: 29px;
  border-radius: 0 4px 4px 4px;
  display: none;
}

.commenBox .emoji-open .emoji-body {
  display: block;
}

.commenBox .emoji-open .emoji-logo {
  border-radius: 4px 4px 0 0;
  border-bottom: none;
}

.commenBox .emoji-open .emoji-logo:hover {
  animation: none;
  -webkit-animation: none;
}

.commenBox .emoji .emoji-items {
  max-height: 197px;
  overflow: scroll;
  font-size: 0;
  padding: 10px;
  z-index: 1;
}

.commenBox .emoji .emoji-items .emoji-item {
  background: #f7f7f7;
  padding: 5px 10px;
  border-radius: 5px;
  display: inline-block;
  margin: 0 10px 12px 0;
  transition: 0.3s;
  line-height: 19px;
  font-size: 20px;
  cursor: pointer;
}

.commenBox .emoji .emoji-items .emoji-item:hover {
  background: #eee;
  box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2),
    0 1px 5px 0 rgba(0, 0, 0, 0.12);
  animation: a 5s infinite ease-in-out;
  -webkit-animation: a 5s infinite ease-in-out;
}

@-webkit-keyframes a {
  2% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  4% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  6% {
    -webkit-transform: translateY(1.5px) rotate(-1.5deg);
    transform: translateY(1.5px) rotate(-1.5deg);
  }
  8% {
    -webkit-transform: translateY(-1.5px) rotate(-1.5deg);
    transform: translateY(-1.5px) rotate(-1.5deg);
  }
  10% {
    -webkit-transform: translateY(2.5px) rotate(1.5deg);
    transform: translateY(2.5px) rotate(1.5deg);
  }
  12% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  14% {
    -webkit-transform: translateY(-1.5px) rotate(1.5deg);
    transform: translateY(-1.5px) rotate(1.5deg);
  }
  16% {
    -webkit-transform: translateY(-0.5px) rotate(-1.5deg);
    transform: translateY(-0.5px) rotate(-1.5deg);
  }
  18% {
    -webkit-transform: translateY(0.5px) rotate(-1.5deg);
    transform: translateY(0.5px) rotate(-1.5deg);
  }
  20% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  22% {
    -webkit-transform: translateY(0.5px) rotate(-1.5deg);
    transform: translateY(0.5px) rotate(-1.5deg);
  }
  24% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  26% {
    -webkit-transform: translateY(0.5px) rotate(0.5deg);
    transform: translateY(0.5px) rotate(0.5deg);
  }
  28% {
    -webkit-transform: translateY(0.5px) rotate(1.5deg);
    transform: translateY(0.5px) rotate(1.5deg);
  }
  30% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  32%,
  34% {
    -webkit-transform: translateY(1.5px) rotate(-0.5deg);
    transform: translateY(1.5px) rotate(-0.5deg);
  }
  36% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  38% {
    -webkit-transform: translateY(1.5px) rotate(-1.5deg);
    transform: translateY(1.5px) rotate(-1.5deg);
  }
  40% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  42% {
    -webkit-transform: translateY(2.5px) rotate(-1.5deg);
    transform: translateY(2.5px) rotate(-1.5deg);
  }
  44% {
    -webkit-transform: translateY(1.5px) rotate(0.5deg);
    transform: translateY(1.5px) rotate(0.5deg);
  }
  46% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  48% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  50% {
    -webkit-transform: translateY(0.5px) rotate(0.5deg);
    transform: translateY(0.5px) rotate(0.5deg);
  }
  52% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  54% {
    -webkit-transform: translateY(-1.5px) rotate(1.5deg);
    transform: translateY(-1.5px) rotate(1.5deg);
  }
  56% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  58% {
    -webkit-transform: translateY(0.5px) rotate(2.5deg);
    transform: translateY(0.5px) rotate(2.5deg);
  }
  60% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  62% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  64% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  66% {
    -webkit-transform: translateY(1.5px) rotate(-0.5deg);
    transform: translateY(1.5px) rotate(-0.5deg);
  }
  68% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  70% {
    -webkit-transform: translateY(1.5px) rotate(0.5deg);
    transform: translateY(1.5px) rotate(0.5deg);
  }
  72% {
    -webkit-transform: translateY(2.5px) rotate(1.5deg);
    transform: translateY(2.5px) rotate(1.5deg);
  }
  74% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  76% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  78% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  80% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  82% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  84% {
    -webkit-transform: translateY(1.5px) rotate(2.5deg);
    transform: translateY(1.5px) rotate(2.5deg);
  }
  86% {
    -webkit-transform: translateY(-1.5px) rotate(-1.5deg);
    transform: translateY(-1.5px) rotate(-1.5deg);
  }
  88% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  90% {
    -webkit-transform: translateY(2.5px) rotate(-0.5deg);
    transform: translateY(2.5px) rotate(-0.5deg);
  }
  92% {
    -webkit-transform: translateY(0.5px) rotate(-0.5deg);
    transform: translateY(0.5px) rotate(-0.5deg);
  }
  94% {
    -webkit-transform: translateY(2.5px) rotate(0.5deg);
    transform: translateY(2.5px) rotate(0.5deg);
  }
  96% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  98% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  0%,
  to {
    -webkit-transform: translate(0) rotate(0deg);
    transform: translate(0) rotate(0deg);
  }
}

@keyframes a {
  2% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  4% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  6% {
    -webkit-transform: translateY(1.5px) rotate(-1.5deg);
    transform: translateY(1.5px) rotate(-1.5deg);
  }
  8% {
    -webkit-transform: translateY(-1.5px) rotate(-1.5deg);
    transform: translateY(-1.5px) rotate(-1.5deg);
  }
  10% {
    -webkit-transform: translateY(2.5px) rotate(1.5deg);
    transform: translateY(2.5px) rotate(1.5deg);
  }
  12% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  14% {
    -webkit-transform: translateY(-1.5px) rotate(1.5deg);
    transform: translateY(-1.5px) rotate(1.5deg);
  }
  16% {
    -webkit-transform: translateY(-0.5px) rotate(-1.5deg);
    transform: translateY(-0.5px) rotate(-1.5deg);
  }
  18% {
    -webkit-transform: translateY(0.5px) rotate(-1.5deg);
    transform: translateY(0.5px) rotate(-1.5deg);
  }
  20% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  22% {
    -webkit-transform: translateY(0.5px) rotate(-1.5deg);
    transform: translateY(0.5px) rotate(-1.5deg);
  }
  24% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  26% {
    -webkit-transform: translateY(0.5px) rotate(0.5deg);
    transform: translateY(0.5px) rotate(0.5deg);
  }
  28% {
    -webkit-transform: translateY(0.5px) rotate(1.5deg);
    transform: translateY(0.5px) rotate(1.5deg);
  }
  30% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  32%,
  34% {
    -webkit-transform: translateY(1.5px) rotate(-0.5deg);
    transform: translateY(1.5px) rotate(-0.5deg);
  }
  36% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  38% {
    -webkit-transform: translateY(1.5px) rotate(-1.5deg);
    transform: translateY(1.5px) rotate(-1.5deg);
  }
  40% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  42% {
    -webkit-transform: translateY(2.5px) rotate(-1.5deg);
    transform: translateY(2.5px) rotate(-1.5deg);
  }
  44% {
    -webkit-transform: translateY(1.5px) rotate(0.5deg);
    transform: translateY(1.5px) rotate(0.5deg);
  }
  46% {
    -webkit-transform: translateY(-1.5px) rotate(2.5deg);
    transform: translateY(-1.5px) rotate(2.5deg);
  }
  48% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  50% {
    -webkit-transform: translateY(0.5px) rotate(0.5deg);
    transform: translateY(0.5px) rotate(0.5deg);
  }
  52% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  54% {
    -webkit-transform: translateY(-1.5px) rotate(1.5deg);
    transform: translateY(-1.5px) rotate(1.5deg);
  }
  56% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  58% {
    -webkit-transform: translateY(0.5px) rotate(2.5deg);
    transform: translateY(0.5px) rotate(2.5deg);
  }
  60% {
    -webkit-transform: translateY(2.5px) rotate(2.5deg);
    transform: translateY(2.5px) rotate(2.5deg);
  }
  62% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  64% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  66% {
    -webkit-transform: translateY(1.5px) rotate(-0.5deg);
    transform: translateY(1.5px) rotate(-0.5deg);
  }
  68% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  70% {
    -webkit-transform: translateY(1.5px) rotate(0.5deg);
    transform: translateY(1.5px) rotate(0.5deg);
  }
  72% {
    -webkit-transform: translateY(2.5px) rotate(1.5deg);
    transform: translateY(2.5px) rotate(1.5deg);
  }
  74% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  76% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  78% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  80% {
    -webkit-transform: translateY(1.5px) rotate(1.5deg);
    transform: translateY(1.5px) rotate(1.5deg);
  }
  82% {
    -webkit-transform: translateY(-0.5px) rotate(0.5deg);
    transform: translateY(-0.5px) rotate(0.5deg);
  }
  84% {
    -webkit-transform: translateY(1.5px) rotate(2.5deg);
    transform: translateY(1.5px) rotate(2.5deg);
  }
  86% {
    -webkit-transform: translateY(-1.5px) rotate(-1.5deg);
    transform: translateY(-1.5px) rotate(-1.5deg);
  }
  88% {
    -webkit-transform: translateY(-0.5px) rotate(2.5deg);
    transform: translateY(-0.5px) rotate(2.5deg);
  }
  90% {
    -webkit-transform: translateY(2.5px) rotate(-0.5deg);
    transform: translateY(2.5px) rotate(-0.5deg);
  }
  92% {
    -webkit-transform: translateY(0.5px) rotate(-0.5deg);
    transform: translateY(0.5px) rotate(-0.5deg);
  }
  94% {
    -webkit-transform: translateY(2.5px) rotate(0.5deg);
    transform: translateY(2.5px) rotate(0.5deg);
  }
  96% {
    -webkit-transform: translateY(-0.5px) rotate(1.5deg);
    transform: translateY(-0.5px) rotate(1.5deg);
  }
  98% {
    -webkit-transform: translateY(-1.5px) rotate(-0.5deg);
    transform: translateY(-1.5px) rotate(-0.5deg);
  }
  0%,
  to {
    -webkit-transform: translate(0) rotate(0deg);
    transform: translate(0) rotate(0deg);
  }
}

.commenBox .commentSendBox {
  margin: 10px 0;
}

.commenBox .commentSendBox .commentSendBtn {
  margin: 10px 0;
  color: #fff;
  border-radius: 5px;
  cursor: pointer;
  height: 30px;
  line-height: 30px;
  text-align: center;
  background: #97dffd;
  transition: all .3s ease-in-out;
  user-select:none;
}

.commenBox .commentSendBox .commentSendBtn:hover {
  background: #48456D;
}

.commenBox .commentContent .commentContentTitle {
  display: block;
  border-left: 2px solid #363d4c;
  padding: 0 10px;
  margin: 40px 0;
  font-size: 20px;
  text-align: left;
}

.commenBox .commentContentList {
  margin-bottom: 20px;
  text-align: left;
}

.commenBox .commentContentBody > .commentContentList {
  border-bottom: 1px solid #e5eaed;
}

.commenBox .commentContentListItem {
  border-top: 1px solid #e5eaed;
}

.commenBox .commentContentListItem .commentItem {
  margin: 20px 0;
}

.commenBox .commentContentListItem .commentItem header {
  margin-bottom: 10px;
}

.commenBox .commentContentListItem .commentItem header img {
  width: 65px;
  height: 65px;
  border-radius: 50%;
  float: left;
  transition: all 0.4s ease-in-out;
  -webkit-transition: all 0.4s ease-in-out;
  margin-right: 15px;
  object-fit: cover;
}

.commenBox .commentContentListItem .commentItem header img:hover {
  transform: rotate(360deg);
  -webkit-transform: rotate(360deg);
}

.commenBox .commentContentListItem .commentItem header .i-name {
  font-size: 14px;
  margin: 5px 8px 7px 0;
  color: #444;
  font-weight: bold;
  display: inline-block;
}

.commenBox .commentContentListItem .commentItem header .i-time {
  color: #aaa;
  font-size: 12px;
}

.commenBox .commentContentListItem .commentItem section {
  margin-left: 80px;
}

.commenBox .commentContentListItem .commentItem section p img {
  vertical-align: middle;
}

.commenBox .commentContentListItem .commentItem section .replyBtn {
  display: inline-block;
  margin: 10px 0;
  font-size: 12px;
  color: #64609e;
  cursor: pointer;
}
</style>

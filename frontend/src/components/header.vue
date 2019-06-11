<template>
<div class="header">
  <el-row :gutter="20" type="flex" justify="center">

    <el-col :lg="5" :md="5" :sm="6" :xs="6">
      <div @click="homeBtn"  class="headerLogo">
        <i class="el-icon-watermelon"></i>MyBlog
      </div>
    </el-col>

    <el-col :lg="15" :md="15" :sm="15" :xs="13">
      <el-row :gutter="3" type="flex" justify="start" class="headerMenu">
        <el-col :lg="3" :md="3" :sm="4">
          <div :class="activeIndex == 1 ? 'active' : ''" @click="homeBtn">
            <i class="el-icon-s-home"></i>首页</div>
        </el-col>
        <el-col :lg="3" :md="4" :sm="4">
          <div :class="activeIndex == 2 ? 'active' : ''">
            <el-dropdown :showTimeout="100" @command="classSelect" placement="bottom">
              <div class="el-dropdown-link">
                <i class="el-icon-s-grid"></i>分类<i class="el-icon-arrow-down el-icon--right"></i>
              </div>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item 
                  v-for="c in classList"
                  :key="c.id"
                  :command="c.id">{{c.name}}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </el-col>
        <el-col :lg="3" :md="3" :sm="4">
          <div :class="activeIndex == 3 ? 'active' : ''" @click="aboutBtn">
            <i class="el-icon-s-custom"></i>关于我</div>
        </el-col>
        <el-col :lg="3" :md="4" :sm="5">
          <div :class="activeIndex == 4 ? 'active' : ''" @click="startBtn">
            <i class="el-icon-s-promotion"></i>开始场景</div>
        </el-col>
      </el-row>
    </el-col>

    <el-col v-if="!this.$store.state.hasLogin" :lg="4" :md="4" :sm="5" :xs="5">
      <i class="el-icon-user-solid"></i>
      <div class="headerLogin" @click="loginBtn(0)">
        登录
      </div>&nbsp;&nbsp;|
      <div class="headerLogin" @click="loginBtn(1)">
        注册
      </div>
    </el-col>
    <el-col v-else :lg="4" :md="4" :sm="5" :xs="5">
      <el-dropdown 
        :showTimeout="100" 
        @command="userSelect" 
        placement="bottom"
        class="userHeaderItem">
        <div class="el-dropdown-link"> 
          <img :src="this.$store.state.userInfo.userAvatar" class="userHeaderImg"/>
        </div>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item 
            v-for="c in userOpList"
            :key="c.command"
            :command="c.command">{{c.name}}</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </el-col>
    
  </el-row>
</div>
</template>

<script>
export default {
  data() {
    return {
      classList: [{
        id: "0",
        name: "技术分享"
      }, {
        id: "1",
        name: "生活随笔"
      }, {
        id: "2",
        name: "读书分享"
      }],
      userOpList: [{
        name: "个人中心",
        command: "/UserInfo"
      }, {
        name: "收藏列表",
        command: "/Collect"
      }, {
        name: "喜欢列表",
        command: "/Like"
      }, {
        name: "登出",
        command: "logout"
      }],
      isAdmin: this.$store.state.isAdmin
    };
  },

  props: {
    activeIndex: Number
  },

  methods: {
    homeBtn() {
      this.$router.push ({
        path: '/'
      });
    },

    classSelect(command) {
      this.$router.push ({
        path: `/Class/${command}`
      });
    },

    userSelect(command) {
      if (command === "logout") {
        this.Logout()
      } else {
        this.$router.push ({
          path: `${command}`
        });
      }
    },

    aboutBtn() {
      this.$router.push ({
        path: '/Aboutme'
      });
    },

    startBtn() {
      window.location.href = 'http://localhost/start'
    },

    loginBtn(t) {
      if (t == 0) {
        this.$router.push ({
          path: '/Login'
        });
      } else {
        this.$router.push ({
          path: '/Register'
        });
      }
    },

    Logout() {
      localStorage.removeItem("currentUser")
      this.$store.commit("logout")
      this.$router.push({
        path: "/"
      })
    },

    getData() {
      if (this.isAdmin) {
        this.userOpList.unshift({
          name: "管理博客",
          command: "/Admin"
        })
      }
    }
  },

  created () {
    this.getData()
  },
}
</script>

<style>
.header {
  position: fixed;
  top: 0;
  left: 0;
  margin: 0;
  width: 100%;
  background-color: rgba(40, 42, 44, 0.6);
  color: white;
  height: 50px;
  line-height: 50px;
  min-width: 750px;
	font-family: Arial, sans-serif;
  font-size: 16px;
  font-weight: 700;
  /* forbid select text */
  -moz-user-select: none; 
  -o-user-select:none; 
  -khtml-user-select:none; 
  -webkit-user-select:none; 
  -ms-user-select:none; 
  user-select:none;
  z-index: 10080;
}

.header .el-row {
  height: 100%;
}

.headerMenu .el-col .el-dropdown {
  font-size: 16px;
  font-weight: 700;
  color: white;
  width: 100%;
  height: 100%;
}

.userHeaderItem.el-dropdown,
.userHeaderItem.el-dropdown .el-dropdown-link {
  height: 100%;
}

.headerMenu .el-col .el-dropdown .el-dropdown-link {
  width: 100%;
  height: 100%;
}

ul.el-dropdown-menu {
  background-color: rgba(40, 42, 44, 0.6);
  border: 0;
  border-radius: 0%;
  margin: 0;
  padding: 0;
  top: 38px !important;
}

ul.el-dropdown-menu .popper__arrow {
  display: none;
}

ul.el-dropdown-menu li { 
  color: white !important;
  font-size: 16px !important;
}

.header .headerLogo {
  font-size: 30px;
  color: antiquewhite;
  cursor: pointer;
  display: inline-block;
}

.header .headerMenu .el-col:hover div,
.header .headerMenu .el-col:hover a,
ul.el-dropdown-menu li:hover,
.header .active {
  background-color: #48456C !important;
  color: gray !important;
}

.header .headerMenu .el-col {
  cursor: pointer;
}

.header .headerMenu .el-col a {
  color: white;
  text-decoration: none;
}

.header .headerLogin {
  cursor: pointer;
  display: inline-block;
}

.header .headerLogin:hover {
  color: gray
}

.header .userHeaderImg {
  width: 40px;
  height: 40px;
  border-radius: 100%;
  cursor: pointer;
  margin-top: 5px;
  object-fit: cover;
}
</style>

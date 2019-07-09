<template>
<div class="header">
  <el-row
    :gutter="20" 
    type="flex" 
    justify="center"
    class="pcHeader">

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
                  :key="c._id"
                  :command="c._id">{{c.name}}</el-dropdown-item>
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

  <div class="mobileHeader">
    <div class="hideMenu">
      <i @click="showMenu=!showMenu" class="el-icon-menu"></i>
      <el-collapse-transition>
        <el-menu
          class="mlistmenu"
          v-show="showMenu"
          theme="dark"
          :unique-opened="true"
          :router="true">
          <el-menu-item index="/">
            <i class="fa fa-wa fa-home"></i> 首页
          </el-menu-item>
          <el-submenu index="none1">
            <template slot="title">
              <i class="fa fa-wa fa-archive"></i> 分类
            </template>
            <el-menu-item 
              v-for="c in classList"
              :key="c._id"
              :index="'/Class/' + c._id">
              {{c.name}}</el-menu-item>
          </el-submenu>
          <el-menu-item index="/Aboutme">
            <i class="fa fa-wa fa-vcard"></i> 关于我
          </el-menu-item>
          <el-menu-item
            v-show="!this.$store.state.hasLogin"
            index=""
            @click="loginBtn(0)">登录</el-menu-item>
          <el-menu-item
            v-show="!this.$store.state.hasLogin"
            index=""
            @click="loginBtn(1)">注册</el-menu-item>
          <el-submenu v-show="this.$store.state.hasLogin" index="none2">
            <template slot="title">
              <i class="fa fa-wa fa-user-circle-o"></i> 我的
            </template>
            <el-menu-item
              v-for="c in userOpListMobile"
              :key="c.command"
              :index="c.command">{{c.name}}</el-menu-item>
            <el-menu-item @click="Logout">登出</el-menu-item>
          </el-submenu>
        </el-menu>
      </el-collapse-transition>
      <div @click="homeBtn" class="mobileHeaderLogo">
        <i class="el-icon-watermelon"></i>MyBlog
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  data() {
    return {
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
      userOpListMobile: [{
        name: "个人中心",
        command: "/UserInfo"
      }, {
        name: "收藏列表",
        command: "/Collect"
      }, {
        name: "喜欢列表",
        command: "/Like"
      }],
      isAdmin: this.$store.state.isAdmin,
      showMenu: false
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
      window.location.href = 'http://47.100.225.59/start'
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
    },

    routeChange() {
      this.showMenu = false
    }
  },

  created () {
    this.getData()
  },

  computed: {
    classList () {
      return this.$store.state.classList
    }
  },

	watch: {
		'$route': 'routeChange'
	}
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

/* 移动端 */
.mobileHeader {
	position: relative;
	height: 38px;
	line-height: 38px;
	color: #fff;
  display: none;
}

.mobileHeaderLogo {
  font-size: 22px;
  color: antiquewhite;
  display: inline-block;
}

.hideMenu {
	position: relative;
	width: 100%;
	height: 100%;
	line-height: 38px;
}

.hideMenu ul.mlistmenu {
	width: 100%;
	position: absolute;
	left: 0;
	top: 100%;
	box-sizing: border-box;
	z-index: 999;
	box-shadow: 0 2px 6px 0 rgba(0, 0, 0, .12), 0 0 8px 0 rgba(0, 0, 0, .04);
	color: #fff!important;
  background: #48456C!important;
	border-right: none;
}

.hideMenu .el-submenu .el-menu {
	background: #64609E;
}

.hideMenu .el-menu-item,
.hideMenu .el-submenu__title {
	color: #fff!important;
  background: #48456C!important;
}

.hideMenu>i {
	position: absolute;
	left: 10px;
	top: 12px;
	width: 30px;
	height: 30px;
	font-size: 16px;
	color: #fff;
	cursor: pointer;
}

.hideMenu .el-menu-item,
.el-submenu__title {
	height: 40px;
	line-height: 40px;
}

.hideMenu ul.mlistmenu .el-submenu__icon-arrow,
.mobileBox li.el-menu-item a {
	color: #fff;
}

@media screen and (max-width:800px) {
  .mobileHeader {
    display: block;
    width: 100%;
  }
  .pcHeader {
    display: none!important;
  }
  .header {
    height: 38px;
    line-height: 38px;
  }
}
</style>

<template>
<div class="header">
  <el-row :gutter="20" type="flex" justify="center">
    <el-col :lg="5" :md="5" :sm="6" :xs="6" class="headerLogo" @click="homeBtn">
      <i class="el-icon-watermelon"></i>MyBlog</el-col>
    <el-col :lg="15" :md="15" :sm="16" :xs="13">
      <el-row :gutter="3" type="flex" justify="start" class="headerMenu">
        <el-col :lg="3" :md="3" :sm="4">
          <div :class="activeIndex == 1 ? 'active' : ''" @click="homeBtn">
            <i class="el-icon-s-home"></i>首页</div>
        </el-col>
        <el-col :lg="3" :md="3" :sm="4">
          <div :class="activeIndex == 2 ? 'active' : ''">
            <el-dropdown :showTimeout="100" @command="classSelect" placement="bottom-start">
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
        <el-col :lg="3" :md="4" :sm="4">
          <div :class="activeIndex == 3 ? 'active' : ''" @click="aboutBtn">
            <i class="el-icon-s-custom"></i>关于我</div>
        </el-col>
        <el-col :lg="3" :md="4" :sm="5">
          <div :class="activeIndex == 4 ? 'active' : ''" @click="startBtn">
            <i class="el-icon-s-promotion"></i>开始场景</div>
        </el-col>
      </el-row>
    </el-col>
    <el-col :lg="4" :md="4" :sm="4" :xs="5"><div class="headerLogin" @click="loginBtn">
      <i class="el-icon-user-solid"></i>登录|注册</div></el-col>
  </el-row>
</div>
</template>

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
  z-index: 10;
}

.headerMenu .el-col .el-dropdown {
  font-size: 16px;
  font-weight: 700;
  color: white;
  width: 100%;
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

.headerLogo {
  font-size: 30px;
  color: antiquewhite;
  cursor: pointer;
}

.headerMenu .el-col:hover div,
.headerMenu .el-col:hover a,
ul.el-dropdown-menu li:hover,
.active {
  background-color: #48456C !important;
  color: gray !important;
}

.headerMenu .el-col {
  cursor: pointer;
}

.headerMenu .el-col a {
  color: white;
  text-decoration: none;
}

.headerLogin {
  cursor: pointer;
}
</style>

<script>
export default {
  data() {
    return {
      classList: [{
        id: 0,
        name: "技术分享"
      }, {
        id: 1,
        name: "生活随笔"
      }, {
        id: 2,
        name: "读书分享"
      }]
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
    aboutBtn() {
      this.$router.push ({
        path: '/Aboutme'
      });
    },
    startBtn() {
      window.location.href = 'http://localhost/start'
    },
    loginBtn() {
      this.$router.push ({
        path: '/Login'
      });
    },
    getData() {

    }
  }
}
</script>

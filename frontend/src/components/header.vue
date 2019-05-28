<template>
<div class="header">
  <el-row :gutter="20" type="flex" justify="center">
    <el-col :span="5" class="logo" @click="homeBtn">
      <i class="el-icon-watermelon"></i>MyBlog</el-col>
    <el-col :span="14">
      <el-row :gutter="3" type="flex" justify="start" class="menu">
        <el-col :span="3">
          <div :class="activeIndex == 1 ? 'active' : ''" @click="homeBtn">
            <i class="el-icon-s-home"></i>首页</div>
        </el-col>
        <el-col :span="3">
          <div :class="activeIndex == 2 ? 'active' : ''" @click="classBtn">
            <i class="el-icon-s-grid"></i>分类</div>
        </el-col>
        <el-col :span="3">
          <div :class="activeIndex == 3 ? 'active' : ''" @click="aboutBtn">
            <i class="el-icon-s-custom"></i>关于我</div>
        </el-col>
        <el-col :span="4">
          <div :class="activeIndex == 4 ? 'active' : ''" @click="startBtn">
            <i class="el-icon-s-promotion"></i>开始场景</div>
        </el-col>
      </el-row>
    </el-col>
    <el-col :span="4"><div class="login" @click="loginBtn">
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
  font-weight: 700;
  /* forbid select text */
  -moz-user-select: none; 
  -o-user-select:none; 
  -khtml-user-select:none; 
  -webkit-user-select:none; 
  -ms-user-select:none; 
  user-select:none;
}

.logo {
  font-size: 30px;
  cursor: pointer;
}

.menu .el-col:hover div,
.menu .el-col:hover a,
.active {
  background-color: #48456C;
  color: gray;
}

.menu .el-col {
  cursor: pointer;
}

.menu .el-col a {
  color: white;
  text-decoration: none;
}

.login {
  cursor: pointer;
}
</style>

<script>
export default {
  data() {
    return {
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
    classBtn() {
      this.$router.push ({
        path: '/class'
      });
    },
    aboutBtn() {
      this.$router.push ({
        path: '/aboutme'
      });
    },
    startBtn() {
      window.location.href = 'http://localhost/start'
    },
    loginBtn() {
      this.$router.push ({
        path: '/Login'
      });
    }
  }
}
</script>

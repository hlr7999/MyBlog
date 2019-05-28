<template>
<div class="header">
  <el-row :gutter="20" type="flex" justify="center">
    <el-col :span="5" class="logo" @click="homeBtn">MyBlog</el-col>
    <el-col :span="14">
      <el-row :gutter="20" type="flex" justify="start" class="menu">
        <el-col :span="4">
          <div :class="activeIndex == 1 ? 'active' : ''" @click="homeBtn">首页</div>
        </el-col>
        <el-col :span="4">
          <div :class="activeIndex == 2 ? 'active' : ''" @click="classBtn">分类</div>
        </el-col>
        <el-col :span="4">
          <div :class="activeIndex == 3 ? 'active' : ''" @click="aboutBtn">关于我</div>
        </el-col>
        <el-col :span="4">
          <div :class="activeIndex == 4 ? 'active' : ''" @click="startBtn">开始场景</div>
        </el-col>
      </el-row>
    </el-col>
    <el-col :span="5"><div class="login" @click="loginBtn">登录|注册</div></el-col>
  </el-row>
</div>
</template>

<style>
.header {
  position: fixed;
  margin: 0;
  width: 100%;
  background-color: rgba(40, 42, 44, 0.6);
  color: white;
  height: 60px;
  line-height: 60px;
  min-width: 900px;
}

.logo {
  font-size: 30px;
  font-family: 'Microsoft YaHei';
  cursor: pointer;
}

.menu .el-col:hover,
.menu .el-col:hover a,
.active {
  background-color: #48456C;
  color: gray;
}

.menu .el-col {
  margin: 0 5px;
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
      window.location.href = 'http://localhost/ocean'
    },
    loginBtn() {
      this.$router.push ({
        path: '/Login'
      });
    }
  }
}
</script>

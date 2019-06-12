<template>
<div>
  <blog-header></blog-header>

  <div v-if="this.$store.state.hasLogin">
    <el-dialog
      title="您已登录"
      :visible="true"
      width="30%"
      :show-close="false">
      <span>您已经登录了</span>
      <span slot="footer">
        <el-button @click="goHome">确 定</el-button>
      </span>
    </el-dialog>
  </div>

  <div v-else class="login" v-loading.fullscreen.lock="this.$store.state.loading">
    <el-card shadow="always">
      <el-tabs v-model="activeName" @tab-click="handleClick" :stretch="true">
        <el-tab-pane label="注册" name="first"></el-tab-pane>
        <el-tab-pane label="登录" name="second"></el-tab-pane>
      </el-tabs>
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
        <el-form-item label="" prop="usernameORemail">
          <el-input v-model="ruleForm.username" 
            placeholder="请输入用户名" spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="" prop="pass" class="passwordItem">  
          <el-input type='password' v-model="ruleForm.pass" 
            placeholder="请输入密码" spellcheck="false"></el-input>
        </el-form-item>
        <!-- <el-form-item class="forget">
          <a href="javascript:void(0)">忘记密码?</a>
        </el-form-item> -->
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>

</div>
</template>

<script>
import { UserLogin } from "../api/api"
import header from "../components/header.vue"

export default {
  name: "login",
  components: {
    'blog-header': header
  },
  data() {
    return {
      ruleForm: {
        username: '',
        pass: '',
        type: '1'
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { max: 30, message: '输入过长', trigger: 'blur'}
        ],
        pass: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { max: 16, message: '输入过长', trigger: 'blur'}
        ]
      },
      activeName: "second"
    }
  },
  methods: {
    submitForm(formName) {
      var that = this
      this.$refs[formName].validate((valid) => {
        if (valid) { 
          this.$store.commit("changeLoading", true)
          UserLogin(this.ruleForm)
          .then(res => {
            this.$message({
              message: '登录成功',
              type: 'success',
              showClose: true,
              duration: 1500
            })
            this.$store.commit("login", res.data)
            localStorage.setItem("currentUser", JSON.stringify(res.data))
            this.$router.push({
              path: '/'
            })
            this.$store.commit("changeLoading", false)
          })
          .catch(() => {
            this.$message.error({
              message: '用户名或密码错误',
              type: 'error',
              showClose: true,
              duration: 2000
            })
            this.$store.commit("changeLoading", false)
          })
        }
      })
    },

    handleClick(tab, event) {
      if (tab.name === 'first') {
        this.$router.push({
          path: '/Register'
        })
      }
    },

    goHome() {
      this.$router.push({
        path: "/"
      })
    }
  }
}
</script>

<style>
.login .el-tabs {
  padding-top: 10px;
  margin-bottom: 50px;
  font-family: 'Hiragino Sans GB';
}

.login .el-tabs .el-tabs__item {
  font-size: 26px!important;
}

.login .el-card {
  margin: 120px auto;
  max-width: 320px;
  border-width: 1.5px;
}

.login .el-form .el-form-item .el-button {
  width: 100%!important;
}

.login .el-form .el-form-item.passwordItem {
  margin-bottom: 0;
}
/* 
.login .el-form .el-form-item.forget {
  text-align: right!important;
  margin-right: 10px!important;
  margin-top: 0 !important;
  margin-bottom: 0 !important;
}

.login .el-form .el-form-item.forget a {
  color: gray;
} */
</style>

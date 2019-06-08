<template>
<div>
  <blog-header></blog-header>
  <div class="register" v-loading.fullscreen.lock="this.$store.state.loading">
    <el-card shadow="always">
      <el-tabs v-model="activeName" @tab-click="handleClick" :stretch="true">
        <el-tab-pane label="注册" name="first"></el-tab-pane>
        <el-tab-pane label="登录" name="second"></el-tab-pane>
      </el-tabs>
      <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
        <el-form-item label="" prop="username">
          <el-input v-model="ruleForm.username" 
            placeholder="请输入用户名" spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="" prop="email">
          <el-input v-model="ruleForm.email" 
            placeholder="请输入邮箱" spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="" prop="pass">
          <el-input type='password' v-model="ruleForm.pass" 
            placeholder="请输入8到16位密码" spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item label="" prop="passAgain">
          <el-input type='password' v-model="ruleForm.passAgain" 
            placeholder="请再次输入密码" spellcheck="false"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">注册</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</div>
</template>

<style>
.register .el-tabs {
  padding-top:10px;
  margin-bottom: 50px;
  font-family: 'Hiragino Sans GB'
}

.register .el-tabs .el-tabs__item {
  font-size: 26px!important;
}

.register .el-card {
  margin: 120px auto;
  max-width: 320px;
  border-width: 1.5px;
}

.el-form .el-form-item .el-button {
  width: 100%!important;
}
</style>

<script>
import { UserRegister } from '../api/api'
import header from '../components/header.vue'

export default {
  name: "register",
  components: {
    'blog-header': header
  },
  data() {
    var checkPass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        username: '',
        email: '',
        pass: '',
        passAgain: '',
        type: '0'
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { max: 30, message: '用户名过长', trigger: 'blur'}
        ],
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱', trigger: ['blur', 'change'] },
          { max: 30, message: '邮箱过长', trigger: 'blur'}
        ],
        pass: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, max: 16, message: '请输入8到16位密码', trigger: ['blur', 'change']}
        ],
        passAgain: [
          { required: true, message: '请再次输入密码', trigger: 'blur' },
          { validator: checkPass, trigger: ['blur', 'change'] }
        ]
      },
      activeName: "first"
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) { 
          this.$store.commit("changeLoading", true)
          UserRegister(this.ruleForm)
          .then(res => {
            this.$message({
              message: '注册成功',
              type: 'success',
              showClose: true,
              duration: 1500
            })
            this.$router.push({
              path: '/Login'
            })
            this.$store.commit("changeLoading", false)
          })
          .catch(res => {
            var err
            try {
              err = res.response.data.errorCode
            } catch(e) {
              this.$message.error({
                message: '未知错误',
                type: 'error',
                showClose: true,
                duration: 2000
              })
              this.$store.commit("changeLoading", false)
              return
            }
            if (err === "0") {
              this.$message.error({
                message: '用户名已被注册',
                type: 'error',
                showClose: true,
                duration: 2000
              })
            } else if (err === "1") {
              this.$message.error({
                message: '邮箱已被注册',
                type: 'error',
                showClose: true,
                duration: 2000
              })
            } else {
              this.$message.error({
                message: '未知错误',
                type: 'error',
                showClose: true,
                duration: 2000
              })
            }
            this.$store.commit("changeLoading", false)
          })
        }
      });
    },
    handleClick(tab, event) {
      if (tab.name === 'second') {
        this.$router.push({
          path: '/Login'
        });
      }
    }
  }
};
</script>

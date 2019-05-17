<template>
<div class="login" v-loading.fullscreen.lock="this.$store.state.loading">
  <el-card shadow="always">
    <el-tabs v-model="activeName" @tab-click="handleClick" :stretch="true">
      <el-tab-pane label="注册" name="first"></el-tab-pane>
      <el-tab-pane label="登录" name="second"></el-tab-pane>
    </el-tabs>
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm">
      <el-form-item label="" prop="usernameORemail">
        <el-input v-model="ruleForm.usernameORemail" 
          placeholder="请输入用户名或邮箱" spellcheck="false"></el-input>
      </el-form-item>
      <el-form-item label="" prop="pass">
        <el-input type='password' v-model="ruleForm.pass" 
          placeholder="请输入密码" spellcheck="false"></el-input>
      </el-form-item>
      <el-form-item>
        <el-checkbox v-model="rememberMe">记住我</el-checkbox>
        <a href="javascript:void(0)" class="forget">忘记密码?</a>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">登录</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</div>
</template>

<style>
.login .el-tabs {
  padding-top:20px;
  margin-bottom: 50px;
  font-family: 'Hiragino Sans GB'
}

.login .el-tabs .el-tabs__item {
  font-size: 26px!important;
}

.login .el-card {
  margin: 60px auto;
  max-width: 320px;
  border-width: 1.5px;
}

.el-form .el-form-item .el-button {
  width: 100%!important;
}

.el-form .el-form-item .el-checkbox {
  float: left!important;
  margin-left: 10px!important;
}

.el-form .el-form-item .forget {
  float: right!important;
  margin-right: 10px!important;
}
</style>

<script>
import {loginServer} from '../api/login'
export default {
  data() {
    return {
      ruleForm: {
        usernameORemail: '',
        pass: '',
        type: '1'
      },
      rules: {
        usernameORemail: [
          { required: true, message: '请输入用户名或邮箱', trigger: 'blur' },
          { max: 30, message: '输入过长', trigger: 'blur'}
        ],
        pass: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { max: 16, message: '输入过长', trigger: 'blur'}
        ]
      },
      activeName: "second",
      rememberMe: false
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          var that = this;
          loginServer(JSON.stringify(this.ruleForm), function(res){
            var code = res.data;
            switch(code) {
              case 1000:
                that.$message({
                  message: '登录成功',
                  type: 'success',
                  showClose: true,
                  duration: 1000
                });
                that.$router.push({
                  path: '/'
                });
                break;
              default:
                that.$message.error({
                message: '用户名或密码错误',
                type: 'error',
                showClose: true,
                duration: 2000
              });
              that.$refs[formName].resetFields();
            }
          });
        }
      });
    },
    handleClick(tab, event) {
      if (tab.name === 'first') {
        this.$router.push({
          path: '/Register'
        });
      }
    }
  }
};
</script>

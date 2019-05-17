<template>
<div class="header">
  <el-button @click="login()" type="primary"
    round>登录</el-button>
  <el-button @click="register()" type="primary"
    round>注册</el-button>
</div>
</template>

<style>

</style>

<script>
export default {
  data() {
    return {
      activeIndex: '1'
    };
  },
  methods: {
    login() {
      this.$router.push ({
        path: '/Login'
      });
    },
    register() {
      this.$router.push ({
        path: '/Register'
      });
    }
  }
}
</script>

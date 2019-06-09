<template>
<div>
  <div v-if="this.$store.state.hasLogin">
    <el-dialog
      title="Forbidden"
      :visible="true"
      width="30%"
      :show-close="false">
      <span class="warningMsg">您没有权限访问</span>
      <span slot="footer">
        <el-button @click="goHome">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</div>
</template>

<script>
export default {
  name: "Forbidden",
  
  data() {
    return {
      
    }
  },

  methods: {
    goHome() {
      this.$router.push({
        path: "/"
      })
    }
  }
}
</script>

<style>
.warningMsg {
  font-size: 16px;
  color: red;
}
</style>

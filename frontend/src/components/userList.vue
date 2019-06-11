<template>
  <div>
    <div class="articleItem" v-for="user in users" :key="user.id">
      <user-card :userInfo="user"></user-card>
    </div>
    <el-pagination
      @current-change="handleCurrentChange"
      :current-page.sync="currentPage"
      :page-size="9"
      layout="total, prev, pager, next"
      :total="totalNum"
    ></el-pagination>
  </div>
</template>

<script>
import userCard from "../components/userCard.vue"

import { GetUsers } from "../api/api";

export default {
  components: {
    "user-card": userCard
  },

  data() {
    return {
      totalUsers: [],
      users: [],
      currentPage: 1,
      totalNum: 0
    };
  },

  methods: {
    handleCurrentChange() {
      this.users = this.totalUsers.slice(
        (this.currentPage - 1) * 9,
        this.currentPage * 9
      );
      this.router.push({
        path: "/Admin/page/" + this.currentPage
      });
    },

    changePage() {
      if (this.$route.params.page) {
        this.currentPage = Number(this.$route.params.page)
        if (this.currentPage == 1) {
          this.$router.push({
            path: "/Admin"
          })
        }
      } else {
        this.currentPage = 1
      }
      
      if (this.currentPage > (this.totalNum-1)/6 + 1) {
        this.currentPage = 1
        this.$router.push({
          path: "/Admin"
        })
      }
    },

    getData() {
      GetUsers()
      .then(res => {
        if (res.data) {
          this.totalUsers = res.data.users;
          this.totalNum = this.users.length;
          this.users = this.users.slice(
            (this.currentPage - 1) * 9,
            this.currentPage * 9
          );
        }
      })
      .catch(() => {
        this.$message({
          message: "未知错误",
          type: "error",
          duration: 1500
        });
      });
    }
  },

  watch: {
    $route: "changePage"
  },

  created() {
    this.getData();
  }
};
</script>

<style>
</style>

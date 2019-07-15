<template>
  <div class="userListBox">
    <el-table
      :data="users"
      style="width: 100%"
      stripe>
      <el-table-column label="用户名">
        <template slot-scope="scope">
          <i class="el-icon-user-solid"></i>
          <span style="margin-left: 10px">{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Email">
        <template slot-scope="scope">
          <i class="el-icon-message"></i>
          <span style="margin-left: 10px">{{ scope.row.email }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      @current-change="handleCurrentChange"
      :current-page.sync="currentPage"
      :page-size="9"
      layout="total, prev, pager, next"
      :total="totalNum"
    ></el-pagination>

    <el-dialog
      title="删除"
      :visible.sync="dialogVisible"
      width="30%"
      :before-close="cancelDelete">
      <span>是否删除用户 {{deleteUserInfo.username}}</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancelDelete">取 消</el-button>
        <el-button type="primary" @click="deleteUser">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { GetUsers, DeleteUser } from "../api/api";

export default {
  data() {
    return {
      totalUsers: [],
      users: [],
      currentPage: 1,
      totalNum: 0,
      deleteUserInfo: "",
      dialogVisible: false
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
          this.totalUsers = res.data;
          this.totalNum = this.totalUsers.length;
          this.users = this.totalUsers.slice(
            (this.currentPage - 1) * 9,
            this.currentPage * 9
          );

        }
      })
      .catch(() => {
        this.$message.error({
          message: "未知错误",
          type: "error",
          duration: 1500
        });
      });
    },

    cancelDelete() {
      this.dialogVisible = false
      this.deleteUserInfo = ""
    },

    handleDelete(userInfo) {
      this.deleteUserInfo = userInfo
      this.dialogVisible = true
    },

    deleteUser() {
      this.dialogVisible = false
      DeleteUser(this.deleteUserInfo._id)
      .then(res => {
        this.$message({
          message: "删除成功",
          type: "success",
          duration: 1500
        })
        this.getData()
      })
      .catch(() => {
        this.$message({
          message: "删除失败",
          type: "error",
          duration: 1500
        })
      })
      this.deleteUserInfo = ""
    }
  },

  watch: {
    $route: "changePage"
  },

  mounted () {
    this.getData();
  }
};
</script>

<style>
.userListBox .el-table {
  padding: 5px 20px;
}
</style>

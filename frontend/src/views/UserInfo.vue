<template>
  <div>
    <blog-header></blog-header>

    <div class="container">
      <div v-show="isEdit" class="userInfoContainer">
        <h1>编辑个人资料</h1>

        <ul class="userInfoBox">
          <li class="avatarlist">
            <span class="leftTitle">头像</span>
            <el-upload
              class="avatar-uploader"
              action=""
              accept="image/jpeg,image/gif,image/png"
              :show-file-list="false"
              :before-upload="beforeAvatarUpload"
              :http-request="uploadAvatar"
            >
              <img
                v-if="userInfo.avatar"
                :src="userInfo.avatar"
                class="avatar"
              >
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
              <div slot="tip" class="el-upload__tip">点击上传头像(jpg/png格式,不超过1M)</div>
            </el-upload>
          </li>
          <li class="username">
            <span class="leftTitle">用户名</span>
            <el-input v-model="userInfo.username" placeholder="用户名"></el-input>
            <i class="fa fa-wa fa-asterisk"></i>
          </li>
          <li>
            <span class="leftTitle">电子邮件</span>
            <span>{{ userInfo.email }}</span>
          </li>
        </ul>

        <div class="saveInfobtn">
          <a class="cancel" href="javascript:void(0);" @click="isEdit=!isEdit">返 回</a>
          <a class="submit" href="javascript:void(0);" @click="saveInfoFun">保 存</a>
        </div>
      </div>

      <div v-show="!isEdit" class="userInfoContainer">
        <h1>
          个人中心
          <span class="gotoEdit" @click="isEdit=!isEdit">
            <i class="fa fa-wa fa-edit"></i>编辑
          </span>
        </h1>

        <ul class="userInfoBox">
          <li class="avatarlist">
            <span class="leftTitle">头像</span>
            <div class="avatar-uploader">
              <img :src="userInfo.avatar" class="avatar">
            </div>
          </li>
          <li class="username">
            <span class="leftTitle">昵称</span>
            <span>{{userInfo.username}}</span>
          </li>
          <li>
            <span class="leftTitle">电子邮件</span>
            <span>{{userInfo.email}}</span>
          </li>
        </ul>
      </div>
    </div>

    <blog-footer></blog-footer>
  </div>
</template>

<script>
import header from "../components/header.vue"
import footer from "../components/footer.vue"

import { GetUserInfo, UploadAvatar, UpdateUser } from "../api/api"

export default {
  name: "UserInfo",
  
  components: {
    "blog-header": header,
    "blog-footer": footer
  },

  data() {
    return {
      isEdit: false,
      userInfo: "",
      oldUsername: ""
    };
  },
  methods: {
    getData() {
      if (!this.$store.state.hasLogin) {
        this.$message.error("您未登录")
        this.$router.push({
          path: "/Forbidden"
        })
        return
      }

      GetUserInfo(this.$store.state.userInfo)
      .then(res => {
        this.userInfo = res.data
        this.oldUsername = res.data.username
      })
      .catch(() => {
        this.$message.error("未知错误")
      })
    },

    beforeAvatarUpload(file) {
      const isJPG =
        file.type == "image/png" ||
        file.type == "image/jpg" ||
        file.type == "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 1;
      
      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG/JPEG/PNG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 1MB!");
      }
      return isJPG && isLt2M;
    },

    uploadAvatar(param) {
      var fileObj = param.file
      var form = new FormData()
      form.append("file", fileObj)

      UploadAvatar(form)
      .then(res => {
        this.userInfo.avatar = "http://localhost/img/avatar/admin.jpg"
        this.userInfo.avatar = res.data
        var user = JSON.parse(localStorage.getItem("currentUser"))
        user.avatar = res.data
        localStorage.setItem("currentUser", JSON.stringify(user))
        this.$store.commit("changeAvatar", "http://localhost/img/avatar/admin.jpg")
        this.$store.commit("changeAvatar", res.data)
        this.$message({
          message: "上传成功",
          type: "success",
          duration: 1500
        })
      })
      .catch(() => {
        this.$message({
          message: "上传失败",
          type: "error",
          duration: 1500
        })
      })
    },

    saveInfoFun: function() {
      if (this.userInfo.username != this.oldUsername) {
        UpdateUser(this.userInfo)
        .then(res => {
          this.oldUsername = this.userInfo.username
          this.$message.success("保存成功")
          this.isEdit = false
        })
        .catch(res => {
          var err
          try {
            err = res.response.data.error
          } catch(e) {
            this.$message.error({
              message: '未知错误',
              type: 'error',
              showClose: true,
              duration: 2000
            })
            return
          }
          if (err === "1") {
            this.$message.error({
              message: '用户名已存在',
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
        })
      } else {
        this.isEdit = false
      }
    }
  },
  
  created() {
    this.getData();
  }
};
</script>

<style>
.userInfoContainer {
  position: relative;
  background: #fff;
  color: black;
  padding: 15px;
  border-radius: 5px;
  margin-bottom: 40px;
  font-size: 16px;
}

.userInfoContainer h1 {
  margin: 10px 0;
  font-size: 25px;
  font-weight: 700;
  text-align: center;
  line-height: 30px;
}

ul.userInfoBox {
  padding: 0;
  text-align: left;
}

.userInfoBox .avatarlist {
  position: relative;
}

.avatar-uploader {
  display: inline-block;
  vertical-align: top;
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 5%;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  width: 120px;
  height: 120px;
}

.avatar-uploader .el-upload:hover {
  border-color: #20a0ff;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 120px;
  height: 120px;
  line-height: 120px;
  text-align: center;
  position: absolute;
  top:0;
  left:0;
}

.avatar {
  width: 120px;
  height: 120px;
  display: block;
  border-radius: 5%;
  object-fit: cover;
}

.gotoEdit {
  font-size: 15px;
  float: right;
  cursor: pointer;
  color: #999;
}

.gotoEdit:hover {
  color: #333;
}

.userInfoBox .leftTitle {
  display: inline-block;
  width: 100px;
  padding: 10px 0;
}

.userInfoBox .rightInner {
  display: inline-block;
  max-width: calc(100% - 140px);
  vertical-align: top;
}

.userInfoBox li {
  padding: 20px;
  border-bottom: 1px solid #ddd;
}

.userInfoBox li:last-child {
  border-bottom: 1px solid transparent;
}

.userInfoBox .el-input,
.userInfoBox .el-textarea {
  max-width: 300px;
  min-width: 100px;
}

.userInfoBox .el-input__inner {
  border-radius: 4px;
}

.userInfoBox .el-textarea {
  vertical-align: top;
}

.saveInfobtn {
  margin: 20px 0;
  text-align: center;
}

.saveInfobtn a {
  color: #fff;
  padding: 6px 20px;
  margin: 5px 10px;
  border-radius: 5px;
  font-size: 14px;
}

.userInfoBox .fa-asterisk {
  color: #df2050;
}

.saveInfobtn .submit {
  background-color: #97dffd;
}

.saveInfobtn .cancel {
  background-color: #c0c4cc;
}

.saveInfobtn .submit:hover {
  color: #409eff
}

.saveInfobtn .cancel:hover {
  color: #409eff
}
</style>

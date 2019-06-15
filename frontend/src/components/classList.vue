<template>
  <div class="classListBox">
    <el-button class="newFirstClassBtn" @click="newFirstClass">
      新建一级分类</el-button>
    <el-button class="newFirstClassBtn" @click="newSecondClass">
      新建二级分类</el-button>

    <el-tree
      ref="indicatorTree"
      :props="indicatorLabels"
      :data="classTree"
      :highlight-current="false"
      node-key="_id"
      class="classTreeBox">
      <span class="custom-tree-node" slot-scope="{ node, data }">
        <span>{{ data.name }}</span>
        <span>
          <el-button
            type="text"
            @click="() => deleteClassHandle(data)">
            删除
          </el-button>
        </span>
      </span>
    </el-tree>

    <el-dialog
      :title = "'新建'+newClassLevel+'级分类'"
      :visible.sync="newDialogVisible"
      width="30%"
      class="newClassDialog">
      <el-select 
        v-if="newClassLevel == '二'"
        v-model="newClassInfo.father" 
        placeholder="请选择一级分类">
        <el-option
          v-for="item in classTree"
          :key="item._id"
          :label="item.name"
          :value="item._id">
        </el-option>
      </el-select>
      <el-input placeholder="输入分类名" v-model="newClassInfo.name"></el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="newDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="newClass">确 定</el-button>
      </span>
    </el-dialog>
    
    <el-dialog
      title="删除"
      :visible.sync="deleteDialogVisible"
      width="30%"
      :before-close="cancelDelete">
      <span>是否删除分类 {{deleteClassInfo.name}} (及其子分类)</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancelDelete">取 消</el-button>
        <el-button type="primary" @click="deleteClass">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { GetClasses, DeleteClass, NewFirstClass, NewSecondClass } from "../api/api";

export default {
  data() {
    return {
      classTree: [],
      deleteClassInfo: {},
      newClassInfo: {
        father: "",
        name: ""
      },
      newClassLevel: "",
      deleteDialogVisible: false,
      newDialogVisible: false,
      indicatorLabels: {
        label: "name",
        children: "child"
      }
    };
  },

  methods: {
    getData() {
      GetClasses()
      .then(res => {
        if (res.data) {
          this.classTree = res.data
        }
      })
      .catch(() => {
        this.$message({
          message: "错误",
          type: "error"
        })
      })
    },

    newFirstClass() {
      this.newClassLevel = "一"
      this.newClassInfo.father = ""
      this.newClassInfo.name = ""
      this.newDialogVisible = true
    },

    newSecondClass() {
      this.newClassLevel = "二"
      this.newClassInfo.father = ""
      this.newClassInfo.name = ""
      this.newDialogVisible = true
    },

    newClass() {
      if (this.newClassLevel == "二") {
        if (!this.newClassInfo.father) {
          this.$message.error("请选择一级分类")
          return
        }
        if (!this.newClassInfo.name) {
          this.$message.error("请输入分类名")
          return
        }

        NewSecondClass(this.newClassInfo.name, this.newClassInfo.father)
        .then(res => {
          this.$message.success("创建成功")
          this.getData()
        })
        .catch(res => {
          var err
          if (res.response.data && res.response.data.error) {
            err = res.response.data.error
          }
          if (err === "1") {
            this.$message.error('分类已存在')
          } else {
            this.$message.error('未知错误')
          }
        })
      } else if (this.newClassLevel == "一") {
        if (!this.newClassInfo.name) {
          this.$message.error("请输入分类名")
          return
        }

        NewFirstClass(this.newClassInfo.name)
        .then(res => {
          this.$message.success("创建成功")
          this.$store.commit("getClassList")
          this.getData()
        })
        .catch(res => {
          var err
          if (res.response.data && res.response.data.error) {
            err = res.response.data.error
          }
          if (err === "1") {
            this.$message.error('分类已存在')
          } else {
            this.$message.error('未知错误')
          }
        })
      }

      this.newDialogVisible = false
    },

    cancelDelete() {
      this.deleteDialogVisible = false
      this.deleteClass = ""
    },

    deleteClassHandle(data) {
      this.deleteClassInfo = data
      this.deleteDialogVisible = true
    },

    deleteClass() {
      this.deleteDialogVisible = false
      DeleteClass(this.deleteClassInfo._id)
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
      this.deleteClassInfo = ""
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
.classListBox {
  background-color: white;
  border-radius: 5px;
  padding: 10px;
}

.classListBox .newFirstClassBtn {
  width: 80%;
  height: 40px;
  font-size: 16px;
  font-weight: 600;
  margin: 10px auto;
}

.classListBox .classTreeBox {
  width: 80%;
  font-size: 16px;
  margin: 10px auto;
}

.classListBox .custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8px;
}

.classListBox .newClassDialog .el-select {
  width: 100%;
  margin-bottom: 20px;
}

.classListBox .custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-right: 8px;
}
</style>

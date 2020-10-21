<template>
  <el-table
      :data="tableData"
      style="width: 100%"
      :row-class-name="tableRowClassName">
    <el-table-column
        prop="Id"
        label="Id">
    </el-table-column>
    <el-table-column
        prop="title"
        label="书名">
    </el-table-column>
    <el-table-column
        prop="author"
        label="作者">
    </el-table-column>
    <el-table-column
        prop="translator"
        label="译者">
    </el-table-column>
    <el-table-column
        prop="press"
        label="出版社">
    </el-table-column>
    <el-table-column
        prop="date"
        label="出版日期">
    </el-table-column>
    <el-table-column
        prop="price"
        label="价格">
    </el-table-column>
    <el-table-column
        prop="star"
        label="星级">
    </el-table-column>
    <el-table-column
        prop="comment"
        label="评价人数">
    </el-table-column>
    <el-table-column
        prop="quote"
        label="名句">
    </el-table-column>
  </el-table>
</template>

<script>
import axios from 'axios'
export default {
  name: "ApiData",
  data() {
    return {
      book: null,
    }
  },
  mounted() {
    axios.get('http://localhost:8000/douban').then(response => (this.book = response.data))
  }
}
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>
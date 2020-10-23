<template>
  <div class="book-body">
    <el-table :data="booksData" border style="width: 100%">
      <el-table-column prop="Id" label="Id"></el-table-column>
      <el-table-column prop="title" label="书名"></el-table-column>
      <el-table-column prop="author" label="作者"></el-table-column>
      <el-table-column prop="translator" label="译者"></el-table-column>
      <el-table-column prop="press" label="出版社"></el-table-column>
      <el-table-column prop="date" label="出版日期"></el-table-column>
      <el-table-column prop="price" label="价格"></el-table-column>
      <el-table-column prop="star" label="星级"></el-table-column>
      <el-table-column prop="comment" label="评价人数"></el-table-column>
      <el-table-column prop="quote" label="名句"></el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      booksData: [],
    }
  },
  created() {
    fetch('http://localhost:8000/douban')
        .then(response => response.json())
        .then(json => {
          this.booksData = json
        })
  },
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
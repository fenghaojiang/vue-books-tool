<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Vue-Books-Tool"/>
  <BooksList>
    <el-table :data="this.booksData">
      <el-table-column prop="Id" label="Id"></el-table-column>
      <el-table-column prop="Title" label="书名"></el-table-column>
      <el-table-column prop="Author" label="作者"></el-table-column>
      <el-table-column prop="Translator" label="译者"></el-table-column>
      <el-table-column prop="Press" label="出版社"></el-table-column>
      <el-table-column prop="Date" label="日期"></el-table-column>
      <el-table-column prop="Price" label="价格"></el-table-column>
      <el-table-column prop="Star" label="星级"></el-table-column>
      <el-table-column prop="Comment" label="评价人数"></el-table-column>
      <el-table-column prop="Quote" label="引言"></el-table-column>
    </el-table>
  </BooksList>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import BooksList from "@/components/BooksList";


export default {
  name: 'App',
  data() {
    return {
      booksData:BooksList.data().booksData
    }
  },
  components: {
    HelloWorld,
    BooksList,
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Vue-Books-Tool"/>
  <el-table
      :data="booksData"
      style="width: 100%">
    <el-table-column
        prop="Id"
        label="Id"
        width="180">
    </el-table-column>
    <el-table-column
        prop="Title"
        label="书名"
        width="180">
    </el-table-column>
    <el-table-column
        prop="Author"
        label="作者">
    </el-table-column>
    <el-table-column
        prop="Translator"
        label="译者">
    </el-table-column>
    <el-table-column
        prop="Press"
        label="出版社">
    </el-table-column>
    <el-table-column
        prop="Date"
        label="日期">
    </el-table-column>
    <el-table-column
        prop="Price"
        label="价格">
    </el-table-column>
    <el-table-column
        prop="Star"
        label="星级">
    </el-table-column>
    <el-table-column
        prop="Comment"
        label="评价人数">
    </el-table-column>
    <el-table-column
        prop="Quote"
        label="引言">
    </el-table-column>
  </el-table>
  <BooksList :data="booksData">
    <table>
      <thead>
      <tr>
        <td>Id</td>
        <td>Title</td>
        <td>Author</td>
        <td>Translator</td>
        <td>Press</td>
        <td>Date</td>
        <td>Price</td>
        <td>Star</td>
        <td>Comment</td>
        <td>Quote</td>
      </tr>
      </thead>
      <tbody id="book-body">
      <tr v-for="books in booksData" :key="books">
        <td>{{books.Id}}</td>
        <td>{{books.Title}}</td>
        <td>{{books.Author}}</td>
        <td>{{books.Translator}}</td>
        <td>{{books.Press}}</td>
        <td>{{books.Date}}</td>
        <td>{{books.Star}}</td>
        <td>{{books.Comment}}</td>
        <td>{{books.Quote}}</td>
      </tr>
      </tbody>
    </table>
  </BooksList>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import BooksList from "@/components/BooksList";


export default {
  name: 'App',
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

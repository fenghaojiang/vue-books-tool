<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Vue-Books-Tool"/>
  <a-table :datasrc="booksData" :loading="loading">

  </a-table>
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
      <tr v-for="books in data" :key="books">
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
import 'ant-design-vue/dist/antd.css';

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

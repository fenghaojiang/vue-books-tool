<template>
  <div class="books-data" :data="booksData">
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
        <td>{{books.Price}}</td>
        <td>{{books.Star}}</td>
        <td>{{books.Comment}}</td>
        <td>{{books.Quote}}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "BooksList",
  data() {
    return {
      booksData: []
    }
  },
  mounted() {
    axios.get('http://localhost:8000/douban').then(response => this.booksData = response.data)
  }
}
</script>

<style scoped>

</style>
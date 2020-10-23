<template>
  <div class="books-body">
    {{booksData}}
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "BooksList",
  booksData: [],
  data() {
    return this.booksData
  },
  mounted() {
    axios.get('http://localhost:8000/douban').then(response => this.booksData = response)
  }
}
</script>

<style scoped>

</style>
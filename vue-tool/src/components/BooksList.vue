<template>
  <div class="books-body" :data="booksData">
    {{booksData}}
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "BooksList",
  booksData: null,
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
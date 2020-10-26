<template>
  <div class="books-data" :data="booksData">
    {{booksData}}
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
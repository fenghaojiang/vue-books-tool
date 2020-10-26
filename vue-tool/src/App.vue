<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Vue-Books-Tool"/>
  <BooksList>
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
  data() {
    return {
      booksData:BooksList.data().booksData
    }
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

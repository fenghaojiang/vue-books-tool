<template>
  <img alt="Vue logo" src="./assets/logo.png" width="50" height="50">
  <HelloWorld msg="Welcome to Vue-Books-Tool"/>
  <BooksList></BooksList>
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

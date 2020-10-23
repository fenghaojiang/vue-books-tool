<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <BooksList>
    <v-table :width="1000" :columns="columns">
    </v-table>
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
  }
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

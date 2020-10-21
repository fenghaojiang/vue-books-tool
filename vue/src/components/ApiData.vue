<template>

</template>

<script>
import HelloWorld from "@/components/HelloWorld";
import Router from 'vue-router'
import Book from '@/components/ApiData'

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/douban',
      name: 'Book',
      component: Book
    },
  ]

})
</script>

<style scoped>

</style>
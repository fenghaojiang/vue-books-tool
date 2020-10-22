<template>
   <ApiData>
   </ApiData>
</template>

<script>
import Router from 'vue-router';
import ApiData from "@/ApiData";

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Book',
      component: ApiData
    },
  ]

})
</script>

<style scoped>

</style>
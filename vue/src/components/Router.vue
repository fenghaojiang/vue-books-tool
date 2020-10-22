<template>
 <el-table>
   <ApiData>
     {{}}
   </ApiData>
 </el-table>
</template>

<script>
import HelloWorld from "@/components/HelloWorld";
import Router from 'vue-router';
import ApiData from "@/ApiData";

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
      component: ApiData
    },
  ]

})
</script>

<style scoped>

</style>
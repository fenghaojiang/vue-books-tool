import Vue from 'vue'

import ApiData from "@/ApiData";

Vue.config.productionTip = false

new Vue({
  render: h => h(ApiData),
}).$mount('#apiData')

<script>
import axios from 'axios'
export default {
  name: "ApiData",
  data() {
    return {
      book: null,
    }
  },
  mounted() {
    axios.get('http://localhost:8000/douban').then(response => (this.book = response.data))
  }
}
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>
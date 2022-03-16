<template>
  <div class="table" id="app">
    <h1><b>Menu List</b></h1>
    <center>
      <table>
        <tr>
          <th>ID</th>
          <th>Nama Menu</th>
          <th>Harga</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.name }}</td>
          <td>{{ log.price }}</td>
        </tr>
      </table>
    </center>
  </div>
</template>

<script>
import axios from "axios";
export default {
  mounted() {
    this.fetchData();
  },
  data() {
      return {
        data: [],
      }
  },
  methods: {
    async fetchData() {
      const res = await axios.get("/ChefManager/allmenu");
      this.data = res.data.data[0];
    },
  },
};
</script>

<style>
#app {
  text-align: center;
  margin-top: 30px;
}
</style>

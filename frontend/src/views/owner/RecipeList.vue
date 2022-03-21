<template>
  <div class="table" id="app">
    <h1><b>Recipe List</b></h1>
    <center>
      <table>
        <tr>
          <th>ID</th>
          <th>Description</th>
          <th>Material</th>
          <th>Quantity</th>
          <th>Unit</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.description }}</td>
          <td>{{ log.material }}</td>
          <td>{{ log.quantity }}</td>
          <td>{{ log.unit }}</td>
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
      const res = await axios.get("/OwnerManager/allrecipe");
      this.data = res.data.data[0];
    },
  },
};
</script>

<style>

</style>

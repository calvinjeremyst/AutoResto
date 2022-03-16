<template>
  <div class="table" id="app">
    <h1><b>Material List</b></h1>
    <center>
      <table>
        <tr>
          <th>ID</th>
          <th>Nama Bahan Baku</th>
          <th>Jumlah</th>
          <th>Satuan</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.name }}</td>
          <td>{{ log.quantity }}</td>
          <td>{{ log.alamat }}</td>
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
      const res = await axios.get("/InventoryManager/allmaterial");
      this.data = res.data.data[0];
    },
  },
};
</script>

<style>

</style>
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
          <th>Data Option</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.description }}</td>
          <td>{{ log.material }}</td>
          <td>{{ log.quantity }}</td>
          <td>{{ log.unit }}</td>
        </tr>
        <tr>
          <td>1</td>
          <td>Nasi Goreng Spesial dengan Daging</td>
          <td>Beras</td>
          <td>200</td>
          <td>Gr</td>
          <td><button name="edit" class="btnUpdate">Edit</button><button name="delete" class="btnDelete">Delete</button></td>
        </tr>
        <tr>
          <td></td>
          <td></td>
          <td>Daging</td>
          <td>100</td>
          <td>Gram</td>
        </tr>
        <tr>
          <td></td>
          <td></td>
          <td>Cabai</td>
          <td>2</td>
          <td>Pcs</td>
        </tr>
        <tr>
          <td></td>
          <td></td>
          <td>Telur</td>
          <td>2</td>
          <td>Butir</td>
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

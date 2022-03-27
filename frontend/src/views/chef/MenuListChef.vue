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
data() {
    return {
      data: [],
    }
},
mounted() {
  this.fetchData();
},
methods: {
  async fetchData() {
    try{
      const res = await axios.get("/ChefManager/allmenu");
      this.data = res.data.data;
   }
    catch(error){
      console.log(error)
    }

  }
},
};
</script>

<style>

</style>

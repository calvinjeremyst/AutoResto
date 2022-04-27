<template>
  <div class="table" id="app">
    <h1><b>Menu List</b></h1>
    <center>
      <table>
        <tr>
          <th>No </th>
          <th>Nama Menu</th>
          <th>Description</th>
        </tr>
        <tr v-for="(log,index) in data" :key="log.id">
          <td>{{ index + 1}}</td>
          <td>{{ log.Menu.name }}</td>
          <td>{{ log.description }}</td>
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
        const res = await axios.get("/ChefManager/allmenuchef");
        this.data = res.data.data;
      }catch(error){
        console.log(error)
      }
    }
  },
};
</script>

<template>
  <div class="table" id="app">
    <h1><b>Recipe List</b></h1>
    <center>
      <table>
        <tr>
          <th>ID</th>
          <th>Description</th>
          <th>ID Material</th>
          <th>Material Name</th>
          <th>Quantity</th>
          <th>Unit</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.idRecipe.description }}</td>
          <td>{{log.idMaterial.Id}}</td>
          <td>{{ log.idMaterial.Name }}</td>
          <td>{{ log.quality }}</td>
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
    try{
      const res = await axios.get("/ChefManager/alldetailrecipe");
      this.data = res.data.data;
    }
    catch(error){
      console.log(error)
    }
  },
},
};
</script>

<style>

</style>

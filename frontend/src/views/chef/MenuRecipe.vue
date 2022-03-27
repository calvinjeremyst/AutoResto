<template>
  <div class="table" id="app">
    <center>
     <h1><b>Recipe Menu List</b></h1>
      <table>
        <tr>
          <th>ID Recipe</th>
          <th>Description</th>
          <th>ID Food</th>
          <th>Food Name </th>
          <th>Food Price</th>
        </tr>
        <tr v-for="log in data" :key="log.id">
          <td>{{ log.id }}</td>
          <td>{{ log.description }}</td>
          <td>{{log.Menu.id}}</td>
          <td>{{log.Menu.name}}</td>
          <td>{{log.Menu.price}}</td>
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
      const res = await axios.get(`ChefManager/1`);
      this.data = res.data.data;
   }
    catch(error){
      console.log(error)
    }
  }
},
};
</script>

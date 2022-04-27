<template>
  <div class="table" id="app">
    <h1><b>Recipe List</b></h1>
    <center>
      <table>
        <tr>
          <th> No</th>
          <th>Description</th>
          <th>ID Material</th>
          <th>Material Name</th>
          <th>Quantity</th>
          <th>Unit</th>
        </tr>
        <tr v-for="(log, index) in data" :key="log.id">
          <td>{{ index + 1}}</td>
          <td>{{ log.recipe.description }}</td>
          <td>{{log.material.Id}}</td>
          <td>{{ log.material.Name }}</td>
          <td>{{ log.quantity}}</td>
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

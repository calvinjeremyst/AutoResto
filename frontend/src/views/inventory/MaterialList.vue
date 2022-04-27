<template>
  <div class="table" id="app">
    <h1><b>Material List</b></h1>
    <center>
      <table>
        <tr>
          <th>No</th>
          <th>Nama Bahan Baku</th>
          <th>Jumlah</th>
          <th>Satuan</th>
          <th>Data Option</th>
        </tr>
        <tr v-for="(log,index) in data" :key="log.Id">
          <td>{{ index + 1}}</td>
          <td>{{ log.Name }}</td>
          <td>{{ log.Quantity }}</td>
          <td>{{ log.Unit }}</td>
          <td>
            <router-link :to="{name:'UpdateMaterial',params:{'id':log.Id}}">
              <button name="edit" class="btnUpdate">Edit</button>
            </router-link>
             <button name="delete" class="btnDelete" @click="deleteMaterial(log.Id)">Delete</button>
          </td>
        </tr>
      </table>
    </center>
  </div>
</template>

<script>
import axios from "axios";
export default{
  mounted(){
    this.fetchData();
  },
  data(){
    return{
      data : [],
    }
  },
  methods : {
      async fetchData(){
      try{
        const res = await axios.get("InventoryManager/allmaterial");
        if(res.data.data == null){
            alert("Inventory Kosong")
        }else{
            alert("Showing material list")
            this.data = res.data.data;
            console.log(res,this.data)
        }
       
       
      } 
      catch(error){
          console.log(error)
      }
    },
    async deleteMaterial(Id){
      try{
          await axios.delete(`InventoryManager/${Id}`);
          alert("Delete material success")
          this.fetchData()
      }
      catch(error){
        console.log(error)
     }
     
    }
 }
};
  
 
</script>

<style>
  .table{
    text-align: center;
  }

  .btnUpdate{
    margin-right: 10px;
  }
</style>

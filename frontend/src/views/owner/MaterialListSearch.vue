<template>
    <div class="table" id="app">
        <h1><b>Material List Searched</b></h1>
        <center>
        <table>
            <tr>
                <th>Number</th>
                <th>Nama Bahan Baku</th>
                <th>Jumlah</th>
                <th>Satuan</th>
            </tr>
            <tr v-for="log in data" :key="log.Name">
                <td>{{ log.Id }}</td>
                <td>{{ log.Name }}</td>
                <td>{{ log.Quantity }}</td>
                <td>{{ log.Unit }}</td>
            </tr>
        </table>
        </center>
    </div>
</template>

<script>
import axios from 'axios'
export default{
    data(){
        return{
            data : []
        }
    },
    mounted(){
        this.fetchByName()
    },
    methods:{
        async fetchByName(){
            try{
                const res = await axios.get('OwnerManager/material/' + this.$route.params.name)
                if (res.data.data == null){
                    alert("Material tidak ditemukan")
                }else{
                    alert("Material get sukses")
                      this.data = res.data.data
                }
               

                console.log(res,this.data)
            }
            catch(error){
                console.log(error)
            }
           
        }
    }
}
</script>

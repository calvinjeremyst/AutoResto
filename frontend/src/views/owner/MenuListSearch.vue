<template>
    <div class="table" id="app">
        <h1><b>Menu List Searched</b></h1>
        <center>
        <table>
            <tr>
                <th>Menu Name</th>
                <th>Price</th>
            </tr>
            <tr v-for="log in data" :key="log.name">
                <td>{{ log.name }}</td>
                <td>{{ log.price}}</td>
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
                const res = await axios.get('OwnerManager/menu/' + this.$route.params.name)
                if(res.data.data == null){
                    alert("Menu Tidak Ditemukan")
                }else{
                    alert("Menu Ditemukan")
                    this.data = res.data.data
                    console.log(res,this.data)
                }
               
            }
            catch(error){
                console.log(error)
            }
           
        }
    }
};
</script>
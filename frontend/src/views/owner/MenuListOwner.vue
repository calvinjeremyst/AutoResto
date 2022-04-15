<template>
    <div class="table" id="app">
        <h1><b> Menu List</b></h1>
        <center>
            <table>
                <tr>
                    <th>ID</th>
                    <th>Description</th>
                    <th>Menu Name</th>
                    <th>Menu Price</th>
                    <th>Data Option</th>
                </tr>
                <tr v-for="log in data" :key="log.id">
                    <td>{{log.id}}</td>
                    <td>{{log.description}}</td>
                    <td>{{log.Menu.name}}</td>
                    <td>{{log.Menu.price}}</td>
                    <td>
                        <router-link :to="{name : 'UpdateMenu',params:{'id':log.id}}">
                            <button name="edit" class="btnUpdate">Update</button>
                        </router-link>
                            <button name="delete" class="btnDelete" @click="deleteMenu(log.id)">Delete</button>
                    </td>
                </tr>     
            </table>
        </center>
    </div>
</template>

<script>
import axios from "axios"
export default {
data(){
    return{
        data : [],
    }
},
mounted(){
    this.fetchData()
},

methods:{
    async fetchData(){
        try{
            const res = await axios.get('OwnerManager/allmenuowner');
            if(res.data.data == null){
                alert("Menu Empty")
            }else{
                alert("Showing list menu")
                this.data = res.data.data
                console.log(res,this.data)
            }
            
        }
        catch(error){
            console.log(error)
        }
    },
    async deleteMenu(id){
        try{
            await axios.delete(`OwnerManager/${id}`)
            this.fetchData()
        }
        catch(error){
            console.log(error)
        }
    }
}

}
</script>

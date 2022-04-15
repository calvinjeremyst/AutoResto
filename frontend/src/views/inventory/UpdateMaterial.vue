<template>
<div class="bg-addmaterial">
    <div class="container">
        <form @submit.prevent="UpdateMaterial"> 
            <div class="headcard" style="padding :1rem">
                <h2 class = "title">Update Material</h2>
            </div>
                <div class = "namemtdiv" style = "padding : 2rem">
                    <b><label for = "name" class = "materialname" style = "text-size:25px">Name:</label></b>
                       <input type = "name" v-model ="Name" class = "isiname" id = "name"><br>
                </div>
                <div class="qtymtdiv" style = "padding : 2rem">
                    <b><label for = "quantity" class = "materialquantity">Quantity:</label></b>
                    <input type = "number" v-model ="Quantity" class = "isiquty" id = "quantity"><br>
                </div>
                <div class="unitdiv" style="padding : 2rem">
                    <b><label for = "unit" class = "materialunit">Unit:</label></b>
                    <input type = "text" v-model ="Unit" class = "isiunit" id = "unit"><br>
                </div>
           
                <div class="buttons"> 
                    <button name = "insertmaterial" class = "btninsert">Edit</button>
                </div>
        </form>
    </div>
</div>
</template>
<script>

import axios from "axios"
export default{
    
    name : 'UpdateMaterial',
    data(){
        return {
            Name : '',
            Quantity : '',
            Unit : '',
        }
    },
    mounted(){
        this.fetchById()
    },
    methods:{
        async fetchById(){
            try{
                const res = await axios.get('InventoryManager/'+this.$route.params.id);
                this.Name = res.data.data[0].Name
                this.Quantity = res.data.data[0].Quantity
                this.Unit = res.data.data[0].Unit
                console.log(res)
                
            }catch(error){
                console.log(error)
            }
        },
        async UpdateMaterial(){
            try{
                const response = await axios.post('InventoryManager/' + this.$route.params.id,
                {Name : this.Name, Quantity : parseInt(this.Quantity) ,Unit : this.Unit});
                console.log(response)
                alert("Update Material Success")
            }
            catch(error){
                console.log(error)
                alert("Update Material Gagal")
            }
        }
    }
};


</script>
<style>
    .container{
        padding: 2000x;
        width: 100%;
        height: 100%;
        margin: 2% auto;
        background-color: rgba(250, 250, 250, 0.285);
    }

    .title{
       text-align: center;
    }

    .isiname{
        margin-left: 25px;
        width: 150px;
        height: 30px;
    }

    .isiquty{
        margin-left: 8px;
        width: 150px;
        height: 30px;
    }

    .isiunit{
        margin-left: 38px;
        width: 150px;
        height: 30px;
    }
    
    .namemtdiv{
        text-align: center;
    }

    .qtymtdiv{
        text-align: center;
    }

    .unitdiv{
        text-align: center;
    }

    .btninsert{
      margin-left: 40rem;
      width : 75px;
    }

</style>
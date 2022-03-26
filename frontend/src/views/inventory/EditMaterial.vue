<template>
<div class="bg-addmaterial">
    <div class="container">
        <form> 
            <div class="headcard" style="padding :1rem">
                <h2 class = "title">Update Material</h2>
            </div>
            <div class="cardbody">
                <div class = "namemtdiv" style = "padding : 2rem">
                    <b><label for = "name" class = "materialname" style = "text-size:25px">Name:</label></b>
                    <input type = "name" v-model = "material.name" class = "isiname" id = "name"><br>
                </div>
                <div class="qtymtdiv" style = "padding : 2rem">
                    <b><label for = "quantity" class = "materialquantity">Quantity:</label></b>
                    <input type = "text" v-model = "material.quantity" class = "isiquty" id = "quantity"><br>
                </div>
                <div class="unitdiv" style="padding : 2rem">
                    <b><label for = "unit" class = "materialunit">Unit:</label></b>
                    <input type = "text" v-model = "material.unit" class = "isiunit" id = "unit"><br>
                </div>
                <div class="buttons"> 
                    <button name = "insertmaterial" class = "btninsert">Edit</button>
                </div>
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
            'material':{
                'name':'',
                'quantity':'',
                'unit': '',
            }
        }
    },
    mounted(){
        this.fetchById()
    },
    methods:{
        async fetchById(){
            try{
                const res = await axios.get('InventoryManager/'+this.$route.params.id);
                this.material.name = res.data.name,
                this.material.quantity = res.data.quantity,
                this.material.unit = res.data.unit
                console.log(res)
                
            }catch(error){
                console.log(error)
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
<template>
<div class="bg-addmaterial">
    <div class="container">
        <form @submit.prevent="InsertMaterial"> 
            <div class="headcard" style="padding :1rem">
                <h2 class = "title">Insert Material</h2>
            </div>
            <div class="cardbody">
                <div class = "namemtdiv" style = "padding : 2rem">
                    <b><label for = "name" class = "materialname" style = "text-size:25px">Name:</label></b>
                    <input type = "name" v-model ="name" class = "isiname" id = "name" required><br>
                </div>
                <div class="qtymtdiv" style = "padding : 2rem">
                    <b><label for = "quantity" class = "materialquantity">Quantity:</label></b>
                    <input type = "number" required v-model ="quantity" class = "isiquty" id = "quantity"><br>
                </div>
                <div class="unitdiv" style="padding : 2rem">
                    <b><label for = "unit" class = "materialunit">Unit:</label></b>
                    <input type = "text" v-model = "unit" class = "isiunit" id = "unit" required><br>
                </div>
                <div class="buttons"> 
                    <button name = "insertmaterial" class = "btninsert">Insert</button>
                </div>
            </div>
        </form>
    </div>
</div>
</template>

<script>
    /*function buildFormData(formData, data, parentKey) {
        if (data && typeof data === 'object' && !(data instanceof Date) && !(data instanceof File)) {
            Object.keys(data).forEach(key => {
            buildFormData(formData, data[key], parentKey ? `${parentKey}[${key}]` : key);
            });
        } else {
            const value = data == null ? '' : data;

            formData.append(parentKey, value);
        }
    }

    function jsonToFormData(data){
        const formData = new FormData();
        buildFormData(formData,data);
        return formData
    }*/

    import axios from "axios";
    export default{
        name : 'AddMaterial',
        data(){
            return {
                name: '',
                quantity : 1000,
                unit: '',        
            };
        },
        
        methods:{
            async InsertMaterial(){ 
                //this.quantity = toString(this.quantity)
                try{
                   
                    const response = await axios.post('InventoryManager/insert',
                    {name : this.name,quantity : this.quantity,unit : this.unit});
                    console.log(response,this.material)
                    alert("Add Material Berhasil")
                }
                catch(error){
                    console.log(error)
                    alert("Add Material Gagal")
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
        margin-left: 50px;
        width: 150px;
        height: 30px;
    }

    .materialquantity{
        margin-left: 10px;
    }

    .isiunit{
        margin-left: 48px;
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
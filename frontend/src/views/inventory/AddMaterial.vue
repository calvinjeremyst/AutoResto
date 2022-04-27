<template>
<div class="bg-addmaterial">
    <div class="container">
        <v-form @submit.prevent="InsertMaterial"> 
            <div class="headcard" style="padding :1rem">
                <h2 class = "title">Insert Material</h2>
            </div>
            <div class="cardbody">
                <div class = "namemtdiv">
                    <b><label for = "name" class = "materialname">Name:</label></b>
                    <input type = "name" v-model = "name" class = "isiname" id = "name" required><br>
                </div>
                <div class="qtymtdiv">
                    <b><label for = "quantity" class = "materialquantity">Quantity:</label></b>
                    <input type = "number" v-model = "quantity" class = "isiqty" id = "quantity" value = "10" min = "10" max = "100" required><br>
                </div>
                <div class="unitdiv">
                    <b><label for = "unit" class = "materialunit">Unit:</label></b>
                    <!-- <input type = "text" v-model = "unit" class = "isiunit" id = "unit" required><br> -->
                    <select v-model = "unit" class="isiunit" id = "unit">
                        <option v-for="unit in items" :key="unit.unit_name" v-bind:value ="unit.unit_name">{{unit.unit_name}}</option>
                    </select>
                </div>
            </div>
            <button name = "insertmaterial" class = "btninsert">Insert</button>
        </v-form>
    </div>
</div>
</template>

<script>
    import axios from "axios";
    export default{
        name : 'AddMaterial',
        data(){
            return {
                name: '',
                quantity : 0,
                unit: '',   
            items: [
                { unit_name: 'Kg' },
                { unit_name: 'Gr' },
                { unit_name: 'Lembar' },
                { unit_name: 'ML' },
                { unit_name: 'Ons' },
                { unit_name: 'Biji' },
            ]     
            };
        },
        methods:{
            async InsertMaterial(){ 
                console.log(this.name,this.quantity,this.unit)
                try{
                    const response = await axios.post('InventoryManager/insert',{
                        name : this.name,
                        quantity : parseInt(this.quantity), 
                        unit : this.unit
                    });
                    console.log(response,this.material)
                    alert("Add Material Berhasil")
                } catch(error){
                    console.log(error)
                    alert("Add Material Gagal")
                }
            }
        }
    };
</script>

<style>
    .namemtdiv{
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .materialquantity{
        margin-left: 32px;
    }

    .qtymtdiv{
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .unitdiv{
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .isiqty{
        height: 30px;
        width: 150px;
        margin-left: 65px;
    }

    .materialunit{
        margin: 0px;
        margin-left: 33px;
        margin-right: -22px;
    }

    .title{
       text-align: center;
    }

    .btninsert{
        margin-top: 20px;
        margin-bottom: 20px;
    }

</style>
<template>
    <div class="bg-menu">
        <div class="container">
            <form @submit.prevent="InsertMenu">
                <div class="headcard">
                    <h2 class="title">Insert Menu</h2>
                </div>
                <div class="cardbody">
                    
                    <div class="desc">
                        <b><label for="description">Description</label></b>
                        <input type="text" v-model="data.description" class="isidescription"><br>
                    </div>

                    <div class="menuname">
                        <b><label for="menu" class="labelmenu">Menu Name</label></b>
                        <input type="text" v-model="data.name" class="isimenuname"><br>
                    </div>

                    <div class="price">
                        <b><label for="price">Price</label></b>
                        <input type="text" required v-model="data.price" class="isiprice"><br>
                    </div>

                    <div class="material">
                        <b><label for="material" class="labelmaterial">Material Name</label></b>
                        <input type="text" v-model="data.material" class="isimaterial"><br>
                    </div>

                    <div class="unitrecipe">
                        <b><label for="unitrecpe" class="untrecipe">Unit Recipe</label></b>
                        <input type="text" v-model="data.unit_recipe" class="isiunitrecipe"><br>
                    </div>

                    <div class="quantityrecipe">
                        <b><label for="quantityrecipe" class="qtyrecipe">Quantity Recipe</label></b>
                        <input type="text" v-model="data.quantity_recipe" class="isiquantityrecipe"><br>
                    </div>

                    <div class="quantitymaterial">
                        <b><label for="quantitymaterial" class="qtymaterial">Quantity Material</label></b>
                        <input type="text" v-model="data.quantity_material" class="isiquantitymaterial"><br>
                    </div>

                    <div class="unitmaterial">
                        <b><label for="unitmaterial" class="untmaterial">Unit Material</label></b>
                        <input type="text" v-model="data.unit_material" class="isiunitmaterial">
                    </div>

                    <div class="buttons">
                        <button name="insertmenu" class="btn-insertmenu">Insert</button>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>
<script>

function buildFormData(formData, data, parentKey) {
    if (data && typeof data === 'object' && !(data instanceof Date) && !(data instanceof File)) {
        Object.keys(data).forEach(key => {
        buildFormData(formData, data[key], parentKey ? `${parentKey}[${key}]` : key);
        });
    } 
    else {
        const value = data == null ? '' : data;
        formData.append(parentKey, value);
    }
}

function jsonToFormData(data){
    const formData = new FormData();
    buildFormData(formData,data);
    return formData
}

import axios from 'axios'
export default{
name:'AddMenu',
data(){
    return{
        'data':{
            'description' : '',
            'name' : '',
            'price' : '',
            'material' : '',
            'unit_recipe' : '',
            'quantity_recipe': '',
            'quantity_material': '',
            'unit_material' : '',
        }
    };
},
methods:{
    async InsertMenu(){
        this.data = jsonToFormData(this.data)
        try{
            const response = await axios.post('/OwnerManager/insert',this.data);
            console.log(response,this.data)
        }
        catch(error){
            console.log(error)
        }
    }
}


};
</script>

<style>
    .desc{
        padding: 2rem;
        text-align: center;
    }

    .isidescription{
        margin-left: 30px;
        height: 30px;
    }

    .menuname{
        padding: 2rem;
        text-align: center;
        margin-right: 10px;
    }

    .isimenuname{
        margin-left: 30px;
        height: 30px;
    }

    .labelmenu{
        margin-left: 20px;
    }

    .price{
        padding: 2rem;
        text-align: center;
    }

    .isiprice{
        margin-left: 88px;
        height: 30px;
    }

    .material{
        padding : 2rem;
        text-align: center;
    }
    .isimaterial{
        margin-left : 30px;
        height : 30px

    }

    .unitrecipe{
        padding : 2rem;
        text-align : center;
    }

    .isiunitrecipe{
        margin-left : 52px;
        height : 30px
    }

    .quantityrecipe{
        padding : 2rem;
        text-align : center;
    }

    .isiquantityrecipe{
        margin-left:25px;
        height : 30px;
    }

    .quantitymaterial{
        padding : 2rem;
        text-align : center;
    }

    .isiquantitymaterial{
        margin-left: 20px;
        height: 30px;
    }

    .unitmaterial{
        padding : 2rem;
        text-align: center;
    }

    .isiunitmaterial{
        margin-left: 49px;
        height: 30px;
    }

    .btn-insertmenu{
        margin-left: 40rem;
        width : 75px;
    }



</style>
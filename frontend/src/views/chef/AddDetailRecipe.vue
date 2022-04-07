<template>
    <div class="bg-menu">
        <div class="container">
            <form @submit.prevent="InsertDetailRecipe">
                <div class="headcard">
                    <h2 class="title">Insert Detail Recipe</h2>
                </div>
                <div class="cardbody">
                    <div class="qty">
                        <b><label for="Quantity">Quantity Material</label></b>
                        <input type="text" v-model="data.quantity_recipe" class="isiquantity"><br>
                    </div>

                    <div class="materialname">
                        <b><label for="material" class="labelmaterial">Material Name</label></b>
                        <input type="text" v-model="data.material" class="isimaterialname"><br>
                    </div>

                    <div class="desc">
                         <b><label for="description" class="labeldesc">Description Menu</label></b>
                        <input type="text" v-model="data.description" class="isidescription"><br>
                    </div>
                    <div class="unt">
                        <b><label for="unit" class="labelunit">Unit</label></b>
                        <input type="text" required v-model="data.unit_recipe" class="isiunit"><br>
                    </div>
                    <div class="buttons">
                        <button name="insertrecipedetail" class="btn-insertrecipedetail">Insert</button>
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
name : 'AddDetailRecipe',
data(){
    return{
        data :{
            'quantity_recipe':'',
            'material' : '',
            'description':'',
            'unit_recipe': '',

        }
    };
},


methods:{
    async InsertDetailRecipe(){
        this.data = jsonToFormData(this.data)
        try{
            const response = await axios.post('/ChefManager/insertdetailrecipe',this.data);
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
 .qty{
        padding: 2rem;
        text-align: center;
    }

    .isiquantity{
        margin-left: 10px;
        height: 30px;
    }

    .materialname{
        padding: 2rem;
        text-align: center;
        margin-right: 10px;
    }

    .isimaterialname{
        margin-left: 30px;
        height: 30px;
    }

    .labelmaterial{
        margin-left: 20px;
    }

    .desc{
        padding : 2rem;
        text-align: center;
        margin-right : 12px;
    }

  

    .unit{
        padding: 5rem;
        margin-left: 50px;
        text-align: center;
    }
    .labelunit{
        margin-left: 520px;
    }
    .isiunit{
        margin-left: 120px;
        height: 30px;
    }

    .buttons{
        padding: 3rem;
    }
    .btn-insertrecipedetail{
    
        margin-left: 40rem;
        width : 75px;
    }

</style>
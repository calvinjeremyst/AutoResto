<template>
    <div class="bg-menu">
        <div class="container">
            <form @submit.prevent="InsertDetailRecipe">
                <div class="headcard">
                    <h2 class="title">Insert Detail Recipe</h2>
                </div>
                <div class="cardbody">
                    <div class="description">
                        <b><label for="desc" class="labeldesc">Menu Food</label></b>
                        <select v-model= "recipe" class="isidescription">
                            <option v-for="recipe in data2" :key="recipe.id" v-bind:value ="recipe.id">{{recipe.description}}</option>
                        </select>
                    </div>
                    <div class="materialname">
                        <b><label for="material" class="labelmaterial">Material Name</label></b>
                        <select v-model = "material" class="isimaterialname">
                            <option v-for="material in data" :key="material.Id" v-bind:value ="material.Id">{{material.Name}}</option>
                        </select>
                    </div>
                    <div class="qty">
                        <b><label for="Quantity">Quantity Material</label></b>
                        <input type = "number" v-model = "quantity" class = "isiquantity" min="1" max="1000"><br>
                    </div>
                    <div class="unt">
                        <b><label for="unit" class="labelunit">Unit</label></b>
                        <select v-model = "unit" class="isiunit">
                            <option v-for="unit in items" :key="unit.unit_name" v-bind:value ="unit.unit_name">{{unit.unit_name}}</option>
                        </select>
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
import axios from 'axios'
export default{
    name : 'AddDetailRecipe',
    data(){
        return{
            data : [],
            data2 : [],
            quantity : '',
            material : {
                Name : ''
            },
            recipe : {
                description : ''
            },
            unit : '',
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

    mounted(){
        this.GetMaterialName(),
        this.GetDescription()
    },

    methods:{
        async GetMaterialName(){
            try{
                const res = await axios.get('ChefManager/showmaterialname')
                this.data = res.data.data
                console.log(this.data)
                console.log(res)
            }
            catch(error){
                console.log(error)
            }
        },
        
        async GetDescription(){
            try{
                const res = await axios.get('ChefManager/showdesc')
                this.data2 = res.data.data
                console.log(this.data2)
                console.log(res)
            }
            catch(error){
                console.log(error)
            }
        },

        async InsertDetailRecipe(){
            try{
                const response = await axios.post('/ChefManager/insertdetailrecipe',
                {quantity : parseInt(this.quantity), material:{Name : this.material.material.Name}, recipe:{description : this.recipe.recipe.description}, unit : this.unit });
                console.log(response,this.data)
                alert("Insert Recipe Detail Success")
            }
            catch(error){
                alert("Insert Recipe Detail Failed")
                console.log(error)
            }
        },  
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
        width: 150px;
    }
    
    .materialname{
        padding: 2rem;
        text-align: center;
        margin-right: 30px;
    }

    .isimaterialname{
        margin-left: 40px;
        height: 30px;
        width: 150px;
    }

    .unt{
        padding: 2rem;
        text-align: center;
    }

    .description{
        padding: 2rem;
        text-align: center;
    }

    .isidescription{
        width: 150px;
        height: 30px;
    }

    .labelmaterial{
        margin-left: 30px;
        margin-right: -9px;
    }

    .labeldesc{
        margin-right: 30px;
    }

    .desc{
        padding : 2rem;
        text-align: center;
        margin-right : 12px;
    }

    .labelunit{
        margin-right: -10px;
    }

    .isiunit{
        margin-left: 120px;
        height: 30px;
    }

    .buttons{
        padding: 3rem;
        margin-left: 43%;
    }

    .btn-insertrecipedetail{
        width : 75px;
    }
</style>
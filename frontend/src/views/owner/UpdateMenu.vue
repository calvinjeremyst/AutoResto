<template>
    <div class="bg-menu">
        <div class="container">
            <form @submit.prevent="UpdateMenu">
                <div class="headcard">
                    <h2 class="title">Insert Menu</h2>
                </div>
                <div class="cardbody">
                    
                    <div class="desc">
                        <b><label for="description">Description</label></b>
                        <input type="text" v-model="description" class="isidescription" required><br>
                    </div>

                    <div class="menuname">
                        <b><label for="menu" class="labelmenu">Menu Name</label></b>
                        <input type="text" v-model="menu.name" class="isimenuname" required><br>
                    </div>

                    <div class="price">
                        <b><label for="price">Price</label></b>
                        <input type="number" v-model="menu.price" class="isiprice" required><br>
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
import axios from "axios"
export default{
    name : 'UpdateMenu',
    data(){
        return{
            description : '',
            menu:{
                name : '',
                price : ''
            }
          
        }
    },
    mounted(){
        this.fetchById()
    },
    methods:{
        async fetchById(){
            try{
                
                const res = await axios.get('OwnerManager/showmenu_id/'+this.$route.params.id);
                this.description = res.data.data[0].description
                this.menu.name = res.data.data[0].Menu.name
                this.menu.price = res.data.data[0].Menu.price
                console.log(res)
            }catch(error){
                console.log(error)
            }

        },

        async UpdateMenu(){
            try{
                const res = await axios.post('OwnerManager/updatemenu/' + this.$route.params.id,
                {description : this.description, menu : {name : this.menu.name, price : parseFloat(this.menu.price)}})
                console.log(res)
                alert("Update menu success")
            }
            catch(error){

                console.log(error)
                alert("Update menu failed")
            }
          
        }
    }

}


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
    .btn-insertmenu{
        margin-left: 40rem;
        width : 75px;
    }



</style>

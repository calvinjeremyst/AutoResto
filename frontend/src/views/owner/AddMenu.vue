<template>
    <div class="bg-menu">
        <div class="container">
            <form @submit.prevent="InsertMenu">
                <div class="headcard">
                    <h2 class="title">Insert Menu</h2>
                </div>
                <div class="cardbody">
                    <div class="menuname">
                        <b><label for="menu" class="label_menu">Menu Name</label></b>
                        <input type="text" v-model="data.name" class="isimenuname" required><br>
                    </div>
                    <div class="desc">
                        <b><label for="description" class="label_desc">Description</label></b>
                        <input type="text" v-model="data.description" class="isidescription" required><br>
                    </div>
                    <div class="price">
                        <b><label for="price" class="label_price">Price</label></b>
                        <input type="text" v-model="data.price" class="isiprice" required><br>
                    </div>
                    <button name="insertmenu" class="btn-insertmenu">Insert</button>
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
                'description': '',
                'name' : '',
                'price' : '',
            }
        };
    },
    methods:{
        async InsertMenu(){
            this.data = jsonToFormData(this.data)
            try{
                const response = await axios.post('/OwnerManager/insertmenu',this.data);
                console.log(response,this.data)
                alert("Add Menu Berhasil")
            }
            catch(error){
                alert("Menu sudah ada")
                console.log(error)
            }
        }
    }
};

</script>

<style>

.label_price{
    margin-left: 10px;
}

.label_desc{
    margin-right: 10px;
}

.isidescription{
    width: 150px;
}

.isimenuname{
    width: 150px;
}

.isiprice{
    width: 150px;
    margin-right: 20px;
}

.btn-insertmenu{
    margin-top: 20px;
    margin-bottom: 20px;
}

</style>
<template>
    <form @submit.prevent="fetchData">
        <label for="id">Input ID Recipe</label>
        <input type="text" id="idrecipe" v-model="recipeid">
        <router-link :to="{name:'MenuRecipe',params:{'id':recipeid}}">
            <button>submit</button>
        </router-link>
    </form>
</template>


<script>
export default{
    data(){
        return{
            recipeid : null,
        }
    }
}


</script>


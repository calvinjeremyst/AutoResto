<template>
    <div class="bg-searchmaterial">
       <div class="cardbody">
           <div class="title">
               <h2>Form Menu Recipe</h2>
            </div>
            <form @submit.prevent="fetchData">  
                <div class="keyword">
                    <h3><b><label for="id" class="labelkeyword">Input ID Recipe</label></b></h3><br>
                    <input type="text" class="searchkeyword" id="idrecipe" v-model="recipeid">
                </div>
                <div class="buttonsubmit">
                <router-link :to="{name:'MenuRecipe',params:{'id':recipeid}}">
                    <button class="button-submit">Submit</button>
                 </router-link>
                </div>
            </form>
        </div>
    </div>
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

<style>

    .buttonsubmit{
        text-align: center;
        padding: 0.5rem;
    }

    .button-submit{
        padding: 0.5rem;
    }

</style>

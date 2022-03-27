<template>
     <form @submit.prevent="fetchData">
        <label for="id">Input ID Recipe</label>
        <input type="text" id="idrecipe" >
        <router-link to="MenuRecipe">
            <button>submit</button>
        </router-link>
    </form>
</template>


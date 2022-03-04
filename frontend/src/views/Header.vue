<template>
    <header>
        <h3>{{ title }} </h3>
    </header>
    
</template>

<script>
    export default {
        data() {
            return {
                title: "AutoResto"
            }
        },
    }
</script>

<style scoped>
    header {
        background-color: darkslategrey;
        padding: 10px;
        height: 50px;
        text-align: center;
        color: white;
    }
</style>

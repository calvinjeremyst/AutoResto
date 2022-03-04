<template>
    <footer>
        <p> {{ footer-text }} </p>
    </footer>
</template>

<script>
    export default {
        data() {
            return {
                footerText: "Copyright by AutoResto"
            }
        }
    }
</script>

<style scoped>
    footer {
        background-color: blanchedalmond;
        padding: 10px;
        text-align: center;
    }
</style>
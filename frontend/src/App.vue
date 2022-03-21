<template>
  <div class="app">
    <div v-if="idRole == 'Guest'">
      <Header />
    </div>
    <div v-if="idRole == 0">
      <ChefNavbar />
    </div>
    <div v-else-if="idRole == 1">
      <OwnerNavbar />
    </div>
    <div v-else-if="idRole == 2">
      <InventoryNavbar />
    </div>
    <router-view />
    <Footer />
  </div>
</template>

<script>
    import OwnerNavbar from "./components/OwnerNavbar";
    import InventoryNavbar from "./components/InventoryNavbar";
    import ChefNavbar from "./components/ChefNavbar";
    import Login from "./services/Login";
    import Header from "./components/Header"
    import Footer from "./components/Footer";
import OwnerNavbar1 from "./components/OwnerNavbar.vue";

    export default {
        mounted() {
            this.fetchData();
        },
        data: () => {
            const data = [];
            return {
                idRole: null,
                loginService: new Login(),
                data,
            };
        },
        components: {
    OwnerNavbar,
    InventoryNavbar,
    ChefNavbar,
    Footer,
    Header,
    OwnerNavbar1
},
        methods: {
            async fetchData() {
                this.idRole = this.loginService.getCurrentUserType();
            },
            async getUserType() {
                return this.idRole;
            },
        },
    };
</script>

<style lang="scss">
@import "../scss/main.scss";

</style>

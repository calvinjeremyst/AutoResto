<template>
  <div class="app">
      <Header />
    <div v-if="userType == 'chef'">
      <ChefNavbar />
    </div>
    <div v-else-if="userType == 'owner'">
      <OwnerNavbar />
    </div>
    <div v-else-if="userType == 'inventory'">
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
    import AddMaterial from "./views/inventory/AddMaterial.vue"

    export default {
        mounted() {
            this.fetchData();
        },
        data: () => {
            const data = [];
            return {
                userType: "",
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
    AddMaterial
},
        methods: {
            async fetchData() {
                this.userType = this.loginService.getCurrentUserType();
            },
            async getUserType() {
                return this.userType;
            },
        },
    };
</script>

<style lang="scss">
@import "../scss/main.scss";

.app {
  background-color: #686868;
  height: 100%;
}
</style>

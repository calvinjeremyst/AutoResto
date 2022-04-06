<template>
  <div class="app">
    <div v-if="userType == 'Guest'">
      <Header />
    </div>
    <div v-else-if="userType == 1">
      <OwnerNavbar />
    </div>
    <div v-else-if="userType == 2">
      <ChefNavbar />
    </div>
    <div v-else-if="userType == 3">
      <InventoryNavbar />
    </div>
    <router-view />
    <Footer />
  </div>
</template>

<script>
    //components navbar
    import OwnerNavbar from "./components/OwnerNavbar";
    import InventoryNavbar from "./components/InventoryNavbar";
    import ChefNavbar from "./components/ChefNavbar";


    import Login from "./services/Login";
    import Header from "./components/Header"
    import Footer from "./components/Footer";
  
    export default {
        mounted() {
            this.fetchData();
        },
        data: () => {
            const data = [];
            return {
                userType: null,
                loginService: new Login(),
                data,
            };
        },
        components: {
          OwnerNavbar,
          InventoryNavbar,
          ChefNavbar,

          Footer,
          Header
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

</style>

<template>
  <div class="app">
    <div>
      <Header />
    </div>
    <div v-if="userType == 0">
      <ChefNavbar />
      <HomePageChef />
    </div>
    <div v-else-if="userType == 1">
      <HomePageOwner />
      <OwnerNavbar />
    </div>
    <div v-else-if="userType == 2">
      <HomePageInventory />
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

//import HomePageChef from './components/ChefHomePage.vue'
//import HomePageInventory from './components/InventoryHomePage.vue'
//import HomePageOwner from './components/OwnerHomePage.vue'

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

          //HomePageInventory,
          //HomePageChef,
          //HomePageOwner,
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

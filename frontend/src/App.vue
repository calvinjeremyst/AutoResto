<template>
  <div class="app">
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
    import InventoryNavbar from "./components/StaffNavbar";
    import ChefNavbar from "./components/ChefNavbar";
    import Login from "./services/Login";
    import Footer from "./components/Footer";

    export default {
        mounted() {
            this.fetchData();
        },
        data: () => {
            const data = [];
            return {
                userType: "",
                loginService: new LoginService(),
                data,
            };
        },
        components: {
            Navbar,
            OwnerNavbar,
            StaffNavbar,
            ChefNavbar,
            DeliveryNavbar,
            Footer,
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
  background-color: #ececec;
  height: 100%;
}
</style>

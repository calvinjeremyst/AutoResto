<template>
  <div class="app">
    <div v-if="userType == 1">
      <HomePageOwner/>
    </div>
    <div v-else-if="userType == 2">
      <HomePageChef/>
    </div>
    <div v-else-if="userType == 3">
      <HomePageInventory/>
    </div>
    <router-view />
  </div>
</template>
<script>

import HomePageChef from './chef/ChefHomePage.vue'
import HomePageInventory from './inventory/InventoryHomePage.vue'
import HomePageOwner from './owner/OwnerHomePage.vue'
import Login from "../services/Login";

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
    HomePageChef,
    HomePageInventory,
    HomePageOwner,
  },

  methods: {
    async fetchData() {
      this.userType = this.loginService.getCurrentUserType();
      //owner = this.owner
    },
    async getUserType() {
      return this.userType;
    },
  },
};
</script>

<template>
  <b-navbar style="background-color: white; padding: 16px" shadow>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/AutoResto.png" />
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item href="/">
        <b-icon pack="fas" icon="home" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Home</span>
      </b-navbar-item>
      <b-navbar-item href="/inventory/process">
        <b-icon
          pack="fas"
          icon="shopping-cart"
          size="is-small"
          type="is-primary"
        >
        </b-icon>
        <span style="margin-left: 0.5rem">Process Orders</span>
      </b-navbar-item>
      <b-navbar-item href="/inventory/view">
        <b-icon
          pack="fas"
          icon="shopping-cart"
          size="is-small"
          type="is-primary"
        >
        </b-icon>
        <span style="margin-left: 0.5rem">View Order</span>
      </b-navbar-item>
      <b-navbar-item href="">
        <b-icon pack="fas" icon="user" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem" v-on:click="logout"
          ><button class="btn">Logout</button></span
        >
      </b-navbar-item>
    </template>
  </b-navbar>
</template>
<script>
import axios from "axios";
import LoginService from "@/services/LoginService";

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
  methods: {
    async fetchData() {
      await axios.get("/logout");
    },
    async logout() {
      console.log("test");
      this.loginService.removeUserType();
      location.replace("/");
    },
  },
};
</script>
<style scoped>
.btn {
  background-color: white;
  border: none;
  font-size: 15px;
  font-weight: normal;
  margin-left: -0.7%;
}
</style>
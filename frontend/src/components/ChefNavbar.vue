<template>
  <b-navbar style="background-color: white; padding: 16px" shadow>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/AutoResto.png" />
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item href="/menuRecipe">
        <b-icon pack="fas" icon="book" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Menu and Recipe</span>
      </b-navbar-item>
      <b-navbar-item href="">
        <b-icon pack="fas" icon="user" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem" v-on:click="logout">
            <button class="btn">Logout</button>
        </span>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>
<script>
import axios from "axios";
import Login from "@/services/Login";

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
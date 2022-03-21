<template>
  <b-navbar style="background-color: white; padding: 16px" shadow>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/AutoResto.png" />
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item href="/add_menu">
        <b-icon pack="fas" icon="plus-square" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Add Menu</span>
      </b-navbar-item>
      <b-navbar-item href="/recipe_list">
        <b-icon pack="fas" icon="sticky-note" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Recipe List</span>
      </b-navbar-item>
      <b-navbar-item href="/material_list">
        <b-icon pack="fas" icon="sticky-note-o" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Material List</span>
      </b-navbar-item>
      <b-navbar-item href="/menu_list">
        <b-icon pack="fas" icon="book" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Menu List</span>
      </b-navbar-item>
      <b-navbar-item href="/search_material">
        <b-icon pack="fas" icon="search" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Search Material</span>
      </b-navbar-item>
      <b-navbar-item href="/search_menu">
        <b-icon pack="fas" icon="search" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Search Menu</span>
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
<template>
  <b-navbar style="background-color:lightgray; padding: 16px" shadow>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/AutoResto.png" />
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item href="/">
        <span style="margin-left: 0.5rem; color:darkgoldenrod; font-size: 110%"><b>AutoResto</b></span>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>
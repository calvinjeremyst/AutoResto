<template>
  <b-navbar style="background-color: white; padding: 16px" shadow>
    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="@/assets/AutoResto.png" />
      </b-navbar-item>
    </template>
    <template slot="start">
      <b-navbar-item href="/menu">
        <b-icon pack="fas" icon="pizza-slice" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Menu</span>
      </b-navbar-item>
      <b-navbar-item href="/order">
        <b-icon
          pack="fas"
          icon="shopping-cart"
          size="is-small"
          type="is-primary"
        >
        </b-icon>
        <span style="margin-left: 0.5rem">Order Now</span>
      </b-navbar-item>
      <b-navbar-item href="/about">
        <b-icon pack="fas" icon="store" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">About Us</span>
      </b-navbar-item>
      <b-navbar-item href="/tracking">
        <b-icon pack="fas" icon="car" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Tracking</span>
      </b-navbar-item>
      <b-navbar-item href="/login">
        <b-icon pack="fas" icon="user" size="is-small" type="is-primary">
        </b-icon>
        <span style="margin-left: 0.5rem">Login</span>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>
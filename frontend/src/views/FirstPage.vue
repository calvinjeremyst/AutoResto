<template>
  <div id="app">
    <router-link to="/login">
      <transition-group appear name="slide-in" class="grid-wrapper" tag="div">
        <Card
          v-for="(card, index) in cards"
          :key="'card' + index"
          :item="card"
          :index="index"
        />
      </transition-group>
    </router-link>
  </div>
</template>

<script>
import Card from "../components/FirstPageLayout";
export default {
  data: () => {
    return {
      cards: [
        { img: "/home.png", grid: "card-3-1" },
      ],
    };
  },
  components: {
    Card,
  },
};
</script>

<style scoped>
#app {
  position: relative;
  width: 100%;
}

.slide-in-enter {
  opacity: 0;
  transform: scale(0.5);
}

.slide-in-enter-active {
  transition: all 0.4s ease;
  transition-delay: calc(0.1s * var(--i));
}
</style>

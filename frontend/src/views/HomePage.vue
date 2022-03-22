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
        { img: "/addMaterial.png", grid: "card-1-1" },
        { img: "/materialList.png", grid: "card-1-1" },
        { img: "/searchMaterial.png", grid: "card-1-1" },
        { img: "/searchMenu.png", grid: "card-1-1" },
        { img: "/menuList.png", grid: "card-1-1" },
        { img: "/recipeList.png", grid: "card-1-1" },
        { img: "/menuRecipe.png", grid: "card-1-1" },
      ],
    };
  },
  components: {
    Card,
  },
};
</script>

<style scoped>
.column {
  padding: 0;
}

@media screen and (max-width: 667px) {
  .font {
    font-size: 25px !important;
  }

  .hero-body {
    padding-left: 24px !important;
  }
}

@media screen and (min-width: 668px) and (max-width: 1023px) {
  .font {
    font-size: 35px !important;
  }

  .hero-body {
    padding-left: 36px !important;
  }
}

#app {
  position: relative;
  width: 100%;
  text-align: center;
}
.grid-wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-gap: 10px;
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

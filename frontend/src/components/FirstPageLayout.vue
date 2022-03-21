<template>
  <div class="card" :class="gridClass" :style="gridStyle"></div>
</template>

<script>
export default {
  props: ["item", "index"],
  computed: {
    gridStyle() {
      return {
        "--i": this.index,
        "background-image": "url(" + require("../assets" + this.item.img) + ")",
      };
    },
    gridClass() {
      return {
        "card-3-1": this.item.grid === "card-3-1",
      };
    },
  },
};
</script>

<style lang="scss" scoped>
.card {
  width: 100%;
  height: 500px;
  background-position: center;
  background-size: cover;
}

.card-3-1 {
  grid-column: span 3;
}
</style>

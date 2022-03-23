<template>
  <div class="card" :style="gridStyle"></div>
</template>

<script>
export default {
  props: ["item", "index"],
  computed: {
    gridStyle() {
      return {
        "--i": this.index,
        "background-image": "url(" + require("../assets/home" + this.item.img) + ")",
      };
    },
  },
};
</script>

<style lang="scss" scoped>
.card {
  width: 300px;
  height: 200px;
  background-position: center;
  background-size: cover;
}
</style>
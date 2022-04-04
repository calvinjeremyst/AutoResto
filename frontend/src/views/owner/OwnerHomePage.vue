<template>
    <div id="app">
        <div class="grid-wrapper">
            <div class="grid-item">
                <router-link to="/add_menu">
                    <Card
                        v-for="(card, index) in card1"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/menu_list_owner">
                    <Card
                        v-for="(card, index) in card2"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/search_menu">
                    <Card
                        v-for="(card, index) in card3"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/search_material">
                    <Card
                        v-for="(card, index) in card4"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/material_list">
                    <Card
                        v-for="(card, index) in card5"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/recipe_list">
                    <Card
                        v-for="(card, index) in card6"
                        :key="'card' + index"
                        :item="card"
                        :index="index"
                    />
                </router-link>
            </div>
        </div>
    </div>
</template>

<script>
import Card from "../../components/HomeLayout.vue";
export default{
    data: ()=>{
        return{
            card1 :[
                { img: "/addMenu.png", grid: "card-1-1"}
            ],
            card2 : [
                {img : "/menuList.png",grid: "card-1-1"}
            ],
            card3 : [
                {img : "/searchMenu.png",grid: "card-1-1"}
            ],
            card4 : [
               {img : "/searchMaterial.png",grid: "card-1-1"}
            ],
            card5 : [
                {img : "/materialList.png",grid : "card-1-1"}
            ],
            card6 : [
                {img : "/recipeList.png",grid : "card-1-1"}
            ]

        }
    },
    components : {
        Card,
    }
}
</script>

<style>
    column {
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
  grid-template-columns: auto auto auto;
  padding: 10px;
  margin-left: 6%;
}

.grid-item {
  padding: 10px;
}

</style>
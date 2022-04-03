<template>
    <div id="app">
        <div class="grid-wrapper">
             <div class="grid-item">
                <router-link to="/add_material">
                    <Card
                    v-for="(card, index) in card1"
                    :key="'card' + index"
                    :item="card"
                    :index="index"
                />
                </router-link>
            </div>
            <div class="grid-item">
                <router-link to="/material_list">
                    <Card
                    v-for="(card, index) in card2"
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
    import Card from "../components/HomeLayout.vue";
export default{
    data: ()=>{
        return{
            card1 :[
                { img: "/addMaterial.png", grid: "card-1-1"}
            ],
            card2 : [
                {img : "/materialList.png",grid: "card-1-1"}
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


import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

//Component
import Login from '../views/Login.vue';
import FirstPage from '../views/FirstPage.vue';
import HomePage from '../views/HomePage.vue';
import AddMaterial from '../views/inventory/AddMaterial.vue'
import MaterialList from '../views/MaterialList.vue'
import MenuList from '../views/MenuList.vue'
import RecipeList from '../views/owner/RecipeList.vue'
import SearchMaterial from '../views/owner/SearchMaterial.vue'
import SearchMenu from '../views/owner/SearchMenu.vue'
import AddMenu from '../views/owner/AddMenu'
import MenuRecipe from '../views/chef/MenuRecipe'

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
      path: '/homepage',
      name: 'HomePage',
      component: HomePage,
    },
    {
      path: '/firstpage',
      name: 'FirstPage',
      component: FirstPage,
    },
    {
      path: '/addMaterial',
      name: 'Add Material',
      component: AddMaterial,
    },
    {
      path: '/materialList',
      name: 'MaterialList',
      component: MaterialList,
    },
    {
      path: '/menuList',
      name: 'MenuList',
      component: MenuList,
    },
    {
      path: '/recipelist',
      name: 'RecipeList',
      component: RecipeList,
    },
    {
      path: '/searchMaterial',
      name: 'SearchMaterial',
      component: SearchMaterial,
    },
    {
      path: '/searchMenu',
      name: 'SearchMenu',
      component: SearchMenu,
    },
    {
      path: '/addMenu',
      name: 'AddMenu',
      component: AddMenu,
    },
    {
      path: '/menuRecipe',
      name: 'MenuRecipe',
      component: MenuRecipe,
    },
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    linkActiveClass: 'is-active',
    routes,
    scrollBehavior() {
      window.scrollTo(0, 0);
    },
  });
  
  export default router;
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
        path: '/',
        name: 'Firstpage',
        component: FirstPage,
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
    },
    {
      path: '/home',
      name: 'HomePage',
      component: HomePage,
    },
    {
      path: '/add_material',
      name: 'Add Material',
      component: AddMaterial,
    },
    {
      path: '/material_list',
      name: 'MaterialList',
      component: MaterialList,
    },
    {
      path: '/menu_list',
      name: 'MenuList',
      component: MenuList,
    },
    {
      path: '/recipe_list',
      name: 'RecipeList',
      component: RecipeList,
    },
    {
      path: '/search_material',
      name: 'SearchMaterial',
      component: SearchMaterial,
    },
    {
      path: '/search_menu',
      name: 'SearchMenu',
      component: SearchMenu,
    },
    {
      path: '/add_menu',
      name: 'AddMenu',
      component: AddMenu,
    },
    {
      path: '/menu_recipe',
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
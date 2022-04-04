import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

//Component
import Login from '../views/Login.vue';
import FirstPage from '../views/FirstPage.vue';
import HomePage from '../views/HomePage.vue';
//Component Material
import AddMaterial from '../views/inventory/AddMaterial.vue'
import MaterialList from '../views/MaterialList.vue'
import MaterialListSearch from '../views/owner/MaterialListSearch.vue'
import SearchMaterial from '../views/owner/SearchMaterial.vue'
import EditMaterial from '../views/inventory/UpdateMaterial'


//Component Menu
import MenuListChef from '../views/chef/MenuListChef'
import MenuListOwner from '../views/owner/MenuListOwner'
import SearchMenu from '../views/owner/SearchMenu.vue'
import AddMenu from '../views/owner/AddMenu'
import MenuListSearch from '../views/owner/MenuListSearch.vue'

//Component Recipe
import RecipeList from '../views/owner/RecipeList.vue'
import MenuRecipe from '../views/chef/MenuRecipe'
import FormRecipe from '../views/chef/FormRecipe'


//Component Homepage
import HomePageChef from '../views/chef/ChefHomePage.vue'
import HomePageInventory from '../views/inventory/InventoryHomePage.vue'
import HomePageOwner from '../views/owner/OwnerHomePage.vue'

const routes = [
    {
        path: '/welcome',
        name: 'FirstPage',
        component: FirstPage,
        meta: {
          title: "Welcome to AutoResto"
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
          title: "Login"
        }
    },
    {
      path: '/home',
      name: 'HomePage',
      component: HomePage,
      meta: {
        title: "Home"
      }
    },
    {
      path: '/add_material',
      name: 'AddMaterial',
      component: AddMaterial,
      meta: {
        title: "Add Material"
      }
    },
    {
      path: '/material_list',
      name: 'MaterialList',
      component: MaterialList,
      meta: {
        title: "Material List"
      }
    },
    {
      path : '/material_list_search/material/:name',
      name : 'MaterialListSearch',
      component : MaterialListSearch,
      meta: {
        title : "Material List Searched"
      }
    },
    {
      path : '/menu_list_search/menu/:name',
      name : 'MenuListSearch',
      component : MenuListSearch,
      meta : {
        title : "Menu List Search"
      }
    },

    {
      path: '/menu_list_chef',
      name: 'MenuListChef',
      component: MenuListChef,
      meta: {
        title: "Menu List Chef"
      }
    },
    {
      path : '/menu_list_owner',
      name : 'MenuListOwner',
      component : MenuListOwner,
      meta:{
        title : "Menu List Owner"
      }
    },
    {
      path: '/recipe_list',
      name: 'RecipeList',
      component: RecipeList,
      meta: {
        title: "Recipe List"
      }
    },
    {
      path: '/search_material',
      name: 'SearchMaterial',
      component: SearchMaterial,
      meta: {
        title: "Search Material"
      }
    },
    {
      path: '/search_menu',
      name: 'SearchMenu',
      component: SearchMenu,
      meta: {
        title: "Search Menu"
      }
    },
    {
      path: '/add_menu/',
      name: 'AddMenu',
      component: AddMenu,
      meta: {
        title: "Add Menu"
      }
    },
    {
      path: '/menu_recipe/:id',
      name: 'MenuRecipe',
      component: MenuRecipe,
      meta: {
        title: "Menu Recipe"
      }
    },

    {
      path : '/update_material/:id',
      name : 'UpdateMaterial',
      component : EditMaterial,
      meta: {
        title : "Update Material"
      }
    },
    {
      path : '/form_recipe_menu',
      name : 'FormRecipe',
      component : FormRecipe,
      meta : {
        title : "Form Recipe Menu"
      }
    },
    {
      path : '/home_chef',
      name : 'HomePageChef',
      component : HomePageChef,
      meta : {
        title : "Home"
      }
    },
    {
      path : '/home_inventory',
      name : 'HomePageInventory',
      component : HomePageInventory,
      meta : {
          title : "Home"
      }
    },
    {
      path : '/home_owner',
      name : 'HomePageOwner',
      component : HomePageOwner,
      meta : {
        title : "Home"
      }
    }

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

// This callback runs before every route change, including on page load.
router.beforeEach((to, from, next) => {
  // This goes through the matched routes from last to first, finding the closest route with a title.
  // e.g., if we have `/some/deep/nested/route` and `/some`, `/deep`, and `/nested` have titles,
  // `/nested`'s will be chosen.
  const nearestWithTitle = to.matched.slice().reverse().find(r => r.meta && r.meta.title);

  // Find the nearest route element with meta tags.
  const nearestWithMeta = to.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);

  const previousNearestWithMeta = from.matched.slice().reverse().find(r => r.meta && r.meta.metaTags);

  // If a route with a title was found, set the document (page) title to that value.
  if(nearestWithTitle) {
    document.title = nearestWithTitle.meta.title;
  } else if(previousNearestWithMeta) {
    document.title = previousNearestWithMeta.meta.title;
  }

  // Remove any stale meta tags from the document using the key attribute we set below.
  Array.from(document.querySelectorAll('[data-vue-router-controlled]')).map(el => el.parentNode.removeChild(el));

  // Skip rendering meta tags if there are none.
  if(!nearestWithMeta) return next();

  // Turn the meta tag definitions into actual elements in the head.
  nearestWithMeta.meta.metaTags.map(tagDef => {
    const tag = document.createElement('meta');

    Object.keys(tagDef).forEach(key => {
      tag.setAttribute(key, tagDef[key]);
    });

    // We use this to track which meta tags we create so we don't interfere with other ones.
    tag.setAttribute('data-vue-router-controlled', '');

    return tag;
  })
  // Add the meta tags to the document head.
  .forEach(tag => document.head.appendChild(tag));

  next();
});

export default router;
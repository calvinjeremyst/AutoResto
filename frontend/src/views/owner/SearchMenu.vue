<template>
   <div class="bg-searchmenu">
       <div class="cardbody">
           <div class="title">
               <h2>Pencarian Menu</h2>
            </div>
            <form>
                <div class="keyword">
                    <h3><b><label for="keywords" class="labelkeyword">Kata Kunci</label></b></h3><br>
                    <input type="text" class="searchkeyword" v-model="namemenu">
                </div>
                <div class="buttonsearch">
                     <router-link :to="{name:'MenuListSearch', params:{'name':namemenu}}">
                        <button class="button-search">Search</button>
                     </router-link>
                </div>
            </form>

        </div>
    </div>
</template>

<script>

export default{
    data(){
        return{
            namemenu : null,
        }
    }
}

</script>

<style>
    .bg-searchmenu{
         padding: 2000x;
        width: 100%;
        height: 100%;
        margin: 2% auto;
        background-color: rgba(250, 250, 250, 0.285);
    }

    .keyword{
        padding: 0.5rem;
        text-align: center;
    }

    .searchkeyword{
        margin-top: 0.1px;
        width: 20%;
    }

    .buttonsearch{
        margin-left: 48%;
        padding: 0.5rem;
    }
</style>

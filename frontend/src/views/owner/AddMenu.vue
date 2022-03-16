<template>
    <div class="bg-menu">
        <div class="container">
            <form>
                <div class="headcard">
                    <h2 class="title">Insert Menu</h2>
                </div>

                <div class="cardbody">
                    <div class="foodname">
                        <b><label for="foodname">Name:</label></b>
                        <input type="text" required v-model="menuname" class="isifoodname"><br>
                    </div>

                    <div class="pricemenu">
                        <b><label for="pricemenu" class="labelprice">Price</label></b>
                        <input type="text" required v-model="menuprice" class="isifoodprice"><br>
                    </div>

                    <div class="idrecipemenu">
                        <b><label for="idrecipe">ID Recipe</label></b>
                        <input type="text" required v-model="idrecipe" class="isiidrecipe"><br>
                    </div>

                    <div class="buttons">
                        <button name="insertmenu" class="btn-insertmenu">Insert</button>
                    </div>
                </div>

            </form>
        </div>
    </div>
</template>
<script>
export default{
    name:'AddMenu',
    data(){
        return{
            menuname:'',
            menuprice:'',
            idrecipe:''

        }
    }
}
</script>

<style>

.foodname{
    padding: 2rem;
    text-align: center;
}
.isifoodname{
    margin-left: 30px;
    height: 30px;
}
.pricemenu{
    padding: 2rem;
    text-align: center;
    margin-right: 10px;
}

.isifoodprice{
    margin-left: 40px;
    height: 30px;
}
.labelprice{
    margin-left: 20px;
}

.idrecipemenu{
    padding: 2rem;
    text-align: center;
    margin-right: 30px;
}

.isiidrecipe{
    margin-left: 50px;
    height: 30px;
}
 .btn-insertmenu{
      margin-left: 210px;
      width : 75px;
  }

</style>
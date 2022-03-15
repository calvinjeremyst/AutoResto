<template>
<div class="bg-addmaterial">
<div class="container">
    <form>
        <div class="headcard" style="padding :1rem">
            <h2 class = "title">Insert Material</h2>
        </div>

        <div class="cardbody">
            <div class = "namemtdiv" style = "padding : 2rem">
                <b><label for="name" class="materialname" style="text-size:25px">Name:</label></b>
                <input type="name" required v-model="name" class="isiname"><br>
            </div>
            <div class="qtymtdiv" style = "padding : 2rem">
                <b><label for = "quantity" class="materialquantity">Quantity:</label></b>
                <input type="text" required v-model="quantity" class="isiquty"><br>
            </div>

            <div class="unitdiv" style="padding : 2rem">
                <b><label for="unit" class = "materialunit">Unit:</label></b>
                <input type = "text" required v-model="unit" class="isiunit"><br>
            </div>
            <div class="buttons"> 
                <button name="insertmaterial" class="btninsert">Insert</button>
            </div>
            
        </div>
    </form>
</div>
</div>
</template>

<script>
     export default{
         name : 'AddMaterial',
         data(){
             return{
                 name: '',
                 quantity: '',
                 unit:''

             }
         }
     }
</script>

<style>
    .container{
        padding: 2000x;
        width: 100%;
        height: 100%;
        margin: 2% auto;
        background-color: rgba(250, 250, 250, 0.285);

    }
   .title{
       text-size : 200%;
       text-align: center;
   }
    .isiname{
        margin-left: 25px;
        width: 150px;
        height: 30px;
      
    }
    .isiquty{
        margin-left: 8px;
        width: 150px;
        height: 30px;
    }

     .isiunit{
        margin-left: 38px;
        width: 150px;
        height: 30px;
    }
    
    .namemtdiv{
        text-align: center;
    }

    .qtymtdiv{
        text-align: center;
    }

    .unitdiv{
        text-align: center;
    }


  .btninsert{
      margin-left: 210px;
      width : 75px;
  }

 

</style>

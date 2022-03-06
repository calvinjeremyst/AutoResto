<template>
    <div class="bgAbout">
        <div class="content">
            <div class="column">
                <h1 style="padding: 12px 48px" class="has-text-centered title">
                    Halaman Sambutan
                </h1>
            </div>

            <center>
                <img src="require(`@/assets/AutoResto.png`)" width="200" >
                <div class="cnt">
                    <p class="has-text-left intro title is-4 *x-6">
                        Hai! Selamat Datang di AutoResto
                    </p>
                </div>
            </center>
        </div>
    </div>
</template>

<style scoped>
.bgAbout {
  background-image: url("../assets/about/about_bg.png");
  background-repeat: no-repeat;
  overflow: hidden;
  background-size: cover;
}
.content {
  background-color: white;
  opacity: 95%;
}
.intro {
  color: red;
}
.cnt {
  width: 80%;
  height: 400px;
}
</style>

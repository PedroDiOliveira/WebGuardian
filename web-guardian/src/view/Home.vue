<script>
  window.addEventListener('scroll', function() {
            const scrollAmount = window.scrollY;
            const backgroundPositionY = scrollAmount * 0.5;
            document.querySelector('.home').style.backgroundPositionY = `-${backgroundPositionY}px`;
        });

</script>

<template >
    <div class="home">
        <div class="titulo">
            <h1>Web Guadian</h1>
        </div>
        <div class="container">
            <RouterLink to="/guardian" ><button class="botao">Comecar</button></RouterLink>
        </div>
    </div>
</template>

<style>
    .home{
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        width: 100vw;
        height: 120vh;
        background-image: url(../assets/image/background.jpg);
        background-position: center;
        transition: background-position 0.4s ease;
        background-repeat: no-repeat;
    }
    
    h1{
        font-family: "Rubik", sans-serif;
        color: white;
    }

    p{
        font-family: "Rubik", sans-serif;
        color: white;
        padding: 4%;
    }
    
    .botao{
        height: 4em;
        width: 10em;
        background-color:white;
        border-radius: 20px;
        border: 6px white solid;
    }

    .titulo{
        height: 100px;
        position: relative;
        width: 50%;
        display: flex;
        flex-direction: column;
        padding: 5% 0 0 5%;
    }

    .container{
        width: 100%;
        position: absolute;
        bottom: 0;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: end;
    }
</style>
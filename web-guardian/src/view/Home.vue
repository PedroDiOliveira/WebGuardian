<script>

</script>

<template >
    <div class="home">
        <div class="titulo">
            <h1>Bem vindo ao WebGuadian</h1>
        </div>
        <div class="container">
            <p>Iniciar monitoramento</p>
            <RouterLink ><button class="botao"></button></RouterLink>
        </div>
    </div>
</template>

<style>
    .home{
        background-image: url(../assets/image/background.jpg);
        background-size: cover;
        background-position: center;
        height: 100%;
        padding-block: 10%;

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
        height: 5em;
        width: 5em;
        background-color:white;
        border-radius: 50%;
        border: 6px white gray;
    }

    .titulo{
        height: 50%;
        
        width: 100%;
        display: flex;
        flex-direction: column;
    }

    .container{
        height: 100%;
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: end;
    }
</style>
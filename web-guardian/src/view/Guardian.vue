<script setup>
const formulario = document.getElementById("form")

const data = ""
const options = {
        method: "POST",
        headers:{
            'Content-Type':'application/json'
        },
        body: JSON.stringify(data)
    }

const RunService = () =>{
    fetch("127.0.0.1:8080",options).then(response => data)
}

    
    
</script>

<template>
    <div class="container">
        <form id="form" @submit.prevent="RunService">
            <label for="url"><strong>URL:</strong></label>
            <input v-model="data" id="url" type="text">
            <input type="submit">
        </form>
    </div>
</template>

<style scoped>
    .container{
        height: 100vh;
        width: 100vw;
        background-color: gray;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    form{
        background-color: black;
        height: 400px;
        width:400px ;
        border-radius: 20px;
        padding: 10px;
        display: flex;
        flex-direction: column;
        align-items: center
    }

    label{
        color: white;
        font-size: larger;
    }

    input{
        height: 30px;
        width: 320px;
    }

</style>
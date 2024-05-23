<template>
    <div class="container">
        <form class="form" @submit.prevent="handleSubmit    ">
            <label for="url"><strong>URL:</strong></label>
            <input name="url" id="url" type="text">
            <input type="submit" value="Submit">
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

<script setup>

const handleSubmit = () => {
  const formEl = document.querySelector(".form");
  const formData = new FormData(formEl);
  console.log(formData.get('url'));
  const data = Object.fromEntries(formData);
  fetch("http://127.0.0.1:8080/check", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  }).then(response => response.json())
    .then(result => {
      console.log(result);
    })
    .catch(error => {
      console.error('Erro:', error);
    });
}
</script>

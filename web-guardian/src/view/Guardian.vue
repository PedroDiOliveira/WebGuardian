<template>
<div class="container">
    <div class="body">
      <div class="titulo">
        <RouterLink class="router" to="/"><h1>Guardian</h1></RouterLink>
      </div>
        <div class="content">
          <form class="form" @submit.prevent="handleSubmit">
            <label for="url">Digite a URL desejada:</label>
            <input name="url" id="url" type="text">
            <input type="submit" value="Submit">
          </form>
          <div class="response">
            <div id="code"></div>
            <div id="desc"></div>
          </div>
        </div>
    </div>
  </div>
</template>

<style scoped>
    .container{
        height: 100vh;
        width: 100vw;
        background-color: white;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .content{
      display: flex;
      flex-direction: column;
      gap: 20px;
    }

    .body{
      background-color: gray;
      border-radius: 20px;
      height: 85%;
      width: 35%;
      display: flex;
      flex-direction: column ;
      justify-content: space-between;
      align-items: center;
    }

    p{
      width: 90px;
    }

    .response{
      display: flex;
      justify-content: space-between;
      width: 100%;
      padding-bottom: 20%;
    }

    form{
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 20px;
    }

    .titulo{
      display: flex;
      align-items: center;
    }

    label{
        color: white;
        font-size: larger;
    }

    input{
        height: 40px;
        width: 320px;
    }

    .router{
      text-decoration: none;
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
  }).then(req => req.json())
    .then(response => {
      DisplayElements(response)
    })
    .catch(error => {
      console.error('Erro:', error);
    });
}

function DisplayElements(data){
  const codeDiv = document.getElementById('code');
  const descDiv = document.getElementById('desc');

  //limpa as divs para receber novos dados
  codeDiv.innerHTML = '';
  descDiv.innerHTML = '';

  //codigo
  const code = document.createElement('h4');
  code.textContent = `Code : ${data.code}`;
  codeDiv.appendChild(code)

  //desc
  const desc = document.createElement('p');
  desc.textContent = `Desc : ${data.desc}`;
  descDiv.appendChild(desc)
}
</script>

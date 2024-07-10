<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="container">
        <div class="logo">
            <RouterLink class="router" to="/"><h1>Guardian</h1></RouterLink>
        </div>
        <div class="links">
            <RouterLink class="routerlink" to="/guardian" ><p>Guardian</p></RouterLink>
            <RouterLink class="routerlink" to="/guardian" ><p>Class Room</p></RouterLink>
            <RouterLink class="routerlink" to="/guardian" ><p>Guardian</p></RouterLink>
        </div> 
    </div>
</template>

<style scoped>

    .container{
        display:flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 0% 3%;
        height: 5em;
    }

    .links{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        width: 2%;
        text-decoration: none;
    }

    .router{
        text-decoration: none;
        color: black;
    }

    .routerlink{
        text-decoration: none;
        color: black;
    }


</style>

<scrip>

</scrip>
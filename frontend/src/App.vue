<script setup>
import {ref} from "vue"
import DefaultPage from "@/pages/DefaultPage.vue";
import CustomizePage from "@/pages/CustomizePage.vue";
import DynamicBackground from "@/components/DynamicBackground.vue";


const page = ref("customizePage");
const btnTxt = ref("定制")

function switchPage() {
  btnTxt.value = btnTxt.value === "定制" ? "默认" : "定制"
  page.value = page.value === "defaultPage" ? "customizePage" : "defaultPage";
}



</script>
<template>
  <DynamicBackground/>
  <n-message-provider>
  <div class="container">
    <main class="main">
      <n-button class="floating-button" @click="switchPage"> {{ btnTxt }}</n-button>
      <transition name="flip" mode="out-in">
        <component :is="page === 'customizePage'?DefaultPage:CustomizePage"/>
      </transition>
    </main>
  </div>
</n-message-provider>
</template>

<style scoped>
.container {
  position: relative;
}

.flip-enter-active, .flip-leave-active {
  transition: transform 0.6s, opacity 0.1s;
  transform-style: preserve-3d;
}

.flip-enter, .flip-leave-to {
  transform: rotateY(90deg);
}


.floating-button {
  position: fixed;
  right: 20px;
  bottom: 20px;
  width: 50px;
  height: 50px;
  border-radius: 30%;
  background-color: #007bff;
  color: white;
  border: none;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 10px;
  transition: background-color 0.3s ease;
  float: inherit;
  z-index: 10;
  opacity: 0.6;
}

.floating-button:hover {
  background-color: #0056b3;
}
</style>

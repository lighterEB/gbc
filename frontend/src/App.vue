<script setup>
import {ref} from "vue"
import DefaultPage from "@/pages/DefaultPage.vue";
import CustomizePage from "@/pages/CustomizePage.vue";
let page = ref("defaultPage");
let btnTxt = ref("定制")
function switchPage() {
  page.value = page.value === "defaultPage" ? "customizePage" : "defaultPage";
  btnTxt.value = btnTxt.value === "定制"?"默认":"定制"
}
</script>
<template>
  <div class="container">
    <n-button class="floating-button" @click="switchPage"> {{ btnTxt }} </n-button>
    <default-page v-if="page==='defaultPage'"/>
    <customize-page v-if="page==='customizePage'"/>
  </div>

</template>
<style scoped>
.container{
  position: relative;
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

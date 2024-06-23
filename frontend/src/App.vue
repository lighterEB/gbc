<template>
  <div class="app">
    <div class="page-container">
      <button @click="toggleComponents">切换组件</button>
      <component v-show="currentComponent"></component>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import ComponentA from './components/ComponentA.vue';
import ComponentB from './components/ComponentB.vue';

let currentComponent = ref("ComponentA");

const toggleComponents = () => {
  currentComponent = (currentComponent === "ComponentA" ? "ComponentB" : "ComponentA");
};
</script>

<style scoped>
.page-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>

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
        <n-button class="floating-button" @click="switchPage" :class="{ 'button-active': page === 'defaultPage' }">
          {{ btnTxt }}
        </n-button>
        <div class="page-wrapper">
          <transition name="page-transition" mode="out-in">
            <component :is="page === 'customizePage' ? DefaultPage : CustomizePage"/>
          </transition>
        </div>
      </main>
    </div>
  </n-message-provider>
</template>

<style scoped>
.container {
  position: relative;
  min-height: 100vh;
  overflow-x: hidden;
}

.main {
  position: relative;
  width: 100%;
  min-height: 100vh;
}

.page-wrapper {
  position: relative;
  width: 100%;
  min-height: 100vh;
  overflow: hidden;
}

/* 页面切换动画 */
.page-transition-enter-active,
.page-transition-leave-active {
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  position: absolute;
  width: 100%;
  left: 0;
  right: 0;
}

.page-transition-enter-from {
  opacity: 0;
  transform: translateX(30px) scale(0.95);
}

.page-transition-leave-to {
  opacity: 0;
  transform: translateX(-30px) scale(0.95);
}

.page-transition-enter-to,
.page-transition-leave-from {
  opacity: 1;
  transform: translateX(0) scale(1);
}

/* 浮动按钮样式 */
.floating-button {
  position: fixed;
  right: 20px;
  bottom: 20px;
  width: 50px;
  height: 50px;
  border-radius: 25px;
  background-color: rgba(0, 123, 255, 0.8);
  color: white;
  border: none;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 0.9em;
  font-weight: 500;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 10;
  backdrop-filter: blur(4px);
}

.floating-button:hover {
  background-color: rgba(0, 86, 179, 0.9);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.2);
}

.floating-button:active {
  transform: translateY(0);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.button-active {
  background-color: rgba(220, 53, 69, 0.8);
}

.button-active:hover {
  background-color: rgba(200, 35, 51, 0.9);
}

:global(body) {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}
</style>

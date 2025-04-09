<script setup>
import {onMounted, ref} from 'vue';
import {NGi, NGrid} from 'naive-ui';
import Cards from "@/components/Cards.vue";
import {QueryData} from "../../wailsjs/go/main/App"

let productsList = ref([])
function queryData() {
  QueryData().then((res) => {
    if (res.code == 200) {
      productsList.value = res.data===undefined ? [] : res.data["products"]
    }
  })
}

const dealt = ref(false);
const isHovered = ref(Array(productsList.value.length).fill(false));

const getCardStyle = (index) => {
  return {
    animationDelay: dealt.value ? `${index * 0.1}s` : '0s'
  };
};

const hoverStyle = (hovered) => ({
      // transform: hovered ? 'translateY(-10px)' : '',
      bottom: hovered?'5%':'0px',
      position: 'relative',
      // transitionDelay: hovered ? `0.2s` : `${(index + 1) * 0.1}s`
    }
)

onMounted(() => {
  requestAnimationFrame(() => {
    dealt.value = true;
  });
  queryData();
});

</script>

<template>
  <div class="grid-container">
    <n-grid :x-gap="32" :y-gap="32" cols="4" responsive="screen">
      <n-gi v-for="(item, index) in productsList" :key="index"
            class="grid-item"
            :class="{'dealt': dealt}"
            :style="getCardStyle(index)"
      >
        <Cards :item="item"/>
      </n-gi>
    </n-grid>
  </div>
</template>

<style scoped>
.grid-container {
  width: 90%;
  max-width: 1200px;
  margin: 2rem auto;
  position: relative;
}

@keyframes dealCard {
  0% {
    opacity: 0;
    transform: translateY(40px) scale(0.95);
  }
  60% {
    opacity: 1;
    transform: translateY(-5px) scale(1);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.grid-item {
  opacity: 0;
  transform: translateY(40px) scale(0.95);
  will-change: transform, opacity;
}

.grid-item.dealt {
  animation: dealCard 0.8s cubic-bezier(0.34, 1.56, 0.64, 1) forwards;
}

/* 防止动画过程中出现闪烁 */
.grid-item, .grid-item * {
  -webkit-transform-style: preserve-3d;
  transform-style: preserve-3d;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
}

@media (max-width: 1200px) {
  .grid-container {
    width: 95%;
  }
}

@media (max-width: 768px) {
  .grid-container {
    width: 100%;
    padding: 0 1rem;
  }
}
</style>
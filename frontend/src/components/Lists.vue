<script setup>
import {onMounted, ref} from 'vue';
import {NGi, NGrid} from 'naive-ui';
import Cards from "@/components/Cards.vue";
import {QueryData} from "../../wailsjs/go/main/App"

let productsList = ref([])
function queryData() {
  QueryData().then((res) => {
    console.log(res)
    if (res.code == 200) {
      productsList.value = res.data===undefined ? [] : res.data["products"]
    }
  })
}

const dealt = ref(false);
const isHovered = ref(Array(productsList.value.length).fill(false));
const dealCards = () => {
  dealt.value = false; // Reset the dealt state
  setTimeout(() => {
    dealt.value = true; // Trigger the animation
  }, 100); // Small delay to ensure reset
};

const getCardStyle = (index) => {
  const delay = (index + 1) * 0.1
  return dealt.value ? {transitionDelay: `${delay}s`} : {}
};

const hoverStyle = (hovered) => ({
      // transform: hovered ? 'translateY(-10px)' : '',
      bottom: hovered?'5%':'0px',
      position: 'relative',
      // transitionDelay: hovered ? `0.2s` : `${(index + 1) * 0.1}s`
    }
)

onMounted(() => {
  dealCards();
  queryData();
});

</script>

<template>
  <n-grid :x-gap="30" :y-gap="30" cols="4" style="width: 90%; margin-left: 5%; margin-top: 5%">
    <n-gi v-for="(item, index) in productsList" :key="index"
          :class="['grid-item', {'dealt': dealt}]" :style="[getCardStyle(index), hoverStyle(isHovered[index])]"
          @mouseenter="isHovered[index] = true"
          @mouseleave="isHovered[index] = false"
    >
      <Cards :item="item"/>
    </n-gi>
  </n-grid>
</template>

<style scoped>
.grid-item {
  width: 100%;
  height: 200px;
  transform: perspective(300px) rotateX(180deg);
  opacity: 0;
  transition: transform 0.8s ease, opacity 1s ease;
}

.grid-item.dealt {
  transForm: translateY(0);
  opacity: 1;
}
</style>
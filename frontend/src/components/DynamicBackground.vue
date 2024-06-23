<template>
  <canvas ref="canvas" class="background-canvas"></canvas>
</template>

<script setup>
import {ref, onMounted, onUnmounted} from 'vue';

const canvas = ref(null);
const stars = ref([]);

const createStars = (numStars, width, height) => {
  const newStars = [];
  for (let i = 0; i < numStars; i++) {
    newStars.push({
      x: Math.random() * width,
      y: Math.random() * height,
      radius: Math.random() * 1.5,
      dx: (Math.random() - 0.5) * 0.5, // 水平速度
      dy: (Math.random() - 0.5) * 0.5, // 垂直速度
    });
  }
  return newStars;
};

const drawStars = (ctx, width, height) => {
  ctx.clearRect(0, 0, width, height);
  stars.value.forEach(star => {
    ctx.beginPath();
    ctx.arc(star.x, star.y, star.radius, 0, Math.PI * 2);
    ctx.fillStyle = 'white';
    ctx.fill();
  });
};

const updateStars = (width, height) => {
  stars.value.forEach(star => {
    star.x += star.dx;
    star.y += star.dy;

    // 如果星星超出了边界，则将其移到另一边
    if (star.x > width) star.x = 0;
    if (star.x < 0) star.x = width;
    if (star.y > height) star.y = 0;
    if (star.y < 0) star.y = height;
  });
};

const animate = (ctx, width, height) => {
  drawStars(ctx, width, height);
  updateStars(width, height);
  requestAnimationFrame(() => animate(ctx, width, height));
};

onMounted(() => {
  const canvasEl = canvas.value;
  const ctx = canvasEl.getContext('2d');
  const resizeCanvas = () => {
    canvasEl.width = window.innerWidth;
    canvasEl.height = window.innerHeight;
    stars.value = createStars(100, canvasEl.width, canvasEl.height);
  };
  resizeCanvas();
  window.addEventListener('resize', resizeCanvas);
  animate(ctx, canvasEl.width, canvasEl.height);

  onUnmounted(() => {
    window.removeEventListener('resize', resizeCanvas);
  });
});
</script>

<style scoped>
.background-canvas {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1; /* 使背景在所有内容之下 */
  background-color: black; /* 背景颜色 */
}
</style>

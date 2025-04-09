<script setup lang="ts">
import { useMessage, NAlert } from 'naive-ui';
import { ref, h, reactive, computed } from 'vue'
import { KeyGenRequest } from '../../wailsjs/go/main/App'

// Types
interface Product {
  code: string;
  fallbackDate: string;
  paidUpTo: string;
  extended: string;
}

interface LicenseParams {
  licenseId: string;
  licenseeName: string;
  assigneeName: string;
  assigneeEmail: string;
  products: Product[];
}

// Constants
const PRODUCT_CODES = {
  CLion: 'CL',
  DataGrip: 'DG',
  GoLand: 'GO',
  IDEA: 'II',
  PyCharm: 'PC',
  Rider: 'RD',
  RustRover: 'RR',
  WebStorm: 'WS'
} as const;

const props = defineProps({
  item: {
    type: Object,
    required: true
  }
})

const title = ref(props.item["name"]);
const pic = ref(props.item["pic"]);
const isHovered = ref(false);
const text = ref("****************")
const msgInfo = useMessage()
const isLoading = ref(false);

const params = reactive<LicenseParams>({
  licenseId: "DADDYHU726",
  licenseeName: "lighterEB",
  assigneeName: "daddyhu",
  assigneeEmail: "hellohus726@126.com",
  products: [
    { code: "PCWMP", fallbackDate: "2099-12-31", paidUpTo: "2099-12-31", extended: "true" },
    { code: "PSI", fallbackDate: "2099-12-31", paidUpTo: "2099-12-31", extended: "true" }
  ]
})

const productTemplate: Product = {
  code: "",
  fallbackDate: "2099-12-31",
  paidUpTo: "2099-12-31",
  extended: "false"
}

const getProductCode = computed(() => {
  return PRODUCT_CODES[props.item["name"]] || '';
});

const rendMessage = (props) => {
  const { type } = props;
  return h(
    NAlert,
    {
      closable: props.closable,
      onclose: props.onclose,
      type: type === "reloading" ? "default" : type,
      title: "获取激活码",
      style: {
        opacity: "0.8",
        backgroundColor: "pink",
      },
    },
    {
      default: () => props.content
    }
  );
};

async function copyCode() {
  try {
    isLoading.value = true;
    const newProduct = { ...productTemplate, code: getProductCode.value };
    params.products.push(newProduct);

    const response = await KeyGenRequest(JSON.stringify(params));
    
    if (response.code === 200) {
      await navigator.clipboard.writeText(response.data);
      msgInfo.success("你已经获取激活码，现在可以直接粘贴啦！", {
        render: rendMessage,
        closable: true,
        duration: 1000
      });
    } else {
      throw new Error('激活码生成失败');
    }
  } catch (error) {
    msgInfo.error("糟糕~激活码丢了~~！", {
      render: rendMessage,
      closable: true,
      duration: 1000
    });
  } finally {
    isLoading.value = false;
    // Remove the added product
    params.products.pop();
  }
}
</script>

<template>
  <div class="card">
    <div class="card-header">{{ title }}</div>
    <div class="card-content">
      <n-avatar :size="48" class="avatar">
        <div v-html="pic"></div>
      </n-avatar>
      <div 
        class="action-area"
        @mouseenter="isHovered = true" 
        @mouseleave="isHovered = false"
      >
        <transition name="fade-slide" mode="out-in">
          <n-text class="text-area" v-if="!isHovered" key="text">
            {{ text }}
          </n-text>
          <n-button 
            class="btn-copy" 
            text 
            @click="copyCode()" 
            :loading="isLoading"
            :disabled="isLoading"
            v-else
            key="button"
          >
            {{ isLoading ? '生成中...' : '复制' }}
          </n-button>
        </transition>
      </div>
    </div>
  </div>
</template>

<style scoped>
.card {
  border-radius: 1rem;
  height: 160px;
  background-color: var(--primary-color, #9a49e5);
  display: flex;
  flex-direction: column;
  position: relative;
  isolation: isolate;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  transition: all 0.3s cubic-bezier(0.23, 1, 0.32, 1);
}

.card::before {
  content: '';
  position: absolute;
  inset: 0;
  background-color: var(--primary-color, #9a49e5);
  border-radius: inherit;
  z-index: -1;
}

.card-header {
  padding: 0.75rem 1rem;
  color: white;
  font-weight: 500;
  font-size: 1.1em;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  flex: 1;
  padding: 1rem;
  gap: 0.5rem;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 12px 20px -10px rgba(255, 255, 255, 0.3);
}

.avatar {
  background-color: transparent;
  transition: transform 0.3s cubic-bezier(0.23, 1, 0.32, 1);
}

.card:hover .avatar {
  transform: scale(1.1);
}

.action-area {
  width: 100%;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow: hidden;
}

.text-area {
  width: 100%;
  text-align: center;
  color: rgba(255, 255, 255, 0.8);
  font-family: monospace;
  font-size: 0.85em;
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-copy {
  border: none;
  border-radius: 999px;
  width: 70%;
  height: 32px;
  padding: 0.5rem;
  color: var(--primary-color, #9a49e5);
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(255, 255, 255, 0.95);
  font-size: 0.9em;
  font-weight: 500;
  position: absolute;
  inset: 0;
  margin: auto;
  transition: transform 0.2s cubic-bezier(0.23, 1, 0.32, 1);
}

.btn-copy:hover:not(:disabled) {
  background-color: white;
  transform: scale(1.05);
}

.btn-copy:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

/* 淡入淡出和滑动过渡效果 */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

.fade-slide-enter-to,
.fade-slide-leave-from {
  opacity: 1;
  transform: translateY(0);
}
</style>
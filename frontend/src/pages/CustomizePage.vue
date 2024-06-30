<script setup>
import { ref } from "vue"
const timestamp = ref(null)
</script>

<template>
  <n-flex class="container" vertical>
    <n-list class="info-list">
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">产品Code: </n-thing>
          <n-input placeholder="请输入产品Code"/>
          <n-tooltip trigger="hover">
            <template #trigger>
              <n-tag type="info">
                ？
              </n-tag>
            </template>
              产品Code可以从这里找
          </n-tooltip>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">账号: </n-thing>
          <n-input placeholder="请输入账号" />
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing class="n-thing">邮箱: </n-thing>
          <n-input placeholder="请输入邮箱"/>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space>
          <n-thing>过期时间: </n-thing>
          <n-date-picker v-model:value="timestamp" size="medium" type="date"  placeholder="请选择日期"/>
        </n-space>
      </n-list-item>
      <n-list-item>
        <n-space style="padding-left: 40%;">
          <n-button>生成</n-button>
        </n-space>
      </n-list-item>
    </n-list>
  </n-flex>
</template>

<style scoped>
.container {
  position: relative;
  background-color: #f1f2e5;
  padding: 30px 0 50px 0;
}

.info-list {
  margin: auto;
  width: 30%;
  padding: 10% 300px 260px 300px;
}
.n-thing {
  padding-top: 6px;
}
</style>
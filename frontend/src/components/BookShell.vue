<script setup>
const bookList = [
  {id: 1, name: '《红楼梦》', author: '曹雪芹'},
  {id: 2, name: '《三国演义》', author: '罗贯中'},
  {id: 3, name: '《西游记》', author: '吴承恩'},
  {id: 4, name: '《水浒传》', author: '施耐庵'}
]
</script>

<template>
  <div class="bookshelf">
  </div>
</template>

<style scoped>
.bookshelf {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  color: mediumslateblue;
}
</style>
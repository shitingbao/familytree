<template>
  <div class="tag">
    <a-tooltip placement="bottomLeft">
      <template #title>
        <span>man</span>
      </template>
      <a-button class="btn1 font">M</a-button>
    </a-tooltip>
    <a-tooltip placement="bottomLeft">
      <template #title>
        <span>woman</span>
      </template>
      <a-button class="btn2">W</a-button>
    </a-tooltip>
    <a-tooltip placement="bottomLeft">
      <template #title>
        <span>配偶</span>
      </template>
      <a-button class="btn3">S</a-button>
    </a-tooltip>
    <a-tooltip placement="bottomLeft">
      <template #title>
        <span>本级范围</span>
      </template>
      <a-button class="btn4"> *</a-button>
    </a-tooltip>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.tag {
  display: flex;
}

.btn1 {
  background-color: rgb(83, 128, 235);
}

.btn2 {
  background-color: rgb(253, 171, 48);
}

.btn3 {
  background-color: blueviolet;
}

.btn4 {
  background-color: aqua;
}

.font {
  color: aliceblue;
}
</style>

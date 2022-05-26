<template>
  <h1>{{ msg }}</h1>
  <div class="body">
    <Node class="node" :formState="formState" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from "axios";
import { Member } from "../model/member";
import Node from "./Node.vue";

defineProps<{ msg: string }>();

const formState = ref<Member>(new Member());

function getMember(id: string) {
  console.log("getMember start");
  formState.value = new Member();
  const formData = new FormData();
  formData.append("id", id);
  axios
    .post("http://localhost:6200/v1/member/last", formData)
    .then((response) => {
      formState.value.getHeader(response.data.data);
    });
}

onMounted(() => {
  getMember("6");
});
</script>

<style scoped>
.body {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
}
</style>

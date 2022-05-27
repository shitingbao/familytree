<template>
  <h1>{{ msg }}</h1>
  <div class="body">
    <a-button v-if="isNull" type="primary" @click="createRoot">
      createRoot
    </a-button>
    <Node v-if="!isNull" class="node" :formState="formState" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from "axios";
import { Member } from "../model/member";
import Node from "./Node.vue";
import bus from "../libs/bus";

defineProps<{ msg: string }>();

const formState = ref<Member>(new Member());

const isNull = ref(false);

function reloadMember(): void {
  console.log("getMember start");
  getMember("6");
}

bus.on("reloadMember", reloadMember);

function getMember(id: string) {
  formState.value = new Member();
  const formData = new FormData();
  formData.append("id", id);
  axios
    .post("http://localhost:6200/v1/member/last", formData)
    .then((response) => {
      formState.value.getHeader(response.data.data);
      if (!(response.data.data as []).length) {
        console.log("=========:", !response.data.data.length);
        isNull.value = true;
      }
    });
}

function createRoot() {
  const formData = new FormData();
  formData.append("name", "newRoot");
  axios
    .post("http://localhost:6200/v1/member/create", formData)
    .then((response) => {
      console.log(response);
      isNull.value = false;
    });
}

onMounted(() => {
  getMember("6");
});
</script>

<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>

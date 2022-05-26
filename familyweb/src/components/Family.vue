<template>
  <h1>{{ msg }}</h1>
  <div class="body"><Node class="node" :formState="formState" /></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, reactive } from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
import axios from "axios";
import { Member } from "../model/member";
import Node from "./Node.vue";

defineProps<{ msg: string }>();

const formState = ref<Member>(new Member());

const ll = ref<Member>({ name: "ll" });
const kk = ref<Member>({ name: "kk" });
const jj = ref<Member>({ name: "jj" });
const ii = ref<Member>({ name: "ii", children: [jj.value] });
const hh = ref<Member>({ name: "hh", children: [ii.value] });
const gg = ref<Member>({
  name: "gg",
  children: [hh.value, kk.value, ll.value],
});
const ff = ref<Member>({ name: "ff" });
const ee = ref<Member>({ name: "ee" });

const cc = ref<Member>({ name: "cc" });
const dd = ref<Member>({ name: "dd", children: [gg.value] });
const bb = ref<Member>({ name: "bb", children: [dd.value] });

const aa = ref<Member>({
  name: "aa",
  children: [ff.value, ee.value, bb.value, cc.value],
});

function getMember(id: string = "0") {
  // const formData = new FormData();
  // formData.append("id", id);
  // axios
  //   .post("http://localhost:6200/v1/member/list", formData)
  //   .then((response) => {
  //     formState.value.initData(response.data.data[0]);
  //     console.log(" response.data response.data", response.data);
  //   });
  formState.value = aa.value;
  console.log(formState.value);
}

onMounted(() => {
  getMember("0");
});
</script>

<style scoped>
/* .node {
  display: flex;
  flex-direction: row;
} */
.body {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
}
</style>

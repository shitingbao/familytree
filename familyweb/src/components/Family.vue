<template>
  <h1>族谱</h1>
  <a-button class="save" @click="keepPicture">保存族谱</a-button>
  <div class="body" ref="tupuPicture">
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
import { getCurrentInstance } from "@vue/runtime-core";
import html2canvas from "html2canvas";

const currentInstance = getCurrentInstance();
const formState = ref<Member>(new Member());

const isNull = ref(false);

function reloadMember(): void {
  console.log("getMember start");
  getMember();
}

bus.on("reloadMember", reloadMember);

function getMember(id: string = "") {
  formState.value = new Member();
  const formData = new FormData();
  formData.append("id", id);
  axios
    .post("http://localhost:6200/v1/member/last", formData)
    .then((response) => {
      formState.value.getHeader(response.data.data);
      if (!(response.data.data as []).length) {
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
      getMember();
      console.log(response);
      isNull.value = false;
    });
}

//保存图片
function keepPicture() {
  html2canvas(currentInstance.ctx.$refs.tupuPicture).then((canvas) => {
    // 获取生成的图片的url,并设置为png格式
    const imgUrl = canvas.toDataURL("png");
    let creatDom = document.createElement("a");
    document.body.appendChild(creatDom);
    creatDom.href = imgUrl;
    //设置生成图片的名称
    creatDom.download = "图片名字";
    creatDom.click();
  });
}

onMounted(() => {
  getMember();
});
</script>

<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
.node {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.save {
  position: absolute;
  top: 10%;
  right: 40%;
}
</style>

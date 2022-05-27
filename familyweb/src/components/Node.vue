<template>
  <div class="node">
    <a-button type="primary" @click="showModal">
      <ManOutlined v-if="formState.sex === 1" />
      <WomanOutlined v-else />
      {{ formState.name }}
    </a-button>
    <a-modal v-model:visible="visible" title="成员信息" :footer="null">
      <Edit :row="formState" @visible="editClose" />
    </a-modal>
    <div class="children">
      <Node
        class="node-child"
        v-for="(item, idx) in formState.children"
        :key="idx"
        :formState="item"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, PropType } from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
import axios from "axios";
import { Member } from "../model/member";
import Edit from "./Edit.vue";
import Node from "./Node.vue";
import bus from "../libs/bus";

const props = defineProps({
  formState: {
    type: Object as PropType<Member>,
    default: () => {
      return new Member();
    },
  },
});

const visible = ref<boolean>(false);

const showModal = () => {
  visible.value = true;
};

function editClose() {
  console.log("editClose start");
  visible.value = false;
  bus.emit("reloadMember");
}
</script>

<style scoped>
.node {
  display: flex;
  flex-direction: column;
  width: 100px;
  padding: 10px 10px 0 0px;
}
.children {
  display: flex;
  flex-direction: row;
  width: 100px;
}
</style>

<template>
  <div class="node">
    <div class="node-content">
      <a-button
        :class="formState.sex == '男' ? 'btn1' : 'btn2'"
        type="primary"
        @click="showModal(1)"
      >
        <ManOutlined v-if="formState.sex === '男'" />
        <WomanOutlined v-else />
        {{ formState.name }}
      </a-button>

      <!-- 未赋值的不显示，标示无配偶 -->
      <a-button
        v-if="formState.marryMember?.parentId == -1"
        class="btnM"
        type="primary"
        @click="showModal(2)"
      >
        <ManOutlined v-if="formState.marryMember.sex === '男'" />
        <WomanOutlined v-else />
        {{ formState.marryMember.name }}
      </a-button>
    </div>

    <a-modal v-model:visible="visible" title="成员信息" :footer="null">
      <Edit :row="selectMember" @visible="editClose" />
    </a-modal>

    <div class="children">
      <Node
        class="node-child"
        v-for="(item, idx) in formState.memberChildren"
        :key="idx"
        :formState="item"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, PropType } from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
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

const selectMember = ref<Member>(props.formState);

function showModal(flag: number = 1) {
  if (flag == 1) {
    selectMember.value = props.formState;
  } else {
    selectMember.value = props.formState.marryMember;
  }
  visible.value = true;
}

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
  padding: 10px 10px 0 0px;
}
.children {
  display: flex;
  flex-direction: row;
}
.btn1 {
  background-color: rgb(83, 128, 235);
}

.btn2 {
  background-color: rgb(253, 171, 48);
}

.btnM {
  background-color: blueviolet;
}

.node-content {
  background-color: aqua;
}
</style>

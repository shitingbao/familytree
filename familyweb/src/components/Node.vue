<template>
  <div class="node">
    <!-- {{ formState }}
    {{ formState.name }} -->
    <a-button type="primary" @click="showModal">
      <ManOutlined v-if="formState.sex === 1" />
      <WomanOutlined v-else />
      {{ formState.name }}
    </a-button>
    <a-modal v-model:visible="visible" title="成员信息" @ok="handleOk">
      <Edit :row="formState" />
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

const props = defineProps({
  formState: {
    type: Object as PropType<Member>,
    default: () => {
      return new Member();
    },
  },
});

// const formState = ref<Member>(new Member());

const visible = ref<boolean>(false);

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 12 },
};

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "${label} is not a valid email!",
    number: "${label} is not a valid number!",
  },
  number: {
    range: "${label} must be between ${min} and ${max}",
  },
};

const onFinish = (values: any) => {
  console.log("Success:", values);
};

const showModal = () => {
  visible.value = true;
};

const handleOk = (e: MouseEvent) => {
  console.log(e);
  visible.value = false;
};
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

<template>
  <div>
    <h1>家族列表</h1>
    <a-button
      class="editable-add-btn"
      style="margin-bottom: 8px"
      @click="handleAdd"
    >
      添加一个新家族
    </a-button>
    <a-table
      class="list"
      bordered
      :data-source="dataSource"
      :columns="columns"
      :pagination="{ pageSize: 5 }"
    >
      <template #bodyCell="{ column, text, record }">
        <template v-if="column.dataIndex === 'name'">
          <div class="editable-cell">
            <div class="editable-cell-text-wrapper">
              {{ text || " " }}
            </div>
          </div>
        </template>
        <template v-else-if="column.dataIndex === 'operation'">
          <span class="right">
            <a @click="onDelete(record.memberId)">Delete</a>
          </span>
          <span>
            <a @click="showMemberList(record.memberId)">show</a>
          </span>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup lang="ts">
import {
  ref,
  onMounted,
  watch,
  PropType,
  Ref,
  UnwrapRef,
  computed,
  reactive,
} from "vue";
import { ManOutlined, WomanOutlined } from "@ant-design/icons-vue";
import axios from "axios";
import { Member } from "../model/member";
import Edit from "./Edit.vue";
import Node from "./Node.vue";
import bus from "../libs/bus";
import { CheckOutlined, EditOutlined } from "@ant-design/icons-vue";
import { cloneDeep } from "lodash-es";
import { SHOW_ALL } from "ant-design-vue/lib/vc-tree-select/utils/strategyUtil";

interface DataItem {
  memberId: string;
  name: string;
}

const props = defineProps({
  formState: {
    type: Object as PropType<Member>,
    default: () => {
      return new Member();
    },
  },
});

bus.on("reloadList", MemberList);

const columns = [
  {
    title: "last",
    dataIndex: "name",
    width: "30%",
  },
  {
    title: "operation",
    dataIndex: "operation",
  },
];

const dataSource: Ref<DataItem[]> = ref([]);
const editableData: UnwrapRef<Record<string, DataItem>> = reactive({});

function showMemberList(memberId: string) {
  bus.emit("reloadMember", memberId);
}

function onDelete(memberId: string) {
  const formData = new FormData();
  formData.append("id", memberId);

  axios
    .post("http://localhost:6200/v1/member/delete", formData)
    .then((response) => {
      console.log(response);
      MemberList();
      bus.emit("reloadMember", ""); // 必须传一个空，为了刷新列表
    });
}

function handleAdd() {
  const formData = new FormData();
  formData.append("name", "newRoot");
  axios
    .post("http://localhost:6200/v1/member/create", formData)
    .then((response) => {
      MemberList();
      console.log(response);
    });
}

function MemberList() {
  dataSource.value = [];
  const formData = new FormData();
  formData.append("id", "0");
  axios
    .post("http://localhost:6200/v1/member/list", formData)
    .then((response) => {
      response.data.data.forEach((e: any) => {
        dataSource.value.push({
          memberId: e.id,
          name: e.name,
        });
      });
      console.log("list:", response.data);
    });
}

onMounted(() => {
  MemberList();
});
</script>

<style scoped>
.right {
  padding-right: 10px;
}
</style>

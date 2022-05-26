<template>
  <modal :show="showModal" @close="onCloseModal()">
    <div class="table">
      <div class="row">
        <label>性别:</label>
        <div class="radioGroup">
          <label
            ><input
              type="radio"
              name="sex"
              value="1"
              v-model="currentPerson.sex"
            />
            男
          </label>
          <label>
            <input
              type="radio"
              name="sex"
              value="2"
              v-model="currentPerson.sex"
            />
            女
          </label>
        </div>
      </div>
      <div class="row">
        <label for="name">姓名:</label>
        <input type="text" id="name" v-model="currentPerson.name" />
      </div>

      <div class="row">
        <label for="dateBirth">出生日期:</label>
        <input type="date" id="dateBirth" v-model="currentPerson.dateBirth" />
      </div>
      <div class="row">
        <label for="dateMarry">结婚日期:</label>
        <input type="date" id="dateMarry" v-model="currentPerson.dateMarry" />
      </div>
      <div class="row">
        <label for="placeBirth">出生地:</label>
        <input type="text" id="placeBirth" v-model="currentPerson.placeBirth" />
      </div>
      <div class="row">
        <label for="dateDeath">死亡时间:</label>
        <input type="date" id="dateDeath" v-model="currentPerson.dateDeath" />
      </div>
      <div class="row">
        <label for="placeDeath">死亡地点:</label>
        <input type="text" id="placeDeath" v-model="currentPerson.placeDeath" />
      </div>
      <div class="row">
        <label for="content">生平简介:</label>
        <input type="text" id="content" v-model="currentPerson.content" />
      </div>
      <div class="row">
        <label for="honor">荣誉:</label>
        <input type="text" id="honor" v-model="currentPerson.honor" />
      </div>
    </div>
    <div class="buttonrow">
      <button @click="onAddChild()">add child</button>
      <button @click="onCloseModal()">close</button>
    </div>
  </modal>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from "vue";

import { IPerson, newPerson } from "../misc/Person";
import { EditMode } from "../misc/Utils";

import Modal from "./Modal.vue";

const EMIT_CLOSE = "close";

export default defineComponent({
  name: "EditPerson",
  components: {
    Modal,
  },
  props: {
    person: {
      type: Object as PropType<IPerson>,
      required: true,
    },
    mode: {
      type: String as PropType<EditMode>,
      required: true,
    },
  },
  emits: [EMIT_CLOSE],
  setup: (props, { emit }) => {
    const base64 = ref("");
    const showModal = ref(true);
    const currentPerson = ref<IPerson>(props.person);
    const editMode = ref(props.mode);

    // 照片选择
    // function onFileChange(event: Event): void {
    //   const fileInput: HTMLInputElement = event.target as HTMLInputElement;
    //   if (fileInput.files && fileInput.files.length > 0) {
    //     const file = fileInput.files.item(0);
    //     if (file) {
    //       currentPerson.value.portrait = {
    //         fileName: file.name,
    //         mimeType: file.type,
    //       };
    //       base64.value = URL.createObjectURL(file);
    //     }
    //     fileInput.value = "";
    //   }
    // }

    function onLoad(): void {
      URL.revokeObjectURL(base64.value);
    }

    function onAddChild(): void {
      const child = newPerson();
      currentPerson.value?.children.push(child);
      currentPerson.value = child;
      editMode.value = EditMode.CHILD;
      base64.value = "";
    }

    function onCloseModal(): void {
      showModal.value = false;
      emit(EMIT_CLOSE);
    }
    return {
      base64,
      currentPerson,
      EditMode,
      editMode,
      showModal,
      // onFileChange,
      onLoad,
      onAddChild,
      onCloseModal,
    };
  },
});
</script>

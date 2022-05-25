<template>
  <modal :show="showModal" @close="onCloseModal()">
    <div class="table">
      <div class="row">
        <label>性别:</label>
        <div class="radioGroup">
          <label
            ><input
              type="radio"
              name="gender"
              value="male"
              v-model="currentPerson.gender"
            />
            男
          </label>
          <label>
            <input
              type="radio"
              name="gender"
              value="female"
              v-model="currentPerson.gender"
            />
            女
          </label>
        </div>
      </div>
      <div class="row">
        <label for="firstName">姓名:</label>
        <input type="text" id="firstName" v-model="currentPerson.firstName" />
      </div>

      <div class="row">
        <label for="dateOfBirth">出生日期:</label>
        <input
          type="date"
          id="dateOfBirth"
          v-model="currentPerson.dateOfBirth"
        />
      </div>
      <div class="row">
        <label for="dateOfBirth">结婚日期:</label>
        <input
          type="date"
          id="dateOfBirth"
          v-model="currentPerson.dateOfBirth"
        />
      </div>
      <div class="row">
        <label for="placeOfBirth">出生地:</label>
        <input
          type="text"
          id="placeOfBirth"
          v-model="currentPerson.placeOfBirth"
        />
      </div>
      <div class="row">
        <label for="dateOfDeath">死亡时间:</label>
        <input
          type="date"
          id="dateOfDeath"
          v-model="currentPerson.dateOfDeath"
        />
      </div>
      <div class="row">
        <label for="placeOfDeath">死亡地点:</label>
        <input
          type="text"
          id="placeOfDeath"
          v-model="currentPerson.placeOfDeath"
        />
      </div>
      <div class="row">
        <label for="placeOfDeath">生平简介:</label>
        <input
          type="text"
          id="placeOfDeath"
          v-model="currentPerson.placeOfDeath"
        />
      </div>
      <div class="row">
        <label for="placeOfDeath">荣誉:</label>
        <input
          type="text"
          id="placeOfDeath"
          v-model="currentPerson.placeOfDeath"
        />
      </div>
    </div>
    <div class="buttonrow">
      <button
        v-if="
          (editMode === EditMode.PERSON || editMode === EditMode.CHILD) &&
          !currentPerson.spouse
        "
        @click="onAddPartner()"
      >
        add partner
      </button>
      <button v-if="editMode === EditMode.PARENT" @click="onAddChild()">
        add child
      </button>
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

    function onAddName(): void {
      currentPerson.value.middleNames.push("");
    }
    function onRemoveName(index: number): void {
      currentPerson.value.middleNames.splice(index, 1);
    }

    function onFileChange(event: Event): void {
      const fileInput: HTMLInputElement = event.target as HTMLInputElement;
      if (fileInput.files && fileInput.files.length > 0) {
        const file = fileInput.files.item(0);
        if (file) {
          currentPerson.value.portrait = {
            fileName: file.name,
            mimeType: file.type,
          };
          base64.value = URL.createObjectURL(file);
        }
        fileInput.value = "";
      }
    }
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
    function onAddPartner(): void {
      if (currentPerson.value) {
        const partner = newPerson();
        if (currentPerson.value.parent) {
          currentPerson.value.spouse = partner;
          editMode.value = EditMode.SPOUSE;
        } else {
          currentPerson.value.parent = partner;
          editMode.value = EditMode.PARENT;
        }
        currentPerson.value = partner;
        base64.value = "";
      }
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
      onAddName,
      onRemoveName,
      onFileChange,
      onLoad,
      onAddChild,
      onAddPartner,
      onCloseModal,
    };
  },
});
</script>

<script setup>

const props = defineProps({
  name: String,
  modelValue: Boolean
})
const emit = defineEmits(['update:modelValue'])

</script>

<template>
  <div class="switcher">
    <span>
      <slot></slot>
    </span>
    <input :value="props.modelValue" type="checkbox" :id="props.name" @change="emit('update:modelValue', !props.modelValue)"/>
    <label :for="props.name"></label>
  </div>
</template>

<style lang="scss" scoped>
.switcher {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 5px;
  & span {
    color: var(--gray-900, #111928);
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 125%; /* 17.5px */
  }
}
input[type=checkbox]{
  height: 0;
  width: 0;
  visibility: hidden;
}

label {
  cursor: pointer;
  text-indent: -9999px;
  width: 44px;
  height: 24px;
  background: grey;
  display: block;
  border-radius: 100px;
  position: relative;
}

label:after {
  content: '';
  position: absolute;
  top: 4px;
  left: 5px;
  width: 16px;
  height: 16px;
  background: #fff;
  border-radius: 90px;
  transition: 0.3s;
}

input:checked + label {
  background: #003791;
}

input:checked + label:after {
  left: calc(100% - 5px);
  transform: translateX(-100%);
}

label:active:after {
  width: 20px;
}

[dark=true] {
    input:checked + label {
        background: #3F83F8;
    }
    .switcher {
      & span {
        color: var(--gray-900, #FFF);
      }
    }
}

</style>
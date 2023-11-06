<script setup>
import 'v-calendar/style.css';
import { ref } from 'vue';
import { useUser } from '@/stores/user'
import { DatePicker } from 'v-calendar';

const userStore = useUser();

const props = defineProps({
    dates: {
        type: Array,
        required: false
    }
})

const date = new Date();

const rules = ref({
  hours: 0,
  minutes: 0,
  seconds: 0,
  milliseconds: 0,
});

const attributes = ref([
  {
    key: 'today',
    highlight: {
      color: 'blue',
      fillMode: 'outline',
    },
    dates: [date.value]
  },
  {
    dot: true, 
    dates: [props.dates]
  },
]);


</script>

<template>
    <DatePicker
        v-model="date"
        :is-dark="userStore.theme"
        :attributes="attributes"
        :rules="rules"
        @dayclick="$emit('datePick', date)"
    />
</template>

<style lang="scss" scoped>

</style>
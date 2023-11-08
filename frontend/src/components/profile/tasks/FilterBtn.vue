<script setup>
import TheButton from '@/components/layouts/TheButton.vue';
const props = defineProps({
    btnInfo: {
        type: Object,
        required: true
    },
    filterActiveStatus: String,
})


</script>

<template>
    <div class="tasks__header-item">
        <the-button
            :styles="['btn_grey-lite']"
            :type="'button'"
            :class="{active: props.filterActiveStatus === props.btnInfo.type}"
            @click="$emit('update:filterActiveStatus', props.btnInfo.type)"
        >
            <!-- Текущие -->
            {{ btnInfo.text }}

        </the-button>
    </div>
</template>

<style lang="scss" scoped>
.tasks__header-item {
    & .btn {
        &.active {
            color: var(--primary-2-default, #003791);
            border: 1px solid var(--primary-2-default, #003791);
            &:hover {
                background: var(--var-blue);
                color: #fff !important;
            }
        }
    }
}

[dark=true] {
    .tasks__header-item {
        & .btn {
            color: var(--gray-300, #D1D5DB);
            border: 1px solid var(--gray-700, #374151);
            &.active {
                border: 1px solid var(--var-blue) !important;
                color: var(--var-blue) !important;
                &:hover {
                    background: var(--var-blue);
                    color: #fff !important;
                }
            }
        }
    }
}
</style>
<script setup>
import { computed } from 'vue'
const props = defineProps({
    first_name: {
        type: String,
        required: false
    },
    last_name: {
        type: String,
        required: false
    }
})
const initialsUser = computed(() => {
    if(props.first_name && props.last_name) {
        return props.first_name[0] + props.last_name[0]
    }
    return ''
})

</script>

<template>
    <div class="avatar">
        {{ initialsUser }}
    </div>
</template>

<style lang="scss" scoped>
.avatar {
    border-radius: 100%;
    background: var(--gray-100, #F3F4F6);
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    min-width: 32px;
    min-height: 32px;
    color: var(--gray-900, #111928);
    font-family: Inter;
    font-size: 12px;
    font-style: normal;
    font-weight: 500;
    line-height: 12px; /* 100% */
    transition: .2s;
    width: fit-content;
}

[dark=true] {
    & .avatar {
        background: #374151;
        color: #F9FAFB;
    }
}
</style>
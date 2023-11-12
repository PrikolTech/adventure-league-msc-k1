<script setup>
import { ref } from "vue";
import TheStep from "./TheStep.vue";
const props = defineProps({
    course: {
        type: Object,
        required: true
    }
})

const steps = ref([
    {
        text: 'Личный кабинет',
        active: true
    },
    {
        text: 'Рассмотрение',
        active: false
    },
    {
        text: 'Решение',
        active: false
    },
])

let photoNumber = ref(getRandomNumberInRange(0,3))

function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

</script>

<template>
    <router-link :to="{ path: `/courses/${props.course.title}` }" class="me__courses-item">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
        <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
        <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        <div class="me__courses-item-body">
            <div class="me__courses-item-header">
                <div class="me__courses-item-info">
                    <p>
                        {{ props.course.date }}
                    </p>
                    <p>
                        {{ props.course.format }}
                    </p>
                </div>
                <p>
                    {{ props.course.title }}
                </p>
                <span class="status">
                    {{ props.course.status }}
                </span>
                <div class="steps">
                    <the-step
                        v-for="(step, index) of steps" :key="index" :step="step"
                    />
                </div>
            </div>
        </div>
    </router-link>
</template>

<style lang="scss" scoped>
.me__courses-item {
    padding-left: 40%;
    overflow: visible;
    display: block;
    @media (max-width: 1023px) {
        position: relative;
    }
    @media (max-width: 767px) {
        padding-left: 30%;
    }
}
.me__courses-item .pic {
    max-width: 50%;
}

.me__courses-item-info {
    display: flex;
    flex-wrap: wrap;
    gap: 12px;
    margin-bottom: 20px;
    & p {
        color: var(--gray-600, #4B5563);
        font-family: Roboto;
        font-size: 12px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 18px */
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 4px 8px;
        border-radius: var(--rounded-lg, 8px);
        background: var(--gray-200, #E5E7EB);
        transition: .2s;
    }
}

.me__courses-item-content {
    gap: 15px;
}
.me__courses-data {
    font-size: 14px;
}
.me__courses-progress-list {
    display: flex;
    gap: 6px;
    flex-wrap: wrap;
}

.status {
    color: var(--gray-400, #9CA3AF);
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%; /* 21px */
    margin-bottom: 15px;
    margin-top: 15px;
    display: block;
}
.steps {
    display: flex;
    align-items: center;
    gap: 5px;
    @media (max-width: 539px) {
        flex-direction: column;
        align-items: flex-start;
    }
    &__item {

    }

    &__line {

    }
}


[dark=true] {
    .me__courses-item-info {
        & p {
            color: var(--gray-200, #E5E7EB);
            background: var(--gray-500, #6B7280);
        }
    }
    & .me__courses-progress-item {
        background: var(--gray-300, #6B7280);
    }
}
</style>
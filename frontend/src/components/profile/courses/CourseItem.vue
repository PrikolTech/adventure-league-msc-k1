<script setup>
import { ref } from "vue";
import CourseTask from "./CourseTask.vue";
const props = defineProps({
    course: {
        type: Object,
        required: true
    }
})

let photoNumber = ref(getRandomNumberInRange(0,3))

function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}



</script>

<template>
    <router-link to="/courses/1" class="me__courses-item">
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
            </div>
            <div class="me__courses-progress-list">
                <course-task
                    v-for="(task, index) of [1,2,3,4,5,6,7,8,9,10,11,12,13,]" :key="index"
                />
            </div>
            <div class="me__courses-item-content">
                <div class="me__courses-data">
                    <span>
                        {{ props.course.lecture.passed }} / {{ props.course.lecture.total }} лекций
                    </span>
                </div>
                <div class="me__courses-data">
                    <span>
                        {{ props.course.tasks.passed }} / {{ props.course.tasks.total }} парктических заданий
                    </span>
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
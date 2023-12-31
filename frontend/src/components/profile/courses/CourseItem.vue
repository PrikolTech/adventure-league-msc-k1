<script setup>
import { onMounted, ref } from "vue";
import CourseTask from "./CourseTask.vue";
import { useUser } from '@/stores/user' 
import router from "../../../router";

const userStore = useUser()
const props = defineProps({
    course: {
        type: Object,
        required: true
    }
})

//дял фоток
let photoNumber = ref(getRandomNumberInRange(0,3))
function getRandomNumberInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

const lectures = ref([])
const getCourseInfo = async () => {
    let url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${props.course.id}/lectures`
    if(userStore.checkRole('student')) {
        url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${props.course.id}/lectures?user_id=${userStore.user.id}`
    }
    try {
        const response = await fetch(url, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()

        lectures.value = [...data]
    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    getCourseInfo()
})
</script>

<template>
    <router-link :to="{ path: `/courses/${props.course.id}` }" class="me__courses-item">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 0">
        <img class="pic" src="@/assets/images/program-disain.png" v-if="photoNumber === 1">
        <img class="pic" src="@/assets/images/card-steps.png" v-if="photoNumber === 2">
        <img class="pic" src="@/assets/images/program-finance.png" v-if="photoNumber === 3">
        <div class="me__courses-item-body">
            <div class="me__courses-item-header">
                <div class="me__courses-item-info">
                    <p>
                        {{ props.course.period.starts_at }} {{ props.course.period.ends_at }}
                    </p>
                    <p>
                        {{ props.course.education_form.name }}
                    </p>
                </div>
                <p>
                    {{ props.course.name }}
                </p>
            </div>
            <div class="me__courses-progress-list">
                <course-task
                    v-for="(job, index) of lectures" :key="index" :job="job"
                />
            </div>
            <div class="me__courses-item-content">
                <div class="me__courses-data">
                    <span>
                        <!-- {{ props.course.lecture.passed }} / {{ props.course.lecture.total }} лекций -->
                        Лекций: {{ lectures.length }} 
                    </span>
                </div>
                <div class="me__courses-data">
                    <span>
                        <!-- {{ props.course.tasks.passed }} / {{ props.course.tasks.total }} парктических заданий -->
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
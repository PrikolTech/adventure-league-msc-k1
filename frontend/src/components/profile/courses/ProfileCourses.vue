<script setup>
import { onMounted, ref } from "vue";
import CourseItem from "./CourseItem.vue";

const courses = ref([
    {
        title: 'Финансовая грамотность',
        date: '06.11.23 - 10.01.24',
        format: 'Онлайн',
        lecture: {
            passed: 0,
            total: 5,
        },
        tasks: {
            passed: 0,
            total: 3,
        },
    },
    {
        title: 'Дизайн система как смысл жизни',
        date: '06.11.23 - 10.01.24',
        format: 'Онлайн',
        lecture: {
            passed: 4,
            total: 9,
        },
        tasks: {
            passed: 2,
            total: 3,
        },
    },
    {
        title: 'Дизайн система как смысл жизни',
        date: '06.11.23 - 10.01.24',
        format: 'Онлайн',
        lecture: {
            passed: 4,
            total: 9,
        },
        tasks: {
            passed: 2,
            total: 3,
        },
    },
])

const getCourses = async() => {
    // const response = [...courses.value]
}

onMounted(() => {
    getCourses()
})
</script>

<template>
    <div class="profile__courses profile__me profile__block">
        <div class="profile__me-body">
            <div class="profile__me-content">
                <div class="profile__block-header">
                    <div class="profile__title title">
                        Курсы
                    </div>
                    <div class="profile__header-line line"></div>
                </div>
                <div class="profile__courses-header">
                    
                </div>
                <div class="profile__courses-list">
                    <course-item
                        v-for="(course, index) of courses" :key="index" :course="course"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.me__courses-item {
    padding-left: 40%;
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

[dark=true] {
    .me__courses-item-info {
    & p {
        color: var(--gray-200, #E5E7EB);
        background: var(--gray-500, #6B7280);
    }
}
}
</style>
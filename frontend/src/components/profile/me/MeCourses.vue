<script setup>
import { onMounted, ref } from "vue";
import MeCourseItem from '@/components/profile/me/MeCourseItem.vue'
import { useUser } from '@/stores/user'

const userStore = useUser()



const courses = ref([])

// const courses = ref([
//     {
//         title: 'Финансовая грамотность',
//         desc: 'Вы прошли половину обучения.',
//         lecture: {
//             passed: 0,
//             total: 5,
//         },
//         tasks: {
//             passed: 0,
//             total: 3,
//         },
//     },
//     {
//         title: 'Финансовая грамотность',
//         desc: 'Вы прошли половину обучения.',
//         lecture: {
//             passed: 4,
//             total: 9,
//         },
//         tasks: {
//             passed: 2,
//             total: 3,
//         },
//     },
//     {
//         title: 'Финансовая грамотность',
//         desc: 'Вы прошли половину обучения.',
//         lecture: {
//             passed: 4,
//             total: 9,
//         },
//         tasks: {
//             passed: 3,
//             total: 3,
//         },
//     },
// ])


const getUserCourse = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses?user_id=${userStore.user.id}`, {
            method: "GET",
        })
        
        const data = await response.json()

        
        console.log(data)
        courses.value = [...data]

    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    getUserCourse()
})
</script>


<template>
    <div class="me__courses">
        <div class="me__courses-header">

        </div>
        <div class="me__courses-list" v-if="courses.length">
            <me-course-item
                v-for="(course, index) of courses" :key="index" :course="course"
            />
        </div>
        <div class="text" v-else>
            У вас еще нет курсов
        </div>
    </div>
</template>

<style lang="scss">
.me__courses-header {
}
.me__courses-item-inside {
}
.me {

    &__courses {

    }

    &__courses-list {
        max-height: 600px;
        overflow: auto;
        border-radius: 20px;
    }

    &__courses-item {
        border-radius: 20px;
        background: var(--gray-50, #F9FAFB);
        position: relative;
        overflow: hidden;
        padding-left: 150px;
        @media (max-width: 767px) {
            padding-left: 75px;
        }
        &:not(:last-child) {
            margin-bottom: 30px;
            @media (max-width: 767px) {
                margin-bottom: 15px;
            }
        }
        & .pic {
            position: absolute;
            height: 100%;
            width: 100%;
            object-fit: cover;
            border-radius: var(--rounded-3xl, 24px);
            left: 0;
            top: 0;
            max-width: 202px;
        }
    }

    &__courses-item-body {
        padding: 30px;
        z-index: 2;
        position: relative;
        display: flex;
        flex-direction: column;
        gap: 20px;
        background: rgba(243, 244, 246, 0.80);
        backdrop-filter: blur(12.5px);
        border-radius: 30px 20px 20px 30px;
        @media (max-width: 767px) {
            gap: 10px;
            padding: 15px;
        }
    }

    &__courses-item-header {
        & p {
            color: var(--gray-900, #111928);
            font-family: Inter;
            font-size: 18px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 27px */
            @media (max-width: 767px) {
                font-size: 15px;
            }
        }
        & span {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__courses-item-content {
        display: flex;
        gap: 40px;
    }

    &__courses-data {
        & p {
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%;
            @media (max-width: 767px) {
                font-size: 13px;
            }
            & b {
                color: var(--unnamed, #1F222E);
            }
        }
    }
}
[dark=true] {
    & .me__courses-item {
        background: #1A2537;
    }

    & .me__courses-item-body {
        background: rgba(55, 65, 81, 0.90);
    }

    & .me__courses-item-header {
        & p {
            color: var(--gray-100, #F3F4F6);

        }
        & span {
            color: var(--gray-100, #F3F4F6);

        }
    }

    & .me__courses-data {
        & p {
            color: var(--gray-400, #9CA3AF);
        }
        & span {
            color: var(--gray-400, #9CA3AF);
            & b {
                color: var(--gray-100, #F3F4F6);
            }
        }
    }
}
</style>
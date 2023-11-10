<script setup>
import { onMounted, onBeforeUnmount, ref, watch } from 'vue';
import TheNavigation from '@/components/layouts/TheNavigation.vue';
import ProfileNav from '@/components/profile/ProfileNav.vue';
import AsideLesson from '@/components/course/AsideLesson.vue';
import TheLesson from '@/components/course/TheLesson.vue';
import TaskLesson from '@/components/course/TaskLesson.vue';
import TestLesson from '@/components/course/TestLesson.vue';
import { useRouter, useRoute } from 'vue-router'
import { useUser } from '@/stores/user'
import TheButton from '@/components/layouts/TheButton.vue';
import { usePopups } from '@/stores/popups';
import CreateLessonPopup from '@/components/course/CreateLessonPopup.vue';

const userStore = useUser()
const router = useRouter()
const route = useRoute()
const course = ref({})
const popupStore = usePopups()

// const course = ref({
//     name: 'Финансовая грамотность',
//     teacher: 'Кураева Анна',
//     curator: 'Зайцев Денис',
//     lessons: [
//         {
//             title: 'Урок 1. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'video',
//             completed: true,
//         },
//         {
//             title: 'Урок 2. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'file',
//             completed: true,
//         },
//         {
//             title: 'Урок 3. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'video',
//             completed: false,
//         },
//         {
//             title: 'Урок 4. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'file',
//             completed: false,
//         },
//         {
//             title: 'Урок 5. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'video',
//             completed: false,
//         },
//         {
//             title: 'Урок 6. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'file',
//             completed: false,
//         },
//         {
//             title: 'Урок 7. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'video',
//             completed: false,
//         },
//         {
//             title: 'Урок 8. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'file',
//             completed: false,
//         },
//         {
//             title: 'Урок 9. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'video',
//             completed: false,
//         },
//         {
//             title: 'Урок 10. Основы и понятия',
//             date: '08.11.23 в 20:00',
//             format: 'file',
//             completed: false,
//         },
//     ]
// })


const currentLesson = ref({})
// const test_currentLesson = ref({
//     name: 'Урок 1. Введение в финансовую грамотность.',
//     video: 'https://www.youtube.com/embed/__-vp0g_BhA?si=ou5xfV8arD_ccw6Q',
//     desc: 'В этом уроке вы познакомитесь с основами и понятиями финансовой грамотности. Более подробно разберете важность изучения финансовой грамотности и многое другое.',
//     links: {
//         presentation: 'Ссылка на презентацю',
//         abstract: 'Ссылка на конспект',
//     },
//     job: {
//         test: 'Ссылка на тест',
//         task: 'Ссылка на задание'
//     },
//     comments: [
//         {
//             first_name: 'Вера',
//             last_name: 'Красулина',
//             type: 'Студент',
//             text: 'Всем привет, подскажите что задали.',
//             date: '25 августа 2023',
//             time: '21:35',
//         },
//         {
//             first_name: 'Анатолий',
//             last_name: 'Видин',
//             type: 'Учитель',
//             text: 'Всем привет, подскажите что задали.',
//             date: '25 августа 2023',
//             time: '21:35',
//         },
//         {
//             first_name: 'Максим',
//             last_name: 'Трушин',
//             type: 'Студент',
//             text: 'Всем привет, подскажите что задали.',
//             date: '25 августа 2023',
//             time: '21:35',
//         }
//     ],
// })

let files = ref({})
let showTest = ref(false)
let activeTab = ref(null)
const navigationLinks = ref([
    {
        href: '/',
        text: 'Главная'
    },
    {
        href: '/profile/courses',
        text: 'Курсы'
    },
    // {
    //     href: `${route.path}`,
    //     text: course.value.name
    // },
])



const getCourseInfo = async () => {
    const courseID = route.params.id
    try {
        
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${courseID}`, {
            method: "GET",
        })

        const data = await response.json()
        course.value = {...data}

        if(navigationLinks.value.length === 3) {
            const link = {...navigationLinks.value[2]}
            navigationLinks.value[2] = {
                href: `${route.path}`,
                text: course.value.name
            }
            navigationLinks.value.push({...link})
        } else {
            navigationLinks.value.push({
                href: `${route.path}`,
                text: course.value.name
            })
        }


        console.log('курс:',data)

        await getCourseLessons()

    } catch(err) {
        console.error(err)
    }
}


const getCourseLessons = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures?user_id=${userStore.user.id}`, {
            method: "GET",
        })
        
        const data = await response.json()
        course.value.lessons = [...data]
        console.log('лекции',data)
    } catch(err) {
        console.error(err)
    }
}

const getLesson = async (id) => {
    let lessonID = null
    if(id) {
        lessonID = id
    } else {
        lessonID = route.query.lesson
    }
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures/${lessonID}`, {
            method: "GET",
        })
        
        const data = await response.json()
        currentLesson.value = {...data}

        console.log('лекция',data)
    } catch(err) {
        console.error(err)
    }

    // const response = {...test_currentLesson.value}

    // currentLesson.value = {...response}

    showTest.value = false
    activeTab.value = 'lesson'
    getLessonFiles(lessonID)
}

const getLessonFiles = async (lessonID) => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures/${lessonID}/contents`, {
            method: "GET",
        })
        
        const data = await response.json()
        files.value = {...data}

        console.log('файлы лекции',data)
    } catch(err) {
        console.error(err)
    }
}

const getTaskOfLesson = async () => {
    if (navigationLinks.value.some(link => link.text === 'Задание к уроку')) {
        return
    }
    if(navigationLinks.value.length === 4) {
        navigationLinks.value.pop()
    }

    const currentParams = { ...route.query };

    if (currentParams.test) {
        delete currentParams.test;
    }

    currentParams.task = 1;

    router.push({ query: currentParams });

    navigationLinks.value.push(
        {
            href: route.fullPath,
            text: 'Задание к уроку'
        },
    )

    activeTab.value = 'task'

}


const getTestOfLesson = async () => {
    if (navigationLinks.value.some(link => link.text === 'Тест к уроку')) {
        return
    }
    if(navigationLinks.value.length === 4) {
        navigationLinks.value.pop()
    }
    const currentParams = { ...route.query };

    if (currentParams.task) {
        delete currentParams.task;
    }

    currentParams.test = 1;

    router.push({ query: currentParams });
    navigationLinks.value.push(
        {
            href: route.fullPath,
            text: 'Тест к уроку'
        },
    )

    activeTab.value = 'test'

}


const watchRouteChanges = watch(
  () => route.query,
  (newParams) => {
    const hasParams = Object.keys(newParams).length > 0;
    if(hasParams) {
        if(newParams.task) {
            getTaskOfLesson()
            return
        }
        if(newParams.test) {
            getTestOfLesson()
            return
        }
        if(newParams.lesson) {
            showTest.value = false
            activeTab.value = 'lesson'
            if(navigationLinks.value.length === 4) {
                navigationLinks.value.pop()
            }
            return
        }
    }
    else if (!hasParams) {
        showTest.value = false
        if(navigationLinks.value.length === 4) {
            navigationLinks.value.pop()
        }
        if(!(Object.keys(newParams).length > 0)) {
            activeTab.value = null
        } else {
            showTest.value = false
            activeTab.value = 'lesson'
            router.push({ query: { lesson: 1 } })
        }

    }
  }
);

const openCreateLessonPopup = () => {
    popupStore.disableScroll('createLesson')
}


const openLesson = (lessonID) => {
    router.push({ query: { lesson: lessonID } })
    getLesson(lessonID)
}


const makeUrlFordownloadFile = (fileName) => {
    return `${import.meta.env.VITE_FILE_COURSE_URL}/api/files/${fileName}`;
}

onBeforeUnmount(() => {
    watchRouteChanges();
});


onMounted(async() => {
    await getCourseInfo()
    if(route.query.lesson) {
        await getLesson()
    }
    if(route.query.task) {
        getTaskOfLesson()
    }
    if(route.query.test) {
        console.log('test')
        getTestOfLesson()
    }
})

</script>

<template>
    <main>
        <div class="profile course container" style="padding-top: 5px;">
            <profile-nav
                :activeLink="'/profile/courses'"
                v-if="!showTest"
            />
            <div class="course__content profile__block">
                <div class="course__header">
                    <div class="course__title title">
                        {{ course.name }}
                    </div>
                    <div class="course__header-line line"></div>
                </div>
                <the-navigation
                    :class="'course__nav'"
                    :links="navigationLinks"
                />
                <div class="course__block">
                    <the-lesson
                        v-if="activeTab === 'lesson' && Object.keys(currentLesson).length !== 0"
                        :lesson="currentLesson"
                    />
                    <task-lesson
                        v-if="activeTab === 'task'"
                    />
                    <test-lesson
                        v-if="activeTab === 'test'"
                        v-model:show-test="showTest"
                    />
                </div>
            </div>
            <div class="course__aside profile__events"
            v-if="!showTest"
            >
                <div class="course__aside-item materials"
                    v-if="files && activeTab"
                >
                    <div class="course__aside-title">
                        Материалы :
                    </div>
                    <div class="course__aside-item-list">
                        <a class="course__aside-link"
                            v-for="(file, index) of files" :key="index"
                            :href="makeUrlFordownloadFile(file.url)"
                        >
                            {{file.url.split('/')[1]}}
                        </a>

                    </div>
                </div>
                <div class="course__aside-item materials"
                    v-if="activeTab && activeTab !=='task' && activeTab !=='test'"
                >
                    <div class="course__aside-title">
                        Практическое задание:
                    </div>
                    <div class="course__aside-item-list">
                        <p class="course__aside-link"
                            v-if="currentLesson"
                            @click="getTaskOfLesson()"
                        >
                            Задание по материалу
                        </p>
                        <p class="course__aside-link"
                            v-if="currentLesson"
                            @click="getTestOfLesson()"
                        >
                            Тест 
                        </p>
                    </div>
                </div>
                <div class="course__aside-item content">
                    <div class="course__aside-title">
                        Содержание программы :
                    </div>
                    <div class="course__aside-item-list">
                        <aside-lesson
                            v-for="(lesson, index) of course.lessons" :key="index" :lesson="lesson"
                            @click="openLesson(lesson.id)"
                        />
                    </div>
                </div>
                <the-button
                    v-if="userStore.user.role === 'teacher'"
                    :styles="['btn_red']"
                    :type="'button'"
                    style="margin-top: 10px; width: 100%;"
                    @click="openCreateLessonPopup()"
                >
                    Добавить занятие
                </the-button>
                <div class="course__aside-item teacher" v-if="course.teacher">
                    <div class="course__aside-title">
                        Вопросы по ДЗ и учебному материалу можно задать преподавателю:
                    </div>
                    <div class="course__aside-item-list">
                        <div class="aside__teacher">
                            <span>
                                Преподаватель : 
                            </span>
                            <b>
                                <!-- Кураева Анна  -->
                                {{ course.teacher }}
                            </b>
                        </div>
                    </div>
                </div>
                <div class="course__aside-item teacher" v-if="course.curator">
                    <div class="course__aside-title">
                        Все остальные вопросы можно задать куратору :
                    </div>
                    <div class="course__aside-item-list">
                        <div class="aside__teacher">
                            <span>
                                Куратор : 
                            </span>
                            <b>
                                <!-- Зайцев Денис -->
                                {{ course.curator }}
                            </b>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    <create-lesson-popup/>
    </main>
</template>

<style lang="scss">
.course {

    &__body {
        display: flex;
        justify-content: space-between;
        gap: 70px;
    }

    &__content {
        flex: 1;
        @media (max-width: 1023px) {
            flex: 0 0 100%;
        }
    }
    &__header {
        display: flex;
        gap: 40px;
        align-items: center;
        margin-bottom: 40px;
        @media (max-width: 1023px) {
            margin-bottom: 20px;
        }
    }

    &__title {
    }

    &__header-line {
        flex: 1;
    }

    &__nav {
        margin-bottom: 60px;
        @media (max-width: 1023px) {
            margin-bottom: 20px;
        }
    }

    &__block {
    }

    &__aside {
        border-radius: 20px;
        background: var(--gray-200, #E5E7EB);
        padding: 30px;
        flex: 0 0 360px;
        height: fit-content;
        @media (max-width: 539px) {
            flex: 0 1 320px;
        }
    }

    &__aside-item {
        border-radius: var(--rounded-xl, 12px);
        background: #FFF;
        padding: 16px;
        &:not(:last-child) {
            margin-bottom: 25px;
        }
    }

    &__aside-title {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
    }

    &__aside-item-list {
    }

    &__aside-link {
        color: var(--primary-2-default, #003791);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        line-height: 167%; /* 26.72px */
        display: block;
        cursor: pointer;
        font-weight: 400;
    }
}

.aside {

    &__lesson {
        display: flex;
        gap: 12px;
        align-items: center;
        position: relative;
        &:not(:last-child) {
            margin-bottom: 15px;
        }
    }

    &__lesson-action {
        min-width: 90px;
        height: 50px;
        border-radius: var(--rounded-xl, 12px);
        background: var(--gray-100, #F3F4F6);
        display: flex;
        justify-content: center;
        align-items: center;
        position: relative;
        cursor: pointer;
        & .icon {
            display: flex;
        }
        & .completed {
            position: absolute;
            left: 5px;
            top: 2px;
            color: var(--gray-500, #6B7280);
            font-family: Roboto;
            font-size: 8px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 9px */
            & div {
                display: flex;
                align-items: center;
                gap: 3px;
            }
        }
    }

    &__lesson-info {
        & p {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 18px */
        }
        & span {
            color: var(--gray-600, #4B5563);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
        }
    }

    &__teacher {
    }
}
.content {
    & .course__aside-item-list {
        padding-top: 10px;
        max-height: 400px;
        overflow: auto;
    }
}
.teacher {
    font-family: Roboto;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 150%; /* 21px */
    & .course__aside-title {
        margin-bottom: 5px;
    }
    & span {
        color: var(--gray-600, #4B5563);
    }
    & b {
    font-weight: 400;
        color: var(--gray-800, #1F2A37);
    }
}

[dark=true] {
    & .course__aside-item {
        background: var(--gray-700, #374151);
    }

    & .course__aside-title {
        color: var(--gray-50, #F9FAFB);
    }

    & .course__aside-link {
        color: var(--primary-400, #76A9FA);
    }

    & .aside__lesson-info {
        & p {
            color: var(--gray-50, #F9FAFB);
        }
        
        & span {
            color: var(--gray-400, #9CA3AF);
        }
    }

    & .aside__lesson-action {
        background: var(--gray-600, #4B5563);
        & .icon {
            & path {
            }
        }
    }

    & .teacher span {
        color: var(--gray-600, #9CA3AF);
    }
    
    & .teacher b {
        color: var(--gray-600, #F9FAFB);
    }

    & .completed {
        color: var(--gray-300, #D1D5DB);
        & svg {
            & path {
                fill: #D1D5DB;
            }
        }
    }
}

</style>
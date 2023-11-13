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
import CreateTaskPopup from '@/components/course/CreateTaskPopup.vue';

const userStore = useUser()
const router = useRouter()
const route = useRoute()
const course = ref({})
const popupStore = usePopups()


const currentLesson = ref({})


let files = ref({})
let tests = ref({})
let homes = ref({})
let currentTest = ref({})
let currentTask = ref({})
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
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })

        const data = await response.json()
        // course.value.length = 0
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



        await getCourseLessons()

    } catch(err) {
        console.error(err)
    }
}


const getCourseLessons = async () => {
    let url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures?user_id=${userStore.user.id}`
    if(userStore.checkRole('student')) {
        url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures?user_id=${userStore.user.id}`
    } else {
        url = `${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures`
    }
    try {
        const response = await fetch(url, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
            // credentials: 'include',
        })
        
        const data = await response.json()
        if(course.value.lessons) {
            course.value.lessons.length = 0
        }
        course.value.lessons = [...data]
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
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()
        currentLesson.value = {...data}

    } catch(err) {
        console.error(err)
    }

    // const response = {...test_currentLesson.value}

    // currentLesson.value = {...response}

    showTest.value = false
    activeTab.value = 'lesson'
    getLessonFiles(lessonID)
    getHomeLesson(lessonID)
    getTestLesson(lessonID)
}

const getLessonFiles = async (lessonID) => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${course.value.id}/lectures/${lessonID}/contents`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()
        files.value.length = 0
        files.value = [...data]

    } catch(err) {
        console.error(err)
    }
}


const getTestLesson = async (lessonID) => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${lessonID}/tests`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        });

        const data = await response.json()
        tests.value = { ...data }
    } catch (err) {
        console.log(err);
    }
}

const getHomeLesson = async (lessonID) => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${lessonID}/homeworks`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        });
        const data = await response.json()
        homes.value.length = 0
        homes.value = { ...data }
    } catch (err) {
        console.log(err);
    }
}

const getTaskOfLesson = async (taskInfo) => {
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

    if(taskInfo) {
        currentParams.task = taskInfo.id;
    }

    router.push({ query: currentParams });

    navigationLinks.value.push(
        {
            href: route.fullPath,
            text: 'Задание к уроку'
        },
    )

    activeTab.value = 'task'
    currentTask.value = { ...taskInfo }

}

const getTestOfLesson = async (testInfo) => {
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
    if(testInfo) {
        currentParams.test = testInfo.id;
    }

    router.push({ query: currentParams });
    navigationLinks.value.push(
        {
            href: route.fullPath,
            text: 'Тест к уроку'
        },
    )

    activeTab.value = 'test'

    currentTest.value = { ...testInfo }

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

const openCreateHomeWorkPopup = () => {
    popupStore.disableScroll('createTask')
}

const openLesson = (lessonID) => {
    router.push({ query: { lesson: lessonID } })
    getLesson(lessonID)
}


const makeUrlFordownloadFile = (fileName) => {
    return `${import.meta.env.VITE_SERVICE_FILE_URL}/files/${fileName}`;
}

const deleteFile = async(fileID) => {
    const courseID = route.params.id
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${courseID}/lectures/${currentLesson.value.id}/contents/${fileID} `, {
            method: "DELETE",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        });
        if(response.ok) {
            files.value.forEach((el) => {
                if (el.id === fileID) {
                    // Находим индекс элемента с нужным id и удаляем его из массива
                    const elementIndex = files.value.findIndex(item => item.id === fileID);
                    if (elementIndex !== -1) {
                        files.value.splice(elementIndex, 1);
                        // Можно также обновить исходный объект files.value, если требуется
                        files.value = files.value.reduce((acc, curr) => {
                            acc[curr.id] = curr;
                            return acc;
                        }, {});
                    }
                }
            });
        }
    } catch (err) {
        console.log(err);
    }
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
                        :course="course"
                        @uploadFile="getLessonFiles(route.query.lesson)"
                        @updateLecture="getCourseInfo()"
                    />
                    <task-lesson
                        v-if="activeTab === 'task'"
                        :task="currentTask"
                        :lesson="currentLesson"
                    />
                    <test-lesson
                        v-if="activeTab === 'test'"
                        :lesson="currentLesson"
                        :test="currentTest"
                        v-model:show-test="showTest"
                    />
                </div>
            </div>
            <div class="course__aside profile__events"
                v-if="!showTest"
            >
                <div class="course__aside-item materials"
                    v-if="files && activeTab === 'lesson'"
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
                            <svg width="20" height="21" viewBox="0 0 20 21" fill="none" xmlns="http://www.w3.org/2000/svg"
                                v-if="userStore.checkRole('teacher')"
                                @click.stop.prevent="deleteFile(file.id)"
                            >
                                <path d="M16.6667 5.67546H13.3333V3.92107C13.3333 3.45578 13.1577 3.00955 12.8452 2.68053C12.5326 2.35152 12.1087 2.16669 11.6667 2.16669H8.33333C7.89131 2.16669 7.46738 2.35152 7.15482 2.68053C6.84226 3.00955 6.66667 3.45578 6.66667 3.92107V5.67546H3.33333C3.11232 5.67546 2.90036 5.76788 2.74408 5.93238C2.5878 6.09689 2.5 6.32001 2.5 6.55265C2.5 6.7853 2.5878 7.00842 2.74408 7.17292C2.90036 7.33743 3.11232 7.42985 3.33333 7.42985H4.16667V17.079C4.16667 17.5443 4.34226 17.9905 4.65482 18.3195C4.96738 18.6485 5.39131 18.8334 5.83333 18.8334H14.1667C14.6087 18.8334 15.0326 18.6485 15.3452 18.3195C15.6577 17.9905 15.8333 17.5443 15.8333 17.079V7.42985H16.6667C16.8877 7.42985 17.0996 7.33743 17.2559 7.17292C17.4122 7.00842 17.5 6.7853 17.5 6.55265C17.5 6.32001 17.4122 6.09689 17.2559 5.93238C17.0996 5.76788 16.8877 5.67546 16.6667 5.67546ZM8.33333 3.92107H11.6667V5.67546H8.33333V3.92107ZM14.1667 17.079H5.83333V7.42985H14.1667V17.079Z" fill="#9CA3AF"/>
                                <path d="M8.33333 8.30704C8.11232 8.30704 7.90036 8.39946 7.74408 8.56396C7.5878 8.72847 7.5 8.95159 7.5 9.18423V15.3246C7.5 15.5572 7.5878 15.7803 7.74408 15.9449C7.90036 16.1094 8.11232 16.2018 8.33333 16.2018C8.55435 16.2018 8.76631 16.1094 8.92259 15.9449C9.07887 15.7803 9.16667 15.5572 9.16667 15.3246V9.18423C9.16667 8.95159 9.07887 8.72847 8.92259 8.56396C8.76631 8.39946 8.55435 8.30704 8.33333 8.30704Z" fill="#9CA3AF"/>
                                <path d="M11.6667 8.30704C11.4457 8.30704 11.2337 8.39946 11.0774 8.56396C10.9211 8.72847 10.8333 8.95159 10.8333 9.18423V15.3246C10.8333 15.5572 10.9211 15.7803 11.0774 15.9449C11.2337 16.1094 11.4457 16.2018 11.6667 16.2018C11.8877 16.2018 12.0996 16.1094 12.2559 15.9449C12.4122 15.7803 12.5 15.5572 12.5 15.3246V9.18423C12.5 8.95159 12.4122 8.72847 12.2559 8.56396C12.0996 8.39946 11.8877 8.30704 11.6667 8.30704Z" fill="#9CA3AF"/>
                            </svg>
                        </a>
                        <div class="text" style="font-size: 15px;" v-if="!Object.keys(files).length > 0">
                            Материалов нет
                        </div>
                    </div>
                </div>
                <div class="course__aside-item materials"
                    v-if="activeTab && activeTab !=='task' && activeTab !=='test' && ( userStore.checkRole('teacher') || userStore.checkRole('student') )"
                >
                    <div class="course__aside-title">
                        Практическое задание:
                    </div>
                    <div class="text"
                        v-if="!Object.keys(homes).length > 0 || !Object.keys(tests).length > 0"
                    >
                        Заданий еще нет
                    </div>
                    <div class="course__aside-item-list">
                        <!-- <p class="course__aside-link"
                            @click="getTaskOfLesson()"
                        >
                            Задание по материалу
                        </p> -->
                        <span class="text"
                            v-if="Object.keys(homes).length > 0"
                        >
                            Домашняя работа
                        </span>
                        <p class="course__aside-link"
                            v-for="home of homes" :key="home.id"
                            @click="getTaskOfLesson(home)"
                        >
                            <!-- Задание по материалу -->
                            {{ home.name }}
                        </p>
                        <span class="text"
                            v-if="Object.keys(tests).length > 0"
                        >
                            Тестирование
                        </span>
                        <p class="course__aside-link"
                            v-for="test of tests" :key="test.id"
                            @click="getTestOfLesson(test)"
                        >
                            {{ test.name }}
                            <!-- Тест -->
                        </p>
                        <!-- <p class="course__aside-link"
                            v-if="currentLesson"
                            @click="getTestOfLesson()"
                        >
                            Тест 
                        </p> -->
                    </div>
                </div>
                <the-button
                    v-if="userStore.checkRole('teacher') && activeTab === 'lesson'"
                    :styles="['btn_red']"
                    :type="'button'"
                    style="margin-top: 10px; margin-bottom: 25px; width: 100%;"
                    @click="openCreateHomeWorkPopup()"
                >
                    Добавить задание
                </the-button>
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
                    v-if="userStore.checkRole('tutor')"
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
    <create-lesson-popup
        @createdLesoon="getCourseLessons()"
    />
    <create-task-popup
        @createdTask="getHomeLesson(route.query.lesson)"
    />
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
        display: flex;
        align-items: center;
        gap: 10px;
        justify-content: space-between;
        transition: .2s;
        &:hover {
            text-decoration: underline;
        }
        & svg {
            & path {

            }
        }
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
        & svg {
            & path {
                fill: #6B7280;
            }
        }
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
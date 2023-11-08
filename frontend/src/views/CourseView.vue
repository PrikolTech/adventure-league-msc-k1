<script setup>
import { onMounted, onBeforeUnmount, ref, watch } from 'vue';
import TheNavigation from '@/components/layouts/TheNavigation.vue';
import ProfileNav from '@/components/profile/ProfileNav.vue';
import AsideLesson from '@/components/course/AsideLesson.vue';
import TheLesson from '@/components/course/TheLesson.vue';
import TaskLesson from '@/components/course/TaskLesson.vue';
import TestLesson from '@/components/course/TestLesson.vue';
import { useRouter, useRoute } from 'vue-router'


const router = useRouter()
const route = useRoute()


const course = ref({
    name: 'Финансовая грамотность',
    teacher: 'Кураева Анна',
    curator: 'Зайцев Денис',
    lessons: [
        {
            title: 'Урок 1. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'video',
            completed: true,
        },
        {
            title: 'Урок 2. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'file',
            completed: true,
        },
        {
            title: 'Урок 3. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'video',
            completed: false,
        },
        {
            title: 'Урок 4. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'file',
            completed: false,
        },
        {
            title: 'Урок 5. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'video',
            completed: false,
        },
        {
            title: 'Урок 6. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'file',
            completed: false,
        },
        {
            title: 'Урок 7. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'video',
            completed: false,
        },
        {
            title: 'Урок 8. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'file',
            completed: false,
        },
        {
            title: 'Урок 9. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'video',
            completed: false,
        },
        {
            title: 'Урок 10. Основы и понятия',
            date: '08.11.23 в 20:00',
            format: 'file',
            completed: false,
        },
    ]
})

const navigationLinks = ref([
    {
        href: '/',
        text: 'Главная'
    },
    {
        href: '/profile/courses',
        text: 'Курсы'
    },
    {
        href: `${route.path}`,
        text: course.value.name
    },
])

let activeTab = ref(null)

const openLesson = () => {
    router.push({ query: { lesson: 1 } })
    getLesson()
}

const test_currentLesson = ref({
    name: 'Урок 1. Введение в финансовую грамотность.',
    video: 'https://www.youtube.com/embed/__-vp0g_BhA?si=ou5xfV8arD_ccw6Q',
    desc: 'В этом уроке вы познакомитесь с основами и понятиями финансовой грамотности. Более подробно разберете важность изучения финансовой грамотности и многое другое.',
    links: {
        presentation: 'Ссылка на презентацю',
        abstract: 'Ссылка на конспект',
    },
    job: {
        test: 'Ссылка на тест',
        task: 'Ссылка на задание'
    },
    comments: [
        {
            first_name: 'Вера',
            last_name: 'Красулина',
            type: 'Студент',
            text: 'Всем привет, подскажите что задали.',
            date: '25 августа 2023',
            time: '21:35',
        },
        {
            first_name: 'Анатолий',
            last_name: 'Видин',
            type: 'Учитель',
            text: 'Всем привет, подскажите что задали.',
            date: '25 августа 2023',
            time: '21:35',
        },
        {
            first_name: 'Максим',
            last_name: 'Трушин',
            type: 'Студент',
            text: 'Всем привет, подскажите что задали.',
            date: '25 августа 2023',
            time: '21:35',
        }
    ],
})

const currentLesson = ref({})
const getLesson = async () => {

    const response = {...test_currentLesson.value}

    currentLesson.value = {...response}


    activeTab.value = 'lesson'
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

    if (!hasParams) {
        if(navigationLinks.value.length === 4) {
            navigationLinks.value.pop()
        }
        activeTab.value = 'lesson'
        router.push({ query: { lesson: 1 } })
    }
  }
);

onBeforeUnmount(() => {
    watchRouteChanges();
});

onMounted(() => {
    if(route.query.lesson) {
        getLesson()
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
                    />
                </div>
            </div>
            <div class="course__aside profile__events">
                <div class="course__aside-item materials"
                    v-if="currentLesson.links"
                >
                    <div class="course__aside-title">
                        Материалы :
                    </div>
                    <div class="course__aside-item-list">
                        <p class="course__aside-link"
                            v-if="currentLesson.links.presentation"
                        >
                            Презентация
                        </p>
                        <p class="course__aside-link"
                            v-if="currentLesson.links.abstract"
                        >
                            Конспект лекции. Введение. 
                        </p>
                    </div>
                </div>
                <div class="course__aside-item materials"
                    v-if="currentLesson.job"
                >
                    <div class="course__aside-title">
                        Практическое задание:
                    </div>
                    <div class="course__aside-item-list">
                        <p class="course__aside-link"
                            v-if="currentLesson.job.task"
                            @click="getTaskOfLesson()"
                        >
                            Задание по материалу
                        </p>
                        <p class="course__aside-link"
                            v-if="currentLesson.job.test"
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
                            @click="openLesson"
                        />
                    </div>
                </div>
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
    }

    &__title {
    }

    &__header-line {
        flex: 1;
    }

    &__nav {
        margin-bottom: 60px;
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
            flex: 0 0 320px;
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

</style>
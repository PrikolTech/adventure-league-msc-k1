<script setup>
import { onMounted, ref, watch } from 'vue';
import TheTask from './tasks/TheTask.vue';
import FilterBtn from './tasks/FilterBtn.vue';
import TheSort from "@/components/layouts/TheSort.vue";
import TheSwitcher from "@/components/layouts/TheSwitcher.vue";


const filterBtns = ref([
    {
        text: 'Текущее',
        type: 'current'
    },
    {
        text: 'На проверке',
        type: 'inspection'
    },
    {
        text: 'Завершенные',
        type: 'completed'
    },
])

const sortSelectors = ref(['Все','Финансовая грамотность','Дизайн как смысл жизни','Управление командой','Общение в команде'])
let filterSelect = ref('Все')
let selectorIsActive = ref(false)
let onlyExams = ref(false)
let filterActiveStatus = ref('current')

const filterTasks = () => {
    filtredTasks.value = [...tasks.value.filter(el => {
        if(filterSelect.value === 'Все') {
            if(onlyExams.value) {
                return (el.status === filterActiveStatus.value && el.format === 'Экзамен')
            } else {
                return (el.status === filterActiveStatus.value)
            }
        } else {
            if(onlyExams.value) {
                return (el.status === filterActiveStatus.value && el.format === 'Экзамен' && el.desc === filterSelect.value)
            } else {
                return (el.status === filterActiveStatus.value && el.desc === filterSelect.value)
            }
        }
    })]
}

const tasks = ref([])
const filtredTasks = ref([])
const test_tasks = ref([
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'ДЗ',
        desc: 'Финансовая грамотность',
        time: 'Сегодня до 20:00',
        status: 'current',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'ДЗ',
        desc: 'Дизайн как смысл жизни',
        time: '21 марта 2023 до 20:00',
        status: 'current',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Экзамен',
        desc: 'Финансовая грамотность',
        time: '09 ноября 2023 до 20:00',
        status: 'current',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'ДЗ',
        desc: 'Общение в команде',
        time: '10 января 2024 до 20:00',
        status: 'current',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'ДЗ',
        desc: 'Дизайн как смысл жизни',
        time: 'Сегодня до 20:00',
        status: 'inspection',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'ДЗ',
        desc: 'Общение в команде',
        time: 'Сегодня до 20:00',
        status: 'inspection',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Тест',
        desc: 'Дизайн как смысл жизни',
        time: 'Сегодня до 20:00',
        status: 'inspection',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Тест',
        desc: 'Управление командой',
        time: 'Сегодня до 20:00',
        status: 'inspection',
        mark: null,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Экзамен',
        desc: 'Дизайн как смысл жизни',
        time: 'Сегодня до 20:00',
        status: 'completed',
        mark: 5,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Экзамен',
        desc: 'Финансовая грамотность',
        time: 'Сегодня до 20:00',
        status: 'completed',
        mark: 4,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Экзамен',
        desc: 'Дизайн как смысл жизни',
        time: 'Сегодня до 20:00',
        status: 'completed',
        mark: 3,
    },
    {
        title: 'Сдать домашнее задание к уроку 3. Название урока.',
        format: 'Экзамен',
        desc: 'Управление командой',
        time: 'Сегодня до 20:00',
        status: 'completed',
        mark: 2,
    },
])

const getTasks = async() => {
    const response = [...test_tasks.value]

    tasks.value = [...response]
    filtredTasks.value = [...response]

}

watch(filterActiveStatus, () => {
    filterTasks()
});

watch(onlyExams, () => {
    filterTasks()
});

watch(filterSelect, () => {
    filterTasks()
});

onMounted(() => {
    getTasks()
    filterTasks()
})

</script>

<template>
    <div class="profile__settings profile__block">
        <div class="profile__block-header">
            <div class="profile__title title">
                Задания
            </div>
            <div class="profile__header-line line"></div>
        </div>
        <div class="tasks">
            <div class="tasks__header">
                <filter-btn
                    v-for="(btn, index) of filterBtns" :key="index" :btnInfo="btn" v-model:filterActiveStatus="filterActiveStatus"
                />
            </div>
            <div class="tasks__filter">
                <the-sort
                    :selectors="sortSelectors"
                    v-model="selectorIsActive"
                    @select="(modelValue) => filterSelect = modelValue"
                />
                <the-switcher
                    :name="'onlyExams'"
                    v-model="onlyExams"
                >
                    Только экзамены
                </the-switcher>
            </div>
            <div class="tasks__list">
                <the-task
                    v-for="(task, index) of filtredTasks" :key="index" :task="task"
                />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.tasks {

    &__header {
        display: flex;
        gap: 15px;
        flex-wrap: wrap;
        margin-bottom: 40px;
    }

    &__header-item {

    }

    &__filter {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 20px;
        flex-wrap: wrap;
        margin-bottom: 30px;
    }

    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 15px;
    }
}

[dark=true] {

}

</style>
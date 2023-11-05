<script setup>
import { ref } from 'vue';
import TheButton from '../layouts/TheButton.vue';
import ProgramItem from './ProgramItem.vue';

let filterProgram = ref([
  {text: 'Все', value: 'all'},
  {text: 'Soft skills', value: 'Soft skills'},
  {text: 'Hard skills', value: 'Hard skills'},
  {text: 'Digital skills', value: 'Digital skills'},
  {text: 'Портфольные программы', value: 'Портфольные программы'},
])
let currentFilterProgram = ref('all')

let programs = ref([
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Онлайн, очно',
        number: '1'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Очно',
        number: '2'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Онлайн, очно',
        number: '3'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Очно',
        number: '4'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Онлайн, очно',
        number: '1'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Очно',
        number: '2'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Онлайн, очно',
        number: '3'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        type: 'Очно',
        number: '4'
    },
])

</script>

<template>
    <div class="program container">
        <div class="program__header">
            <div class="program__title title">
                Программы
            </div>
            <div class="program__line line"></div>
            <div class="program__filter">
                <div class="program__filter-item"
                    v-for="(item, index) of filterProgram" :key="index"
                >
                    <the-button
                        :styles="['btn_grey-lite', {active: currentFilterProgram === item.value}]"
                        :type="'button'"
                    >
                        {{ item.text }}
                    </the-button>
                </div>
            </div>
        </div>
        <div class="program__list">
            <program-item
                v-for="(program, index) of programs" :key="index"
                :program="program"
                :number="index"
            />
        </div>
        <div class="program__footer">
            <div class="program__footer-line line"></div>
            <the-button
                    :styles="['btn_grey']"
                    :type="'button'"
                >
                Показать ещё
            </the-button>
            <div class="program__footer-line line"></div>
        </div>
    </div>
</template>

<style lang="scss" >
.program {

    &__header {
        display: flex;
        gap: 40px;
        align-items: center;
        margin-bottom: 66px;
    }

    &__title {
    }

    &__filter {
        display: flex;
        gap: 18px;
        align-items: center;
    }

    &__filter-item {
        & button {
            &.active {
                border: 1px solid var(--var-blue) !important;
                color: var(--var-blue) !important;
            }
        }
    }

    &__line {
        flex: 1;
    }

    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 40px;
        margin-bottom: 60px;
    }


    &__footer {
        display: flex;
        gap: 47px;
        align-items: center;
        justify-content: space-between;
    }

    &__footer-line {
        flex: 1;
    }
}

[dark=true] {
    & .program__filter-item {
        & .btn_grey-lite {
            color: var(--gray-300, #D1D5DB);
            border: 1px solid var(--gray-700, #374151);
        }
    }
}

</style>
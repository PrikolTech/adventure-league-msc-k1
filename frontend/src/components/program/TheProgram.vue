<script setup>
import { onMounted, ref } from 'vue';
import TheButton from '../layouts/TheButton.vue';
import ProgramItem from './ProgramItem.vue';

let filterProgram = ref([
    //   {text: 'Все', value: 'all'},
    //   {text: 'Soft skills', value: 'Soft skills'},
    //   {text: 'Hard skills', value: 'Hard skills'},
    //   {text: 'Digital skills', value: 'Digital skills'},
    //   {text: 'Портфольные программы', value: 'Портфольные программы'},
])

let currentFilterProgram = ref('all')
let quantityProgramsToShow = ref(7)
let programs = ref([
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. ',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '1',
        type: 'Soft skills'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '2',
        type: 'Hard skills'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '3',
        type: 'Digital skills'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '4',
        type: 'Портфольные программы'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '1',
        type: 'Soft skills'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '2',
        type: 'Hard skills'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '3',
        type: 'Digital skills'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '4',
        type: 'Портфольные программы'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности. ',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '1',
        type: 'Soft skills'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '2',
        type: 'Hard skills'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '3',
        type: 'Digital skills'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '4',
        type: 'Портфольные программы'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '1',
        type: 'Soft skills'
    },
    {
        title: 'Дизайн система как смысл жизни',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '2',
        type: 'Hard skills'
    },
    {
        title: 'Финансовая грамотность',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Онлайн, очно',
        number: '3',
        type: 'Digital skills'
    },
    {
        title: 'Лидерский интенсив',
        description: 'Курс по финансовой грамотности обучает управлению личными финансами, включая бюджетирование, инвестирование и планирование будущего для финансовой стабильности.',
        date: '19 декабря',
        format: 'Очно',
        number: '4',
        type: 'Портфольные программы'
    },
])

let filteredPrograms = ref([])
const getPrograms = async () => {

    const response = programs.value;
    filteredPrograms.value = [...programs.value];

    filterProgram.value = [
        { text: 'Все', value: 'all' },
        ...[...new Set(response.map(program => program.type))].map(type => ({ text: type, value: type }))
    ];

}

const filterPrograms = (value) => {
    currentFilterProgram.value = value

    if(value === 'all') {
        filteredPrograms.value = [...programs.value]
        return
    }

    filteredPrograms.value = [...programs.value.filter(el => {
        return el.type === value;
    })];
}

onMounted(() => {
    getPrograms()
})

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
                        @click="filterPrograms(item.value)"
                    >
                        {{ item.text }}
                    </the-button>
                </div>
            </div>
        </div>
        <div class="program__list">
            <program-item
                v-for="(program, index) of filteredPrograms" :key="index"
                :program="program"
                :number="index"
                :quantityProgramsToShow="quantityProgramsToShow"
            />
        </div>
        <div class="program__footer"
            v-if="(quantityProgramsToShow < filteredPrograms.length - 1)"
        >
            <div class="program__footer-line line"></div>
            <the-button
                    :styles="['btn_grey']"
                    :type="'button'"
                    @click="quantityProgramsToShow += 8"
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
        @media (max-width: 1023px) {
            flex-wrap: wrap;
            gap: 20px;
            margin-bottom: 30px;
        }
        @media (max-width: 539px) {
            gap: 10px;
            margin-bottom: 15px;
        }
    }

    &__title {
    }

    &__filter {
        display: flex;
        gap: 18px;
        align-items: center;
        @media (max-width: 1023px) {
            flex: 0 0 100%;
            width: 100%;
            overflow: auto;
        }
    }

    &__filter-item {
        & button {
            @media (max-width: 1023px) {
                white-space: nowrap;
            }
            &.active {
                border: 1px solid var(--var-blue) !important;
                color: var(--var-blue) !important;
                &:hover {
                    background: var(--var-blue);
                    color: #fff !important;
                }
            }
        }
    }

    &__line {
        flex: 1;
        @media (max-width: 1023px) {
            display: none;
        }
    }

    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: 40px;
        margin-bottom: 60px;
        @media (max-width: 899px) {
            margin-bottom: 40px;
            gap: 20px;
        }
        @media (max-width: 539px) {
            margin-bottom: 25px;
        }
    }


    &__footer {
        display: flex;
        gap: 47px;
        align-items: center;
        justify-content: space-between;
        @media (max-width: 539px) {
            justify-content: center;
        }
    }

    &__footer-line {
        flex: 1;
        @media (max-width: 539px) {
            display: none;
        }
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
<script setup>
import { onMounted, ref, watch } from "vue";
import TheSort from "@/components/layouts/TheSort.vue";
import TheSwitcher from "@/components/layouts/TheSwitcher.vue";

const sortSelectors = ref(['Все'])
let filterSelect = ref('Все')

let selectorIsActive = ref(false)
let onlyExams = ref(false)

let marks = ref([
    {
        title: "Подготовка к старту проекта",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
        format: "Экзамен",
        mark: '5'
    },
    {
        title: "Финансовая грамотность",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
        format: "Урок",
        mark: '4'
    },
    {
        title: "Дизайн как смысл жизни",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
        format: "Домашнее задание",
        mark: '3'
    },
    {
        title: "Управление командой",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
        format: "Экзамен",
        mark: '2'
    },
    {
        title: "Общение в команде",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
        format: "Экзамен",
        mark: '5'
    },
])

let filteredMarks = ref([])

const getMarks = () => {
    const response = [...marks.value]


    filteredMarks.value = [...response]

    filteredMarks.value.forEach(el => {
        if(!sortSelectors.value.includes(el.title)) {
            sortSelectors.value.push(el.title)
        }
    });
}

const filterMarks = (filterValue) => {

    if(filterValue) {
        filterSelect.value = filterValue
    }

    filteredMarks.value = [...marks.value.filter(mark => {
        if(onlyExams.value) {
            if(filterSelect.value === 'Все') {
                return mark.format === 'Экзамен'
            } else {
                return (mark.format === 'Экзамен' && mark.title === filterSelect.value)
            }
        } else {
            if(filterSelect.value === 'Все') {
                return true
            } else {
                return mark.title === filterSelect.value
            }
        }
    })]

}

watch(onlyExams, () => {
    filterMarks()
});


onMounted(() => {
    getMarks()
})


</script>

<template>
    <div class="me__marks">
        <div class="me__marks-header">
            <the-sort
                :selectors="sortSelectors"
                v-model="selectorIsActive"
                @select="(modelValue) => filterMarks(modelValue)"
            />
            <the-switcher
                :name="'onlyExams'"
                v-model="onlyExams"
            >
                Только экзамены
            </the-switcher>
        </div>
        <div class="me__marks-list">
            <div class="me__marks-item"
                v-for="(event, index) of filteredMarks" :key="index"
            >
                <div class="me__marks-item-header">
                    <p>
                        <!-- # Экзамен -->
                        # {{ event.format }}
                    </p>
                    <span>
                        <!-- 25 августа 2023 -->
                        {{ event.date }}
                    </span>
                </div>
                <p class="me__marks-item-title">
                    <!-- Подготовка к старту проекта -->
                    {{ event.title }}
                </p>
                <p class="me__marks-item-course">
                    <!-- Курс : Agile, Scrum.  -->
                    {{ event.desc }}
                </p>
                <div class="me__marks-item-score">
                    <span
                        :class="[{great: event.mark === '5'},{fine: event.mark === '4'},{satisfactorily: event.mark === '3'},{bad: event.mark === '2'}]"
                    >
                        <!-- 5 -->
                        {{ event.mark }}
                    </span>
                    <p>
                        Оценка преподавателя
                    </p>
                </div>
                <div class="me__marks-item-comment">
                    <textarea class="comment" readonly>Наталья, отличная работа! Молодец)</textarea>
                </div>
            </div>
        </div>
        <div class="text" v-if="filteredMarks.length === 0">
            По заданному фильтру ничего не найдено...
        </div>
    </div>
</template>

<style lang="scss" >
.me {

    &__marks {

    }

    &__marks-header {
        margin-bottom: 40px;
        display: flex;
        align-items: center;
        gap: 20px;
        flex-wrap: wrap;
        justify-content: space-between;
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
    }

    &__marks-list {
        max-height: 600px;
        overflow: auto;
        border-radius: 20px;
    }

    &__marks-item {
        border-radius: 20px;
        background: var(--gray-50, #F9FAFB);
        padding: 30px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        @media (max-width: 767px) {
            padding: 15px;
            gap: 10px;
        }

        &:not(:last-child) {
            margin-bottom: 30px;
            @media (max-width: 767px) {
                margin-bottom: 15px;
            }
        }
    }

    &__marks-item-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 10px;
        flex-wrap: wrap;
        & p {
            border-radius: 8px;
            background: var(--gray-200, #E5E7EB);
            color: var(--gray-600, #4B5563);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
            padding: 4px 8px;
            display: flex;
            align-items: center;
            justify-content: center;
            text-align: center;
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__marks-item-title {
        color: var(--unnamed, #1F222E);
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        display: flex;
        align-items: center;
        justify-content: space-between;
        gap: 10px;
        @media (max-width: 767px) {
            font-size: 15px;
        }
        & span {
            color: var(--unnamed, #7C8092);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
            @media (max-width: 767px) {
                font-size: 13px;
            }
        }
    }

    &__marks-item-course {
        color: var(--gray-500, #6B7280);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
    }

    &__marks-item-score {
        display: flex;
        gap: 12px;
        align-items: center;
        & span {
            display: flex;
            justify-content: center;
            align-items: center;
            min-width: 34px;
            min-height: 34px;
            background: #DDF7DA;
            border-radius: 50%;
            color: var(--green-500, var(--green-500, #0E9F6E));
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 24px */
            &.great {
                background: #DDF7DA;
                color: var(--green-500, var(--green-500, #0E9F6E));
            }
            &.fine {
                background: #DDF7DA;
                color: var(--green-500, var(--green-500, #0E9F6E));
            }
            &.satisfactorily {
                background: #eef7da;
                color: var(--green-500, var(--green-500, #959f0e));
            }
            &.bad {
                background: #f7dada;
                color: var(--green-500, var(--green-500, #9f0e0e));
            }
        }
        & p {
            color: var(--unnamed, #1F222E);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
    }

    &__marks-item-comment {

    }
}


[dark=true] {
    & .me__marks-item {
        background: #1A2537;
    }

    & .me__marks-item-header {
        & p {
            color: var(--gray-200, #E5E7EB);
            background: var(--gray-600, #4B5563);
        }
        & span {
            color: var(--gray-500, #6B7280);
        }
    }

    & .me__marks-item-title {
        color: var(--gray-100, #F3F4F6);
    }

    & .me__marks-item-course {
        color: var(--gray-500, #6B7280);
    }

    & .me__marks-item-score {
        & p {
            color: var(--gray-200, #E5E7EB);
        }
    }
}
</style>
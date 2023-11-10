<script setup>
import { onMounted, ref } from 'vue';

const props = defineProps({
    question: {
        type: Object,
        required: true
    },
    number: Number,
    testID: Number,
    lessonID: Number,
})

const questionFullInfo = ref({})

const getQuestionInfo = async() => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/api/jobs/${props.lessonID}/tests/${props.testID}/questions/${props.question.id}`, {
            method: "GET",
        })
        
        const data = await response.json()
        console.log('ВОПРОС:', data)
        questionFullInfo.value = { ...data }
    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    getQuestionInfo()
})


</script>

<template>
    <div class="test__item">
        <div class="test__item-aside">
            <div class="test__preview">
                Вопрос {{ props.number + 1 }}
            </div>
            <div class="test__ball">
                Балл : 1
            </div>
        </div>
        <div class="test__item-content">
            <div class="test__item-title">
                <!-- При выборе банка необходимо в первую очередь обратить внимание на: -->
                {{ props.question.body }}
            </div>
            <div class="test__item-questions-list">
                <!-- <div class="test__item-question">
                    <label class="b-contain">
                        <span>Его рейтинг и отзывы в интернете</span>
                        <input name="radio" type="radio" />
                        <div class="b-input"></div>
                    </label>
                </div> -->
                <div class="test__item-question"
                    v-for="answer of questionFullInfo.answers" :key="answer.id"
                >
                    <label class="b-contain">
                        <!-- <span>Наличие лицензии, выданной Банком России</span> -->
                        <span>{{ answer.body }}</span>
                        <input type="checkbox"
                            v-if="questionFullInfo.is_multiple_answer"
                            :name="questionFullInfo.id"
                            :value="answer.id"
                        />
                        <input type="radio"
                            v-else
                            :name="questionFullInfo.id"
                            :value="answer.id"
                        />
                        <div class="b-input"></div>
                    </label>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.test {

    &__item {
        display: flex;
        column-gap: 70px;
        row-gap: 10px;
        @media (max-width: 767px) {
            flex-wrap: wrap;
        }
        &:not(:last-child) {
            margin-bottom: 50px;
            @media (max-width: 767px) {
                margin-bottom: 30px;
            }
        }
    }

    &__item-content {
        
    }

    &__item-aside {
        @media (max-width: 767px) {
            display: flex;
            gap: 10px;
            align-items: center;
        }
    }

    &__preview {

    }

    &__ball {
        color: #000;
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
        margin-top: 10px;
        @media (max-width: 767px) {
            margin-top: 0px;
        }
    }

    &__item-title {
        margin-bottom: 30px;
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        @media (max-width: 767px) {
            margin-bottom: 10px;
            font-size: 16px;
        }
    }

    &__item-questions-list {
        display: flex;
        flex-direction: column;
        gap: 25px;
        @media (max-width: 767px) {
            gap: 10px;
        }
    }

    &__item-question {

    }
}

[dark=true] {
    & .test__ball {
        color: #fff;
    }
}
</style>
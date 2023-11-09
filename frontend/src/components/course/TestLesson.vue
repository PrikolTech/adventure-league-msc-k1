<script setup>
import TestQuestion from '@/components/course/TestQuestion.vue'
import TheButton from '@/components/layouts/TheButton.vue';

const props = defineProps({
    showTest: {
        type: Boolean,
        required: true
    }
})

</script>

<template>
    <div class="test">
        <div class="test-before"
            v-if="!props.showTest"
        >
            <div class="test__before-title">
                Тест к уроку 1. Введение в финансовую грамотность. 
            </div>
            <div class="test__before-text">
                <p>
                    Тест к уроку 1 содержит в себе весь пройденный материал, обсуждённый в уроке “Введение в финансовую грамотность”. Все вопросы читайте вниматель и не торопитесь. 
                </p>
                <p>
                    <b>
                        У Вас есть 3 попытки для успешного завершения теста. Тест считается пройденным по достижению больше половины правильных ответов. Если оценка теста Вас не устраивает то Вы можете перепройти тест, пока у Вас не закончатся попытки. Желаем удачи!
                    </b>
                </p>
            </div>
            <div class="test__before-footer">
                <b>
                    Попытка 0/3
                </b>
                <the-button class=""
                    :styles="['btn_red']"
                    :type="'button'"
                    @click="$emit('update:showTest', true)"
                >
                    Пройти тест
                </the-button>
            </div>
        </div>
        <div class="test"
            v-else
        >
            <div class="test__header">
                <div class="test__preview">
                    Вопросов 6
                </div>
                <div class="test__title">
                    Тест к уроку 1. Введение в финансовую грамотность. 
                </div>
            </div>
            <div class="test__content">
                <div class="test__list">
                    <test-question
                        v-for="(question, index) of [1,2,3,4]" :key="index" :question="question"
                    />
                </div>
            </div>
            <the-button class="test__btn-send"
                :styles="['btn_blue']"
                :type="'button'"
            >
                Завершить попытку
            </the-button>
        </div>
    </div>
</template>

<style lang="scss" >
.test {
    
    &__btn-send {
        margin-left: 166px;
        @media (max-width: 767px) {
            margin-left: 0px;
        }
    }

    &__header {
        display: flex;
        align-items: center;
        column-gap: 50px;
        margin-bottom: 35px;
        row-gap: 10px;
        @media (max-width: 767px) {
            flex-wrap: wrap-reverse;
        }
    }

    &__preview {
        border-radius: 8px;
        background: var(--blue-50, #EBF5FF);
        color: var(--primary-2-default, #003791);
        font-family: Roboto;
        font-size: 18px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 27px */
        padding: 6px 12px;
        width: fit-content;
        white-space: nowrap;
        @media (max-width: 539px) {
            font-size: 16px;
            padding: 5px 10px;
        }
    }

    &__title {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 30px */
        @media (max-width: 539px) {
            font-size: 16px;
        }
    }

    &__content {
        margin-bottom: 40px;
    }

    &__list {
    }
}

.test-before {
}
.test {

    &__before-title {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 30px */
        margin-bottom: 40px;
    }

    &__before-text {
        margin-bottom: 30px;
        & p {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
            &:not(:last-child) {
                margin-bottom: 20px;
            }
        }
    }

    &__before-footer {
        & b {
            color: #000;
            font-family: Roboto;
            font-size: 18px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 27px */
            display: block;
            margin-bottom: 15px;
        }
        & .btn {
            width: 100%;
        }
    }

    &__btn-send {
    }
}


[dark=true] {
    & .test__preview {

    }
}

</style>
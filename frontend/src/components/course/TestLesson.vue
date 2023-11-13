<script setup>
import TestQuestion from '@/components/course/TestQuestion.vue';
import TheButton from '@/components/layouts/TheButton.vue';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useUser } from '@/stores/user';

const userStore = useUser()
const route = useRoute()
const props = defineProps({
    showTest: {
        type: Boolean,
        required: true
    },
    test: {
        type: Object,
        required: true
    },
    lesson: {
        type: Object,
        required: true
    }
})

const emit = defineEmits(['update:showTest'])

const questions = ref({})

const getTestInfo = async() => {
    emit('update:showTest', true)
    let testID = null
    if(Object.keys(props.test).length > 0) {
        testID = props.test.id
    } else {
        testID = route.query.test
    }
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${props.lesson.id}/tests/${testID}`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()
        questions.value = { ...data }
        console.log('Тест:', data)
        
    } catch(err) {
        console.error(err)
    }
}
let form = ref(null)
const completeTest = async(e) => {
    const formData = new FormData(e.target);

    let testID = null
    if(Object.keys(props.test).length > 0) {
        testID = props.test.id
    } else {
        testID = route.query.test
    }
    const answers = []
    for (let [key] of formData.entries()) {
        answers.push(
            {
                'answer_id': key
            }
        )
    }
    console.log(answers)
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${props.lesson.id}/tests/${testID}/test_solutions`, {
            method: "POST",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            body: JSON.stringify({
                user_id: userStore.user.id,
                answers
            }),
            mode: 'cors',
        })
        console.log(response)
        const data = await response.json()
        console.log(data)
        
    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    // getTestInfo()
})

</script>

<template>
    <div class="test">
        <div class="test-before"
            v-if="!props.showTest"
        >
            <div class="test__before-title">
                <!-- Тест к уроку 1. Введение в финансовую грамотность.  -->
                {{ test.name }}
            </div>
            <div class="test__before-text" v-if="false">
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
                <!-- <b>
                    Попытка 0/3
                </b> -->
                <the-button class=""
                    :styles="['btn_red']"
                    :type="'button'"
                    @click="getTestInfo()"
                >
                    Пройти тест
                </the-button>
            </div>
        </div>
        <form class="test"
            ref="form"
            @submit.prevent="(event) => completeTest(event)"
            v-else
        >
            <div class="test__header">
                <div class="test__preview">
                    <!-- Вопросов 6 -->
                    Вопросов {{ questions.length }}
                </div>
                <div class="test__title">
                    {{ questions.name }}
                    <!-- Тест к уроку 1. Введение в финансовую грамотность.  -->
                </div>
            </div>
            <div class="test__content">
                <div class="test__list">
                    <test-question
                        v-for="(question, index) of questions.questions" :key="question.id" :question="question" :number="index" :testID="props.test.id" :lessonID="props.lesson.id"
                    />
                </div>
            </div>
            <the-button class="test__btn-send"
                :styles="['btn_blue']"
                :type="'button'"
            >
                Завершить попытку
            </the-button>
        </form>
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
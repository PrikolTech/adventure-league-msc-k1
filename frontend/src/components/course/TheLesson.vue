<script setup>
import { computed, ref } from 'vue';
import TheComment from '@/components/layouts/TheComment.vue';
import MakeComment from '@/components/layouts/MakeComment.vue';

const props = defineProps({
    lesson: {
        type: Object,
        required: true
    }
})

const quantityComments = computed(() => {
    return props.lesson.comments.length
})

let commentInput = ref('')

</script>

<template>
    <div class="lesson">
        <div class="lesson__title">
            <!-- Урок 1. Введение в финансовую грамотность. -->
            {{ props.lesson.name }}
        </div>
        <div class="lesson__material"
            v-if="props.lesson.video"
        >
            <iframe width="100%" height="315"
            :src="props.lesson.video" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
        </div>
        <div class="lesson__desc">
            <div class="lesson__desc-header">
                <p>
                    Описание
                </p>
                <div class="lesson__desc-header-line line"></div>
            </div>
            <p class="lesson__desc-text">
                <!-- В этом уроке вы познакомитесь с основами и понятиями финансовой грамотности. Более подробно разберете важность изучения финансовой грамотности и многое другое. -->
                {{ props.lesson.desc }}
            </p>
        </div>
        <div class="lesson__comments comments">
            <div class="lesson__comments-header comments-header">
                <p>
                    <!-- Комментарии - 3 -->
                    Комментарии - {{ quantityComments }}
                </p>
                <div class="lesson__comments-header-line line"></div>
            </div>
            <div class="lesson__comments-content comments-content">
                <make-comment
                    v-model="commentInput"
                    @update:modelValue="(modelValue) => commentInput=modelValue"
                />
                <div class="comments__list">
                    <the-comment
                        v-for="(comment, index) of props.lesson.comments" :key="index"
                        :comment="comment"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss">
.lesson {

    &__title {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 30px */
        margin-bottom: 30px;
    }

    &__material {
        margin-bottom: 50px;
    }

    &__desc {
        margin-bottom: 50px;
    }

    &__desc-header {
        display: flex;
        align-items: center;
        margin-bottom: 30px;
        gap: 50px;
        & p {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 20px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 30px */
        }
    }

    &__desc-header-line {
        flex: 1;
    }

    &__desc-text {
        color: #000;
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
    }

    &__comments {
    }

    &__comments-header {

    }

    &__comments-header-line {
        flex: 1;
    }

    &__comments-content {
    }
}
.line {
}
.comments {

    &__list {
    }

    &__item {
        &:not(:last-child) {
            margin-bottom: 30px;
        }
    }

    &__item-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 20px;
        flex-wrap: wrap;
        margin-bottom: 20px;
    }

    &__item-info {
        display: flex;
        align-items: center;
        gap: 15px;
        flex-wrap: wrap;
        & p {
            color: #000;
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 500;
            line-height: 150%; /* 24px */
        }
        & span {
            color: var(--blue-600, #1C64F2);
            font-family: Roboto;
            font-size: 12px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 18px */
            border-radius: 8px;
            background: var(--blue-100, #E1EFFE);
            padding: 4px 8px;
        }
    }

    &__item-date {
        color: var(--gray-400, #9CA3AF);
        font-family: Roboto;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        line-height: 150%; /* 21px */
    }

    &__item-text {
        color: #000;
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
    }
}
.comments-header {
    display: flex;
    align-items: center;
    margin-bottom: 30px;
    gap: 50px;
    & p {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 30px */
    }
}
.comments-content {
}


</style>
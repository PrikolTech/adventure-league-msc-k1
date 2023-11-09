<script setup>
import TheAvatar from '@/components/layouts/TheAvatar.vue';
import { useUser } from '@/stores/user'
import TheButton from '@/components/layouts/TheButton.vue';
const userStore = useUser()

const props = defineProps({
    modelValue: String
})


</script>

<template>
    <div class="make__comment">
        <div class="make__comment-preview">
            <the-avatar
                :first_name="userStore.user.first_name"
                :last_name="userStore.user.last_name"
            />
            <input
                placeholder="Ваш комментарий"
                :value="props.modelValue"
                @input="(e) => $emit('update:modelValue', e.target.value)"
            />
        </div>
        <div class="make__comment-hidden"
            v-if="props.modelValue"
        >
            <the-button
                :styles="['btn_red']"
                :type="'button'"
            >
                Отправить
            </the-button>
            <the-button
                :styles="['btn_grey']"
                :type="'button'"
                @click="$emit('update:modelValue', '')"
            >
                <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 12 12" fill="none">
                    <path d="M1 11L11 1L1 11ZM1 1L11 11L1 1Z" fill="#9CA3AF"/>
                    <path d="M1 1L11 11M1 11L11 1L1 11Z" stroke="#9CA3AF" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
            </the-button>
        </div>
    </div>
</template>

<style lang="scss">
.make {
    &__comment {
        margin-bottom: 40px;
        @media (max-width: 1023px) {
            margin-bottom: 20px;
        }
    }
    &__comment-preview {
        display: flex;
        align-items: center;
        gap: 15px;
        margin-bottom: 10px;
        & input {
            border-radius: var(--rounded-xl, 12px);
            border: 1px solid var(--backgrounds-day-gray-background-for-dies, #E1E1E9);
            background: var(--elements-day-text-icons-on-substrates, #FFF);
            flex: 1;
            color: var(--gray-400, #9CA3AF);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
    }
    &__comment-hidden {
        display: flex;
        gap: 15px;
        & .btn_red {
            flex: 1;
        }
    }
}

[dark=true] {
    & .make__comment-preview {
        & input {
            color: var(--gray-500, #6B7280);
            border: 1px solid var(--gray-700, #374151);
            background: var(--gray-700, #374151);
        }
    }
}
</style>
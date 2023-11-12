<script setup>
import TheAvatar from '@/components/layouts/TheAvatar.vue';
import { computed } from 'vue';
import { useUser } from '@/stores/user'
const userStore = useUser()

const props = defineProps({
    comment: {
        type: Object,
        required: true,
    }
})

const dateComment = computed(() => {
    const dateObj = new Date(props.comment.created_at);
    const options = { year: 'numeric', month: '2-digit', day: '2-digit' };
    return dateObj.toLocaleString('ru', options);
});

const timeComment = computed(() => {
    const dateObj = new Date(props.comment.created_at);
    const options = { hour: '2-digit', minute: '2-digit', second: '2-digit' };
    return dateObj.toLocaleTimeString('ru', options);
});
</script>

<template>
    <div class="comments__item">
        <div class="comments__item-header">
            <div class="comments__item-info">
                <the-avatar
                    v-if="userStore.user.id === props.comment.user_id"
                    :first_name="userStore.user.first_name"
                    :last_name="userStore.user.last_name"
                />
                <the-avatar
                    v-else-if="userStore.user.user_id === props.comment.user_id"
                    :first_name="props.comment.first_name"
                    :last_name="props.comment.last_name"
                />
                <p>
                    <!-- Вера Красулина -->
                    {{ props.comment.first_name }} {{ props.comment.last_name }}
                </p>
                <span v-if="userStore.user.user_id !== props.comment.user_id">
                    Преподаватель
                </span>
                <span v-if="userStore.checkRole('student') && userStore.user.user_id === props.comment.user_id">
                    Студент
                </span>
            </div>
            <p class="comments__item-date">
                <!-- 25 августа 2023 в 21:35 -->
                {{ dateComment}} в {{ timeComment }}
            </p>
        </div>
        <p class="comments__item-text">
            <!-- Всем привет, подскажите что задали. -->
            {{ props.comment.body }} 
        </p>
    </div>
</template>

<style lang="scss" scoped>

</style>
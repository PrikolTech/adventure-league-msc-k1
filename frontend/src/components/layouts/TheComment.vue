<script setup>
import TheAvatar from '@/components/layouts/TheAvatar.vue';
import { computed, onMounted, ref } from 'vue';
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
const userInitials = ref({})
const getCommnetInfo = async() => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_USER_URL}/user/${props.comment.user_id}`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`,
            },
            mode: 'cors',
            credentials: 'include'
        });
        const data = await response.json()
        userInitials.value = {
            ...data
        }
    } catch(err) {
        console.error(err)
    }
}

onMounted(() => {
    getCommnetInfo()
})
</script>

<template>
    <div class="comments__item">
        <div class="comments__item-header">
            <div class="comments__item-info">
                <the-avatar
                    :first_name="userInitials.first_name"
                    :last_name="userInitials.last_name"
                />
                <!-- <the-avatar
                    v-else-if="userStore.user.user_id === props.comment.user_id"
                    :first_name="props.comment.first_name"
                    :last_name="props.comment.last_name"
                /> -->
                <p>
                    <!-- Вера Красулина -->
                    {{ props.comment.first_name }} {{ props.comment.last_name }}
                </p>
                <span v-if="userStore.user.user_id === props.comment.user_id">
                    {{ userStore.user.roles[0].description }}
                </span>
                <span v-else-if="userStore.checkRole('student')">
                    Преподаватель
                </span>
                <span v-else-if="userStore.checkRole('teacher')">
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
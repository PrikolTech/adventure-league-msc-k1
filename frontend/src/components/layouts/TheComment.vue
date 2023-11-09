<script setup>
import TheAvatar from '@/components/layouts/TheAvatar.vue';

const props = defineProps({
    comment: {
        type: Object,
        required: true,
    }
})
</script>

<template>
    <div class="comments__item">
        <div class="comments__item-header">
            <div class="comments__item-info">
                <the-avatar
                    :first_name="props.comment.first_name"
                    :last_name="props.comment.last_name"
                />
                <p>
                    <!-- Вера Красулина -->
                    {{ props.comment.first_name }} {{ props.comment.last_name }}
                </p>
                <span>
                    <!-- Студент -->
                    {{ props.comment.type }}
                </span>
            </div>
            <p class="comments__item-date">
                <!-- 25 августа 2023 в 21:35 -->
                {{ props.comment.date }} в {{ props.comment.time }} 
            </p>
        </div>
        <p class="comments__item-text">
            <!-- Всем привет, подскажите что задали. -->
            {{ props.comment.text }} 
        </p>
    </div>
</template>

<style lang="scss" scoped>

</style>
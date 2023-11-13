<script setup>
import { ref } from "vue";


let comments = ref([
    {
        title: "Подготовка к старту проекта",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
    },
    {
        title: "Подготовка к старту проекта",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
    },
    {
        title: "Подготовка к старту проекта",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
    },
        {
        title: "Подготовка к старту проекта",
        desc: "Курс : Agile, Scruma",
        date: "25 августа 2023",
        text: "Наталья, отличная работа! Молодец)",
    },
])

</script>

<template>
    <div class="me__marks">
        <div class="me__marks-header">

        </div>
        <div class="me__marks-list scroll-custom">
            <div class="me__marks-item"
                v-for="(comment, index) of comments" :key="index"
            >
                <p class="me__marks-item-title">
                    <!-- Подготовка к старту проекта -->
                        {{ comment.title }}
                    <span>
                        <!-- 25 августа 2023 -->
                        {{ comment.date }}
                    </span>
                </p>
                <p class="me__marks-item-course">
                    <!-- Курс : Agile, Scrum.  -->
                    {{ comment.desc }}
                </p>
                <div class="me__marks-item-comment">
                    <textarea class="comment" readonly v-model="comment.text"></textarea>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>

</style>
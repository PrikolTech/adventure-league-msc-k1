<script setup>

</script>

<template>
    <div class="me__marks">
        <div class="me__marks-header">

        </div>
        <div class="me__marks-list">
            <div class="me__marks-item"
                v-for="(mark, index) of [1,2,3]" :key="index"
            >
                <p class="me__marks-item-title">
                    Подготовка к старту проекта
                    <span>
                        25 августа 2023
                    </span>
                </p>
                <p class="me__marks-item-course">
                    Курс : Agile, Scrum. 
                </p>
                <div class="me__marks-item-comment">
                    <textarea class="comment" readonly>Наталья, отличная работа! Молодец)</textarea>
                </div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>

</style>
<script setup>
import TheModal from "../layouts/TheModal.vue";
const namePopup = 'createLesson'
import { ref } from 'vue';

let nameLesson = ref('Лекция по созданию фронта')
let descLesson = ref('Лекция по созданию фронта описание')
let dateLesson = ref('2023-08-10')
let timeLesson = ref('12:00')


</script>

<template>
    <the-modal
        :name="namePopup"
    >
        <template v-slot:title>
            Создание занятия
        </template>
        <div class="field">
            <p>
                Название
            </p>
            <div class="input-w">
                <input
                    v-model="nameLesson"
                >
            </div>
        </div>
        <div class="field">
            <p>
                Описание занятия
            </p>
            <div class="input-w">
                <textarea v-model="descLesson"></textarea>
            </div>
        </div>
        <div class="field">
            <p>
                Дата
            </p>
            <div class="input-w">
                <input
                    type="date"
                    v-model="dateLesson"
                >
            </div>
        </div>
        <div class="field">
            <div class="input-w">
                <input
                    type="time"
                    v-model="timeLesson"
                >
            </div>
        </div>
        <template v-slot:btnSend>
            Создать
        </template>
    </the-modal>
</template>

<style lang="scss" scoped>

</style>
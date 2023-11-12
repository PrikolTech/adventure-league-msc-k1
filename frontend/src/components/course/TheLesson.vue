<script setup>
import { computed, onMounted, ref } from 'vue';
import TheComment from '@/components/layouts/TheComment.vue';
import MakeComment from '@/components/layouts/MakeComment.vue';
import TheButton from '../layouts/TheButton.vue';
import { useUser } from '@/stores/user'
import UploadFile from '../layouts/UploadFile.vue';

const userStore = useUser()
const props = defineProps({
    lesson: {
        type: Object,
        required: true
    },
    course: {
        type: Object,
        required: true
    }
})

const quantityComments = computed(() => {
    if(props.lesson.comments) {
        return props.lesson.comments.length
    } else {
        return 0
    }
})

let fileInput = ref(null)
let fileName = ref('')
const changeFile = (event) => {
    try {
        const selectedFile = event.target.files[0];
        fileInput.value = selectedFile
        fileName.value = selectedFile.name;
        console.log("Selected file name: " + fileName.value);
    } catch (err) {
        console.log(err);
    }
};

let commentInput = ref('')
const comments = ref([])
const getComments = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COMMENT_URL}/api/comments/lecture/${props.lesson.id}`, {
            method: "GET",
        })

        const data = await response.json()
        console.log('комментарии',data)
        comments.value = [...data]
    } catch (err) {
        console.log(err);
    }
}

const postComment = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COMMENT_URL}/api/comments/lecture/`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                user_id: userStore.user.id,
                body: commentInput.value,
                target_id: props.lesson.id
            })
        });

        const data = await response.json()
        console.log('новый комментарий',data)
        commentInput.value = ''
        if(data) {
            comments.value.push(data)
        }
    } catch (err) {
        console.log(err);
    }
}

const deleteFileFromInput = () => {
    fileInput.value = null
    fileName.value = ''
}

const postFile = async () => {
    const file = fileInput.value;

    try {
        const formData = new FormData();
        formData.append('file', file);
        console.log('test')
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${props.course.id}/lectures/${props.lesson.id}/contents`, {
            method: "POST",
            body: formData
        });

        console.log('testststst',response);
    } catch (err) {
        console.log(err);
    }
}

const updateLecture = async () => {
    try {
        console.log('test');
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COURSE_URL}/courses/${props.course.id}/lectures/${props.lesson.id}`, {
            method: "PUT",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                description: props.lesson.description
            })
        });

        const data = await response.json(); // Need to await the JSON parsing

        console.log('test', data);
    } catch (err) {
        console.log(err);
    }
}


onMounted(() => {
    getComments()
})
</script>

<template>
    <div class="lesson">
        <div class="lesson__title">
            <!-- Урок 1. Введение в финансовую грамотность. -->
            {{ props.lesson.name }}
        </div>
        <div class="lesson__material"
            v-if="true"
        >
            <iframe width="100%" height="315"
            :src="'https://www.youtube.com/embed/__-vp0g_BhA?si=ou5xfV8arD_ccw6Q'" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>
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
                <span v-if="userStore.user.role === 'student'">
                    {{ props.lesson.description }}
                </span>
                <textarea v-model="props.lesson.description" v-if="userStore.user.role === 'teacher'"></textarea>
                <the-button
                    v-if="userStore.user.role === 'teacher'"
                    :styles="['btn_red']"
                    :type="'button'"
                    class="add-file-btn"
                    @click="updateLecture"
                >
                    Сохранить
                </the-button>
            </p>
        </div>
        <!-- <add-file
            :test="fileInput"
        /> -->
        <div class="add-file-w" v-if="userStore.user.role === 'teacher'">
            <label class="add-file">
                <input
                    type="file"
                    @change="((event) => changeFile(event))"
                />
                <svg class="add-file-icon" xmlns="http://www.w3.org/2000/svg" width="40" height="41" viewBox="0 0 40 41" fill="none">
                    <path d="M27.8457 11.1729L21.179 4.33552C21.0242 4.17634 20.8402 4.05004 20.6378 3.96387C20.4353 3.8777 20.2182 3.83334 19.999 3.83334C19.7798 3.83334 19.5627 3.8777 19.3602 3.96387C19.1577 4.05004 18.9738 4.17634 18.819 4.33552L12.1523 11.1729C11.8398 11.4938 11.6644 11.929 11.6647 12.3826C11.665 12.8362 11.841 13.271 12.154 13.5916C12.4669 13.9121 12.8912 14.092 13.3335 14.0916C13.7758 14.0913 14.1998 13.9108 14.5123 13.5899L18.334 9.67035V24.3467C18.334 24.8 18.5096 25.2348 18.8221 25.5554C19.1347 25.8759 19.5586 26.056 20.0007 26.056C20.4427 26.056 20.8666 25.8759 21.1792 25.5554C21.4917 25.2348 21.6673 24.8 21.6673 24.3467V9.67035L25.489 13.5899C25.8033 13.9012 26.2243 14.0735 26.6613 14.0696C27.0983 14.0657 27.5163 13.886 27.8253 13.569C28.1344 13.2521 28.3096 12.8234 28.3134 12.3752C28.3172 11.927 28.1493 11.4952 27.8457 11.1729Z" fill="#D1D5DB"/>
                    <path d="M33.334 23.492H25.0007V24.3467C25.0007 25.7067 24.4739 27.011 23.5362 27.9727C22.5985 28.9344 21.3267 29.4747 20.0007 29.4747C18.6746 29.4747 17.4028 28.9344 16.4651 27.9727C15.5274 27.011 15.0007 25.7067 15.0007 24.3467V23.492H6.66732C5.78326 23.492 4.93542 23.8522 4.31029 24.4933C3.68517 25.1344 3.33398 26.004 3.33398 26.9107V33.748C3.33398 34.6547 3.68517 35.5243 4.31029 36.1654C4.93542 36.8065 5.78326 37.1667 6.66732 37.1667H33.334C34.218 37.1667 35.0659 36.8065 35.691 36.1654C36.3161 35.5243 36.6673 34.6547 36.6673 33.748V26.9107C36.6673 26.004 36.3161 25.1344 35.691 24.4933C35.0659 23.8522 34.218 23.492 33.334 23.492ZM29.1673 33.748C28.6729 33.748 28.1895 33.5976 27.7784 33.3159C27.3673 33.0342 27.0468 32.6337 26.8576 32.1652C26.6684 31.6967 26.6189 31.1812 26.7154 30.6838C26.8118 30.1864 27.0499 29.7296 27.3996 29.371C27.7492 29.0124 28.1946 28.7682 28.6796 28.6693C29.1645 28.5704 29.6672 28.6211 30.124 28.8152C30.5808 29.0093 30.9713 29.3379 31.246 29.7595C31.5207 30.1812 31.6673 30.6769 31.6673 31.184C31.6673 31.864 31.4039 32.5162 30.9351 32.997C30.4662 33.4779 29.8304 33.748 29.1673 33.748Z" fill="#D1D5DB"/>
                </svg>
                <p>
                    Кликните для загрузки или перетащите и отпустите
                </p>
                <span>
                    DOCX, PDF, PNG, JPG or PPTX (до 2 МБ)
                </span>
            </label>
            <the-button
                v-if="fileInput"
                :styles="['btn_red']"
                :type="'button'"
                class="add-file-btn"
                @click="postFile()"
            >
                Добавить
            </the-button>
        </div>
        <upload-file
            :name="fileName"
            @deleteFile="deleteFileFromInput()"
        />
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
                    @send="postComment()"
                />
                <div class="comments__list">
                    <the-comment
                        v-for="(comment, index) of comments" :key="index"
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
        @media (max-width: 539px) {
            margin-bottom: 20px;
        }
    }

    &__desc {
        margin-bottom: 50px;
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
    }

    &__desc-header {
        display: flex;
        align-items: center;
        margin-bottom: 30px;
        gap: 50px;
        @media (max-width: 539px) {
            margin-bottom: 10px;
        }
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
        & textarea {
            width: 100%;
            resize: none;
            height: 70px;
            border-radius: var(--rounded-xl, 12px);
            border: 1px solid var(--backgrounds-day-gray-background-for-dies, #E1E1E9);
            background: var(--elements-day-text-icons-on-substrates, #FFF);
            flex: 1;
            color: var(--gray-400, #9CA3AF);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%;
            padding: 7px 16px;
        }
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
        column-gap: 20px;
        row-gap: 5px;
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
        padding-left: 10px;
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

[dark=true] {
    & .lesson__title {
        color: var(--gray-50, #F9FAFB);
    }

    & .lesson__desc-header {
        & p {
            color: var(--gray-50, #F9FAFB);
        }
    }

    & .lesson__desc-text {
        color: var(--gray-50, #F9FAFB);
    }
    
    & .comments-header p {
        color: var(--gray-50, #F9FAFB);
    }

    & .comments__item-info {
        & p {
            color: var(--gray-100, #F3F4F6);
        }
        
        & span {
            color: var(--blue-200, #C3DDFD);
            background: var(--blue-900, #233876);
        }
    }
    
    & .comments__item-date {
        color: var(--gray-500, #6B7280);
    }

    & .comments__item-text {
        color: var(--gray-100, #F3F4F6);
    }
}


</style>
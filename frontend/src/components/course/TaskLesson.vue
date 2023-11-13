<script setup>
import MakeComment from '@/components/layouts/MakeComment.vue';
import TheButton from '../layouts/TheButton.vue';
import UploadFile from "../layouts/UploadFile.vue";
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useUser } from '@/stores/user';
import TheComment from '@/components/layouts/TheComment.vue';

const userStore = useUser()
const route = useRoute()
const props = defineProps({
    lesson: {
        type: Object,
    },
    task: {
        type: Object,
    }
})

let commentInput = ref('')
// let input = ref(null)
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

const deleteFile = () => {
    fileInput.value = null
    fileName.value = ''
}
const taskInfo = ref({})
const getTaskInfo = async() => {
    let homeWorkID = null
    if(props.task) {
        if(Object.keys(props.task).length > 0) {
            homeWorkID = props.task.id
        }
    } 
    if(!homeWorkID) homeWorkID = route.query.task

    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${props.lesson.id}/homeworks/${homeWorkID}`, {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${userStore.user.access}`
            },
            mode: 'cors',
        })
        
        const data = await response.json()
        console.log('Домашка:', data)
        taskInfo.value = { ...data }
    } catch(err) {
        console.error(err)
    }
}

const comments = ref([])
const getComments = async () => {
    console.log(route.query.lesson)
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COMMENT_URL}/comments/homeworks/${route.query.task}`, {
            method: "GET",
        })

        const data = await response.json()
        console.log('комментарии',data)
        comments.value.length = 0
        comments.value = [...data]
    } catch (err) {
        console.log(err);
    }
}

const postComment = async () => {
    try {
        const response = await fetch(`${import.meta.env.VITE_SERVICE_COMMENT_URL}/comments/homeworks/`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${userStore.user.access}`,
            },
            mode: 'cors',
            body: JSON.stringify({
                user_id: userStore.user.id,
                body: commentInput.value,
                target_id: route.query.task
            })
        });
        console.log(response.ok)
        if(response.ok) {
            commentInput.value = ''
            const data = await response.json()
            console.log('новый комментарий',data)
            comments.value.push(data)
        }
    } catch (err) {
        console.log(err);
    }
}


// const sendFile = async () => {
//     try {
//         const response = await fetch(`${import.meta.env.VITE_SERVICE_JOB_URL}/jobs/${job_id}/homeworks/${homework_id}/homework_solutions `, {
//             method: "POST",
//             headers: {
//                 'Content-Type': 'application/json'
//             },
//             body: JSON.stringify({
//                 user_id: userStore.user.id,
//                 body: commentInput.value,
//                 target_id: props.lesson.id
//             })
//         });

//         const data = await response.json()
//         console.log('новый комментарий',data)
//         commentInput.value = ''
//         if(data) {
//             comments.value.push(data)
//         }
//     } catch (err) {
//         console.log(err);
//     }
// }

onMounted(() => {
    getTaskInfo()
    getComments()
})
</script>

<template>
    <div class="lesson__task">
        <div class="lesson__task-title">
            <!-- Задание к уроку 1. Введение в финансовую грамотность.  -->
            {{ taskInfo.name }}
        </div>
        <div class="lesson__task-text">
            {{ taskInfo.description }}
        </div>
        <div class="add-file-w"
            v-if="userStore.checkRole('student')"
        >
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
                :styles="['btn_red']"
                :type="'button'"
                class="add-file-btn"
                @click="sendFile()"
            >
                Отправить
            </the-button>
        </div>
        <upload-file
            :name="fileName"
            @deleteFile="deleteFile()"
            v-if="fileName"
        />
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
</template>

<style lang="scss" scoped>
ul.numbers {
    padding-left: 15px;
    & li {
        list-style-type: auto;
    }
}
.lesson {

    &__task {
    }

    &__task-title {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 20px;
        font-style: normal;
        font-weight: 500;
        line-height: 150%; /* 30px */
        margin-bottom: 40px;
        @media (max-width: 1023px) {
            margin-bottom: 10px;
        }
    }

    &__task-text {
        color: var(--gray-900, #111928);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
        margin-bottom: 60px;
        @media (max-width: 1023px) {
            margin-bottom: 20px;
        }
    }
}

.add-file-icon {
}

.file-info {
    margin-bottom: 40px;
    &__header {
        display: flex;
        justify-content: space-between;
        gap: 20px;
        margin-bottom: 20px;
        & p {
            color: var(--gray-900, #111928);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
        & span {
            color: var(--gray-400, #9CA3AF);
            font-family: Roboto;
            font-size: 14px;
            font-style: normal;
            font-weight: 400;
            line-height: 150%; /* 21px */
        }
    }

    &__item {
        border-radius: var(--rounded-xl, 12px);
        background: var(--gray-100, #F3F4F6);
        padding: 12px 20px;
        display: flex;
        justify-content: space-between;
        gap: 20px;
        color: var(--gray-600, #4B5563);
        font-family: Roboto;
        font-size: 16px;
        font-style: normal;
        font-weight: 400;
        line-height: 167%; /* 26.72px */
        & svg {
            cursor: pointer;
        }
    }
}

[dark=true] {
    & .lesson__task-title {
        color: var(--gray-50, #F9FAFB);
    }

    & .lesson__task-text {
        color: var(--gray-50, #F9FAFB);
    }


}

</style>
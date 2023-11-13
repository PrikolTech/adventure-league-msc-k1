<script setup>
import PreviewSwiper from '../components/PreviewSwiper.vue';
import TheProgram from '../components/program/TheProgram.vue';
import ThePopup from '@/components/layouts/ThePopup.vue'
import { onMounted, ref } from 'vue';
import { useUser } from '@/stores/user'
import { useAlerts } from '@/stores/alerts'

const alertsStore = useAlerts()
const userStore = useUser()
const currentOpenCourse = ref(null)


const formData = ref({
  initiator_first_name: 'Сергей',
  initiator_last_name: 'Македонский',
  supervisor_first_name: 'Петр',
  supervisor_last_name: 'Волков',
  email: 'marketingdirect24@gmail.com',
  department: 'Невероятное',
  post: 'Директор',
  history: 'учился всю жизнь',
  achievements: 'Поднял компанию с коле',
  motivation: 'я готов пройти курс и стать еще круче',
  birthdate: '2002-11-11',
})


let popupData = ref([
    {
        name: 'Фамилия руководителя *',
        placeHolder: 'Иванов',
        type: 'input',
        typeValue: 'text',
        nameInput: 'supervisor_last_name',
    },
    {
        name: 'Имя руководителя *',
        placeHolder: 'Иван',
        type: 'input',
        typeValue: 'text',
        nameInput: 'supervisor_first_name',
    },
    {
        name: 'Подразделение *',
        placeHolder: 'Placeholder',
        type: 'input',
        typeValue: 'text',
        nameInput: 'department',
    },
    {
        name: 'Должность *',
        placeHolder: 'Менеджер',
        type: 'input',
        typeValue: 'text',
        nameInput: 'post',
    },
    {
        name: 'Ваша история',
        placeHolder: 'Напишите ваши достижения в компании...',
        type: 'textarea',
        typeValue: 'text',
        nameInput: 'history',
    },
    {
        name: 'Личные достижения в компании * (за последние 12 месяцев)',
        placeHolder: 'Напишите ваши достижения в компании...',
        type: 'textarea',
        typeValue: 'text',
        nameInput: 'achievements',
    },
    {
        name: 'Мотивационное письмо *',
        placeHolder: 'Напишите почему Вы хотите учиться...',
        type: 'textarea',
        typeValue: 'text',
        nameInput: 'motivation',
    },
])

const sendForm = async () => {
  let fieldsAreValid = true; // Флаг для проверки пустых полей

  const dateString = formData.value.birthdate
  const date = new Date(dateString + "T00:00:00Z")

  const formattedDate = date.toISOString()

  const preparedFormData = {
    initiator: {
      first_name: formData.value.initiator_first_name,
      last_name: formData.value.initiator_last_name
    },
    supervisor: {
      first_name: formData.value.supervisor_first_name,
      last_name: formData.value.supervisor_last_name
    },
    email: formData.value.email,
    birthdate: formattedDate,
    department: formData.value.department,
    post: formData.value.post,
    history: formData.value.history,
    achievements: formData.value.achievements,
    motivation: formData.value.motivation,
    course_id: currentOpenCourse.value,
  };

  if(userStore.user) {
    console.log('bad')
    preparedFormData.birthdate = userStore.user.birthdate
    preparedFormData.initiator.first_name = userStore.user.first_name
    preparedFormData.initiator.last_name = userStore.user.last_name
    preparedFormData.email = userStore.user.email
    preparedFormData.user_id = userStore.user.user_id
  }

  Object.values(formData.value).forEach(value => {
    if (!value) {
      alertsStore.addAlert('Все поля должны быть заполнены', 'error');
      fieldsAreValid = false;
      return
    }
  });

  if(!fieldsAreValid) return
  try {
    let url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration/append`
    let headers = {
      'Content-Type': 'application/json'
    };

    if(!userStore.user) {
      url = `${import.meta.env.VITE_SERVICE_FORM_URL}/form/registration`
    } else {
      headers['Authorization'] = `Bearer ${userStore.user.access}`;
    }
    const response = await fetch(url, {
      method: 'POST',
      headers,
      mode: 'cors',
      body: JSON.stringify(preparedFormData)
    });

    if(response.ok) {
      alertsStore.addAlert('Вы успешно зарегистрировались на курс', 'success')
    } else {
      alertsStore.addAlert('Произошла ошибка при запись на курс, повторите попытку', 'error')
    }
    console.log(response);
    const data = await response.json();
    console.log(data);
  } catch (err) {
    alertsStore.addAlert('Произошла ошибка при запись на курс, повторите попытку', 'error')
    console.error(err);
  }
};


onMounted(() => {
  if(!userStore.user) {
    popupData.value.unshift(
      {
        name: 'Фамилия *',
        placeHolder: 'Иванов',
        type: 'input',
        typeValue: 'text',
        nameInput: 'initiator_last_name',
      },
      {
        name: 'Имя *',
        placeHolder: 'Иван',
        type: 'input',
        typeValue: 'text',
        nameInput: 'initiator_first_name',
      },
      {
        name: 'Email *',
        placeHolder: 'test@mail.ru',
        type: 'input',
        typeValue: 'text',
        nameInput: 'email',
      },
      {
        name: 'Дата рождения *',
        placeHolder: '',
        type: 'input',
        typeValue: 'date',
        nameInput: 'birthdate',
      },
    )
  }
})

</script>


<template>
    <main>
      <preview-swiper/>
      <the-program
        @openCoursePopup="(id) => currentOpenCourse = id"
      />
      <the-popup
        :data="popupData"
        v-model:formData="formData"
        @sendForm="sendForm()"
      />
      
    </main>
</template>

<style lang="scss">

</style>
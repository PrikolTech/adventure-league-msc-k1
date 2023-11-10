<script setup>
import alertsList from '@/components/AlertsList.vue'
import TheHeader from '@/components/layouts/TheHeader.vue';
import TheFooter from '@/components/layouts/TheFooter.vue';
import ThePopup from '@/components/layouts/ThePopup.vue'
import CreateTaskPopup from './components/course/CreateTaskPopup.vue';
import { useUser } from '@/stores/user'
const user = useUser()
</script>

<template>
  <div class="wrapper" :dark="user.theme">
    <the-header/>
    <RouterView />
    <the-footer/>

    <!-- форма -->
    <the-popup/>
    <create-task-popup/>
    <alerts-list/>
  </div>
</template>

<style scoped>
.wrapper {
  transition: .2s;
  background: var(--var-body);
}
.bg  {
  position: absolute;
  top: 50%;
}

.bg-1 {
  left: 0;
}

.bg-2 {
  right: 0;
}


</style>

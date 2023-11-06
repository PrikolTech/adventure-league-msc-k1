<script setup>
import ProfileNav from '@/components/profile/ProfileNav.vue';

</script>

<template>
    <main>
        <div class="profile container">
            <profile-nav/>
            <RouterView />
        </div>
    </main>
</template>

<style lang="scss" >
.profile {
    padding-top: 5px;
    display: flex;
    gap: 70px;
    &__block {
        flex: 1;
    }
}
.profile {

    &__block-header {
        display: flex;
        gap: 50px;
        align-items: center;
        margin-bottom: 50px;
    }

    &__title {

    }

    &__header-line {
        flex: 1;
    }
}
</style>
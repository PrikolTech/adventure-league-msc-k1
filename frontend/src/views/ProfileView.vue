<script setup>
import ProfileNav from '@/components/profile/ProfileNav.vue';
import TheEvents from '@/components/TheEvents.vue';
</script>

<template>
    <main>
        <div class="profile container">
            <profile-nav/>
            <RouterView />
            <the-events/>
        </div>
    </main>
</template>

<style lang="scss" >
.profile {
    padding-top: 5px;
    display: flex;
    gap: 70px;
    @media (max-width: 1199px) {
        gap: 40px;
    }
    @media (max-width: 1023px) {
        flex-wrap: wrap;
        gap: 40px;
    }
    @media (max-width: 767px) {
        gap: 15px;
    }
    &__block {
        flex: 1;
        @media (max-width: 1023px) {
            flex: 0 0 100%;
        }
    }
}
.profile {

    &__block-header {
        display: flex;
        gap: 50px;
        align-items: center;
        margin-bottom: 50px;
        @media (max-width: 1023px) {
            margin-bottom: 20px;
        }
    }

    &__title {

    }

    &__header-line {
        flex: 1;
    }
}
</style>
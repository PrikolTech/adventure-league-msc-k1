<script setup>
import LogoIcon from '@/components/icons/LogoIcon.vue';
</script>

<template>
    <footer class="footer-w">
        <div class="footer container">
            <router-link class="footer__logo logo" to="/">
                <logo-icon/>
            </router-link>
            <div class="footer__line line"></div>
            <div class="footer__links">
                <p>
                    Контакты:
                </p>
                <div class="footer__links-list">
                    <router-link class="footer__links-item" to="/">
                        <svg xmlns="http://www.w3.org/2000/svg" width="30" height="22" viewBox="0 0 30 22" fill="none">
                            <path d="M29.7012 5C29.7012 5 29.4082 2.93164 28.5059 2.02344C27.3633 0.828125 26.0859 0.822265 25.5 0.751953C21.3047 0.447265 15.0059 0.447266 15.0059 0.447266H14.9941C14.9941 0.447266 8.69531 0.447265 4.5 0.751953C3.91406 0.822265 2.63672 0.828125 1.49414 2.02344C0.591797 2.93164 0.304687 5 0.304687 5C0.304687 5 0 7.43164 0 9.85742V12.1309C0 14.5566 0.298828 16.9883 0.298828 16.9883C0.298828 16.9883 0.591797 19.0566 1.48828 19.9648C2.63086 21.1602 4.13086 21.1191 4.79883 21.248C7.20117 21.4766 15 21.5469 15 21.5469C15 21.5469 21.3047 21.5352 25.5 21.2363C26.0859 21.166 27.3633 21.1602 28.5059 19.9648C29.4082 19.0566 29.7012 16.9883 29.7012 16.9883C29.7012 16.9883 30 14.5625 30 12.1309V9.85742C30 7.43164 29.7012 5 29.7012 5ZM11.9004 14.8906V6.45898L20.0039 10.6895L11.9004 14.8906Z" fill="#1F2A37"/>
                        </svg>
                    </router-link>
                    <router-link class="footer__links-item" to="/">
                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="26" viewBox="0 0 25 26" fill="none">
                            <g clip-path="url(#clip0_1411_41492)">
                                <path fill-rule="evenodd" clip-rule="evenodd" d="M25 13C25 19.9036 19.4036 25.5 12.5 25.5C5.59644 25.5 0 19.9036 0 13C0 6.09644 5.59644 0.5 12.5 0.5C19.4036 0.5 25 6.09644 25 13ZM12.9479 9.72805C11.7321 10.2337 9.30223 11.2804 5.65822 12.868C5.06649 13.1034 4.75652 13.3336 4.7283 13.5586C4.68061 13.9391 5.15699 14.0889 5.8057 14.2928C5.89394 14.3206 5.98537 14.3493 6.0791 14.3798C6.71733 14.5873 7.57586 14.83 8.02217 14.8396C8.42702 14.8484 8.87888 14.6815 9.37775 14.3389C12.7825 12.0406 14.54 10.879 14.6503 10.8539C14.7282 10.8363 14.836 10.814 14.9091 10.879C14.9822 10.944 14.975 11.067 14.9673 11.1C14.9201 11.3012 13.0501 13.0397 12.0824 13.9394C11.7807 14.2198 11.5667 14.4188 11.523 14.4642C11.425 14.566 11.3251 14.6623 11.2291 14.7548C10.6362 15.3264 10.1916 15.755 11.2538 16.455C11.7642 16.7914 12.1727 17.0695 12.5802 17.347C13.0252 17.6501 13.4691 17.9524 14.0433 18.3288C14.1897 18.4247 14.3294 18.5243 14.4655 18.6214C14.9834 18.9906 15.4487 19.3223 16.0235 19.2694C16.3575 19.2387 16.7026 18.9246 16.8778 17.9879C17.2919 15.7741 18.1059 10.9776 18.294 9.00101C18.3105 8.82784 18.2897 8.60622 18.2731 8.50893C18.2564 8.41164 18.2216 8.27303 18.0952 8.17042C17.9454 8.04889 17.7142 8.02327 17.6108 8.02509C17.1406 8.03337 16.4193 8.28418 12.9479 9.72805Z" fill="#1F2A37"/>
                            </g>
                            <defs>
                                <clipPath id="clip0_1411_41492">
                                <rect width="25" height="25" fill="white" transform="translate(0 0.5)"/>
                                </clipPath>
                            </defs>
                        </svg>
                    </router-link>
                    <router-link class="footer__links-item" to="/">
                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="26" viewBox="0 0 25 26" fill="none">
                            <g clip-path="url(#clip0_1411_41493)">
                                <path fill-rule="evenodd" clip-rule="evenodd" d="M1.75736 2.25736C0 4.01472 0 6.84312 0 12.5V13.5C0 19.1569 0 21.9853 1.75736 23.7427C3.51472 25.5 6.34312 25.5 12 25.5H13C18.6569 25.5 21.4853 25.5 23.2427 23.7427C25 21.9853 25 19.1569 25 13.5V12.5C25 6.84312 25 4.01472 23.2427 2.25736C21.4853 0.5 18.6569 0.5 13 0.5H12C6.34312 0.5 3.51472 0.5 1.75736 2.25736ZM4.21881 8.10422C4.35423 14.6042 7.60422 18.5105 13.3021 18.5105H13.6251V14.7917C15.7189 15.0001 17.3021 16.5313 17.9375 18.5105H20.8959C20.0834 15.5521 17.9479 13.9167 16.6145 13.2917C17.9479 12.5209 19.8229 10.6459 20.2708 8.10422H17.5832C16.9999 10.1667 15.2709 12.0417 13.6251 12.2188V8.10422H10.9375V15.3126C9.27083 14.8959 7.16672 12.8751 7.07297 8.10422H4.21881Z" fill="#1F2A37"/>
                            </g>
                            <defs>
                                <clipPath id="clip0_1411_41493">
                                <rect width="25" height="25" fill="white" transform="translate(0 0.5)"/>
                                </clipPath>
                            </defs>
                        </svg>
                    </router-link>
                </div>
            </div>
        </div>
    </footer>
</template>

<style lang="scss" scoped>
.footer-w {
    margin-top: 64px;
    @media (max-width: 1023px) {
        margin-top: 50px;
    }
    @media (max-width: 767px) {
        margin-top: 35px;
    }
    @media (max-width: 539px) {
        margin-top: 20px;
    }
}
.footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 53px;
    padding-top: 45px;
    padding-bottom: 65px;
    @media (max-width: 1023px) {
        padding-top: 30px;
        padding-bottom: 40px;
    }
    @media (max-width: 767px) {
        padding-top: 25px;
        padding-bottom: 30px;
    }
    @media (max-width: 539px) {
        flex-wrap: wrap;
        gap: 15px;
        padding-top: 20px;
        padding-bottom: 20px;
    }
    &__line {
        flex: 1;
        @media (max-width: 539px) {
            display: none;
        }
    }

    &__logo {
        display: flex;
        justify-content: center;
    }

    &__links {
        display: flex;
        gap: 16px;
        align-items: center;
        @media (max-width: 539px) {
            flex: 0 0 100%;
        }
        & p {
            color: var(--var-text);
            font-family: Roboto;
            font-size: 16px;
            font-style: normal;
            font-weight: 400;
            line-height: 167%; /* 26.72px */
        }
    }

    &__links-list {
        display: flex;
        gap: 14px;
        align-items: center;
    }

    &__links-item {

    }
}

[dark=true] {
& .footer__links-item {
    & svg {
        & path {
            fill: #fff;
        }
    }
}
}

</style>
<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  type: {
    type: String,
    required: true
  },
  link: {
    type: String,
    required: false
  },
  styles: {
    type: Array,
    required: false
  }
});
</script>

<template>
  <router-link class="btn" :class="props.styles" :to="props.link" v-if="props.type === 'link'">
      <slot></slot>
  </router-link>
  <button class="btn" :class="props.styles" v-else>
      <slot></slot>
      <svg class="icon-arrow" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path d="M20.9021 12.5341C20.9995 12.2789 21.0249 11.998 20.9753 11.7271C20.9256 11.4562 20.803 11.2073 20.6231 11.012L15.4799 5.42631C15.3613 5.29294 15.2194 5.18656 15.0626 5.11337C14.9057 5.04019 14.737 5.00166 14.5662 5.00005C14.3955 4.99844 14.2262 5.03377 14.0682 5.10399C13.9102 5.1742 13.7666 5.27789 13.6459 5.409C13.5252 5.54012 13.4297 5.69603 13.365 5.86765C13.3004 6.03926 13.2678 6.22315 13.2693 6.40856C13.2708 6.59398 13.3063 6.77722 13.3737 6.94759C13.4411 7.11796 13.539 7.27205 13.6618 7.40086L16.6114 10.6043H4.2858C3.94478 10.6043 3.61773 10.7514 3.3766 11.0133C3.13547 11.2752 3 11.6303 3 12.0007C3 12.3711 3.13547 12.7262 3.3766 12.9881C3.61773 13.25 3.94478 13.3971 4.2858 13.3971H16.6114L13.6631 16.5991C13.5403 16.728 13.4423 16.882 13.375 17.0524C13.3076 17.2228 13.2721 17.406 13.2706 17.5914C13.2691 17.7769 13.3017 17.9607 13.3663 18.1324C13.431 18.304 13.5264 18.4599 13.6472 18.591C13.7679 18.7221 13.9115 18.8258 14.0695 18.896C14.2275 18.9662 14.3968 19.0016 14.5675 18.9999C14.7383 18.9983 14.907 18.9598 15.0639 18.8866C15.2207 18.8134 15.3626 18.7071 15.4812 18.5737L20.6244 12.988C20.7435 12.858 20.8379 12.7038 20.9021 12.5341Z" fill="#4B5563"/>
      </svg>
  </button>
</template>

<style lang="scss">

.btn {
  font-family: Roboto;
  font-size: 16px;
  font-style: normal;
  font-weight: 500;
  line-height: 150%; /* 24px */
  display: flex;
  width: fit-content;
  justify-content: center;
  align-items: center;
  padding: 10px 16px;
  border-radius: var(--rounded-xl, 12px);
  position: relative;
  transition: .2s;
  & .icon-arrow {
    display: none;
    transition: .2s;
  }
  &_arrow {
    & .icon-arrow {
        & path {
          transition: .2s;
        }
        display: block;
    }
  }
  &_red-border {
    color: var(--primary-1-default, #FF4B5F);
    border: 1px solid #FF4B5F;
    &:hover {
      color: var(--primary-1-default, #e43f52);
      border: 1px solid #e43f52;
      background: rgba(0,0,0,0.01);
    }
  }
  &_red {
    background: #FF4B5F;
    color: #fff;
    &:hover {
      background: #e43f52;
    }
  }
  &_blue {
    background: #003791;
    color: #fff;
    &:hover {
      background: #032b6b;
    }
  }
  &_grey {
    color: var(--var-grey-dark-font);
    border: 1px solid var(--gray-400, #9CA3AF);
    &:hover {
      background: rgba(0,0,0,0.05);
    }
  }
  &_grey-lite {
    border-radius: var(--rounded-xl, 12px);
    border: 1px solid var(--gray-300, #D1D5DB);
    color: var(--gray-800, #1F2A37);
    &:hover {
      background: rgba(0,0,0,0.05);
    }
  }
  &_arrow {
    gap: 5px;
  }
}

[dark=true] {
  .btn {
    & .icon-arrow {
        & path {
            fill: #fff;
        }
    }
    &_red-border {
      color: var(--primary-1-default, #FF4B5F);
      border: 1px solid #FF4B5F;
      &:hover {
        color: var(--primary-1-default, #e43f52);
        border: 1px solid #e43f52;
        background: rgba(0,0,0,0.01);
      }
    }
    &_red {
      background: #FF4B5F;
      color: #fff;
      &:hover {
        background: #e43f52;
      }
    }
    &_grey {
      border: 1px solid var(--gray-100, #F3F4F6);
      &:hover {
        background: rgba(0,0,0,0.1);
        border: 1px solid var(--gray-100, #89909e);
        color: #89909e;
      }
    }
    &_arrow {
      gap: 5px;
      &:hover {
        & svg {
          & path {
            fill: #89909e;
          }
        }
      }
    }
  }
}
</style>

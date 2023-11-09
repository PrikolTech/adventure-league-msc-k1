<script>
import TheButton from '@/components/layouts/TheButton.vue';
// import Swiper core and required modules
import { Navigation, Pagination} from 'swiper/modules';

// Import Swiper Vue.js components
import { Swiper, SwiperSlide } from 'swiper/vue';

// Import Swiper styles
import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';

// Import Swiper styles
export default {
  components: {
    Swiper,
    SwiperSlide,
    TheButton
  },
  setup() {
    const onSwiper = (swiper) => {
      console.log(swiper);
    };
    const onSlideChange = () => {
      console.log('slide change');
    };
    return {
      onSwiper,
      onSlideChange,
      modules: [Navigation, Pagination],
    };
  },
};
</script>

<template>
    <div class="preview container">
        <swiper
            :modules="modules"
            :slides-per-view="1"
            :space-between="50"
            navigation
            :pagination="{ clickable: true }"
            :scrollbar="{ draggable: true }"
            @swiper="onSwiper"
            @slideChange="onSlideChange"
            class="preview__swiper"
        >
            <swiper-slide class="swiper__slide">
                <div class="swiper__slide-text">
                    <div class="swiper__slide-date">
                        <div class="swiper__slide-date-item">
                        20 ноября
                        </div>
                        <div class="swiper__slide-date-item">
                        Онлайн, очно
                        </div>
                    </div>
                    <h2 class="swiper__slide-title">
                      Гармония в Команде: Искусство Эффективного Общения
                    </h2>
                    <p class="swiper__slide-desc">
                      Научим навыкам эффективного общения и улучшению взаимодействия в команде для достижения общих целей
                    </p>
                    <div class="swiper__slide-btns">
                      <the-button
                        :styles="['swiper__slide-btn btn btn_red']"
                        :type="'button'"
                      >
                        Записаться
                      </the-button>
                      <the-button
                        :styles="['swiper__slide-btn btn btn_grey btn_arrow']"
                        :type="'button'"
                      >
                      Подробнее
                      </the-button>
                    </div>
                </div>
                <div class="swiper__slide-pic">
                    <img src="@/assets/images/messages.png">
                </div>
            </swiper-slide>
            <swiper-slide class="swiper__slide">
                <div class="swiper__slide-text">
                    <div class="swiper__slide-date">
                        <div class="swiper__slide-date-item">
                          20 ноября
                        </div>
                        <div class="swiper__slide-date-item">
                          Онлайн, очно
                        </div>
                    </div>
                    <h2 class="swiper__slide-title">
                    Гармония в Команде: Искусство Эффективного Общения
                    </h2>
                    <p class="swiper__slide-desc">
                    Научим навыкам эффективного общения и улучшению взаимодействия в команде для достижения общих целей
                    </p>
                    <div class="swiper__slide-btns">
                      <the-button
                        :styles="['swiper__slide-btn btn btn_red']"
                        :type="'button'"
                      >
                        Записаться
                      </the-button>
                      <the-button
                        :styles="['swiper__slide-btn btn btn_grey btn_arrow']"
                        :type="'button'"
                      >
                        Подробнее
                      </the-button>
                    </div>
                </div>
                <div class="swiper__slide-pic">
                    <img src="@/assets/images/messages.png">
                </div>
            </swiper-slide>
        </swiper>
    </div>
</template>

<style lang="scss">
.preview__swiper {
    margin-bottom: 80px;
    border-radius: 40px;
    @media (max-width: 1023px) {
      margin-bottom: 60px;
    }

    @media (max-width: 767px) {
      margin-bottom: 40px;
    }

    @media (max-width: 539px) {
      margin-bottom: 20px;
    }
    & .swiper-pagination-bullet {
        background: #fff;
        opacity: 1;
        width: 8px;
        height: 8px;
        transition: .2s;
    }

    & .swiper-pagination-bullet-active {
        width: 24px;
        border-radius: 90px;
    }

    & .swiper-button-prev, & .swiper-button-next {
        width: 20px;
        height: 20px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 8px;
        color: #fff;
        --swiper-navigation-size: 30px;
        font-weight: bold;
        @media (max-width: 539px) {
          display: none;
        }
    }
}
.preview {

}
.swiper {
  & .swiper-pagination  {
    bottom: 30px;
    @media (max-width: 899px) {
      bottom: 18px;
    }
    @media (max-width: 539px) {
      bottom: 5px;
    }
  }
  &__slide {
    border-radius: 40px;
    background: var(--var-preview-swiper);
    padding: 30px 0px 90px 100px;
    display: flex;
    gap: 20px;
    align-items: center;
    height: auto;
    transition: .2s;
    @media (max-width: 899px) {
      padding: 40px 50px;
    }
    @media (max-width: 539px) {
      padding: 20px 15px 30px;
    }
  }

  &__slide-text {
    flex: 0 0 59%;
    @media (max-width: 899px) {
      flex: 0 0 100%;
    }
    & h2 {
      color: var( --var-grey-lite-font);
      font-family: Montserrat;
      font-size: 40px;
      font-style: normal;
      font-weight: 600;
      line-height: 150%; /* 60px */
      margin-bottom: 30px;
      @media (max-width: 899px) {
        font-size: 30px;
      }
      @media (max-width: 767px) {
        font-size: 25px;
        margin-bottom: 15px;
      }
      @media (max-width: 539px) {
        font-size: 20px;
        margin-bottom: 10px;
      }
    }
    & p {
      color: var( --var-grey-lite-font);
      font-family: Inter;
      font-size: 20px;
      font-style: normal;
      font-weight: 400;
      line-height: 150%; /* 30px */
      margin-bottom: 40px;
      @media (max-width: 899px) {
        font-size: 17px;
      }
      @media (max-width: 767px) {
        font-size: 16px;
        margin-bottom: 15px;
      }
      @media (max-width: 539px) {
        margin-bottom: 10px;
      }
    }
  }

  &__slide-date {
    display: flex;
    gap: 20px;
    margin-bottom: 30px;
    @media (max-width: 767px) {
        margin-bottom: 15px;
    }
    @media (max-width: 539px) {
      margin-bottom: 10px;
    }
  }

  &__slide-date-item {
    color: var( --var-grey-lite-font);
    font-family: Roboto;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    padding: 10px 10px;
    border-radius: var(--rounded-xl, 12px);
    background: var(--var-grey-lite);
    backdrop-filter: blur(2px);
    transition: .2s;
    @media (max-width: 539px) {
      flex: 1;
    }
    
  }

  &__slide-title {
  }

  &__slide-desc {
  }

  &__slide-btns {
    display: flex;
    gap: 20px;
    @media (max-width: 539px) {
      flex-wrap: wrap;
      gap: 10px;
    }
  }

  &__slide-btn {
    @media (max-width: 539px) {
      flex: 0 0 100%;
    }
  }

  &__slide-pic {
    flex: 1;
    @media (max-width: 899px) {
      display: none;
    }
    & img {
      width: 100%;
    }
  }
}

</style>
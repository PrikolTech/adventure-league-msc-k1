<script>

  import { Swiper, SwiperSlide } from 'swiper/vue';

  import 'swiper/css';
  import 'swiper/css/navigation';
  import 'swiper/css/pagination';

  export default {
    components: {
      Swiper,
      SwiperSlide,
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
                    <button class="swiper__slide-btn btn btn_red">
                        Записаться
                    </button>
                    <button class="swiper__slide-btn btn btn_grey btn_arrow">
                        Подробнее
                        <img src="@/assets/images/icons/arrow-grey.svg">
                    </button>
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
                    <button class="swiper__slide-btn btn btn_red">
                        Записаться
                    </button>
                    <button class="swiper__slide-btn btn btn_grey btn_arrow">
                        Подробнее
                        <img src="@/assets/images/icons/arrow-grey.svg">
                    </button>
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
    border-radius: 40px;
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
    }
}
.preview {

}
.swiper {
  
  &__slide {
    border-radius: 40px;
    background: #CDD9E5;
    padding: 30px 0px 30px 100px;
    display: flex;
    gap: 20px;
    align-items: center;
    height: auto;
  }

  &__slide-text {
    flex: 0 0 59%;
    & h2 {
      color: var(--gray-900, #111928);
      font-family: Montserrat;
      font-size: 40px;
      font-style: normal;
      font-weight: 600;
      line-height: 150%; /* 60px */
      margin-bottom: 30px;
    }
    & p {
      color: var(--gray-900, #111928);
      font-family: Inter;
      font-size: 20px;
      font-style: normal;
      font-weight: 400;
      line-height: 150%; /* 30px */
      margin-bottom: 40px;
    }
  }

  &__slide-date {
    display: flex;
    gap: 20px;
    margin-bottom: 30px;
  }

  &__slide-date-item {
    color: var(--gray-900, #111928);
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
    background: #FFF;
    backdrop-filter: blur(2px);
  }

  &__slide-title {
  }

  &__slide-desc {
  }

  &__slide-btns {
    display: flex;
    gap: 20px;
  }

  &__slide-btn {
  }

  &__slide-pic {
    & img {
      width: 100%;
    }
  }
}

</style>
<template>
    <div class="rm">
        <div class="swiper-container">
            <div class="swiper-wrapper">
                <div class="swiper-slide" v-for="rm in rotationMap" :key="rm.bid">
                    <img :src="rm.url"/>
                </div>
            </div>
            <!-- 如果需要分页器 -->
            <div class="swiper-pagination"></div>
            <!-- 如果需要导航按钮 -->
            <div class="swiper-button-prev"></div>
            <div class="swiper-button-next"></div>
        </div>
    </div>
</template>

<script>
import {api_rotationMap} from "@/store/rotationMmap";
import Swiper from "swiper";

export default {
    name: "RotationMap",
    data() {
        return {
            rotationMap: {}
        }
    },
    async created() {
        api_rotationMap().then(res => {
            this.rotationMap = res.data.data.rotationMap
            console.log(this.rotationMap)
        })
    },
    updated() {
        new Swiper('.swiper-container', {
            //direction: 'vertical', // 垂直切换选项
            loop: true, // 循环模式选项
            autoplay: {
                disableOnInteraction: false,
                delay: 2000,
            },
            // 如果需要分页器
            pagination: {
                el: '.swiper-pagination',
                // 点击小球也可以跳转
                clickable: true
            },
            // 如果需要前进后退按钮
            navigation: {
                nextEl: '.swiper-button-next',
                prevEl: '.swiper-button-prev',
            },
        })
    }
}
</script>

<style scoped>
.rm {
    width: 502px;
    height: 400px;
}

.swiper-container {
    width: 502px;
    height: 400px;
}

img {
    width: 502px;
    height: 400px;
}
</style>
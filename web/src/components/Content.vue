<template>
    <div class="content">
        <div class="classify">
            <div class="cl-left">
                <div>
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-dongtai"></use>
                    </svg>
                    <span>动态</span>
                </div>
                <div>
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-remen"></use>
                    </svg>
                    <span>热门</span>
                </div>
                <div>
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-pindao"></use>
                    </svg>
                    <span>频道</span>
                </div>
            </div>
            <div class="cl-con">
                <ul>
                    <li><span>番剧</span></li>
                    <li><span>国创</span></li>
                    <li><span>综艺</span></li>
                    <li><span>动画</span></li>
                    <li><span>鬼畜</span></li>
                    <li><span>舞蹈</span></li>
                    <li><span>娱乐</span></li>
                    <li><span>科技</span></li>
                    <li><span>美食</span></li>
                    <li><span>电影</span></li>
                    <li><span>电视剧</span></li>
                    <li><span>纪录片</span></li>
                    <li><span>游戏</span></li>
                    <li><span>音乐</span></li>
                    <li><span>影视</span></li>
                    <li><span>知识</span></li>
                    <li><span>资讯</span></li>
                    <li><span>更多</span></li>
                </ul>
            </div>
            <div class="cl-ri">
                <ul>
                    <li><a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-columnszhuanlan"></use>
                        </svg>
                        <span>专栏</span>
                    </a></li>
                    <li><a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-huodong"></use>
                        </svg>
                        <span>活动</span>
                    </a></li>
                    <li style="width: 91px"><a class="max">
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-shequzhongxin"></use>
                        </svg>
                        <span>社区中心</span>
                    </a></li>
                    <li><a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-zhibomaiche"></use>
                        </svg>
                        <span>直播</span>
                    </a></li>
                    <li><a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-ketangpeixun"></use>
                        </svg>
                        <span>课堂</span>
                    </a></li>
                    <li style="width: 91px"><a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-bangdan"></use>
                        </svg>
                        <span>新歌热榜</span>
                    </a></li>
                </ul>
            </div>
        </div>
        <div class="ct">
            <div class="left">
                <RotationMap/>
            </div>
            <div class="right">
                <div class="item">
                    <Item/>
                </div>
                <div class="item">
                    <Item/>
                </div>
                <div class="item">
                    <Item/>
                </div>
                <div class="item">
                    <Item/>
                </div>
            </div>
            <div class="huanyihuan"
                 @click.prevent="refresh">
                <svg
                        class="icon hh"
                        aria-hidden="true"
                        ref="icon"
                >
                    <use xlink:href="#icon-weibiaoti--"></use>
                </svg>
                <span style="margin-top: 6px">
                换一换
            </span>
            </div>
        </div>

    </div>
</template>

<script>
import Item from "@/components/Item.vue";
import RotationMap from "@/components/RotationMap.vue";

export default {
    name: "Content",
    components: {RotationMap, Item},
    methods: {
        refresh() {

        }
    }
}
</script>

<style scoped>.content {
    width: 1144px;
    height: 600px;
    margin: 20px auto;
    position: relative;
    /*background-color: #00aeec;*/
}

.classify {
    width: 1144px;
    height: 100px;
    display: flex;
}

.cl-left {
    width: 168px;
    height: 64px;
    display: flex;
}

.cl-left .icon {
    width: 40px;
    height: 40px;
}

.cl-left > div {
    display: flex;
    flex-direction: column;
    margin: 0 16px 0 0;
    width: 40px;
    height: 64px;
    font-size: 13px;
    justify-content: center;
}

.cl-left > div span {
    width: 40px;
    height: 18px;
    line-height: 18px;
    text-align: center;
}

.cl-con ul {
    width: 722px;
    height: 66px;
    display: flex;
    flex-wrap: wrap;
    border-right: 1px seashell solid;
}


.cl-con li {
    width: 73px;
    height: 28px;
    background-color: #F6F7F8;
    margin-left: 5px;
    border-radius: 5px;
}

.cl-con li span {
    color: #61666D;
    font-size: 13px;
    height: 20px;
    line-height: 20px;
    text-align: center;
    display: flex;
    justify-content: center;
    margin-top: 4px;
}

.cl-ri {
    margin-left: 5px;
}

.cl-ri .icon {
    width: 20px;
    height: 20px;
    margin: 0 4px 0 0;
}

.cl-ri ul {
    display: flex;
    flex-wrap: wrap;
}

.cl-ri li {
    width: 69px;
    height: 28px;
    font-size: 13px;
    justify-content: flex-end;
    margin-bottom: 10px;
}


.ct {
    width: 1144px;
    height: 500px;
    display: flex;
    margin: 0 auto;
    position: relative;
}

.left {
    margin: 10px;
    height: 500px;
}

.right {
    display: flex;
    flex-wrap: wrap;
    height: 500px;
}

.right .item {
    margin-left: 30px;
    margin-top: 10px;
    width: 251px;
    /*+ +5 */
    height: 221px;
}

.huanyihuan {
    font-size: 12px;
    padding: 8px 8px;
    border-radius: 6px;
    border: 1px #999999 solid;
    position: absolute;
    right: 10px;
    top: 10px;
    display: flex;
    writing-mode: vertical-rl;
    justify-content: center;
    color: #18191C;
    cursor: pointer;
    transition: 1s all linear;
}

.huanyihuan:active {
    transform: rotate(360deg);
}


.hh {
    width: 16px;
    height: 16px;
}


</style>
<template>
    <div>
        <div class="item">
            <div class="video">
                <img src="./img.png">
                <div class="video-info">
                    <svg class="icon" aria-hidden="true">
                        <use xlink:href="#icon-shipin"></use>
                    </svg>
                    <span>10.1万</span>
                    <svg class="icon" aria-hidden="true" style="margin-left: 10px">
                        <use xlink:href="#icon-danmu"></use>
                    </svg>
                    <span>2.1万</span>
                    <span>19:87</span>
                </div>
            </div>
            <div class="title">
                <p>因为这本书，地坛成为极少数“轮椅畅行无阻”的景点【史铁生·番外篇】</p>
                <div class="info">
                    <a>
                        <svg class="icon" aria-hidden="true">
                            <use xlink:href="#icon-up"></use>
                        </svg>
                        <span>大象放映室</span><span>· 5-12</span>
                    </a>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
export default {
    name: "Item"
}
</script>

<style scoped>
.item {
    width: 251px;
    /*+ +5 */
    height: 221px;
    cursor: pointer;
}

.video {
    position: relative;
}

.video-info {
    position: absolute;
    bottom: 6px;
    font-size: 0;
    color: white;
}

.video-info svg {
    width: 15px;
    height: 15px;
    vertical-align: bottom;
}

.video-info span {
    height: 15px;
    line-height: 15px;
    font-size: 8px;
}

.video-info span:last-child {
    margin-left: 100px;
}

.video img {
    width: 251px;
    height: 152px;
}


.title {
    margin-top: 5px;
    width: 251px;
    height: 65px;
}

.icon {
    margin-top: 5px;
    width: 16px;
    height: 12px;
}
</style>
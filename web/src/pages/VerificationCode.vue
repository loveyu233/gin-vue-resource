<template>
    <div>
        <div :style="styleVar" id="captcha">
            <div id="bg" ref="bg"></div>
            <div id="refresh" @click="this.refresh"></div>
            <div id="handle" ref="handle" @mousemove="mousemoveHandler" @mouseup="mouseupHandler">
                <span ref="span" @mousedown="mousedownHandler"></span>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "VerificationCode",
    props: ["imgBase64", "bgBase64", "imgWidth", "imgHeight", "bgWidth", "bgHeight", "callback", "y", "refresh"],
    data() {
        return {
            flag: false,
        }
    },
    computed: {
        styleVar() {
            return `--bgBase64: ${this.bgBase64};
            --imgBase64: ${this.imgBase64};
            --width: ${this.bgWidth};
            --height: ${this.bgHeight};
            --puzzle-width: ${this.imgWidth};
            --puzzle-height: ${this.imgHeight};
            --moved: 0;
             --toppx: ${this.y};
            `
        },
        t() {
            return `--bg: #00a1d6;--fs: 30px`
        }
    },
    methods: {
        mousedownHandler() {
            this.flag = true
        },
        mousemoveHandler(e) {
            if (this.flag) {
                let oleft = this.$refs.handle.getBoundingClientRect().left;
                let buttonWidth = this.$refs.span.getBoundingClientRect().width;
                this.$refs.bg.style.setProperty('--moved', `${e.clientX - oleft - buttonWidth / 2}px`)
                this.$refs.span.style.setProperty('--moved', `${e.clientX - oleft - buttonWidth / 2}px`)
            }
        },
        mouseupHandler() {
            if (this.flag) {
                this.flag = !this.flag
            }
            let moved = this.$refs.bg.style.getPropertyValue('--moved')
            this.callback(moved)
            this.$refs.bg.style.setProperty('--moved', "0px")
            this.$refs.span.style.setProperty('--moved', "0px")
        }
    },
    mounted() {
        this.$refs.bg.style.setProperty('--top', `${this.y}px`)
    }
}
</script>

<style>
/*body {*/
/*    !* 背景图 *!*/
/*    --imgBase64: url("data:image/png;base64,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");*/
/*    !* 移动图*!*/
/*    --bgBase64: url("data:image/png;base64,/9j/2wCEAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDIBCQkJDAsMGA0NGDIhHCEyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMv/AABEIADIAMgMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/AFooorA7AooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigD//Z");*/
/*    !* 背景图宽高 *!*/
/*    --width: 310px;*/
/*    --height: 160px;*/
/*    !* 移动图片宽高 *!*/
/*    --puzzle-width: 50px;*/
/*    --puzzle-height: 50px;*/
/*    !* x轴 *!*/
/*    --moved: 0px;*/
/*    !* y轴 *!*/
/*    --toppx: 81px;*/
/*}*/
#refresh {
    width: 30px;
    height: 30px;
    /*background-color: red;*/
    position: absolute;
    right: 0;
    background-image: url("../../public/images/refresh.png");
    background-size: contain;
}

#captcha {
    width: var(--width);
    height: var(--height);
    background-image: var(--imgBase64);
    background-size: cover;
    background-position: center;
    position: relative;
    border-radius: 4px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, .3);
}

#bg {
    width: var(--puzzle-height);
    height: var(--puzzle-height);
    background-image: var(--bgBase64);
    /*background-color: red;*/
    position: absolute;
    top: var(--toppx);
    transform: translateX(clamp(0px,
    var(--moved),
    calc(var(--width) - var(--puzzle-width))))
}

#handle {
    height: 30px;
    width: var(--width);
    border-radius: 18px;
    background-color: #eee;
    box-shadow: inset 0 0 12px rgba(0, 0, 0, .2);
    border: 3px solid #ccc;
    position: absolute;
    bottom: -50px;
}

#handle span {
    display: block;
    width: var(--puzzle-width);
    height: inherit;
    border-radius: inherit;
    background-color: #fff;
    box-shadow: 0 0 6px rgba(0, 0, 0, .25);
    cursor: move;
    transform: translateX(clamp(0px,
    var(--moved),
    calc(var(--width) - var(--puzzle-width))))
}
</style>
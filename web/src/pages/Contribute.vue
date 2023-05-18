<template>
    <div>
        <el-steps :active="active">
            <el-step title="上传封面" icon="el-icon-picture"></el-step>
            <el-step title="上传视频" icon="el-icon-video-camera-solid"></el-step>
            <el-step title="成功" icon="el-icon-upload"></el-step>
        </el-steps>
        <el-button type="primary" @click="next" class="nextClick">下一步</el-button>
        <el-button type="primary" @click="pre" class="preClick">上一步</el-button>

        <div v-if="active===1" class="active1">
            <el-upload
                    action="#"
                    list-type="picture-card"
                    :limit="1"
                    ref="uploadImage"
                    :on-change="handleChangeImage"
                    :auto-upload="false"
            >

                <i slot="default" class="el-icon-plus"></i>
                <div slot="file" slot-scope="{file}">
                    <img
                            class="el-upload-list__item-thumbnail"
                            :src="file.url" alt=""
                    >
                    <span class="el-upload-list__item-actions"><span
                            class="el-upload-list__item-preview"
                            @click="handlePictureCardPreview(file)">
                <i class="el-icon-zoom-in"></i></span>
                <span
                        v-if="!disabled"
                        class="el-upload-list__item-delete"
                        @click="handleDownload(file)"
                >
                <i class="el-icon-download"></i>
                </span>
                <span
                        v-if="!disabled"
                        class="el-upload-list__item-delete"
                        @click="handleRemove(file)"
                >
                  <i class="el-icon-delete"></i>
                </span>
              </span>
                </div>
            </el-upload>
        </div>

        <div v-if="active===2" class="active2">
            <el-upload
                    class="upload-demo"
                    drag
                    action="#"
                    :auto-upload="false"
                    :on-change="handleChangeVideo"
            >
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过2M</div>
            </el-upload>
        </div>


        <div v-if="active===3" class="active3">

        </div>
    </div>
</template>

<script>

export default {
    name: "Contribute",
    data() {
        return {
            active: 1,
            videoInfo: {
                cover: Object,
                video: Object
            },
            dialogImageUrl: '',
            dialogVisible: false,
            disabled: false
        }
    },
    methods: {
        handleRemove(file) {
            let fs = this.$refs.uploadImage.uploadFiles
            for (let i in fs) {
                if (file.url === fs[i].url) {
                    fs.splice(i, 1);
                }
            }
            this.videoInfo.cover = Object
        },
        handlePictureCardPreview(file) {
            this.dialogImageUrl = file.url;
            this.dialogVisible = true;
        },
        handleDownload(file) {
            console.log(file);
        },
        next() {
            this.active++
        },
        pre() {
            this.active--
        },
        handleChangeImage(file) {
            this.videoInfo.cover = file
        },
        handleChangeVideo(file) {
            this.videoInfo.video = file
        }
    }
}
</script>

<style scoped>
.active1, .active2 {
    width: 360px;
    height: 200px;
    margin: 100px auto;
}
</style>
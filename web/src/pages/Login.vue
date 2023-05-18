<template>
    <div class="login" :style="bg.data">
        <div class="close-icon" @click="closelogin"></div>
        <div class="scan-wp">
            <div class="login-scan-title">扫描二维码登录</div>
            <div class="login-scan-hover-wp">
                <img src="../assets/erweima.png" alt="二维码">
            </div>
            <div class="login-scan-desc">
                <p> 请使用 <a href="https://app.bilibili.com/" target="_blank">哔哩哔哩客户端</a></p>
                <p>扫码登录或扫码下载APP</p>
            </div>
        </div>
        <div class="mini-line"></div>
        <div class="login-right">
            <div class="login-tab-wp">
                <div class="login-tab-item" @click="showHandler(1)"> 密码登录</div>
                <div class="login-tab-line"></div>
                <div class="login-tab-item" @click="showHandler(0)"> 邮箱登录</div>
            </div>
            <div class="verify" v-show="verify">
                <div id="captcha">
                    <VerificationCode
                            :imgBase64="code.imageBase"
                            :bgBase64="code.bgBase"
                            :imgWidth="code.imgWidth"
                            :imgHeight="code.imgHeight"
                            :bgWidth="code.bgWidth"
                            :bgHeight="code.bgHeight"
                            :callback="callback"
                            :y="code.y"
                            :refresh="refresh"
                    />
                </div>
            </div>
            <div v-show="show.passwordLogin">
                <div class="log-pwd-wp">
                    <form class="tab-form">
                        <div class="form-item">
                            <div class="form-info">账号</div>
                            <input type="text"
                                   autocomplete='off'
                                   placeholder="请输入账号"
                                   v-model="user.email"
                            >
                        </div>
                        <div class="form-line"></div>
                        <div class="form-item">
                            <div class="form-info">密码</div>
                            <input :type="showType.type" placeholder="请输入密码"
                                   @focus="focusHandler"
                                   @blur="blurHandler"
                                   v-model.trim="user.password"
                                   autocomplete="off"
                            >
                            <div class="eye-btn" style="--thememini-color:#00a1d6;" @click="eyeHandler">
                                <svg data-v-8f158982="" width="20"
                                     height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path data-v-8f158982=""
                                          fill-rule="evenodd"
                                          clip-rule="evenodd"
                                          :d="eye.d"
                                          fill="#9499A0"></path>
                                </svg>
                            </div>
                            <div class="clickable">忘记密码？</div>
                        </div>

                    </form>
                </div>
                <div class="btn_wp">
                    <div class="btn_other">
                        <p>注册</p>
                    </div>
                    <div class="btn_primary" @click="login">
                        <p>登录</p>
                    </div>
                </div>
            </div>
            <div v-show="show.emailLogin">
                <div class="log-pwd-wp">
                    <form class="tab-form">
                        <div class="form-item phone-input">
                            <div class="form-info phone-info">@qq.com</div>
                            <input type="text" placeholder="请输入邮箱" v-model="registerData.email">
                            <div class="phone-line"></div>
                            <div>
                                <button
                                        class="code-button"
                                        :disabled="sendModel.countFlag"
                                        @click.prevent="getNumberCodeHandler"
                                >
                                    {{ sendModel.btnMsg == null ? sendModel.countNum + 's' : sendModel.btnMsg }}
                                </button>
                            </div>
                        </div>
                        <div class="form-line"></div>
                        <div class="form-item">
                            <div class="form-info phone-code">验证码</div>
                            <input :type="showType.type" placeholder="请输入验证码"
                                   @focus="focusHandler"
                                   @blur="blurHandler"
                                   autocomplete="off"
                                   v-model="registerData.code"
                            >
                            <div class="eye-btn" style="--thememini-color:#00a1d6;" @click="eyeHandler">
                                <svg data-v-8f158982="" width="20"
                                     height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path data-v-8f158982=""
                                          fill-rule="evenodd"
                                          clip-rule="evenodd"
                                          :d="eye.d"
                                          fill="#9499A0"></path>
                                </svg>
                            </div>
                        </div>

                    </form>
                </div>
                <div class="btn_wp">
                    <div class="phone-btn_other" @click.prevent="registerHandler">
                        <p>登录/注册</p>
                    </div>
                </div>
            </div>
            <div class="login-sns-wp">
                <div class="login-sns-title">其他方式登录</div>
                <div class="login-sns-content">
                    <div class="login-sns-item"><img
                            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4CAMAAACfWMssAAAAgVBMVEUAAABXu0BYt0BQt0BXu0BWu0BXu0BXu0BXu0BWukBWukBXvEBXu0BXvEBXvEBYukBVukBXu0BWt0BXukBXuEBXu0D////1+/Pq9+fV7s/A5rer3aBsw1jg89uBzHBiwEy14auL0Xug2ZOW1YeBzG93yGSW1YjL6sO14azL6cN2yGP3XpzOAAAAFXRSTlMA3yAQ78+/r5+AUI9w74BgYEBAkHDBb56KAAACF0lEQVRIx52W6XKDIBRGwT3GZmsRUXFP0vb9H7AKGS8aiCXnR0TCmU/gOoh0uJck8jEZwUGYXND/cOPggyz4CE//0HZgKeDIsdSA3Qs1Bk2XejLF7ckGe1fnOT7ZBDsaDxNiZ4Jna4Jnb7rgbeKrK7QnFuzBOxIrYrsJAth9iIdl/9CwLE0pv/elqfoegWpfXdAUYINW9GRkRIBWakBemiOVGRbpE1lpijwaPDANCxvCc8qBbcVF47vq5EQ1YjCK3nyXiXE3QqrpSseeu+jptc96XgWmHSGDEGtCmDHygpK5nUuRdr2MvvfNdMvzXCN+KVNk6RO0qOpr37fXJzFCwdzmT9532THZovmqFHxlF3/WcdWQGUsBI2g3K/G3WG3o4oEVsVqOK4RHaTpfaKkXCVsWzPTL65pPN7X4kxnEX6qIXS4mJqfOH5tVKSJWzJsiXqlcklxe5AI0yuL4RDUpiKkGphRrRFRK+lLk88AQSg4KXVC9TvwSRQ4MU5m1xZ2xlmnEm1LkrqeKTVbU5rcaNtJDCAWqCOutq90CpjiKMTFQZuuah/9Oo+h6ZtPkYTSxI0YKReWLtxFBpJ5bzjOasWYsoBp6HQSRW5R5tz4C4HS0PltjO/H05sH6iQDXtz0d3/94ANPes/9Asjd9572PwE8X6Tm+DPViZMQ5mLUDxGnVCFtqwDH0VlYQS22bcxIGIhn7UXLWWn+10s6FZo+4YQAAAABJRU5ErkJggg=="
                            class="login-sns-item-icon" alt="微信">
                        <span class="login-sns-name">微信登录</span>
                    </div>
                    <div class="login-sns-item"><img
                            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4CAMAAACfWMssAAAAh1BMVEUAAADXQDjTQDzaQzfcRTfbRDfcRDbcQzbbRDbbRDfcQzbaQzbbRDbbRDbaRDfcQzbbRDfaQjjaQjXcQzfYRDfbRDTbRDf////gW1D99PPtoZvyubT76Ob20M3pioLdT0Tvrajkc2n20c3rlo7mfnb0xcD43NrrlY7kcmniZ1343driZ1z0xMEgvW1iAAAAFnRSTlMAIBDf34Dv78+/n1BAj7CvcGBgkHBwDUc+aAAAAmpJREFUSMeVlueCqjAQRgOI0qy7m0YVsN297/98OwEdCUQ05w8WDt/MkIjEhJv8RIFHAS/c/CTkM9w4XFKN5ebwgeaDNWURObYa4s+oMWrG1MOruDV9w9o1eU5A37JwDN6CUksTPWsTPXvTRe89wXBCa2rB+ukd6By3K6ca8WcNVoyxLB26i0ex/mxgzRTiMlx990A6C09PZ6XKwXbpIyP6Fn4C8zKONHVYFrJ4tJWfoUMJ5k2PPBi0igHZrX8j4HVLITPXB7udeK1gHem9ziJjouTwQYmnhCCuxh5UhSJO9qKmWwxrTcZeyh5UpcyFyOH0Bi5SaYNNyG7qYeK95CttWUWlJn6PW5RsSgVtlyqxoEhEQs1rNEPrldf5cI8Qb+hxYRKlaVsSY4NZKqU81aynpQY0kTNF3pS4fsDFvTESx4Gi0b4vMqWKdl4UcHWukmSeZXnRV9Gb5VgcDucIJ3B11JbckSmKyXAC7V7I52hxI2WmyYYk0m7+Dbd8JR7boTYlbshOE9WBsYb+MglBZ6r4p8TxXL+Hi7zov2dCnY1N8v7ViIS4z21V9q3U7Ap9YkzOulnrLAkZLtYTA4f+CpYLHEjFgCMdtwhiTBGedT+D/NwtIHU2V3li4tE9iO5qaILQwvF4LNUtTLsr8InnTX6Pebfq8qopmup/V29W0CkRIaNISDmxJwLqNeEYHwG8repzltVpA/Ua8UmP61ErPHy2xnbiniBflg9WxA0sCnXJAMezadDSRM/aRM/eDMCb4n7Z/glE9rOhq5i8xPFfaz7GmdXIs9aw4O1qZIUxam9IdtuwS/aCaJcYrT+uh9kYccQkXQAAAABJRU5ErkJggg=="
                            class="login-sns-item-icon" alt="微博">
                        <span class="login-sns-name">微博登录</span>
                    </div>
                    <div class="login-sns-item"><img
                            src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADgAAAA4CAMAAACfWMssAAAAb1BMVEUAAABMouVIn+NEn99MouVMouVMo+ZMouVKouRKouJMn+NMouZLouRLo+VMoeNMouZLoeVLoeVNn+RMouX///+83Pal0PL0+v6x1vRireiay/FXqOfS6Pl5uevH4vePxe+Ev+1utOrp9Pzp8/yEv+7XzqLPAAAAE3RSTlMA3yAQ74C/n2BQQK9wz5CPz7BwJ8NfpgAAAbNJREFUSMe1lutygjAQhTeEOyi2AQIC3t//GcvYqk3OYsh0+v1yMN+czc4mQBwyL5ONUDMiLsqc1iGzOFQGYZGu0KLZQkQSrNSQ6I2aGRqkpktxlXJQSc4LNsqJCBhPOCQwwXOa6PmbErxlNpJeVMqD6uWliqOfprPiyN5tcOp0PaObntnmo9hIAV39pFNA9BOogKb+RQN/h9+RCfzR1gbtQqSArtQmemAjUy4QI7GxBTy92qKGJfEshlgp0DO15srmhuIRFuX0Cc/2KO5hUUkF1xt3dxKKuUS3uCPBjA0w4rySwqYyDLCM+EqxVrd45UTtFm81ywlEuzkHXhyhOfY1DArfntg+jRMY/PB80CeODaLhHigpt7fYNSgeuvpgD7k0j9V4Gfbc6RjN7oRE1rCeejVo27vMD4/WFokynB33qUpnUYb+t5zg7+OhNYrVxx5PIzGRJ9ijniwxIC7y3DV3urZ9/DrDfQxvR793a+YnpvRk6+Nt6YXceRQq//rxgKb/h87/fVrtAmKQW2c/JfGkb0PDjBYJomUtMuJQTcRKDUmL0LLiDDWevCzie7LYJGXOWl/BnLhvbq/sWgAAAABJRU5ErkJggg=="
                            class="login-sns-item-icon" alt="QQ">
                        <span class="login-sns-name">QQ登录</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="login-agreement-wp">
            <p> 未注册过哔哩哔哩的手机号，我们将自动帮你注册账号 <!----></p>
            <p> 登录或完成注册即代表你同意 <span> 用户协议 <span
                    class="link_word"> 和 </span></span><span> 隐私政策 <span class="link_word">
                    </span></span></p>
        </div>
    </div>
</template>

<script>
import VerificationCode from "@/pages/VerificationCode.vue";
import {mapState} from "vuex";
import {api_imageCode, api_login, api_numberCode, api_register} from "@/api/user";
import {Message} from "element-ui";
import {mapActions} from "vuex";
import {setToken} from "@/store/user";

export default {
    name: "Login",
    components: {VerificationCode},
    props: ["closeHandler", "loginSuccessHandler"],
    computed: {
        ...mapState
    },
    data() {
        return {
            bg: {
                data: {
                    backgroundImage: "url(../images/22a.png),url(../images/33a.png)"
                },
                up: true
            },
            eye: {
                d: "M17.5753 6.85456C17.7122 6.71896 17.8939 6.63806 18.0866 6.63806C18.7321 6.63806 19.0436 7.42626 18.5748 7.87006C18.1144 8.30554 17.457 8.69885 16.6478 9.03168L18.1457 10.5296C18.2101 10.5941 18.2613 10.6706 18.2962 10.7548C18.331 10.839 18.349 10.9293 18.349 11.0204C18.349 11.1116 18.331 11.2019 18.2962 11.2861C18.2613 11.3703 18.2101 11.4468 18.1457 11.5113C18.0812 11.5757 18.0047 11.6269 17.9205 11.6618C17.8363 11.6967 17.746 11.7146 17.6548 11.7146C17.5637 11.7146 17.4734 11.6967 17.3892 11.6618C17.305 11.6269 17.2284 11.5757 17.164 11.5113L15.3409 9.68819C15.2898 9.63708 15.247 9.57838 15.2141 9.51428C14.4874 9.71293 13.6876 9.87122 12.8344 9.98119C12.8363 9.99011 12.8381 9.99908 12.8397 10.0081L13.2874 12.5472C13.315 12.7266 13.2713 12.9098 13.1656 13.0573C13.0598 13.2049 12.9005 13.3052 12.7217 13.3367C12.5429 13.3683 12.3589 13.3285 12.2091 13.2259C12.0592 13.1234 11.9555 12.9663 11.9202 12.7882L11.4725 10.2491C11.4645 10.2039 11.4611 10.1581 11.4621 10.1125C10.9858 10.1428 10.4976 10.1586 10.0002 10.1586C9.57059 10.1586 9.14778 10.1468 8.73362 10.1241C8.73477 10.1656 8.7322 10.2074 8.72578 10.249L8.27808 12.7881C8.24612 12.9694 8.14345 13.1306 7.99265 13.2362C7.84186 13.3418 7.65528 13.3831 7.47398 13.3512C7.29268 13.3192 7.1315 13.2166 7.0259 13.0658C6.9203 12.915 6.87892 12.7284 6.91088 12.5471L7.35858 10.008C7.35877 10.007 7.35896 10.0061 7.35915 10.0052C6.50085 9.90284 5.6941 9.75191 4.95838 9.56025C4.93012 9.60634 4.89634 9.64933 4.85748 9.68819L3.03438 11.5113C2.96992 11.5757 2.8934 11.6269 2.80918 11.6618C2.72496 11.6967 2.63469 11.7146 2.54353 11.7146C2.45237 11.7146 2.36211 11.6967 2.27789 11.6618C2.19367 11.6269 2.11714 11.5757 2.05268 11.5113C1.98822 11.4468 1.93709 11.3703 1.90221 11.2861C1.86732 11.2019 1.84937 11.1116 1.84937 11.0204C1.84937 10.9293 1.86732 10.839 1.90221 10.7548C1.93709 10.6706 1.98822 10.5941 2.05268 10.5296L3.49373 9.08855C2.6197 8.744 1.91247 8.33062 1.42559 7.87006C0.956591 7.42636 1.26799 6.63816 1.91359 6.63816C2.10629 6.63816 2.28789 6.71896 2.42489 6.85456C2.70009 7.12696 3.19529 7.45886 3.98459 7.77796C5.54429 8.40856 7.73699 8.77016 10.0001 8.77016C12.2632 8.77016 14.4558 8.40856 16.0156 7.77796C16.8049 7.45886 17.3001 7.12696 17.5753 6.85456Z",
                is: true
            },
            showType: {
                type: "password"
            },
            show: {
                passwordLogin: true,
                emailLogin: false
            },
            msg: '',
            verify: false,
            user: {
                email: "",
                password: "",
                x: 0
            },
            code: {
                imageBase: "",
                bgBase: "",
                y: "",
                imgWidth: "50px",
                imgHeight: "50px",
                bgWidth: "310px",
                bgHeight: "160px",
            },
            sendModel: {
                // 倒计时周期
                countNum: 60,
                // 用于倒计时标记，true-正在倒计时
                countFlag: false,
                // 定时器
                intervalBtn: {},
                // 默认按钮的值
                btnMsg: '获取验证码',
            },
            registerData: {
                email: "",
                code: ""
            }
        }
    },
    methods: {
        ...mapActions("login", ["userinfo"]),
        focusHandler() {
            this.bg = {
                data: {
                    backgroundImage: "url(../images/22b.png),url(../images/33b.png)"
                }
            }
        },
        blurHandler() {
            this.bg = {
                data: {
                    backgroundImage: "url(../images/22a.png),url(../images/33a.png)"
                }
            }
        },
        eyeHandler() {
            if (this.eye.is) {
                this.eye = {
                    d: "M2.11069 9.43732C3.21647 7.77542 5.87904 4.58331 9.89458 4.58331C13.8801 4.58331 16.6483 7.72502 17.8345 9.4049C18.0905 9.76747 18.0905 10.2325 17.8345 10.5951C16.6483 12.2749 13.8801 15.4166 9.89458 15.4166C5.87904 15.4166 3.21647 12.2245 2.11069 10.5626C1.88009 10.2161 1.88009 9.7839 2.11069 9.43732ZM9.89458 3.33331C5.19832 3.33331 2.20919 7.03277 1.07001 8.74489C0.560324 9.51091 0.560323 10.4891 1.07001 11.2551C2.20919 12.9672 5.19832 16.6666 9.89458 16.6666C14.5412 16.6666 17.6368 13.0422 18.8556 11.3161C19.4168 10.5213 19.4168 9.4787 18.8556 8.68391C17.6368 6.95774 14.5412 3.33331 9.89458 3.33331ZM7.29165 9.99998C7.29165 8.50421 8.50421 7.29165 9.99998 7.29165C11.4958 7.29165 12.7083 8.50421 12.7083 9.99998C12.7083 11.4958 11.4958 12.7083 9.99998 12.7083C8.50421 12.7083 7.29165 11.4958 7.29165 9.99998ZM9.99998 6.04165C7.81385 6.04165 6.04165 7.81385 6.04165 9.99998C6.04165 12.1861 7.81385 13.9583 9.99998 13.9583C12.1861 13.9583 13.9583 12.1861 13.9583 9.99998C13.9583 7.81385 12.1861 6.04165 9.99998 6.04165Z",
                    is: false
                }
                this.showType = {
                    type: "text"
                }
            } else {
                this.eye = {
                    d: "M17.5753 6.85456C17.7122 6.71896 17.8939 6.63806 18.0866 6.63806C18.7321 6.63806 19.0436 7.42626 18.5748 7.87006C18.1144 8.30554 17.457 8.69885 16.6478 9.03168L18.1457 10.5296C18.2101 10.5941 18.2613 10.6706 18.2962 10.7548C18.331 10.839 18.349 10.9293 18.349 11.0204C18.349 11.1116 18.331 11.2019 18.2962 11.2861C18.2613 11.3703 18.2101 11.4468 18.1457 11.5113C18.0812 11.5757 18.0047 11.6269 17.9205 11.6618C17.8363 11.6967 17.746 11.7146 17.6548 11.7146C17.5637 11.7146 17.4734 11.6967 17.3892 11.6618C17.305 11.6269 17.2284 11.5757 17.164 11.5113L15.3409 9.68819C15.2898 9.63708 15.247 9.57838 15.2141 9.51428C14.4874 9.71293 13.6876 9.87122 12.8344 9.98119C12.8363 9.99011 12.8381 9.99908 12.8397 10.0081L13.2874 12.5472C13.315 12.7266 13.2713 12.9098 13.1656 13.0573C13.0598 13.2049 12.9005 13.3052 12.7217 13.3367C12.5429 13.3683 12.3589 13.3285 12.2091 13.2259C12.0592 13.1234 11.9555 12.9663 11.9202 12.7882L11.4725 10.2491C11.4645 10.2039 11.4611 10.1581 11.4621 10.1125C10.9858 10.1428 10.4976 10.1586 10.0002 10.1586C9.57059 10.1586 9.14778 10.1468 8.73362 10.1241C8.73477 10.1656 8.7322 10.2074 8.72578 10.249L8.27808 12.7881C8.24612 12.9694 8.14345 13.1306 7.99265 13.2362C7.84186 13.3418 7.65528 13.3831 7.47398 13.3512C7.29268 13.3192 7.1315 13.2166 7.0259 13.0658C6.9203 12.915 6.87892 12.7284 6.91088 12.5471L7.35858 10.008C7.35877 10.007 7.35896 10.0061 7.35915 10.0052C6.50085 9.90284 5.6941 9.75191 4.95838 9.56025C4.93012 9.60634 4.89634 9.64933 4.85748 9.68819L3.03438 11.5113C2.96992 11.5757 2.8934 11.6269 2.80918 11.6618C2.72496 11.6967 2.63469 11.7146 2.54353 11.7146C2.45237 11.7146 2.36211 11.6967 2.27789 11.6618C2.19367 11.6269 2.11714 11.5757 2.05268 11.5113C1.98822 11.4468 1.93709 11.3703 1.90221 11.2861C1.86732 11.2019 1.84937 11.1116 1.84937 11.0204C1.84937 10.9293 1.86732 10.839 1.90221 10.7548C1.93709 10.6706 1.98822 10.5941 2.05268 10.5296L3.49373 9.08855C2.6197 8.744 1.91247 8.33062 1.42559 7.87006C0.956591 7.42636 1.26799 6.63816 1.91359 6.63816C2.10629 6.63816 2.28789 6.71896 2.42489 6.85456C2.70009 7.12696 3.19529 7.45886 3.98459 7.77796C5.54429 8.40856 7.73699 8.77016 10.0001 8.77016C12.2632 8.77016 14.4558 8.40856 16.0156 7.77796C16.8049 7.45886 17.3001 7.12696 17.5753 6.85456Z",
                    is: true
                }
                this.showType = {
                    type: "password"
                }
            }

        },
        showHandler(type) {
            if (type === 1) {
                this.show.passwordLogin = true
                this.show.emailLogin = false
            } else {
                this.show.passwordLogin = false
                this.show.emailLogin = true
            }
        },
        login() {
            let emailReg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
            if (this.user.email === "" || this.user.email === undefined || this.user.email == null) {
                Message.error("邮箱不能为空")
            } else if (!emailReg.test(this.user.email)) {
                Message.error("请输入正确的邮箱格式");
            } else if (this.user.password.trim() === "") {
                Message.error("密码不能为空");
            } else {
                this.refresh()
            }
        },
        callback(x) {
            this.user.x = x.substring(0, x.length - 2)
            api_login(this.user).then(res => {
                if (res.data.code === 3002) {
                    this.verify = false
                    Message.error("账号或密码错误")
                } else if (res.data.code === 3006) {
                    Message.error("验证码错误,请重新尝试")
                    this.refresh()
                } else if (res.data.code === 200) {
                    this.loginOK(res.data.data)
                }
            })
        },
        refresh() {
            api_imageCode({email: this.user.email}).then(res => {
                this.code = {...this.code, ...res.data.data}
                this.verify = true
            })
        },
        registerHandler() {
            api_register(this.registerData).then(res => {
                if (res.data["code"] === 3006) {
                    Message.error("验证码错误")
                } else {
                    this.loginOK(res.data.data)
                }
            })
        },
        getNumberCodeHandler() {
            let emailReg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
            if (this.registerData.email === "" || this.registerData.email === undefined || this.registerData.email == null) {
                Message.error("邮箱不能为空")
            } else if (!emailReg.test(this.registerData.email)) {
                Message.error("请输入正确的邮箱格式");
            } else {
                this.countDown()
                api_numberCode({email: this.registerData.email}).then(() => {
                    Message.success("验证码发送成功,请查看邮箱")
                })
            }
        },
        countDown() {
            // 设置btn倒计时时显示的信息
            this.sendModel.btnMsg = null
            // 更改btn状态
            this.sendModel.countFlag = !this.sendModel.countFlag
            // 设置倒计时
            this.sendModel.intervalBtn = setInterval(() => {
                if (this.sendModel.countNum <= 0) {
                    // 重置btn提示信息
                    this.sendModel.btnMsg = '获取验证码'
                    // 清除定时器
                    clearInterval(this.sendModel.intervalBtn)
                    // 更改btn状态
                    this.sendModel.countFlag = !this.sendModel.countFlag
                    // 重置倒计时状态
                    this.sendModel.countNum = 60
                }
                // 倒计时
                this.sendModel.countNum--
            }, 1000)
        },
        closelogin() {
            this.closeHandler()
        },
        loginOK(userinfo) {
            Message.success("登录成功")
            this.verify = false
            setToken(userinfo)
            this.loginSuccessHandler()
        }
    },
    mounted() {

    }
}
</script>
<style scoped>
.login {
    width: 663px;
    height: 349px;
    background-position: 0 100%, 100% 100%;
    background-repeat: no-repeat, no-repeat;
    background-size: 14%;
    position: relative;
    padding: 52px 65px 29px 92px;
    display: flex;
    border-radius: 10px;
}

.verify {
    position: absolute;
    width: 320px;
    height: 170px;
    top: 100px;
    left: 400px;
}

.clickable {
    color: #00a1d6;
    font-size: 14px;
}

.close-icon {
    width: 32px;
    height: 32px;
    background-image: url("../assets/close.svg");
    position: absolute;
    top: 20px;
    right: 20px;
    cursor: pointer;
}

.scan-wp {
    width: 173px;
    height: 349px;
}

.mini-line {
    width: 1px;
    height: 228px;
    background-color: #e3e5e7;
    margin: 43px 44px 0 45px;
}

.login-right {
    width: 400px;
    height: 349px;
}

.login-tab-wp {
    display: flex;
    width: 187px;
    height: 20px;
    margin: 0 auto 24px auto;
}

.login-tab-item {
    width: 72px;
    height: 20px;
    line-height: 20px;
    cursor: pointer;
    font-weight: 500;
    font-size: 18px;
    color: #00a1d6;
}

.login-tab-line {
    width: 1px;
    height: 20px;
    background-color: #e3e5e7;
    margin: 0 18px;
}

.tab-form {
    width: 400px;
    height: 90px;
}

.log-pwd-wp {
    border-radius: 10px;
    border: 1px solid #e3e5e7;
}

.form-item {
    width: 358px;
    height: 20px;
    display: flex;
    padding: 12px 20px;
}

.form-info {
    width: 30px;
    height: 20px;
    line-height: 20px;
    font-size: 14px;
}

.form-item input {
    width: 210px;
    height: 20px;
    margin: 0 0 0 20px;
    font-size: 14px;
}

.eye-btn {
    width: 20px;
    height: 20px;
    margin: 0 6px 0 0;
    cursor: pointer;
}

.clickable {
    width: 70px;
    height: 20px;
    line-height: 20px;
    cursor: pointer;
}

.btn_wp {
    width: 400px;
    height: 40px;
    display: flex;
    justify-content: space-between;
    margin: 10px auto;
}

.btn_wp div {
    width: 190px;
    height: 40px;
    border: 1px solid #e3e5e7;
    border-radius: 10px;
    cursor: pointer;
}

.btn_wp div p {
    height: 40px;
    line-height: 40px;
    text-align: center;
    font-weight: 400;
    font-size: 14px;
}

.btn_wp div:last-child {
    background-color: #00aeec;
}

.btn_wp div:last-child p {
    color: #FFF;
}

.login-sns-wp {
    width: 400px;
    height: 64px;
}

.login-sns-title {
    display: block;
    width: 78px;
    height: 16px;
    margin: 15px auto;
    font-size: 13px;
    color: #9499a0;
}

.login-sns-item {
    width: 88px;
    height: 28px;
    display: flex;
    cursor: pointer;
}

.login-sns-item img {
    width: 28px;
    height: 28px;
}

.login-sns-item span {
    height: 28px;
    line-height: 28px;
    margin-left: 5px;
    color: #9499a0;
}

.login-sns-content {
    display: flex;
    justify-content: space-around;
}

.form-line {
    width: 99%;
    border: 1px solid #e3e5e7;
}

.login-agreement-wp {
    width: 351px;
    height: 38px;
    position: absolute;
    bottom: 30px;
    left: 260px;
    color: #999;
    text-align: center;
    font-size: 13px;
}

.login-agreement-wp span {
    color: #00a1d6;
    cursor: pointer;
}

.login-scan-title {
    font-weight: 500;
    font-size: 18px;
    line-height: 16px;
    color: #505050;
    margin-bottom: 26px;
    text-align: center;
}

.login-scan-hover-wp {
    width: 161px;
    height: 161px;
    border: 1px solid #e3e5e7;
    border-radius: 8px;
    padding: 6px;
}

.login-scan-hover-wp img {
    width: 160px;
    height: 160px;
}

.login-scan-desc {
    width: 142px;
    height: 38px;
    margin-top: 18px;
}

.login-scan-desc p {
    width: 142px;
    height: 19px;
    line-height: 19px;
    text-align: center;
    font-size: 13px;
    margin-left: 18px;
}

a {
    color: #00aeec;
}

.phone-input input {
    width: 185px;
}

.phone-line {
    width: 1px;
    height: 26px;
    border-left: 1px solid #e3e5e7;
    margin-right: 20px;
}

.phone-input div:last-child {
    width: 90px;
    height: 20px;
    color: #c9ccd0 !important;
    cursor: pointer;
    font-size: 14px;
}

.phone-info {
    width: 46px;
    height: 20px;
    line-height: 20px;
}

.phone-code {
    width: 46px;
    height: 20px;
    line-height: 20px;
}

.phone-btn_other {
    margin: 0 auto;
    background-color: #00aeec;
    height: 40px;
    line-height: 40px;
    cursor: pointer;
    color: #fff;
    font-size: 14px;
}

.code-button {
    background-color: white;
    cursor: pointer;
}
</style>
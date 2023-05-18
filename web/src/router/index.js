import VueRouter from "vue-router";
import Vue from "vue";
import User from "@/pages/User.vue";
import Index from "@/components/Index.vue";
import Contribute from "@/pages/Contribute.vue";

Vue.use(VueRouter)

const router = new VueRouter({
    routes: [
        {
            path: "/",
            component: Index
        },
        {
            path: "/user",
            component: User
        },
        {
            path: "/contribute",
            component: Contribute
        }
    ]
})

export default router
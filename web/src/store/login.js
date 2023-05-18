export default {
    namespaced: true,
    actions: {
        userinfo({commit}, user) {
            commit("USERINFO", user)
        },
    },

    mutations: {
        USERINFO(state, user) {
            state.user = {...state.user, ...user}
        }
    },

    state: {
        token: "",                      //jwt
        user: {
            uid: 0,                     // id
            username: "",               //用户名
            background_img: "",         // 背景图
            sculpture_head: "",         // 头像
            introduce_briefly: "",      // 个人简介
            number_coin: "",            // 硬币数量
            number_b_coin: "",          // B币数
            level: "",                  // 等级
            experience: "",             // 当前经验值
            experience_required: "",    // 升到下一级所需经验值
            number_fans: "",            // 粉丝数量
            number_follow: ""           // 关注博主的数量
        }
    }
}
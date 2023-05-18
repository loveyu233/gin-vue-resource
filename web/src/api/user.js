import http from "@/api/http";

let userUrl = "/user"

export function api_login(params) {
    return http.post(`${userUrl}/login`, params)
}

export function api_imageCode(params) {
    return http.get(`${userUrl}/imagecode`, params)
}

export function api_numberCode(params) {
    return http.get(`${userUrl}/numbercode`, params)
}

export function api_register(params) {
    return http.post(`${userUrl}/register`, params)
}

export function api_userinfo(params) {
    return http.post(`${userUrl}/token`, params)
}
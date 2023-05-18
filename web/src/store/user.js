export function getToken() {
    return JSON.parse(localStorage.getItem("user"))
}

export function setToken(value) {
    return localStorage.setItem("user", JSON.stringify(value))
}

export function clearToken() {
    return localStorage.clear()
}
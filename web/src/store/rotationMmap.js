import http from "@/api/http";

export function api_rotationMap() {
    return http.get("/rotationMap/get")
}
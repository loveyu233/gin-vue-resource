<template>
    <div id="app">
        <router-view>

        </router-view>
    </div>
</template>

<script>

import Top from "@/components/Top.vue";
import Content from "@/components/Content.vue";
import Bottom from "@/components/Bottom.vue";

export default {
    name: 'App',
    components: {
        Bottom,
        Content,
        Top,
    }
}
</script>

<style>

</style>

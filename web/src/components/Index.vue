<template>
    <div>
        <Top/>
        <Content/>
        <Bottom/>
    </div>
</template>

<script>
import Content from "@/components/Content.vue";
import Top from "@/components/Top.vue";
import Bottom from "@/components/Bottom.vue";

export default {
    name: "Index",
    components: {Bottom, Top, Content}
}
</script>

<style scoped>

</style>
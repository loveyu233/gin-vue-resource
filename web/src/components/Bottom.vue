<template>
    <div class="b">
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <div class="item">
            <Item/>
        </div>
        <!--        <ul-->
        <!--                class="infinite-list"-->
        <!--                v-infinite-scroll="load"-->
        <!--                style="overflow:auto"-->
        <!--        >-->
        <!--            <li v-for="(c,i) in count" :key="i">-->
        <!--                -->
        <!--            </li>-->
        <!--        </ul>-->
    </div>
</template>

<script>
import Item from "@/components/Item.vue";

export default {
    name: "Bottom",
    components: {Item},
    data() {
        return {
            count: 0
        }
    },
    methods: {
        load() {
            this.count += 10
        },
        test() {
            console.log("abc")
        }
    }
}
</script>

<style scoped>
.b {
    width: 1144px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.item {
    margin-left: 34px;
    margin-top: 34px;
}

.infinite-list {
    height: calc(100vh);
    width: 100vw;
    overflow: auto;
}
</style>
<template>
    <div class='app-body'>
        <router-view></router-view>
    </div>
</template>

<script setup lang="ts">
import NotFound from "./NotFound.vue";

defineOptions({
    routes: [
        { path: "/:pathMatch(.*)*", component: NotFound },
    ],
});
</script>

<style lang="scss" scoped>
.app-body {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;
}
</style>

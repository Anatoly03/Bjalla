<template>
    <div class='app-body'>
        <router-view></router-view>
        <ModalGlobalView />
    </div>
</template>

<script setup lang="ts">
import ModalProfileSettings from "./ModalProfileSettings.vue";
import NotFound from "./NotFound.vue";
import ViewAuth from "./ViewAuth.vue";
import ViewHome from "./home/ViewHome.vue";

defineOptions({
    routes: [
        { name: "Home", path: "/", component: ViewHome },
        { name: "HomeGuild", path: "/:guild", component: ViewHome },
        { name: "HomeChannel", path: "/:guild/:channel", component: ViewHome },
        { name: "Login", path: "/login", component: ViewAuth },
        { name: "NotFound", path: "/:pathMatch(.*)*", component: NotFound },
    ],
    global: [
        {
            path: "profile-settings",
            name: "ProfileSettings",
            component: ModalProfileSettings,
            meta: { modal: true, direct: false },
        },
    ]
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

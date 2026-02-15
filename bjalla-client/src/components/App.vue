<template>
    <div class='app-body'>
        <router-view></router-view>
        <ModalGlobalView />
    </div>
</template>

<script setup lang="ts">
import ModalProfileSettings from "./settings/ModalProfileSettings.vue";
import ModalProfileSettingsGeneral from "./settings/ModalProfileSettingsGeneral.vue";
import ModalProfileSettingsTheme from "./settings/ModalProfileSettingsTheme.vue";
import NotFound from "./NotFound.vue";
import ViewAuth from "./ViewAuth.vue";
import ViewHome from "./home/ViewHome.vue";
import ModalCreateGuild from "./create-guild/ModalCreateGuild.vue";
import ModalGuildSettings from "./guild-settings/ModalGuildSettings.vue";
import ModalGuildSettingsGeneral from "./guild-settings/ModalGuildSettingsGeneral.vue";
import ModalGuildSettingsRoles from "./guild-settings/ModalGuildSettingsRoles.vue";
import ModalViewMemberProfile from "./view-member-profile/ModalViewMemberProfile.vue";

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
            children: [
                { name: "ProfileSettingsGeneral", path: "", component: ModalProfileSettingsGeneral },
                { name: "ProfileSettingsTheme", path: "theme", component: ModalProfileSettingsTheme },
            ],
        },
        {
            path: "create-guild",
            name: "CreateGuild",
            component: ModalCreateGuild,
            meta: { modal: true, direct: false },
        },
        {
            path: "guild-settings",
            name: "GuildSettings",
            component: ModalGuildSettings,
            meta: { modal: true, direct: false },
            children: [
                { name: "GuildSettingsGeneral", path: "", component: ModalGuildSettingsGeneral },
                { name: "GuildSettingsRoles", path: "roles", component: ModalGuildSettingsRoles },
            ],
        },
        {
            path: "view-member-profile",
            name: "ViewMemberProfile",
            component: ModalViewMemberProfile,
            meta: { modal: true, direct: false },
        }
    ]
});
</script>

<style lang="scss" scoped>
@use "../assets/color.scss";

.app-body {
    display: flex;
    flex-direction: row;
    width: 100%;
    height: 100%;

    color: var(--theme-text, #1a1a1a);
    background: var(--theme-bg, #f2e6d6);

    transition: background 160ms ease, color 160ms ease;
}
</style>

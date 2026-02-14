<template>
    <div class="view-home">
        <ViewSidebar />
        <div class="view-home-content">
            <ViewChannel />
            <ViewMemberList :key="$route.params.guild" />
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { useRouter } from "vue-router";

import ViewSidebar from "./ViewSidebar.vue";
import ViewMemberList from "./ViewMemberList.vue";
import ViewChannel from "./ViewChannel.vue";

/**
 * Router instance to redirect unauthenticated users to the login page.
 */
const router = useRouter();

/**
 * Redirect to login if the user is not authenticated.
 */
if (!pb.authStore.isValid) {
    router.push("/login");
}
</script>

<style lang="scss" scoped>
.view-home {
    display: flex;
    flex: 1;
    justify-content: center;
    align-items: center;

    .view-home-content {
        display: flex;
        flex-direction: row;
        flex: 1;
        justify-content: center;
        align-items: center;
        height: 100%;
    }
}
</style>
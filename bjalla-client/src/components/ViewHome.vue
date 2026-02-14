<template>
    <div class="view-home">Logged in as {{ pb.authStore.model?.email }}</div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import { useRouter } from "vue-router";

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
}
</style>
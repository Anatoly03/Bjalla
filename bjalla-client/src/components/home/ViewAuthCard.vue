<template>
    <div class="view-profile">
        <div class="card">
            <div class="card-body" @click="openProfile">
                {{ pb.authStore.model?.name }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import pb from "../../service/pocketbase";
import { useModalRoute } from "@vmrh/core";

const router = useRouter();
const { openModal } = useModalRoute();

function openProfile() {
    const currentPath = router.currentRoute.value.fullPath;

    openModal("ProfileSettings");
    // hotfix: keep the visible URL unchanged while modal is open.
    setTimeout(() => window.history.replaceState(window.history.state, "", currentPath), 0);
}
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-profile {
    position: absolute;
    bottom: 0;

    display: flex;
    flex-direction: column;
    margin-bottom: 1rem;
    width: inherit;
    box-sizing: border-box;

    .card {
        margin: 0 1em;
        padding: 0.25em;
        background-color: white;

        // width: 100%;
        box-sizing: border-box;

        .card-body {
            display: flex;
            align-items: begin;
            padding: 0.5rem;
            border-radius: 0.25rem;
            color: inherit;
            text-decoration: none;
            cursor: pointer;

            &:hover {
                background: #eee;
            }

            &.active {
                background: #ddd;
            }
        }
    }
}
</style>

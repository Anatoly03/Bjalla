<template>
    <div v-if="modelValue" class="modal-background" @click.self="closeModal">
        <div class="modal modal-profile-settings">
            <div class="modal-sidebar">
                <a @click="navigateTo('ProfileSettingsGeneral')" class="modal-sidebar-item">General</a>
                <a @click="navigateTo('ProfileSettingsTheme')" class="modal-sidebar-item">Theme</a>
            </div>
            <div class="modal-content">
                <ModalRouterView />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ModalRouterView, useCurrentModal } from "@vmrh/core";
import { useRoute, useRouter } from "vue-router";
import ModalProfileSettingsGeneral from "./ModalProfileSettingsGeneral.vue";
import ModalProfileSettingsTheme from "./ModalProfileSettingsTheme.vue";

const { close, modelValue } = useCurrentModal();
const router = useRouter();
const route = useRoute();

/**
 * Closes the modal when the background is clicked or when the update
 * is successful.
 */
function closeModal() {
    close();
}

/**
 * Updates the user's name in PocketBase and closes the modal on success.
 */
async function navigateTo(name: string) {
    try {
        await router.push({
            name,
            params: route.params,
            query: route.query,
            hash: route.hash,
        });
    } catch (error) {
        console.error("Navigation error:", error);
    }
}

/**
 * Expose child routes for global modal registration.
 */
defineOptions({
    routes: [
        { name: "ProfileSettingsGeneral", path: "", component: ModalProfileSettingsGeneral },
        { name: "ProfileSettingsTheme", path: "theme", component: ModalProfileSettingsTheme },
    ],
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.modal-background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal {
    display: flex;
    flex-direction: row;

    position: absolute;
    margin: auto;
    width: 80%;
    height: 80%;

    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(4px);
    border-radius: 0.5rem;
    padding: 1rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);

    .modal-content {
        flex: 1;
        padding: 1rem;
        overflow-y: auto;
    }
}

.modal-sidebar {
    width: 200px;
    border-right: 1px solid #ccc;
    padding: 1rem;

    .modal-sidebar-item {
        display: block;
        padding: 0.5rem;
        color: #333;
        text-decoration: none;
        cursor: pointer;

        &.active {
            background-color: #eee;
            font-weight: bold;
        }

        &:hover {
            background-color: #f5f5f5;
        }
    }
}
</style>

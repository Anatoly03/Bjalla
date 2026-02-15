<template>
    <div class="modal-view-general">
        <h2>{{ pb.authStore.record!.name }} <small>({{ pb.authStore.record!.email }})</small></h2>
        <div class="row">
            <input type="text" v-model="newName">
            <button @click="updateName">Update Name</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { useCurrentModal } from "@vmrh/core";
import { ref } from "vue";

const { close } = useCurrentModal();

/**
 * Fields for profile rename functionality.
 */
const newName = ref(pb.authStore.record!.name);

/**
 * Updates the user's name in PocketBase and closes the modal on success.
 */
async function updateName() {
    try {
        await pb.collection("users").update(pb.authStore.record!.id, {
            name: newName.value,
        });
        pb.authStore.record!.name = newName.value;
        close();
    } catch (error) {
        console.error("Failed to update name:", error);
    }
}
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.row {
    display: flex;
    gap: 0.5rem;
    margin-top: 1rem;

    input {
        flex: 1;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 0.25rem;
    }

    button {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 0.25rem;
        background-color: #0a3;
        color: white;
        cursor: pointer;

        &:hover {
            background-color: #0056b3;
        }
    }
}
</style>

<template>
    <div v-if="modelValue" class="modal-background" @click.self="closeModal">
        <div class="modal modal-profile-settings">
            <h2>{{ pb.authStore.record!.name }} <small>({{ pb.authStore.record!.email }})</small></h2>
            <div class="row">
                <input type="text" v-model="newName">
                <button @click="updateName">Update Name</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { useCurrentModal } from "@vmrh/core";
import { ref } from "vue";

const { close, modelValue } = useCurrentModal();

/**
 * Fields for profile rename functionality.
 */
const newName = ref(pb.authStore.record!.name);

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
async function updateName() {
    try {
        await pb.collection("users").update(pb.authStore.record!.id, {
            name: newName.value,
        });
        pb.authStore.record!.name = newName.value;
        closeModal();
    } catch (error) {
        console.error("Failed to update name:", error);
    }
}
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
    position: absolute;
    margin: auto;
    width: 80%;
    height: 80%;

    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(4px);
    border-radius: 0.5rem;
    padding: 1rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);

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
}
</style>

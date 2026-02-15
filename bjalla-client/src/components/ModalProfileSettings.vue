<template>
    <div v-if="modelValue" class="modal-background" @click.self="closeModal">
        <div class="modal modal-profile-settings">
            <h2>Profile Settings</h2>
            <p>This is a modal for profile settings.</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useCurrentModal } from "@vmrh/core";
const { close, modelValue } = useCurrentModal();

function closeModal() {
    close();
}
</script>

<style lang="scss" scoped>
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
}
</style>

<template>
    <div v-if="modelValue" class="modal-background" @click.self="closeModal">
        <div class="modal modal-create-guild card">
            <div class="card-header">
                Propose New Guild
            </div>
            <form class="card-body" @click.prevent="createGuild">
                <div class="card-banner">
                    <div class="guild-avatar" @click="uploadGuildAvatar" :class="{ empty: !imgFile }">
                        <img v-if="imgFile" :src="imgFile" alt="Guild Avatar" />
                        <span class="upload-hint" v-else>
                            <font-awesome-icon icon="fa-regular fa-file-image" />
                        </span>
                    </div>
                </div>
                <input type="text" placeholder="Guild Name" required>
                <div class="alert alert-warning">
                    Guild Creation is currently disabled. At launch you will be
                    able to petition new communities and we will create the guild for you.
                </div>
                <input type="submit" value="Propose Community!" disabled>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useCurrentModal } from "@vmrh/core";
import { ref } from "vue";

const { close, modelValue } = useCurrentModal();

const imgFile = ref<string | null>(null);

/**
 * Closes the modal when the background is clicked or when the update
 * is successful.
 */
function closeModal() {
    // revokes imgFile URL to free up memory
    if (imgFile.value) {
        URL.revokeObjectURL(imgFile.value);
    }

    close();
}

/**
 * Prompts the user to upload a file and replaces the guild avatar
 * with the uploaded image.
 */
function uploadGuildAvatar() {
    const input = document.createElement("input");
    input.type = "file";
    input.accept = "image/*";
    input.onchange = (event) => {
        const target = event.target as HTMLInputElement;
        if (target.files && target.files[0]) {
            imgFile.value = URL.createObjectURL(target.files[0]);
        }
    };
    input.click();
}

/**
 * Handles the form submission to create a new guild.
 */
async function createGuild() {
    console.log("Guild creation is currently disabled.");
}
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";
@use "../../assets/ui.scss";

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
    flex-direction: column;

    position: absolute;
    width: 480px;

    background: var(--theme-bg, #f2e6d6);
    backdrop-filter: blur(4px);
    border-radius: 0.5rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);

    .card-header {
        justify-content: center;
        background: var(--theme-bg-1, #f2e6d6);
    }
}

.card-banner {
    display: flex;
    justify-content: center;
    padding: auto;

    .guild-avatar {
        width: 64px;
        height: 64px;

        &.empty, &:hover {
            background: var(--theme-bg-1, #f2e6d6);
        }

        img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            border-radius: 8px;
        }

        .upload-hint {
            display: flex;
            align-items: center;
            justify-content: center;
            width: 100%;
            height: 100%;
            color: var(--theme-text-secondary, #888);
        }
    }
}
</style>

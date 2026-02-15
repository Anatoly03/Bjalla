<template>
    <div v-if="isActive" class="view-member-profile-modal-background" :style="modalStyle" ref="modalRef" tabindex="-1" @focusout="handleFocusOut">
        <div class="card view-member-profile">
            <div class="profile-banner" :class="{ empty: !user?.avatar }">
                <img v-if="user?.avatar" :src="avatarUrl" alt="avatar" class="user-avatar" />
                <span v-else class="user-avatar-placeholder">{{ user?.name?.[0] ?? "?" }}</span>
            </div>
            <div class="profile-name">
                {{ user?.name ?? "Unknown User" }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { useModal } from "@vmrh/core";
import { computed, nextTick, ref, watch } from "vue";

const { close, isActive } = useModal("ViewMemberProfile");
const modalRef = ref<HTMLDivElement | null>(null);

/**
 * Gets the user ID and guild ID from the query parameters, as well as
 * the click position.
 */
const props = defineProps<{
    user_id: string;
    guild_id: string;
    position?: { x: number; y: number };
}>();

/**
 * User Information
 */
const user = ref<any>(null);
const avatarUrl = computed(() => {
    if (!user.value?.avatar) {
        return "";
    }

    return pb.files.getURL(user.value, user.value.avatar);
});

const fallbackPosition = computed(() => {
    if (typeof window === "undefined") {
        return { x: 0, y: 0 };
    }

    return {
        x: Math.round(window.innerWidth / 2),
        y: Math.round(window.innerHeight / 2),
    };
});

const modalStyle = computed(() => {
    const position = props.position ?? fallbackPosition.value;

    return {
        left: `${position.x}px`,
        top: `${position.y}px`,
    };
});

/**
 * Closes the modal when the background is clicked or when the update
 * is successful.
 */
function closeModal() {
    close();
}

function handleFocusOut(event: FocusEvent) {
    const nextTarget = event.relatedTarget as Node | null;
    if (!nextTarget || !modalRef.value?.contains(nextTarget)) {
        closeModal();
    }
}

watch(isActive, async (active) => {
    if (!active) return;
    await nextTick();
    modalRef.value?.focus();
});

watch(
    () => props.user_id,
    async (userId) => {
        if (!userId) {
            user.value = null;
            return;
        }

        try {
            user.value = await pb.collection("users").getOne(userId, {
                requestKey: `user-profile-${userId}`,
            });
        } catch (error) {
            console.error("Failed to fetch user data:", error);
            user.value = null;
        }
    },
    { immediate: true }
);
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";
@use "../../assets/ui.scss";

.view-member-profile-modal-background {
    position: fixed;
    transform: translateX(-100%);
    transform-origin: top right;

    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    background: var(--theme-bg, #f2e6d6);
    backdrop-filter: blur(4px);
    border-radius: 0.5rem;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);

    .card-header {
        justify-content: center;
        background: var(--theme-bg-1, #f2e6d6);
    }
}

.view-member-profile {
    width: 300px;
    height: 400px;
    padding: 1em;

    box-shadow: 0 10px 22px rgba(0, 0, 0, 0.3);

    .profile-banner {
        display: flex;
        justify-content: center;
        padding: auto;

        .user-avatar {
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

    .profile-name {
        width: 100%;
        justify-content: center;
        margin-top: 1rem;
        font-size: 1.25rem;
        font-weight: bold;
        color: var(--theme-text, #1a1a1a);
    }
}
</style>

<template>
    <div v-if="isActive" class="view-member-profile-modal-background" :style="modalStyle" ref="modalRef" tabindex="-1" @focusout="handleFocusOut">
        <div class="card view-member-profile">
            Card
        </div>
    </div>
</template>

<script setup lang="ts">
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

    box-shadow: 0 10px 22px rgba(0, 0, 0, 0.3);
}
</style>

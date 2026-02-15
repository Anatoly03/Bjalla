<template>
    <div
        class="theme-card"
        :data-theme="theme"
        role="button"
        tabindex="0"
        @click="selectTheme"
        @keydown.enter.prevent="selectTheme"
        @keydown.space.prevent="selectTheme"
    >
        <div class="theme-preview">
            <!-- <span v-if="badge" class="theme-badge">{{ badge }}</span> -->
            <div class="theme-discs" aria-hidden="true">
                <span class="theme-disc"></span>
                <span class="theme-disc"></span>
                <span class="theme-disc"></span>
                <span class="theme-disc"></span>
                <span class="theme-disc"></span>
                <span class="theme-disc"></span>
            </div>
        </div>
        <div class="theme-name">{{ name }}</div>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps<{
    theme: string;
    name: string;
    badge?: string;
}>();

const emit = defineEmits<{
    (event: "select", theme: string): void;
}>();

function selectTheme() {
    emit("select", props.theme);
}
</script>

<style lang="scss" scoped>
@use "../../assets/color.scss";
@use "../../assets/theme.scss";

.theme-card {
    display: flex;
    flex-direction: column;
    height: 128px;

    color: #5b3b2a;
    cursor: pointer;
    user-select: none;

    box-shadow: 0 5px 11px rgba(0, 0, 0, 0.2);
    transition: transform 160ms ease, box-shadow 160ms ease;

    // border: 1px solid black;
    border-radius: 0.5rem;
    outline: none;

    &:hover, &:focus {
        transform: translateY(-4px);
        box-shadow: 0 10px 22px rgba(0, 0, 0, 0.6);
    }
}

.theme-preview {
    display: flex;
    flex: 1;
    align-items: center;
    justify-content: center;
    position: relative;
    padding: 12px 14px;

    border-top-right-radius: inherit;
    border-top-left-radius: inherit;
    overflow: hidden;

    background: var(--theme-bg, inherit);
}

.theme-name {
    padding: 0.5em;
    color: var(--theme-text, inherit);
    background: var(--theme-bg-2, inherit);
    border-bottom-left-radius: inherit;
    border-bottom-right-radius: inherit;
}

.theme-badge {
    position: absolute;
    left: 12px;
    bottom: 10px;
    padding: 2px 10px 3px;
    font-size: 0.7rem;
    letter-spacing: 0.12em;
    border-radius: 999px;
    background: #f7efe4;
    color: #3f2a1c;
    border: 1px solid rgba(63, 42, 28, 0.15);
}

.theme-discs {
    display: flex;
    gap: 0;
    align-items: center;
}

.theme-disc {
    width: 34px;
    height: 34px;
    border-radius: 999px;
    border: 1px solid rgba(0, 0, 0, 0.15);
    box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.25);

    &:not(:nth-child(1)) {
        margin-left: -14px;
    }
}

.theme-disc:nth-child(1) { background-color: var(--primary); }
.theme-disc:nth-child(2) { background-color: var(--primary-1); }
.theme-disc:nth-child(3) { background-color: var(--primary-2); }
.theme-disc:nth-child(4) { background-color: var(--primary-3); }
.theme-disc:nth-child(5) { background-color: var(--primary-4); }
.theme-disc:nth-child(6) { background-color: var(--primary-5); }

@media (max-width: 720px) {
    .theme-preview {
        height: 72px;
    }

    .theme-disc {
        width: 30px;
        height: 30px;
    }
}
</style>

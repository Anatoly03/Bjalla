<template>
    <div class="modal-view-theme">
        <ViewThemePreview theme="blank" name="Blank" @select="applyTheme" />
        <ViewThemePreview theme="european" name="European Blue" badge="NEW" @select="applyTheme" />
        <ViewThemePreview theme="dark" name="Dark Mode" @select="applyTheme" />
    </div>
</template>

<script setup lang="ts">
import ViewThemePreview from './ViewThemePreview.vue';
import { applyTheme } from "../../service/theme";
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.modal-view-theme {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    gap: 1rem;
}

@media (max-width: 960px) {
    .modal-view-theme {
        grid-template-columns: 1fr;
    }
}
</style>

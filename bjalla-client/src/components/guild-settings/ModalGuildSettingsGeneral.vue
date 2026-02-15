<template>
    <div class="modal-view-general">
        Guild Settings
    </div>
</template>

<script setup lang="ts">
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

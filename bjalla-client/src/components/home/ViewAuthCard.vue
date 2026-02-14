<template>
    <div class="view-profile">
        <div class="card">
            <div class="card-body">
                {{ pb.authStore.model?.email }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";

</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-profile {
    display: flex;
    flex-direction: column;
    margin: 1rem;
    width: 100%;
}
</style>

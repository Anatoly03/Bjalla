<template>
    <div class="view-sidebar">
        <div class="view-sidebar-content">
            <ViewJoinedGuildList />
            <div class="separator"></div>
            <ViewChannelList :key="$route.params.guild" />
        </div>
        <ViewAuthCard />
    </div>
</template>

<script setup lang="ts">
import ViewAuthCard from "./ViewAuthCard.vue";
import ViewChannelList from "./ViewChannelList.vue";
import ViewJoinedGuildList from "./ViewJoinedGuildList.vue";
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-sidebar {
    display: flex;
    flex-direction: column;
    height: 100%;
    border-right: 1px solid black;

    width: 300px;

    .view-sidebar-content {
        display: flex;
        flex-direction: row;
        flex: 1;
        overflow: hidden;

        .separator {
            width: 1px;
            background: black;
            height: calc(100% - 1rem);
        }
    }
}
</style>
<template>
    <div class="view-channels">
        <router-link :to="`/${route.params.guild}/${channel.id}`" v-for="channel in channels" :key="channel.id" class="channel-item" :class="{ active: route.params.channel === channel.id }">
            {{ channel.name }}
        </router-link>
        <div @click="openGuildSettings" class="channel-item view-guilds-create">
            <font-awesome-icon icon="fa-solid fa-gear" />
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { useModalRoute } from "@vmrh/core";

const { openModal } = useModalRoute();

/**
 * Route instance to get the current guild and channel from the URL.
 */
const route = useRoute();

/**
 * Channels of the current guild.
 */
const channels = ref<any[]>([]);

/**
 * Opens the guild settings modal.
 */
async function openGuildSettings() {
    await openModal("GuildSettings");
}

onMounted(async () => {
    if (!route.params.guild) return;

    const resultList = await pb.collection("channels").getList(1, 50, {
        requestKey: 'channels-' + route.params.guild, // cache results per guild
        filter: `guild="${route.params.guild}"`,
    });

    console.log('Fetched Channels', resultList.items);

    channels.value = resultList.items;
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-channels {
    display: flex;
    flex: 1;
    flex-direction: column;
    padding: 0.5em;
    gap: 0.25rem;
    overflow-y: scroll;

    .channel-item {
        display: flex;
        align-items: center;
        padding: 0.5rem;
        border-radius: 0.25rem;
        color: inherit;
        text-decoration: none;

        &:hover {
            background: var(--theme-bg-1, #e0c9b7);
        }

        &.active {
            background: var(--theme-bg-2, #ddd);
        }
    }
}
</style>

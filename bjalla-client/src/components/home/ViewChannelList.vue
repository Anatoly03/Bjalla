<template>
    <div class="view-channels">
        <router-link :to="`/${route.params.guild}/${row.expand.channel.id}`" v-for="row in channels" :key="row.id" class="channel-item" :class="{ active: route.params.channel === row.expand.channel.id }">
            {{ row.expand.channel.name }}
        </router-link>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

/**
 * Route instance to get the current guild and channel from the URL.
 */
const route = useRoute();

/**
 * Channels of the current guild.
 */
const channels = ref<any[]>([]);

onMounted(async () => {
    if (!route.params.guild) return;

    const resultList = await pb.collection("guild_channels").getList(1, 50, {
        expand: "channel",
        filter: `guild="${route.params.guild}"`,
    });

    console.log(resultList);

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

    .channel-item {
        display: flex;
        align-items: center;
        padding: 0.5rem;
        border-radius: 0.25rem;
        color: inherit;
        text-decoration: none;

        &:hover {
            background: #eee;
        }

        &.active {
            background: #ddd;
        }
    }
}
</style>

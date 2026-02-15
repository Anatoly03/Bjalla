<template>
    <div class="view-guilds">
        <router-link :to="`/${row.expand.guild.id}`" v-for="row in guilds" :key="row.id" class="view-guilds-item"  :class="{ active: route.params.guild === row.expand.guild.id }">
            <img
                v-if="row.expand.guild.logo"
                :src="pb.files.getURL(row.expand.guild, row.expand.guild.logo)"
                alt="logo"
                class="guild-logo"
                />
            <span v-else class="guild-capital">{{ row.expand.guild.name.charAt(0).toUpperCase() }}</span>
        </router-link>
        <div @click="openGuildProposal" class="view-guilds-item view-guilds-create">
            <font-awesome-icon icon="fa-solid fa-circle-plus" />
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useModalRoute } from "@vmrh/core";

const router = useRouter();
const { openModal } = useModalRoute();

/**
 * Route instance to get the current guild and channel from the URL.
 */
const route = useRoute();

/**
 * Guilds the user is a member of.
 */
const guilds = ref<any[]>([]);

/**
 * Opens the guild proposal modal.
 */
async function openGuildProposal() {
    const currentPath = router.currentRoute.value.fullPath;

    await openModal("CreateGuild");
    // hotfix: keep the visible URL unchanged while modal is open.
    // setTimeout(() => window.history.replaceState(window.history.state, "", currentPath), 0);
}

/**
 * Fetch the guilds the user is a member of on component mount.
 */
onMounted(async () => {
    const resultList = await pb.collection('guild_members').getList(1, 50, {
        requestKey: 'guilds', // cache results
        filter: `user="${pb.authStore.record?.id}"`,
        expand: 'guild',
    });

    guilds.value = resultList.items;
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-guilds {
    display: flex;
    flex-direction: column;
    padding-top: 0.5em;
    margin-left: 0.5em;
    padding-right: 0.5em;
    gap: 0.5rem;
    overflow-y: scroll;

    .view-guilds-item {
        display: flex;
        justify-content: center;
        align-items: center;
        box-sizing: border-box;
        border-radius: 0.5rem;
        text-decoration: none;
        
        width: 4rem;
        aspect-ratio: 1;

        color: inherit;

        &:hover {
            background: var(--theme-bg-1, #e0c9b7);
        }

        &.active {
            background: var(--theme-bg-2, #ddd);
        }

        .guild-logo {
            aspect-ratio: 1;
            border-radius: 4px;
            overflow: hidden;

            .logo-image {
                width: 100%;
                height: 100%;
                object-fit: cover;
            }
        }
    }

    .view-guilds-create {
        border: 1px solid #0a5;
    }
}
</style>

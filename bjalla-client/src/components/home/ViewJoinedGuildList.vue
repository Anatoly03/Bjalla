<template>
    <div class="view-guilds">
        <router-link :to="`/${row.expand.guild.id}`" v-for="row in guilds" :key="row.id" class="view-guilds-item"  :class="{ active: route.params.guild === row.expand.guild.id }">
            <font-awesome-icon icon="fa-regular fa-comment" />
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
 * Guilds the user is a member of.
 */
const guilds = ref<any[]>([]);

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
    border-right: 1px solid black;
    gap: 0.5rem;
    overflow-y: scroll;

    .view-guilds-item {
        display: flex;
        justify-content: center;
        align-items: center;
        box-sizing: border-box;
        border-radius: 0.5rem;
        
        width: 4rem;
        aspect-ratio: 1;

        color: inherit;

        &:hover {
            background: #eee;
        }

        &.active {
            background: #ddd;
        }
    }
}
</style>

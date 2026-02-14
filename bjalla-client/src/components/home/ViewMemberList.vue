<template>
    <div class="view-members">
        <span v-for="row in members" :key="row.id" class="view-member-item">
            {{ row.expand.user.name }}
        </span>
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
 * Members of the current guild.
 */
const members = ref<any[]>([]);

/**
 * Fetch the members of the current guild on component mount.
 */
onMounted(async () => {
    const resultList = await pb.collection('guild_members').getList(1, 50, {
        requestKey: 'members-' + route.params.guild, // cache results per guild
        filter: `guild="${route.params.guild}"`,
        expand: 'user',
    });

    console.log(resultList);

    members.value = resultList.items;
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-members {
    display: flex;
    flex-direction: column;
    padding: 0.5em;
    gap: 0.25rem;

    padding-right: 0.5em;
    border-left: 1px solid black;
    box-sizing: border-box;
    gap: 0.5rem;

    height: 100%;
    width: 240px;

    .view-member-item {
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

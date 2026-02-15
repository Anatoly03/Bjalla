<template>
    <div class="view-members">
        <span v-for="row in members" :key="row.id" class="view-member-item" @click="openMemberInfo($event, row)">
            <span class="avatar">
                <img
                    v-if="row.expand.user.avatar"
                    :src="pb.files.getURL(row.expand.user, row.expand.user.avatar)"
                    alt="avatar"
                    class="avatar-image"
                    />
            </span>
            <span class="username">
                {{ row.expand.user.name }}
            </span>
        </span>
    </div>
</template>

<script setup lang="ts">
import { useModalRoute } from "@vmrh/core";
import pb from "../../service/pocketbase";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const { openModal } = useModalRoute();

/**
 * Route instance to get the current guild and channel from the URL.
 */
const route = useRoute();

/**
 * Members of the current guild.
 */
const members = ref<any[]>([]);

/**
 * Opens the member info modal when a member is clicked.
 */
async function openMemberInfo(event: MouseEvent, member: any) {
    const target = event.currentTarget as HTMLElement | null;
    const rect = target?.getBoundingClientRect();
    const position = rect
        ? { x: rect.left + 22, y: rect.top + rect.height / 2 }
        : { x: event.clientX, y: event.clientY };
    
    await openModal("ViewMemberProfile", {
        data: {
            user_id: member.expand.user.id,
            guild_id: route.params.guild,
            position,
        },
    });
}

/**
 * Fetch the members of the current guild on component mount.
 */
onMounted(async () => {
    const resultList = await pb.collection("guild_members").getList(1, 50, {
        requestKey: "members-" + route.params.guild, // cache results per guild
        filter: `guild="${route.params.guild}"`,
        expand: "user",
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
    overflow-y: scroll;

    border-left: 1px solid black;
    box-sizing: border-box;

    height: 100%;
    width: 240px;

    padding-top: 10px;

    .view-member-item {
        display: flex;
        align-items: center;
        padding: 0.25rem;
        color: inherit;
        text-decoration: none;

        &:hover {
            background: var(--theme-bg-1, #e0c9b7);
        }

        &.active {
            background: var(--theme-bg-2, #ddd);
        }

        .avatar {
            width: 32px;
            height: 32px;
            border-radius: 4px;
            overflow: hidden;
            margin-right: 0.5rem;

            .avatar-image {
                width: 100%;
                height: 100%;
                object-fit: cover;
            }
        }
    }
}
</style>

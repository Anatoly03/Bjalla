<template>
    <div class="view-channel-messages">
        <div v-for="message in messages" :key="message.id" class="message">
            <strong>{{ message.expand?.author?.name ?? '???' }}</strong>
            <span>:</span>
            <span class="fill">{{ message.content }}</span>
              <button v-if="canDeleteMessages || message.expand?.author?.id === pb.authStore.record?.id" @click="deleteMessage(message.id)">Delete</button>
        </div>
        <div class="load-early-messages">
            <button
                ref="loadButton"
                @click="currentPage++; loadMessages()"
                :disabled="currentPage >= lastPage"
                >Load earlier messages</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onBeforeUnmount, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { RecordModel, RecordSubscription, UnsubscribeFunc } from "pocketbase";

/**
 * Route instance to get the current channel from the URL.
 */
const route = useRoute();

/**
 * Messages of the current channel.
 */
const messages = ref<any[]>([]);

// Pagination
const currentPage = ref(1);
const lastPage = ref(2);
const pageSize = 30;

const loadButton = ref<HTMLButtonElement | null>(null);
const isLoading = ref(false);
const loadObserver = ref<IntersectionObserver | null>(null);

/**
 * If true, the user can delete messages in the current channel.
 */
const canDeleteMessages = ref(false);

/**
 * Messages of the current channel.
 */
const unsubscribeFunc = ref<UnsubscribeFunc>(async () => {});

/**
 * Handle real-time message events for the current channel.
 */
function receiveMessage(event: RecordSubscription<RecordModel>) {
    const record = event.record;
    if (record.channel !== route.params.channel) return; // ignore messages from other channels

    switch (event.action) {
        case "create":
            messages.value.unshift(record);
            break;
        case "delete":
            messages.value = messages.value.filter((m) => m.id !== record.id);
            break;
        case "update":
            const index = messages.value.findIndex((m) => m.id === record.id);
            if (index !== -1) {
                messages.value[index] = record;
            }
            break;
    }
}

/**
 * Deletes a message by its ID.
 */
async function deleteMessage(messageId: string) {
    try {
        await pb.collection("messages").delete(messageId);
    } catch (error) {
        console.error("Failed to delete message:", error);
    }
}

async function loadMessages() {
    if (isLoading.value) return;
    isLoading.value = true;
    // Fetch messages for the current channel on mount
    try {
        const resultList = await pb.collection("messages").getList(currentPage.value, pageSize, {
            requestKey: "messages-" + route.params.channel, // cache results per channel
            sort: "-created", // sort by creation date descending
            filter: `channel="${route.params.channel}"`,
            expand: "author",
        });

        lastPage.value = resultList.totalPages;
        messages.value.push(...resultList.items);

        // sort messages by creation date descending so column-reverse shows oldest at the top
        messages.value.sort((a, b) => new Date(b.created).getTime() - new Date(a.created).getTime());
    } finally {
        isLoading.value = false;
    }
}

/**
 * Fetch messages for the current channel on mount and listen
 * for real-time updates to messages in the current channel.
 */
onMounted(async () => {
    await loadMessages();

    loadObserver.value = new IntersectionObserver((entries) => {
        const [entry] = entries;
        if (!entry?.isIntersecting) return;
        if (currentPage.value >= lastPage.value) return;
        if (isLoading.value) return;

        currentPage.value += 1;
        void loadMessages();
    });

    if (loadButton.value) {
        loadObserver.value.observe(loadButton.value);
    }

    // Listen for real-time updates to messages in the current channel.
    unsubscribeFunc.value = await pb.collection("messages").subscribe("*", receiveMessage, {
        expand: "author", // expand author on real-time events as well
    });

    // Test auth user permissions
    const membership = await pb.collection("guild_members").getFirstListItem(`guild="${route.params.guild}" && user="${pb.authStore.record?.id}"`, {
        expand: "roles",
    });

    // Test individual permissions
    const roles = membership.expand?.roles || [];
    canDeleteMessages.value = roles.some((r: any) => r.is_admin);
});

onBeforeUnmount(() => {
    // Unsubscribe from real-time updates when the component is unmounted.
    unsubscribeFunc.value();

    if (loadObserver.value && loadButton.value) {
        loadObserver.value.unobserve(loadButton.value);
    }
    loadObserver.value?.disconnect();
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-channel-messages {
    display: flex;
    flex-direction: column-reverse; // reverse to show newest messages at the bottom
    flex: 1;
    gap: 0.25em;
    overflow: scroll;

    .message {
        display: flex;
        padding: 0.2em 0.5em;

        &:hover {
            background: var(--theme-bg-1, #e0c9b7);
        }

        .fill {
            flex: 1;
            margin-left: 0.5em;
        }
    }

    .load-early-messages {
        display: flex;
        justify-content: center;
        padding: 0.5em;
    }
}
</style>

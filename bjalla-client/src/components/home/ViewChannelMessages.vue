<template>
    <div class="view-channel-messages">
        <div v-for="message in messages" :key="message.id" class="message">
            <strong>{{ message.expand.author.name }}</strong
            >: {{ message.content }}
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onBeforeUnmount, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { RecordModel, RecordSubscription } from "pocketbase";

/**
 * Route instance to get the current channel from the URL.
 */
const route = useRoute();

/**
 * Messages of the current channel.
 */
const messages = ref<any[]>([]);

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
 * Fetch messages for the current channel on mount and listen
 * for real-time updates to messages in the current channel.
 */
onMounted(async () => {
    // Fetch messages for the current channel on mount
    const resultList = await pb.collection("messages").getList(1, 50, {
        requestKey: "messages-" + route.params.channel, // cache results per channel
        sort: "-created", // sort by creation date descending
        filter: `channel="${route.params.channel}"`,
        expand: "author",
    });
    messages.value = resultList.items;

    // Listen for real-time updates to messages in the current channel.
    pb.collection("messages").subscribe("*", receiveMessage, {
        expand: "author", // expand author on real-time events as well
    });
});

onBeforeUnmount(() => {
    // Unsubscribe from real-time updates when the component is unmounted.
    pb.collection("messages").unsubscribe("*");
});
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-channel-messages {
    display: flex;
    flex-direction: column-reverse; // reverse to show newest messages at the bottom
    flex: 1;
    gap: 0.25em;

    .message {
        display: flex;
        padding: 0.2em 0.5em;

        &:hover {
            background: #eee;
        }
    }
}
</style>

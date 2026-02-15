<template>
    <div class="view-channel">
        <ViewChannelMessages :key="$route.params.channel as any" />
        <input type="text" placeholder="Message channel" class="view-channel-input" @keypress.enter="sendMessage" />
    </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import pb from "../../service/pocketbase";
import ViewChannelMessages from "./ViewChannelMessages.vue";

/**
 * Route instance to get the current guild and channel from the URL.
 */
const route = useRoute();

/**
 * Push upstream message to the current channel when the user
 * presses Enter in the input field.
 */
async function sendMessage(event: KeyboardEvent) {
    const input = event.target as HTMLInputElement;
    const content = input.value.trim();
    if (!content) return;

    const channel = route.params.channel;
    const record = await pb.collection("messages").create({
        content,
        channel,
        author: pb.authStore.model?.id,
    });
    
    input.value = "";
    console.debug(record);
}
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.view-channel {
    display: flex;
    flex-direction: column;
    flex: 1;
    height: 100%;
    box-sizing: border-box;
}

.view-channel-input {
    padding: 0.5em;
    margin: 0.5em;
    border: 1px solid #ccc;
    border-radius: 0.25rem;
}
</style>

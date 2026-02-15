<template>
    <div class="view-auth card">
        <div class="card-header">Authenticate</div>
        <div class="card-body">
            <!-- Password Authentication Region -->
            <form class="password-region" v-if="authMethods.password.enabled" @submit.prevent="authWithPassword">
                <input id="username-input" type="text" v-model="username" :placeholder="usernamePlaceholder" autocomplete="username email" />
                <input name="password" type="password" v-model="password" placeholder="password" autocomplete="current-password" />
                <input type="submit" value="Login" />
            </form>

            <!-- OAuth2 Authentication Region -->
            <div class="oauth-region" v-if="authMethods.oauth2.providers.length > 0">
                <div class="oauth-header" v-if="authMethods.password.enabled">Or login with</div>
                <div class="oauth-buttons">
                    <button v-for="provider in authMethods.oauth2.providers" :key="provider.name" @click="authWithOAuth2(provider.name)" class="oauth-button">
                        <font-awesome-icon :icon="'fa-brands fa-' + provider.name.toLowerCase()" />
                    </button>
                </div>
            </div>

            <!-- Error Delivery Region -->
            <div v-if="errorReason" class="error">{{ errorReason }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
import pb from "../service/pocketbase";
import { AuthMethodsList, RecordAuthResponse, RecordModel } from "pocketbase";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";

/**
 * Router instance for navigation after successful authentication.
 */
const router = useRouter();

/**
 * Redirect to home if the user is already authenticated.
 */
if (pb.authStore.isValid) {
    router.push("/");
}

/**
 * Authentication methods, fetched from PocketBase and can be disabled in the
 * production database on the fly.
 */
const authMethods = ref<AuthMethodsList>({
    mfa: { enabled: false, duration: 0 },
    otp: { enabled: false, duration: 0 },
    password: { enabled: false, identityFields: [] },
    oauth2: { enabled: false, providers: [] },
});

/**
 * Placeholder for the username.
 */
const usernamePlaceholder = ref("identity");

/**
 * V-model for username
 */
const username = ref("");

/**
 * V-model for password
 */
const password = ref("");

/**
 * Error handling for authentication, e.g. wrong password, network error, etc.
 */
const errorReason = ref("");

/**
 * Fetch authentication methods from PocketBase on component mount and update the
 * state. (This allows dynamic enabling/disabling of auth methods without code changes.)
 */
onMounted(async () => {
    let methods = await pb.collection("users").listAuthMethods();

    console.debug("Auth methods:", methods);

    authMethods.value = methods;

    // create placeholder based on the identity fields separated by comma, last one is "or" separated
    // e.g. "email, username or phone"
    usernamePlaceholder.value = "" + methods.password.identityFields.slice(0, -1).join(", ") + (methods.password.identityFields.length > 1 ? " or " : "") + methods.password.identityFields.slice(-1);
});

/**
 * Handle form submission for authentication.
 */
async function authWithPassword(_: SubmitEvent) {
    try {
        const authData = await pb.collection("users").authWithPassword(username.value, password.value);
        authenticate(authData);
    } catch (err: any) {
        console.error(err);
        errorReason.value = err.message || "An unknown error occurred during authentication.";
    }
}

/**
 * Handle OAuth2 authentication with the specified provider.
 */
async function authWithOAuth2(provider: string) {
    try {
        const authData = await pb.collection("users").authWithOAuth2({ provider });
        authenticate(authData);
    } catch (err: any) {
        console.error(err);
        errorReason.value = err.message || "An unknown error occurred during authentication.";
    }
}

/**
 * Authenticate the user with the given auth data.
 */
function authenticate(authData: RecordAuthResponse<RecordModel>) {
    console.debug("Authenticated:", authData);
    router.push("/");
}

</script>

<style lang="scss" scoped>
@use "../assets/theme.scss";
@use "../assets/ui.scss";

.view-auth {
    display: flex;
    flex: 1;
    justify-content: center;
    align-items: center;

    .password-region {
        display: flex;
        flex-direction: column;
        gap: 0.3rem;

        input {
            width: 100%;
            padding: 0.5rem;
            box-sizing: border-box;
        }
    }

    .oauth-region {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 0.5rem;

        .oauth-header {
            font-size: 0.9rem;
            color: gray;
        }

        .oauth-buttons {
            display: flex;
            gap: 0.5rem;

            .oauth-button {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 2.5rem;
                height: 2.5rem;
                border-radius: 50%;
                border: 1px solid black;
                background-color: white;
                cursor: pointer;

                &:hover {
                    background-color: #f0f0f0;
                }
            }
        }
    }
}
</style>

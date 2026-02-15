<template>
    <div class="modal-view-roles">
        <h2>Roles</h2>
        <table class="role-matrix">
            <thead>
                <tr class="permissions">
                    <th class="role-name-col" scope="col"></th>
                    <th class="permission-col" scope="col">Admin</th>
                    <th class="permission-col" scope="col">Default</th>
                    <th class="fill-col" scope="col"></th>
                    <th class="delete-col" scope="col"></th>
                </tr>
            </thead>
            <tbody>

                <tr v-for="role in roles" :key="role.id" class="role">
                    <th class="role-name-col">{{ role.name }}</th>
                    <td class="permission-col"><input type="checkbox" :checked="role.is_admin" disabled /></td>
                    <td class="permission-col"><input type="checkbox" :checked="role.is_default" disabled /></td>
                    <td class="fill-col"></td>
                    <td class="delete-col">
                        <button @click="deleteRole(role.id)" :disabled="!canEditRoles">
                            <font-awesome-icon icon="fa-solid fa-trash-can" />
                        </button>
                    </td>
                </tr>
                <tr>
                    <td><input type="text" placeholder="New role name" v-model="newRoleName" @keypress.enter="createRole" :disabled="!canEditRoles" /></td>
                    <td class="permission-col"><input type="checkbox" v-model="newRoleIsAdmin" :disabled="!canEditRoles" /></td>
                    <td class="permission-col"><input type="checkbox" v-model="newRoleIsDefault" :disabled="!canEditRoles" /></td>
                    <td class="fill-col"><button @click="createRole" :disabled="!canEditRoles">Create Role</button></td>
                    <td class="delete-col"></td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import pb from "../../service/pocketbase";
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

/**
 * Roles of the current guild.
 */
const roles = ref<any[]>([]);

/**
 * If true, the user can edit roles in the current guild.
 */
const canEditRoles = ref(false);

/**
 * Proposed new role
 */
const newRoleName = ref("");
const newRoleIsAdmin = ref(false);
const newRoleIsDefault = ref(false);

/**
 * Creates a new role with the name from the input field and default permissions.
 */
async function createRole() {
    const newRole = await pb.collection("guild_roles").create({
        name: newRoleName.value,
        is_admin: newRoleIsAdmin.value,
        is_default: newRoleIsDefault.value,
        guild: route.params.guild
    });

    // reset input fields
    newRoleName.value = "";
    newRoleIsAdmin.value = false;
    newRoleIsDefault.value = false;

    // add new role to the list
    roles.value.push(newRole);
}

/**
 * Deletes the role with the given ID and removes it from the list.
 */
async function deleteRole(roleId: string) {
    await pb.collection("guild_roles").delete(roleId);
    roles.value = roles.value.filter(role => role.id !== roleId);
}

/**
 * Sync roles from PocketBase on component mount.
 */
onMounted(async () => {
    try {
        const rolesList = await pb.collection("guild_roles").getList(1, 50, {
            filter: `guild = "${route.params.guild}"`,
        });
        roles.value = rolesList.items;
    } catch (error) {
        console.error("Failed to fetch roles:", error);
    }

    // Test auth user permissions
    const membership = await pb.collection("guild_members").getFirstListItem(`guild="${route.params.guild}" && user="${pb.authStore.record?.id}"`, {
        expand: "roles",
    });

    // Test individual permissions
    const userRoles = membership.expand?.roles || [];
    canEditRoles.value = userRoles.some((r: any) => r.is_admin);
})
</script>

<style lang="scss" scoped>
@use "../../assets/theme.scss";

.modal-view-roles {
    display: flex;
    flex-direction: column;
    gap: 1em;
}

.role-matrix {
    border-collapse: collapse;
    width: 100%;

    .permissions {
        font-weight: normal;
        color: #666;
    }

    .role {
        &:hover {
            background: var(--theme-bg-1, #e0c9b7);
        }

        &.active {
            background: var(--theme-bg-2, #ddd);
        }
    }

    th, td {
        padding: 0.5em;
        text-align: left;

        &.role-name-col {
            width: 200px;
        }

        &.permission-col {
            width: 64px;
            text-align: center;
        }

        &.fill-col {
            width: auto;
        }
    }
}
</style>

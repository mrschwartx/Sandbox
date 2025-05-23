<script setup>
import { Head, Link, router } from '@inertiajs/vue3';
import { ref } from 'vue';
import DashboardLayout from '@/Layouts/DashboardLayout.vue'
import Pagination from '@/Components/Pagination.vue';
import Swal from 'sweetalert2';

const props = defineProps({
    roles: {
        type: Object
    }
});

const search = ref('' || (new URL(document.location)).searchParams.get('q'));

const handleSearch = () => {
    router.get('/dashboard/roles', {
        q: search.value,
    });
}

const onDestroy = (id) => {
    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete it!'
    })
        .then((result) => {
            if (result.isConfirmed) {

                router.delete(`/dashboard/roles/${id}`);

                Swal.fire({
                    title: 'Deleted!',
                    text: 'Role deleted successfully.',
                    icon: 'success',
                    timer: 2000,
                    showConfirmButton: false,
                });
            }
        })
}
</script>

<template>
    <DashboardLayout>

        <Head title="Roles" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">
                    <div class="row">
                        <div class="col-md-12">
                            <div class="card border-0 rounded-3 shadow border-top-laravel">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-shield-alt"></i> ROLES</span>
                                </div>
                                <div class="card-body">
                                    <form @submit.prevent="handleSearch">
                                        <div class="input-group mb-3">

                                            <Link href="/dashboard/roles/create" v-if="hasAnyPermission(['roles.create'])"
                                                class="btn btn-primary input-group-text"> <i
                                                class="fa fa-plus-circle me-2"></i> NEW</Link>

                                            <input type="text" class="form-control" v-model="search"
                                                placeholder="search by role name...">

                                            <button class="btn btn-primary input-group-text" type="submit"> <i
                                                    class="fa fa-search me-2"></i> SEARCH</button>

                                        </div>
                                    </form>
                                    <table class="table table-striped table-bordered table-hover">
                                        <thead>
                                            <tr>
                                                <th scope="col">Role Name</th>
                                                <th scope="col" style="width:50%">Permissions</th>
                                                <th scope="col" style="width:20%">Actions</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr v-for="(role, index) in roles.data" :key="index">
                                                <td>{{ role.name }}</td>
                                                <td>
                                                    <span v-for="(permission, index) in role.permissions" :key="index"
                                                        class="badge badge-primary shadow border-0 ms-2 mb-2">
                                                        {{ permission.name }}
                                                    </span>
                                                </td>
                                                <td class="text-center">
                                                    <Link :href="`/dashboard/roles/${role.id}/edit`"
                                                        v-if="hasAnyPermission(['roles.edit'])"
                                                        class="btn btn-success btn-sm me-2"><i
                                                        class="fa fa-pencil-alt me-1"></i> EDIT</Link>
                                                    <button @click.prevent="onDestroy(role.id)"
                                                        v-if="hasAnyPermission(['roles.delete'])"
                                                        class="btn btn-danger btn-sm"><i class="fa fa-trash"></i>
                                                        DELETE</button>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                    <Pagination :links="roles.links" align="end" />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </main>

    </DashboardLayout>
</template>

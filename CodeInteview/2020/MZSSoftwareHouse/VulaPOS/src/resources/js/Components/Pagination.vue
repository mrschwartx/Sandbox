<script setup>
import { Link } from '@inertiajs/vue3';

defineProps({
    links: {
        type: Array,
    },
    align: {
        type: String
    }
});

</script>

<template>
    <nav>
        <ul :class="`pagination justify-content-${align} mb-0`">
            <li :class="[
                'page-item',
                link.url == null ? 'disabled' : '',
                link.active ? 'active' : '',
            ]" v-for="(link, index) in links" :key="index">

                <Link class="page-link" :href="link.url === null ? '#' : link.url" v-html="link.label">
                </Link>

            </li>
        </ul>
    </nav>
</template>

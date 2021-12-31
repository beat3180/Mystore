<template>
  <global-header />
  <index-page />
  <global-footer />
  <router-view />
</template>
<script>
import GlobalHeader from "./components/common/area/header.vue";
import GlobalFooter from "./components/common/area/footer.vue";
import IndexPage from "./pages/index.vue";

export default {
  components: {
    GlobalHeader,
    GlobalFooter,
    IndexPage,
  },
};
</script>

<style>
</style>

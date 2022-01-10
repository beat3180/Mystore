<template>
  <router-view />
</template>
<script>
import { onMounted } from "vue";
import axios from "axios";
import { useStore } from "vuex";

export default {
  head() {
    return {
      title: "Top",
    };
  },
  data() {
    return {};
  },

  mounted() {
    this.init();
  },
  methods: {
    init: async function () {},
  },

  setup() {
    const store = useStore();

    onMounted(async () => {
      try {
        // user情報を取得
        // ログイン情報は、Cookieに保存してあるので、
        // リクエストするだけでOK
        const { data } = await axios.get("user");
        // actionsに設定したパラメータ名を設定
        await store.dispatch("setAuth", true);
        await store.dispatch("setUserId", data.ID);
      } catch (e) {
        await store.dispatch("setAuth", false);
        await store.dispatch("setUserId", 0);
      }
    });

    return {};
  },
};
</script>

<style>
</style>

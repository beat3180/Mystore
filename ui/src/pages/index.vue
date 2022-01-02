<template>
  <section id="wrapper">
    <div id="main">aaaa{{ user_id }}</div>
  </section>
</template>

<script>
import { ref, onMounted } from "vue";
import axios from "axios";
import { useStore } from "vuex";

export default {
  name: "Index",

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
    const user_id = ref("");
    const store = useStore();

    onMounted(async () => {
      try {
        // user情報を取得
        // ログイン情報は、Cookieに保存してあるので、
        // リクエストするだけでOK
        const { data } = await axios.get("user");
        user_id.value = data.ID;
        // actionsに設定したパラメータ名を設定
        await store.dispatch("setAuth", true);
      } catch (e) {
        await store.dispatch("setAuth", false);
      }
    });

    return {
      user_id,
    };
  },
};
</script>

<style scoped>
</style>

<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div id="main">aaaa</div>
    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";

export default {
  name: "Index",
  components: {
    Header,
    Footer,
  },

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
};
</script>

<style scoped>
</style>

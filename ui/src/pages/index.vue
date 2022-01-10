<template>
  <section id="wrapper">
    <Header ref="header_area" />

    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";

export default {
  name: "Post",
  components: {
    Header,
    Footer,
  },

  head() {
    return {
      title: "Post",
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
.main {
  height: 100%;
  width: 100%;
}
</style>

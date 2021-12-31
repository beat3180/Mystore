<template>
  <div class="footer">
    <div class="mystore">Mystore</div>
  </div>
</template>
<script>
export default {
  components: {},
  props: [],
  data() {
    return {};
  },

  computed: {},
  watch: {},
  mounted() {
    this.init();
  },

  methods: {
    init: async function () {},
  },
};
</script>
<style scoped>
ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.li {
  display: block;
  padding: 0;
  margin: 0;
}

.footer {
  padding: 20px 10px;
  background-color: #d9f0f7;
}
</style>

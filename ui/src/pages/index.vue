<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div class="main">
      <input rype="text" v-model="address" />
      <button type="button" @click="map_search">検索</button>
      <GMapMap
        :center="center"
        :zoom="7"
        map-type-id="terrain"
        style="width: 500px; height: 300px"
      >
        <GMapCluster>
          <GMapMarker
            :key="index"
            v-for="(m, index) in markers"
            :position="m.position"
            :clickable="true"
            :draggable="true"
            @click="center = m.position"
          />
        </GMapCluster>
      </GMapMap>
    </div>
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
    return {
      center: { lat: 51.093048, lng: 6.84212 },
      markers: [
        {
          position: {
            lat: 51.093048,
            lng: 6.84212,
          },
        },
      ],
    };
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

<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div class="main">
      <div class="content_grid">
        <div class="item_grid">
          <div class="item_title">
            {{ my_store.Title }}
          </div>
          <div class="item_comment">
            {{ my_store.Comment }}
          </div>
          <div class="item_image">
            <img :src="my_store.ImagePath" />
          </div>
          <div class="item_gmap">
            <div
              ref="gmap"
              class="gmap"
              style="height: 500px; width: 800px"
            ></div>
          </div>
          <div></div>
        </div>
      </div>
    </div>

    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";
import axios from "axios";

export default {
  name: "Content",
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
      my_store: {},
      myLatLng: { lat: 0, lng: 0 },
    };
  },
  mounted() {
    this.init();
  },

  methods: {
    init: async function () {
      //コンテンツの情報を取得
      await this.get_post_id();
      //Googleマップを取得
      await this.get_gmap();

      console.log(this.my_store.Lng);
    },
    get_post_id: async function () {
      await axios
        .get("/post/get/content/" + this.$route.params.id.substr(3))
        .then((res) => {
          console.log(res);
          this.my_store = res.data;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        });
    },
    get_gmap: async function () {
      this.myLatLng.lat = this.my_store.Lat;
      this.myLatLng.lng = this.my_store.Lng;

      let timer = setInterval(() => {
        if (window.google) {
          clearInterval(timer);
          const map = new window.google.maps.Map(this.$refs.gmap, {
            center: this.myLatLng,
            zoom: 12,
          });
          new window.google.maps.Marker({
            position: this.myLatLng,
            map,
          });
        }
      }, 500);
    },
  },
};
</script>

<style scoped>
.main {
  height: 100%;
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
}
.content_grid {
  display: grid;
  grid-auto-columns: auto;
  grid-auto-rows: auto;
  text-align: center;
}

.item_grid {
  display: grid;
  grid-template-columns: auto auto;
  grid-auto-rows: auto;
  background-color: white;
  grid-gap: 10px;
  margin: 10px;
}
.item_title {
  grid-column: 1/3;
  grid-row: 1/2;
}
.item_gmap {
  grid-column: 1/3;
  grid-row: 4/5;
}
</style>

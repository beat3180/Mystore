<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div class="main">
      <div class="list_grid">
        <div v-for="(mystore, key) in my_store_list" v-bind:key="key">
          <div class="item_grid">
            <div class="background finger" v-on:click="get_content(mystore.ID)">
              <div>
                {{ mystore.Title }}
              </div>
              <img :src="mystore.ImagePath" />
            </div>
          </div>
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
  name: "Index",
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
      my_store_list: [],
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init: async function () {
      await this.get_mystore();
    },
    get_mystore() {
      axios
        .post("/post/get/mystore", {
          user_id: 1,
        })
        .then((res) => {
          console.log(res);
          this.my_store_list = res.data;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        });
    },
    get_content(id) {
      //トップページに戻る
      this.$router.push("/content/id=" + id);
    },
  },
};
</script>

<style scoped>
.finger {
  cursor: pointer;
}
.main {
  height: 100%;
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
}
.list_grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: auto;
  text-align: center;
}
.item_grid {
  display: grid;
  grid-auto-columns: auto;
  grid-auto-rows: auto;
  background-color: white;
  grid-gap: 10px;
  margin: 10px;
}
</style>

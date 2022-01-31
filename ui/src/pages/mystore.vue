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
          <div>
            <button
              class="btn01"
              v-show="mystore.PublicFlag"
              v-on:click="update_public(mystore.ID)"
            >
              非公開にする
            </button>
            <button
              class="btn03"
              v-show="!mystore.PublicFlag"
              v-on:click="update_hide(mystore.ID)"
            >
              公開する
            </button>
          </div>
          <div>
            <button class="btn02" v-on:click="delete_mystore(mystore.ID)">
              削除
            </button>
          </div>
        </div>
      </div>
    </div>
    <Loader ref="loader" :nowLoading="nowLoading" />
    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Loader from "@/components/common/layer/loader.vue";
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";
import axios from "axios";

export default {
  name: "Index",
  components: {
    Header,
    Footer,
    Loader,
  },

  head() {
    return {
      title: "Post",
    };
  },
  data() {
    return {
      my_store_list: [],
      user_id: 0,
      public_flag: false,
    };
  },
  watch: {
    public_flag: function () {
      this.get_mystore();
    },
  },
  mounted() {
    this.init();
  },
  methods: {
    init: async function () {
      await this.get_user_id();
      await this.get_mystore();
    },
    get_user_id: async function () {
      await axios
        .get("/user")
        .then((res) => {
          console.log(res);
          this.user_id = res.data.ID;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        });
    },
    get_mystore: async function () {
      await axios
        .post("/post/get/mystore", {
          user_id: this.user_id,
        })
        .then((res) => {
          console.log(res);
          this.my_store_list = res.data;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        });
    },

    update_public: async function (content_id) {
      if (this.nowLoading) {
        return;
      }
      this.nowLoading = true;
      await axios
        .post("/post/mystore/public/update", {
          id: content_id,
        })
        .then((res) => {
          console.log(res);

          this.my_store_list.public_flag = res.data.public_flag;
          this.public_flag = !this.public_flag;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        }) //最後に処理する記述
        .finally(() => {
          this.nowLoading = false;
        });
    },
    update_hide: async function (content_id) {
      if (this.nowLoading) {
        return;
      }
      this.nowLoading = true;
      await axios
        .post("/post/mystore/hide/update", {
          id: content_id,
        })
        .then((res) => {
          console.log(res);
          this.my_store_list.public_flag = res.data.public_flag;
          this.public_flag = !this.public_flag;
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        }) //最後に処理する記述
        .finally(() => {
          this.nowLoading = false;
        });
    },
    delete_mystore: async function (content_id) {
      if (this.nowLoading) {
        return;
      }
      this.nowLoading = true;
      await axios
        .post("/post/mystore/delete", {
          id: content_id,
        })
        .then((res) => {
          console.log(res);
          this.get_mystore();
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        }) //最後に処理する記述
        .finally(() => {
          this.nowLoading = false;
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

.btn01 {
  margin: 2px 0;
  padding: 3px 10px;
  min-width: 120px;
  border: 3px solid yellow;
  border-radius: 100px;
  background-color: yellow;
  color: black;
  font-weight: bold;
  transition-duration: 0.3s;
  cursor: pointer;
}

.btn01:hover {
  background-color: #fff;
  color: black;
}

.btn03 {
  margin: 2px 0;
  padding: 3px 10px;
  min-width: 120px;
  border: 3px solid #0000ff;
  border-radius: 100px;
  background-color: #0000ff;
  color: #fff;
  font-weight: bold;
  transition-duration: 0.3s;
  cursor: pointer;
}

.btn03:hover {
  background-color: #fff;
  color: #303030;
}

.btn02 {
  margin: 2px 0;
  padding: 3px 10px;
  min-width: 120px;
  border: 3px solid red;
  border-radius: 100px;
  background-color: #fff;
  color: #303030;
  font-weight: bold;
  transition-duration: 0.3s;
  cursor: pointer;
}

.btn02:hover {
  background-color: red;
  color: #fff;
}
</style>

<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div class="post_container">
      <h1>Mystoreの投稿</h1>
      <p v-if="error_message" class="red_text">{{ error_message }}</p>
      <h2>お店の名前(必須)</h2>
      <input class="title_text" type="text" v-model="title" />
      <h2>コメント</h2>
      <textarea rows="10" cols="60" v-model="comment"></textarea>
      <h2>画像(必須)</h2>
      <input type="file" v-on:change="fileSelected" />
      <div>
        <img class="image" :src="url" v-show="url" />
      </div>
      <h2>お店の場所(必須)</h2>
      <vue-google-autocomplete
        class="gmap_text"
        id="map"
        classname="form-control"
        placeholder="住所を入力してください"
        v-on:placechanged="getAddressData"
        v-model="address_search"
      >
      </vue-google-autocomplete>

      <button type="button" @click="getlatlng">お店を検索</button>
      <div class="gmap_container">
        <div id="gmap" class="gmap" style="height: 500px; width: 800px"></div>
      </div>
    </div>
    <div class="center">
      <button class="btn01 button" v-on:click="mystore_post">投稿</button>
    </div>

    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";
import VueGoogleAutocomplete from "vue-google-autocomplete";
import axios from "axios";

export default {
  name: "Post",
  components: {
    Header,
    Footer,
    VueGoogleAutocomplete,
  },

  head() {
    return {
      title: "Post",
    };
  },
  data() {
    return {
      url: "",
      error_message: "",
      public_flag: true,
      comment: "",
      title: "",
      address: "",
      lng: "",
      lat: "",
      fileInfo: "",
      user: "",
      showImage: false,
      map: {},
      geocoder: {},
      label_name: "",
      address_search: "",
      marker: null,
      options: {
        zoom: 12,
        center: { lng: 139.7690174, lat: 35.6803997 },
        mapTypeId: "roadmap",
      },
    };
  },

  mounted() {
    this.init();

    // this.map = new window.google.maps.Map(document.getElementById("map"), {
    //   center: { lat: -34.397, lng: 150.644 },
    //   zoom: 8,
    // });

    this.map = new window.google.maps.Map(
      document.getElementById("gmap"),
      this.options
    );
    this.geocoder = new window.google.maps.Geocoder();
  },

  methods: {
    init: async function () {
      await this.get_user();
    },
    get_user() {
      axios
        .get("/user")
        .then((res) => {
          this.user = res.data;
          console.log(this.user);
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
        });
    },
    //ファイル情報取得
    fileSelected(event) {
      this.fileInfo = event.target.files[0];
      this.url = URL.createObjectURL(this.fileInfo);
    },
    mystore_post() {
      const posts = {
        user_id: this.user.ID === undefined ? 0 : this.user.ID,
        title: this.title,
        comment: this.comment,
        lng: this.lng,
        lat: this.lat,
        address: this.address,
        public_flag: this.public_flag,
      };
      if (
        !this.title ||
        !this.lng ||
        !this.lat ||
        !this.address ||
        !this.fileInfo
      ) {
        return (this.error_message = "必須項目に入力がありません");
      }

      const formData = new FormData();
      formData.append("file", this.fileInfo);
      const postsData = JSON.stringify(posts);
      formData.append("postsData", postsData);

      axios
        .post("/post", formData, {
          headers: {
            "content-type": "multipart/form-data",
          },
        })
        .then((res) => {
          console.log(res);
          //トップページに戻る
          this.$router.push("/");
        })
        .catch((e) => {
          console.log("axiosエラー e:" + e);
          this.error_message = e.response.data.message;
        });
    },

    getlatlng() {
      var vm = this;
      let adrs = this.address_search;
      //マーカーがある場合は削除
      if (this.marker !== null) {
        this.marker.setMap((this.map = null));
        this.map = new window.google.maps.Map(
          document.getElementById("gmap"),
          this.options
        );
      }

      if (adrs.length || adrs.route) {
        //予測変換した場合
        if (adrs.route) {
          adrs = adrs.route;
        }
        //var rslts = this.adrs_list;
        //var geocoder = new window.google.maps.Geocoder();
        // 第１引数はGeocoderRequest。住所⇒緯度経度座標の変換時はaddressプロパティを入れればOK。
        // 第２引数はコールバック関数。
        this.geocoder.geocode(
          {
            address: adrs,
          },
          function (results, status) {
            console.log(results);
            if (status == window.google.maps.GeocoderStatus.OK) {
              if (results[0].geometry) {
                vm.map.setCenter(results[0].geometry.location);
                results[0].geometry.location.lat();
                results[0].geometry.location.lng();

                // 緯度経度を取得
                var latlng = results[0].geometry.location;

                vm.lat = results[0].geometry.location.lat();
                vm.lng = results[0].geometry.location.lng();
                // 住所を取得(日本の場合だけ「日本, 」を削除)
                vm.address = results[0].formatted_address.replace(
                  /^日本、/,
                  ""
                );
                // rslts.push({ address: address, latlng: latlng });

                // var Options = {
                //   zoom: 12,
                //   center: latlng,
                //   mapTypeId: "roadmap",
                // };
                // var map = new window.google.maps.Map(
                //   document.getElementById("gmap"),
                //   Options
                // );

                // for (let i in rslts) {
                //   var marker = new window.google.maps.Marker({
                //     position: rslts[i].latlng,
                //     map: this.map,
                //   });

                //   let infoWindow = new window.google.maps.InfoWindow({
                //     content:
                //       rslts[i].address + "<br />" + rslts[i].latlng.toString(),
                //   });
                //   marker.addListener("click", function () {
                //     infoWindow.open(this.map, marker);
                //   });
                // }
                // console.log(results);
                // vm.marker = new window.google.maps.Marker({
                //   position: results[0].geometry.location,
                //   map: vm.map,
                // });
                vm.marker = new window.google.maps.Marker({
                  //map: vm.map,
                  position: results[0].geometry.location,
                });
                vm.marker.setMap(vm.map);

                let infoWindow = new window.google.maps.InfoWindow({
                  content: vm.address + "<br />" + latlng.toString(),
                });
                vm.marker.addListener("click", function () {
                  infoWindow.open(vm.map, vm.marker);
                });
              }
            } else {
              console.log("Geocodingに失敗しました。。。");
            }
          }
        );
      }
    },

    getAddressData: function (addressData) {
      this.address_search = addressData;
    },
    // initLocationSearch() {
    //   let self = this,
    //     autocomplete = new window.google.maps.places.Autocomplete(
    //       this.$refs.search[this.counter - 1],
    //       {
    //         fields: ["address_components"], // Reduce API costs by specifying fields
    //       }
    //     );

    //   autocomplete.addListener("place_changed", function () {
    //     let place = autocomplete.getPlace();
    //     if (place && place.address_components) {
    //       console.log(place);
    //       self.setSelectedCities();
    //     }
    //   });
    // },

    // mapSearch() {
    //   this.geocoder.geocode(
    //     {
    //       address: this.address,
    //     },
    //     (results, status) => {
    //       if (status === window.google.maps.GeocoderStatus.OK) {
    //         this.map.setCenter(results[0].geometry.location);
    //         // 緯度経度の取得
    //         // results[0].geometry.location.lat();
    //         // results[0].geometry.location.lng();
    //         this.marker = new window.google.maps.Marker({
    //           map: this.map,
    //           position: results[0].geometry.location,
    //         });
    //       }
    //     }
    //   );
    // },
  },
};
</script>

<style scoped>
.red_text {
  color: red;
}
.post_container {
  width: 100%;
  text-align: center;
  margin: 0 auto;
}

.gmap {
  margin: 0 auto;
}

.background_white {
  background-color: white;
}
.container {
  max-width: 700px;
  margin: 0 auto;
  padding: 20px;
  border: 2px solid #ddd; /*枠線*/
}

.title_text {
  max-width: 500px;
  height: 40px;
  width: 100%;
  font-size: 20px;
  border: 2px solid #ddd; /*枠線*/
  box-sizing: border-box; /*横幅の解釈をpadding, borderまでとする*/
}
.gmap_text {
  max-width: 300px;
  width: 100%;
  border: 2px solid #ddd;
  box-sizing: border-box;
}

.btn01 {
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

.btn01:hover {
  background-color: #fff;
  color: #303030;
}
.center {
  margin: 0 auto;
  text-align: center;
}
.button {
  margin: 30px 0px;
}
.image {
  height: 200px;
}
/* .comment_text {
  max-width: 500px;
  height: 100px;
  width: 100%;
  border: 2px solid #ddd;
  box-sizing: border-box;
}*/
</style>

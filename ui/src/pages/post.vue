<template>
  <section id="wrapper">
    <Header ref="header_area" />
    <div>
      <vue-google-autocomplete
        id="map"
        classname="form-control"
        placeholder="住所を入力してください"
        v-on:placechanged="getAddressData"
        v-model="address_search"
      >
      </vue-google-autocomplete>

      <button type="button" @click="getlatlng">検索</button>
      <div id="gmap" style="height: 500px; width: 800px"></div>
    </div>

    <Footer ref="footer_area" />
  </section>
</template>

<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";
import VueGoogleAutocomplete from "vue-google-autocomplete";

export default {
  name: "Index",
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
      // map: {},
      // marker: null,
      //geocoder: {},
      // address: "",
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
    init: async function () {},
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
            if (status == window.google.maps.GeocoderStatus.OK) {
              if (results[0].geometry) {
                vm.map.setCenter(results[0].geometry.location);
                results[0].geometry.location.lat();
                results[0].geometry.location.lng();

                // 緯度経度を取得
                var latlng = results[0].geometry.location;
                // 住所を取得(日本の場合だけ「日本, 」を削除)
                var address = results[0].formatted_address.replace(
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
                  content: address + "<br />" + latlng.toString(),
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
</style>

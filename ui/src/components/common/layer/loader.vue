<template>
  <div v-if="nowLoading" class="loaderBG">
    <div class="loader">Loading...</div>
    <progress
      v-if="withbar != null"
      class="loader_under"
      ref="bar"
      value="0"
      max="100"
    />
  </div>
</template>

<script>
export default {
  //nowLoading trueの場合表示する
  //withbar 進捗バーを表示する
  // post時にonUploadをフックすることでアップロード時の進捗を表示する
  props: ["nowLoading", "withbar"],
  methods: {
    onUpload(e) {
      this.$refs.bar.value = Math.floor((e.loaded * 100) / e.total);
    },
  },
};
</script>
<style scoped>
.loader,
.loader:after {
  border-radius: 50%;
  width: 10em;
  height: 10em;
}
.loader {
  width: 50px;
  height: 50px;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  position: absolute;
  margin: auto;
  font-size: 10px;
  text-indent: -9999em;
  border-top: 1.1em solid rgba(255, 255, 255, 0.2);
  border-right: 1.1em solid rgba(255, 255, 255, 0.2);
  border-bottom: 1.1em solid rgba(255, 255, 255, 0.2);
  border-left: 1.1em solid #ffffff;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation: load8 1.1s infinite linear;
  animation: load8 1.1s infinite linear;
}
@-webkit-keyframes load8 {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
@keyframes load8 {
  0% {
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
  }
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
.loaderBG {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 10000;
  background-color: rgba(0, 0, 0, 0.5);
  width: 100%;
  height: auto;
  min-height: 100vh;
  overflow: hidden;
  box-sizing: border-box;
  /* display:block; */
}
.loader_under {
  position: fixed;
  width: 80px;
  height: 10px;
  top: 100px;
  bottom: 0;
  left: 0px;
  right: 0;
  margin: auto;
  color: #fff;
}
</style>

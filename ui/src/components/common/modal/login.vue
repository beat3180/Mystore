<template>
  <div>
    <h3 class="bold"></h3>
    <ul>
      <li
        v-for="(error_message, key) in error_messages"
        :key="key"
        v-html="error_message"
      ></li>
    </ul>
    <div class="modal_form">
      メールアドレス<br />
      <input type="email" class="li_email" v-model="email" />
      パスワード<br />
      <input type="password" class="li_password" v-model="password" />
      <div class="center b10">
        <button class="btn04 finger" v-on:click="click_submit">ログイン</button>
      </div>
      <div class="center b10">
        <span class="mini_link">
          <span class="finger" v-on:click="$emit('modal_change', 'password')"
            >パスワードを忘れた場合</span
          >
        </span>
      </div>
      <hr class="dot_g" />
      <div class="center b10">
        <button
          class="btn05 finger"
          v-on:click="$emit('modal_change', 'regist')"
        >
          {{ $t("regist") }}
        </button>
      </div>
    </div>
    <div class="close_btn">
      <img
        src="/img/components/common/modal/close.png"
        class="finger"
        v-on:click="$emit('modal_off')"
      />
    </div>
    <Loader ref="loader" :nowLoading="nowLoading" />
  </div>
</template>
<script>
import Loader from "~/components/common/layer/loader.vue";

export default {
  components: {
    Loader,
  },
  data() {
    return {
      //      modal: false,
      nowLoading: false,

      email: "",
      password: "",

      //エラーメッセージ配列
      error_messages: [],
    };
  },
  methods: {
    init: function () {},
  },
};
</script>
<style scoped>
</style>

<template>
  <div v-if="is_modal" class="modal">
    <div class="modal_background" v-on:click="is_modal = false"></div>
    <div class="modal_inner">
      <div class="modal_msg center">
        <div class="modal_form_wrap">
          <!--  <div v-if="is_modal" id="regist_layer">-->
          <!-- ログインモーダル -->
          <LoginModal
            ref="login_modal"
            v-if="modal_type == 'login'"
            v-on:modal_change="modal_change"
            v-on:modal_off="modal_off"
          />
          <!-- 新規登録モーダル -->
          <RegistModal
            ref="regist_modal"
            v-if="modal_type == 'regist'"
            v-on:modal_change="modal_change"
            v-on:modal_off="modal_off"
          />
          <!-- パスワード忘れたモーダル -->
          <PasswordResetModal
            ref="password_reset_modal"
            v-if="modal_type == 'password'"
            v-on:modal_change="modal_change"
            v-on:modal_off="modal_off"
          />
          <!--  </div>-->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import LoginModal from "@/components/common/modal/login.vue";
import RegistModal from "@/components/common/modal/register.vue";
import PasswordResetModal from "@/components/common/modal/password_reset.vue";
export default {
  props: ["modal_switch", "modal_kind"],
  components: {
    LoginModal,
    RegistModal,
    PasswordResetModal,
  },
  data() {
    return {
      is_modal: false,
      modal_type: "login",
    };
  },
  watch: {
    modal_switch: function () {
      if (this.modal_switch <= 0) {
        this.is_modal = false;
        return;
      }

      this.is_modal = true;
    },
    modal_kind: function () {
      console.log(12345);
      this.modal_type = this.modal_kind;
    },
  },
  mounted() {
    //初期動作
    this.init();
  },
  methods: {
    init: function () {},
    modal_change: function (modal_type) {
      this.modal_type = modal_type;
    },

    modal_off: function () {
      this.is_modal = false;
    },
  },
};
</script>
<style scoped>
#regist_layer {
  line-height: normal;
}

/* モーダル関連 */
.modal {
  position: fixed;
  top: 0;
  left: 0;
  display: grid;
  place-items: center;
  width: 100%;
  min-height: 100%;
  background-color: rgba(0, 0, 0, 0.8);
  z-index: 99;
}
.modal_background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  min-height: 100%;
  z-index: 0;
}

.modal_inner {
  width: 95%;
  max-width: 400px;
  padding: 10px;
  background-color: #fff;
  border-radius: 3px;
}

.modal_title {
  padding: 3px 5px;
  background-color: #607d8b;
  color: #fff;
  font-size: 1.1em;
  font-weight: bold;
  border-radius: 3px;
}

.download_wrap,
.modal_form_wrap {
  position: relative;
}

.modal_msg {
  text-align: center;
}
</style>

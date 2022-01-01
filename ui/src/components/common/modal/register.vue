<template>
  <div>
    <span class="bold">新規登録</span>
    <ul id="error_message" v-if="!is_success && show_error">
      <li v-for="(error, key) in errors" :key="key">{{ error }}</li>
    </ul>
    <div v-if="is_success">
      <p v-html="message"></p>
    </div>
    <form @submit.prevent="submit">
      <div class="modal_form" v-if="!is_success">
        ニックネーム
        <input
          type="text"
          name="su_nickname"
          class="su_nickname"
          v-model="nick_name"
        />
        メールアドレス
        <input type="email" name="su_email" class="su_email" v-model="email" />
        パスワード<br />
        <input
          type="password"
          name="su_password"
          class="su_password"
          v-model="password"
        />
        パスワード確認<br />
        <input
          type="password"
          name="su_password_conf"
          class="su_password_conf"
          v-model="confirm_password"
        />

        <div class="center">
          <button
            class="btn01"
            type="submit"
            v-bind:class="{ signup_btn: !condition_agree }"
          >
            新規登録
          </button>
        </div>
        <hr class="dot_g" />
        <div class="center">
          <button class="btn02" v-on:click="$emit('modal_change', 'login')">
            ログイン
          </button>
        </div>
      </div>
    </form>
    <div class="close_btn">
      <img
        src="@/assets/img/close.png"
        class="finger"
        v-on:click="$emit('modal_off')"
      />
    </div>
    <Loader ref="loader" :nowLoading="nowLoading" />
  </div>
</template>
<script>
import { ref } from "vue";
import axios from "axios";
import Loader from "@/components/common/layer/loader.vue";

export default {
  components: {
    Loader,
  },
  data() {
    return {
      condition_agree: false,
      errors: [],
      message: "",
      show_error: false,

      //登録メール送信成功フラグ
      is_success: false,
    };
  },
  methods: {
    check_agree: function () {
      this.condition_agree = !this.condition_agree;
    },
    show_message: function (message) {
      this.errors = message;
      this.show_error = true;
    },
    // submit_registration: function () {
    //   let error_array = [];

    //   if (this.user_nickname == "") {
    //     error_array.push("ニックネームを入力してください");
    //   }

    //   if (this.user_email == "") {
    //     error_array.push("メールアドレスを入力してください");
    //   } else if (this.user_email.length > 100) {
    //     error_array.push(i18n.tc("erroremailnum"));
    //   } else {
    //     let regexp =
    //       /^[A-Za-z0-9]{1}[A-Za-z0-9\+_.-]*@{1}[A-Za-z0-9_.-]{1,}\.[A-Za-z0-9]{1,}$/;
    //     if (!regexp.test(this.user_email)) {
    //       error_array.push(i18n.tc("erroremailformat"));
    //     }
    //   }

    //   if (this.user_password == "") {
    //     error_array.push(i18n.tc("errorpassword"));
    //   } else if (this.user_password.length > 20) {
    //     error_array.push(i18n.tc("errorpasswordnum"));
    //   } else if (this.user_password != this.user_confirm_password) {
    //     error_array.push(i18n.tc("errorpasswordnotsame"));
    //   }

    //   //      if (this.user_password == "") {
    //   //        error_array.push("・" + i18n.tc("errorpassword"));
    //   //      } else if (this.user_password.length > 20) {
    //   //        error_array.push("・" + i18n.tc("errorpasswordnum"));
    //   //      } else if (this.user_password != this.user_confirm_password) {
    //   //        error_array.push("・" + i18n.tc("errorpasswordnotsame"));
    //   //      }

    //   if (!this.condition_agree) {
    //     error_array.push(i18n.tc("errornocheckterm"));
    //   }
    //   if (error_array.length) {
    //     this.show_message(error_array);
    //     return;
    //   }

    //   this.$axios
    //     .$post("/liver/registration", {
    //       email: this.user_email,
    //       nick_name: this.user_nickname,
    //       production_id: 1,
    //       introducer_id: this.$route.query.i ? this.$route.query.i : 0,
    //       campaign_id: this.campaign_id ? this.campaign_id : 0,
    //       password: this.user_password,
    //       password_confirm: this.user_confirm_password,
    //     })
    //     .then((res) => {
    //       var message =
    //         i18n.tc("message01") +
    //         "<br />" +
    //         i18n.tc("message02") +
    //         "<br />" +
    //         i18n.tc("message03");

    //       this.message = message;

    //       this.is_success = true;
    //     })
    //     .catch((e) => {
    //       console.log("axiosエラー" + e);
    //       if (e.response) {
    //         switch (e.response.status) {
    //           case 409:
    //             error_array.push(e.response.data.error.desc);
    //             this.show_message(error_array);
    //             this.error_proc(e.response.data.error.desc);
    //             break;

    //           default:
    //             this.error_proc(i18n.tc("passwerrorsystemerrorord") + e);
    //             break;
    //         }
    //       } else {
    //         this.error_proc(e);
    //       }
    //     })
    //     .finally(() => {
    //       this.nowLoading = false;
    //     });
    // },
    // error_proc: function (e) {
    //   console.log(e);
    // },
  },

  setup(props, context) {
    // v-modelからフォームに入力された値を格納
    const nick_name = ref("");
    const email = ref("");
    const password = ref("");
    const confirm_password = ref("");
    const nowLoading = ref(false);

    const submit = async () => {
      nowLoading.value = true;
      // Register apiへPOSTリクエスト
      await axios.post("register", {
        nick_name: nick_name.value,
        email: email.value,
        password: password.value,
        password_confirm: confirm_password.value,
      });
      nowLoading.value = false;
      // モーダルを消す
      await context.emit("modal_off");
    };

    return {
      nick_name,
      email,
      password,
      confirm_password,
      nowLoading,
      submit,
    };
  },
};
</script>
<style scoped>
ul#error_message {
  margin-top: 1em;
  text-align: left;
  color: red;
}
.signup_btn {
  cursor: default;
}
.signup_btn:hover {
  background-color: #ff8f00;
  color: white;
}
.bold {
  font-weight: bold;
}

.modal_form {
  width: 300px;
  margin: 15px auto 0px auto;
  text-align: left;
}

.modal_form input {
  display: block;
  width: 100%;
  margin-bottom: 20px;
}

.close_btn {
  position: absolute;
  top: 0px;
  right: 0px;
  cursor: pointer;
}

.close_btn img {
  width: 20px;
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
  text-align: center;
  margin-bottom: 10px;
}

.btn02 {
  padding: 3px 10px;
  min-width: 120px;
  border: 3px solid #0000ff;
  border-radius: 100px;
  background-color: #fff;
  color: #303030;
  font-weight: bold;
  transition-duration: 0.3s;
  cursor: pointer;
}

.btn02:hover {
  background-color: #0000ff;
  color: #fff;
}
</style>

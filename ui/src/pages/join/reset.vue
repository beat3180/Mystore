<template>
  <div>
    <Header ref="header_area" />
    <div class="bold center">パスワードリセット</div>
    <ul id="error_message" v-if="submit_flag === true">
      <li>{{ errors.password }}</li>
      <li>{{ errors.confirm_password }}</li>
      <li>{{ error_message }}</li>
    </ul>
    <form @submit.prevent="submit">
      <div class="modal_form">
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
          <button class="btn01" type="submit">送信</button>
        </div>
      </div>
    </form>
    <Footer ref="footer_area" />
    <Loader ref="loader" :nowLoading="nowLoading" />
  </div>
</template>
<script>
import Header from "@/components/common/area/header.vue";
import Footer from "@/components/common/area/footer.vue";
import { ref, computed, watchEffect } from "vue";
import { useStore } from "vuex";
import { useField, useForm } from "vee-validate";
import * as yup from "yup";
import axios from "axios";
import Loader from "@/components/common/layer/loader.vue";
import { useRoute, useRouter } from "vue-router";

export default {
  name: "Reset",
  components: {
    Loader,
    Header,
    Footer,
  },
  data() {
    return {};
  },

  setup() {
    //送信ボタンを押すとメッセージ表示
    const submit_flag = ref(false);
    //ローディングの値
    const nowLoading = ref(false);
    //エラーメッセージ
    const error_message = ref("");
    const store = useStore();
    const route = useRoute();
    const router = useRouter();
    const auth = computed(() => store.state.auth);
    //ログインチェック
    watchEffect(() => {
      if (auth.value == true) {
        router.push("/");
      } else {
        return;
      }
    });

    //バリデーションしつつフォームから値を格納
    const schema = yup.object({
      password: yup.string().required("パスワードは必須項目です"),
      confirm_password: yup.string().required("パスワード確認は必須項目です"),
    });

    const { errors } = useForm({
      validationSchema: schema,
    });

    const { value: password } = useField("password");
    const { value: confirm_password } = useField("confirm_password");

    const submit = async () => {
      submit_flag.value = true;
      if (
        password.value &&
        confirm_password.value &&
        !errors.value.password &&
        !errors.value.confirm_password
      ) {
        nowLoading.value = true;
        // Register apiへPOSTリクエスト
        await axios
          .post("reset", {
            // urlからtokenを取得
            token: route.params.token,
            password: password.value,
            password_confirm: confirm_password.value,
          })
          .then((res) => {
            submit_flag.value = false;
            console.log(res);

            // ログイン画面へ遷移する
            router.push("/");
          })
          .catch((e) => {
            console.log("axiosエラー e:" + e);
            error_message.value = e.response.data.message;
          })
          .finally(() => {
            nowLoading.value = false;
          });
      }
    };

    return {
      password,
      confirm_password,
      error_message,
      errors,
      submit_flag,
      nowLoading,
      submit,
    };
  },
};
</script>
<style scoped>
ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.li {
  padding: 0;
  margin: 0;
}
ul#error_message {
  margin-top: 1em;
  text-align: center;
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
  margin: 10px 0;
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

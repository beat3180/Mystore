<template>
  <div>
    <span class="bold">新規登録</span>
    <ul id="error_message" v-if="submit_flag === true">
      <li>{{ errors.nick_name }}</li>
      <li>{{ errors.email }}</li>
      <li>{{ errors.password }}</li>
      <li>{{ errors.confirm_password }}</li>
      <li>{{ error_message }}</li>
    </ul>
    <div v-if="is_success">
      <p>会員登録が完了しました</p>
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
          <button class="btn01" type="submit">新規登録</button>
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
import { useField, useForm } from "vee-validate";
import * as yup from "yup";
import axios from "axios";
import Loader from "@/components/common/layer/loader.vue";

export default {
  components: {
    Loader,
  },
  data() {
    return {
      // condition_agree: false,
      // //errors: [],
      // message: "",
      // show_error: false,
      // //登録メール送信成功フラグ
      // is_success: false,
    };
  },

  setup() {
    //送信ボタンを押すとメッセージ表示
    const submit_flag = ref(false);
    //ローディングの値
    const nowLoading = ref(false);
    //会員登録完了
    const is_success = ref(false);
    //エラーメッセージ
    const error_message = ref("");

    //バリデーションしつつフォームから値を格納
    const schema = yup.object({
      nick_name: yup.string().required("ニックネームは必須の項目です"),
      email: yup
        .string()
        .required("メールアドレスは必須の項目です")
        .email("メールアドレスが不正です"),
      password: yup.string().required("パスワードは必須項目です"),
      confirm_password: yup.string().required("パスワード確認は必須項目です"),
    });

    const { errors } = useForm({
      validationSchema: schema,
    });

    const { value: nick_name } = useField("nick_name");
    const { value: email } = useField("email");
    const { value: password } = useField("password");
    const { value: confirm_password } = useField("confirm_password");

    const submit = async () => {
      submit_flag.value = true;
      if (
        nick_name.value &&
        email.value &&
        password.value &&
        confirm_password.value &&
        !errors.value.nick_name &&
        !errors.value.email &&
        !errors.value.password &&
        !errors.value.confirm_password
      ) {
        nowLoading.value = true;
        // Register apiへPOSTリクエスト
        await axios
          .post("register", {
            nick_name: nick_name.value,
            email: email.value,
            password: password.value,
            password_confirm: confirm_password.value,
          })
          .then((res) => {
            submit_flag.value = false;
            console.log(res);

            is_success.value = true;
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
      nick_name,
      email,
      password,
      confirm_password,
      error_message,
      errors,
      submit_flag,
      nowLoading,
      is_success,
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

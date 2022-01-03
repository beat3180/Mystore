<template>
  <div>
    <span class="bold">ログイン</span>
    <ul id="error_message" v-if="submit_flag === true">
      <li>{{ errors.email }}</li>
      <li>{{ errors.password }}</li>
      <li>{{ error_message }}</li>
    </ul>

    <div class="modal_form">
      メールアドレス<br />
      <input type="email" class="li_email" v-model="email" />
      パスワード<br />
      <input type="password" class="li_password" v-model="password" />
      <div class="center b10">
        <button class="btn01" type="submit" @click="submit">ログイン</button>
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
        <button class="btn02" v-on:click="$emit('modal_change', 'regist')">
          新規登録
        </button>
      </div>
    </div>
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
import Loader from "@/components/common/layer/loader.vue";
import { ref } from "vue";
import { useField, useForm } from "vee-validate";
import * as yup from "yup";
import axios from "axios";

export default {
  components: {
    Loader,
  },
  data() {
    return {};
  },
  methods: {
    init: function () {},
  },

  setup(props, context) {
    //ローディングの値
    const nowLoading = ref(false);
    //送信ボタンを押すとメッセージ表示
    const submit_flag = ref(false);
    //エラーメッセージ
    const error_message = ref("");

    //バリデーションしつつフォームから値を格納
    const schema = yup.object({
      email: yup
        .string()
        .required("メールアドレスは必須の項目です")
        .email("メールアドレスが不正です"),
      password: yup.string().required("パスワードは必須項目です"),
    });

    const { errors } = useForm({
      validationSchema: schema,
    });

    const { value: email } = useField("email");
    const { value: password } = useField("password");

    const submit = async () => {
      submit_flag.value = true;
      if (
        email.value &&
        password.value &&
        !errors.value.email &&
        !errors.value.password
      ) {
        nowLoading.value = true;
        // login apiへPOSTリクエスト
        await axios
          .post("login", {
            email: email.value,
            password: password.value,
          })
          .then((res) => {
            submit_flag.value = false;
            console.log(res);

            context.emit("modal_off");
            //リロード
            document.location.reload();
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
      email,
      password,
      errors,
      submit_flag,
      error_message,
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

<template>
  <div>
    <span class="bold">パスワードリセット</span>
    <div v-if="notify.cls == 'success'">
      {{ notify.message }}
    </div>
    <ul id="error_message" v-if="submit_flag === true">
      <li>{{ errors.email }}</li>
      <li>{{ errors.password }}</li>
      <li>{{ error_message }}</li>
    </ul>
    <div v-if="!is_success" class="modal_form">
      メールアドレス<br />
      <input type="email" v-model="email" class="fp_email" />
      メールアドレス確認<br />
      <input type="email" v-model="email_confirm" class="fp_email_conf" />
      <div class="center b10">
        <button class="btn01 finger" @click="submit">送信</button>
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
import { reactive, ref } from "vue";
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
    show_message: function (message) {
      this.message = "<h3>" + message + "</h3>";
    },
    error_proc: function (e) {
      this.$parent.error_proc(e);
    },
  },

  setup() {
    //ローディングの値
    const nowLoading = ref(false);
    //送信ボタンを押すとメッセージ表示
    const submit_flag = ref(false);
    //エラーメッセージ
    const error_message = ref("");
    // メール送信結果を格納
    const notify = reactive({
      cls: "",
      message: "",
    });

    //バリデーションしつつフォームから値を格納
    const schema = yup.object({
      email: yup
        .string()
        .required("メールアドレスは必須の項目です")
        .email("メールアドレスが不正です"),
      email_confirm: yup
        .string()
        .required("メールアドレス確認は必須項目です")
        .email("メールアドレス確認が不正です"),
    });

    const { errors } = useForm({
      validationSchema: schema,
    });

    const { value: email } = useField("email");
    const { value: email_confirm } = useField("email_confirm");

    const submit = async () => {
      submit_flag.value = true;
      if (
        email.value &&
        email_confirm.value &&
        !errors.value.email &&
        !errors.value.email_confirm
      ) {
        nowLoading.value = true;
        // forgot apiへPOSTリクエスト
        await axios
          .post("forgot", {
            email: email.value,
            email_confirm: email_confirm.value,
          })
          .then((res) => {
            submit_flag.value = false;
            console.log(res);

            notify.cls = "success";
            notify.message = "メールを送信しました";
          })
          .catch((e) => {
            console.log("axiosエラー e:" + e);
            error_message.value = e.response.data.message;
            notify.cls = "danger";
            notify.message = "メールの送信ができませんでした";
          })
          .finally(() => {
            nowLoading.value = false;
          });
      }
    };

    return {
      email,
      email_confirm,
      errors,
      submit_flag,
      error_message,
      notify,
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

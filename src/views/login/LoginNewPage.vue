<script setup>
import { userRegisterService, userLoginService } from '@/api/user'
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/modules/user'
import { User, Lock } from '@element-plus/icons-vue'


const userStore = useUserStore()
const router = useRouter()
const isPage = ref(0)
const registerForm = ref(null)

const formModel = ref({
  account: '',
  password: '',
  repassword: '',
})

const rules = {
  account: [
    { required: true, message: '账号不得为空', trigger: 'blur' },
    {
      min: 6,
      max: 10,
      message: '请输入正确账号',
      trigger: 'blur',
    },
  ],
  password: [
    { required: true, message: '密码不得为空', trigger: 'blur' },
    { pattern: /^\S{6,15}$/, message: '密码长度必须是6-15位非空字符', trigger: 'blur' },
  ],
  repassword: [
    {
      validator: (rules, value, callback) => {
        if (value !== formModel.value.password) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
}

const register = async () => {
  if (!registerForm.value) {
    ElMessage.error('表单未正确绑定')
    return
  }

  
  const valid = await registerForm.value.validate()
  if (!valid) {
    ElMessage.error('请正确填写注册信息')
    return
  }
  const response = await userRegisterService(
    formModel.value.account,
    formModel.value.password,
    formModel.value.repassword
  )

  if (response.data.message) {
    ElMessage.success('注册成功')
    isPage.value = 0
  } 
}

const goToHome = () => {
  router.push('/')
}

const login = async () => {
  // 调用登录接口，获取返回值
  const res = await userLoginService(formModel.value.account, formModel.value.password)
  if (res.data.token) {
    ElMessage.success(res.data.message || '登录成功')
    userStore.setToken(res.data.token)
    userStore.setUserId(res.data.userId)
    userStore.setName(res.data.userName)
    goToHome()
  } 
}

watch(isPage, (newValue) => {
  if (newValue !== 1) {
    formModel.value = {
      account: '',
      password: '',
      repassword: '',
    }
  }
})
</script>
<template>
  <el-row class="login-page">
    <el-col :span="8" class="bg"> </el-col>
    <el-col :span="8" class="form">
      <el-form v-if="isPage == 0" :model="formModel" :rules="rules" ref="registerForm">
        <el-form-item><h1>登录</h1></el-form-item>
        <el-form-item prop="account">
          <el-input
            v-model="formModel.account"
            :prefix-icon="User"
            placeholder="请输入账号"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            placeholder="请输入密码"
            show-password
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div class="flex">
            <el-checkbox value="Online activities" name="type"> 记住我 </el-checkbox>
            <el-link type="primary" :underline="false" @click="isPage = 2">忘记密码？</el-link>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="button" @click="login">登录</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :underline="false" @click="isPage = 1">
            注册<el-icon><Right /></el-icon>
          </el-link>
        </el-form-item>
      </el-form>
      <el-form v-else-if="isPage == 1" :model="formModel" :rules="rules" ref="registerForm">
        <el-form-item><h1>注册</h1></el-form-item>
        <el-form-item prop="account"
          ><el-input
            v-model="formModel.account"
            :prefix-icon="User"
            placeholder="请输入账号"
          ></el-input
        ></el-form-item>
        <el-form-item prop="password"
          ><el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            placeholder="请输入密码"
            show-password
          ></el-input
        ></el-form-item>
        <el-form-item prop="repassword"
          ><el-input
            v-model="formModel.repassword"
            :prefix-icon="Lock"
            placeholder="请再次输入密码"
            show-password
          ></el-input
        ></el-form-item>
        <el-form-item prop="captcha"> </el-form-item>
        <el-form-item>
          <el-button type="primary" class="button" @click="register">注册</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :offset="2" :underline="false" @click="isPage = 0"
            ><el-icon><Back /></el-icon>返回</el-link
          >
        </el-form-item>
      </el-form>
      <el-form v-else-if="isPage == 2" :model="formModel" :rules="rules" ref="form">
        <el-form-item><h1>找回密码</h1></el-form-item>
        <el-form-item>
          <el-link type="info" :offset="2" :underline="false" @click="isPage = 0"
            ><el-icon><Back /></el-icon>返回</el-link
          >
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<style scoped>
.login-page {
  height: 100vh;
  min-width: 350px;
  min-height: 350px;
  display: flex;
  justify-content: center; 
  align-items: center;
  background-color: #9cc6e3;
}
.bg {
  min-width: 350px;
  min-height: 350px;
  height: 70vh;
  background-image: url('../../assets/【哲风壁纸】笑嘻嘻-美女.png');
  background-size: cover;
  background-position: center;
  border-radius: 20px 0 0 20px;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
}
.form {
  min-width: 350px;
  min-height: 350px;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 70vh;
  background-color: #fff;
  border-radius: 0 20px 20px 0;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);

  .button {
    width: 100%;
  }

  .flex {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
}

@media screen and (max-width: 768px) {
  .login-page {
    background-image: url(../../assets/0927be6b971a98f94aeefca333c26343.png);
    background-size: 100% 100%;
  }

  .bg {
    display: none;
  }

  .form {
    border-radius: 20px; /* 调整边框圆角 */
    background-color: rgba(250, 250, 250, 0.6);
  }
}

@media screen and (max-height: 600px){
  .form {
    height: 90vh;
  }

  .bg {
    height: 90vh;
  }
}
</style>

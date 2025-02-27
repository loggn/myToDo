<script setup>
import { User, Lock } from '@element-plus/icons-vue'
import { ref, watch } from 'vue'
import { userRegisterService } from '@/api/user.js'
const isRegister = ref(0)
const form = ref()
const formModel = ref({
  phoneNumber: '',
  password: '',
  repassword: '',
  captcha: '',
})

const rules = {
  phoneNumber: [
    { required: true, message: '手机号不得为空', trigger: 'blur' },
    {
      pattern: /^1[3456789]\d{9}$/,
      len: 11,
      message: '请输入正确手机号',
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
      trigger: 'change',
    },
  ],
  captcha: [{ required: true, message: '验证码不得为空', trigger: 'blur' }],
}

const register = async () => {
  await form.value.validate()
  await userRegisterService(
    formModel.value.phoneNumber,
    formModel.value.password,
    formModel.value.repassword,
    formModel.value.captcha,
  )
  ElMessage.success('注册成功')
  isRegister.value = 0
}

watch(isRegister, () => {
  formModel.value = {
    phoneNumber: '',
    password: '',
    repassword: '',
    captcha: '',
  }
})
</script>

<template>
  <el-row class="login-page">
    <el-col :span="8" :offset="4" class="bg"> </el-col>
    <el-col :span="8" class="form">
      <el-form v-if="isRegister == 0" :model="formModel" :rules="rules" ref="form">
        <el-form-item><h1>登录</h1></el-form-item>
        <el-form-item prop="phoneNumber">
          <el-input
            v-model="formModel.phoneNumber"
            :prefix-icon="User"
            placeholder="请输入手机号"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            placeholder="请输入密码"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div class="flex">
            <el-checkbox value="Online activities" name="type"> 记住我 </el-checkbox>
            <el-link type="primary" :underline="false" @click="isRegister = 2">忘记密码？</el-link>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="button">登录</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :underline="false" @click="isRegister = 1">
            注册<el-icon><Right /></el-icon>
          </el-link>
        </el-form-item>
      </el-form>
      <el-form v-else-if="isRegister == 1" :model="formModel" :rules="rules" ref="form">
        <el-form-item><h1>注册</h1></el-form-item>
        <el-form-item prop="phoneNumber"
          ><el-input
            v-model="formModel.phoneNumber"
            :prefix-icon="User"
            placeholder="请输入手机号"
          ></el-input
        ></el-form-item>
        <el-form-item prop="password"
          ><el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            placeholder="请输入密码"
          ></el-input
        ></el-form-item>
        <el-form-item prop="repassword"
          ><el-input
            v-model="formModel.repassword"
            :prefix-icon="Lock"
            placeholder="请再次输入密码"
          ></el-input
        ></el-form-item>
        <el-form-item prop="captcha">
          <div class="flex">
            <el-input v-model="formModel.captcha" placeholder="请输入验证码"></el-input>
            <el-button>发送验证码</el-button>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="button" @click="register">注册</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :offset="2" :underline="false" @click="isRegister = 0"
            ><el-icon><Back /></el-icon>返回</el-link
          >
        </el-form-item>
      </el-form>
      <el-form v-else-if="isRegister == 2" :model="formModel" :rules="rules" ref="form">
        <el-form-item><h1>重置密码</h1></el-form-item>
        <el-form-item prop="phoneNumber"
          ><el-input
            v-model="formModel.phoneNumber"
            :prefix-icon="User"
            placeholder="请输入验证手机号"
          ></el-input
        ></el-form-item>
        <el-form-item prop="captcha">
          <div class="flex">
            <el-input v-model="formModel.captcha" placeholder="请输入验证码"></el-input>
            <el-button>发送验证码</el-button>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="button" @click="isRegister = 3">重置</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :offset="2" :underline="false" @click="isRegister = 0"
            ><el-icon><Back /></el-icon>返回</el-link
          >
        </el-form-item>
      </el-form>
      <el-form v-else :model="formModel" :rules="rules" ref="form">
        <el-form-item><h1>重置密码</h1></el-form-item>
        <el-form-item prop="password"
          ><el-input
            v-model="formModel.password"
            :prefix-icon="Lock"
            placeholder="请输入密码"
          ></el-input
        ></el-form-item>
        <el-form-item prop="repassword"
          ><el-input
            v-model="formModel.repassword"
            :prefix-icon="Lock"
            placeholder="请再次输入密码"
          ></el-input
        ></el-form-item>
        <el-form-item>
          <el-button type="primary" class="button">确认</el-button>
        </el-form-item>
        <el-form-item>
          <el-link type="info" :offset="2" :underline="false" @click="isRegister = 0"
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
  align-items: center;
  background-color: #aae5f1;
}
.bg {
  height: 50vh;
  background-color: pink;
  border-radius: 20px 0 0 20px;
}
.form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50vh;
  background-color: #fff;
  border-radius: 0 20px 20px 0;

  .button {
    width: 100%;
  }

  .flex {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
}
</style>

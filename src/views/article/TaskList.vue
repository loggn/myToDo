<script  setup>
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/modules/user'

const router = useRouter()
const userStore = useUserStore()

const logout = () => {
  goTologin()
  userStore.removeName()
  userStore.removeUserId()
  userStore.removeToken()
}

const goTologin = () => {
  router.push('/login')
}

const toPage = (toUrl) => {
  router.push(toUrl)
}
</script>
<template>
  <div class="Task-common-layout">
    <el-container>
      <el-header class="Task-header">
        <div class="Task-header-left">
          <h1>
            <el-icon><Collection /></el-icon>
            任务
          </h1>
        </div>
        <div class="Task-header-rigth">
          <el-popover placement="bottom" :width="90" trigger="click">
            <template #reference>
          <div>
            <el-icon><TopRight /></el-icon>
          </div>
        </template>
        <template #default>
              <ul>
                <li @click="logout">退出登录</li>
                <li>功能待开发...</li>
              </ul>
        </template>
        </el-popover>
          <!-- <div>
            <el-icon><Opportunity /></el-icon>
          </div> -->
          <el-popover placement="bottom" :width="90" trigger="click">
          <template #reference>
           <el-icon><More /></el-icon>
          </template>
          <template #default>
              <ul>
                <li @click="toPage('/article/myday')">myday</li>
                <li @click="toPage('/article/ImportantToDo')">重要</li>
                <li @click="toPage('/article/TaskInPlan')">计划内</li>
                <li @click="toPage('/article/AssignedToMe')">已分配给我</li>
                <li @click="toPage('/article/TaskList')">已完成</li>
              </ul>
            </template>
          </el-popover>
        </div>
      </el-header>
      <el-main class="Task-body">任务列表</el-main>
      <el-footer class="Task-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.Task-common-layout {
  display: flex;
  background-image: url('../../assets/新背景1.jpg');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
}
.Task-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.Task-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.Task-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>

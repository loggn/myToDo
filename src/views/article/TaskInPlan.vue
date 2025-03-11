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
  <div class="TaskInPlan-common-layout">
    <el-container>
      <el-header class="TaskInPlan-header">
        <div class="TaskInPlan-header-left">
          <h1>
            <el-icon><List /></el-icon>
            计划内
          </h1>
        </div>
        <div class="TaskInPlan-header-rigth">
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
      <el-main class="TaskInPlan-body">任务列表</el-main>
      <el-footer class="TaskInPlan-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.TaskInPlan-common-layout {
  display: flex;
  background-image: url('../../assets/【哲风壁纸】二次元-动漫.png');
  background-size: cover;
  background-position: center;
  height: 100vh;
}
.TaskInPlan-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.TaskInPlan-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.TaskInPlan-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>

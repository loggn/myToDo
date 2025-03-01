<script>
import { useRouter } from 'vue-router'

const router = useRouter()

const goTologin = () => {
  router.push('/login')
}

const toPage = (toUrl) => {
  router.push(toUrl)
}
</script>
<template>
  <div class="Import-common-layout">
    <el-container>
      <el-header class="Import-header">
        <div class="Import-header-left">
          <h1>
            <el-icon><StarFilled /></el-icon>
            重要
          </h1>
        </div>
        <div class="Import-header-rigth">
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
      <el-main class="Import-body">任务列表</el-main>
      <el-footer class="Import-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.Import-common-layout {
  display: flex;
  background-image: url('../../assets/背景3.jpg');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
}
.Import-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.Import-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.Import-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>

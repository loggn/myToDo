<script  setup>
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/modules/user'
import { ref } from 'vue'

const router = useRouter()
const userStore = useUserStore()
const newStyle = ref({
  opacity: 0.5,
})

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
  <div class="AssignedToMe-common-layout">
    <el-container>
      <el-header class="AssignedToMe-header">
        <div class="AssignedToMe-header-left">
          <h1>
            <el-icon><User /></el-icon>
            分配给我的
          </h1>
        </div>
        <div class="AssignedToMe-header-rigth">
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
                <li @click="toPage('/article/TaskList')">历史任务</li>
              </ul>
            </template>
          </el-popover>
        </div>
      </el-header>
      <el-main class="AssignedToMe-body">任务列表</el-main>
      <el-footer class="AssignedToMe-foot">
        <el-input 
        :style="newStyle"
        placeholder="搜索任务"
        >
        </el-input>
      </el-footer>
    </el-container>
  </div>
</template>
<style scoped>
.AssignedToMe-common-layout {
  display: flex;
  background-image: url('../../assets/【哲风壁纸】写真-封疆疆.png');
  background-size: cover;
  background-position: center;
  height: 100vh;
}
.AssignedToMe-header {
  display: flex;
  justify-content: space-between;
  margin: 25px;
}
.AssignedToMe-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.AssignedToMe-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}

@media screen and (max-width: 768px) {
  .AssignedToMe-header {
    margin: 0px;
  }

  .AssignedToMe-header-rigth {
    width: 100px;
  }
}
</style>

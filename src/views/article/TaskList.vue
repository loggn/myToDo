<script  setup>
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/modules/user'
import { computed, ref, onMounted } from 'vue'
import { userGetTask } from '@/api/user'

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

const tasks = ref([])

const getTasks = async () => {
  const res = await userGetTask(userStore.user_id)
  tasks.value = res.data.todos.map((todo) => ({
    id: todo.id,
    text: todo.text,
    isFinish: todo.isFinish,
    isImportant: todo.isImportant,
  }))
}

onMounted(() => {
  getTasks()
})

const finishedTasks = computed(() => {
  return tasks.value.filter(task => task.isFinish)
})
</script>
<template>
  <div class="Task-common-layout">
    <el-container>
      <el-header class="Task-header">
        <div class="Task-header-left">
          <h1>
            <el-icon><Collection /></el-icon>
            历史任务
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
                <li @click="toPage('/article/TaskList')">历史任务</li>
              </ul>
            </template>
          </el-popover>
        </div>
      </el-header>
      <el-main class="Task-body">
        <div
          v-for="(task, index) in finishedTasks"
          :key="index"
          class="task-item"
        >
          <!-- 任务文本 -->
          <h3>
            {{ task.text }}
          </h3>
        </div>
      </el-main>
      <el-footer class="Task-foot"></el-footer>
    </el-container>
  </div>
</template>
<style>
.Task-common-layout {
  display: flex;
  background-image: url('../../assets/【哲风壁纸】完美世界-火灵儿.png');
  background-size: cover;
  background-position: center;
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

.task-item {
  margin: 8px;
  display: flex;
  align-items: center;
  background-color: #f9f9f98b;
  border-radius: 8px;
  padding: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

@media screen and (max-width: 768px) {
  .Task-header {
    margin: 0px;
  }

  .Task-header-rigth {
    width: 100px;
  }
}
</style>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { DocumentAdd, Search, CircleCheck, Star, StarFilled } from '@element-plus/icons-vue'
import {
  userAddTask,
  userGetTask,
  changeTaskIsFinish,
  changeTaskIsImportant,
  changeTaskText,
  deleteTask,
} from '@/api/user'
import { useUserStore } from '../../stores/modules/user'
import { useRouter } from 'vue-router'

const drawerShowTask = ref()
const drawer = ref(false)
const userStore = useUserStore()
const inputValue = ref('')
const icon = ref(DocumentAdd)
const transparency = ref({
  opacity: 0.5,
})
const tasks = ref([])
const isPop = ref(false)
const mouse = ref({ left: 0, top: 0 })
const delTask = ref()
const router = useRouter()

const addTask = async () => {
  await userAddTask(inputValue.value.trim(), userStore.user_id)
  inputValue.value = '' // 清空输入框
}

const getTasks = async () => {
  const res = await userGetTask(userStore.user_id)
  tasks.value = res.data.todos.map((todo) => ({
    id: todo.id,
    text: todo.text,
    isFinish: todo.isFinish,
    isImportant: todo.isImportant,
  }))
}

const inputFocus = () => {
  icon.value = Search
  transparency.value.opacity = 1
}
const inputBlur = () => {
  icon.value = DocumentAdd
  transparency.value.opacity = 0.5
}

const inputEnter = async () => {
  if (inputValue.value == '') {
    console.log('任务不得为空')
  } else {
    await addTask()
    await getTasks()
  }
  icon.value = DocumentAdd
  transparency.value.opacity = 0.5
}

const changeIsFinish = async (task_id, task_isFinish) => {
  const res = await changeTaskIsFinish(task_id, task_isFinish)
  console.log(res)
}
// 切换任务完成状态
const toggleFinishStatus = (task) => {
  task.isFinish = !task.isFinish
  changeIsFinish(task.id, task.isFinish)
}

// 切换任务重要性状态
const toggleImportanceStatus = async (task) => {
  task.isImportant = !task.isImportant
  const res = await changeTaskIsImportant(task.id, task.isImportant)
  console.log(res)
}

const toChangeTaskText = async (drawerShowTask) => {
  await changeTaskText(drawerShowTask.id, drawerShowTask.text)
  getTasks()
}

onMounted(() => {
  getTasks()
})

const openDrawer = (task) => {
  drawerShowTask.value = {
    id: task.id,
    text: task.text,
    isFinish: task.isFinish,
    isImportant: task.isImportant,
  }
  drawer.value = true
}

const pop = (e, task) => {
  mouse.value.left = e.pageX
  mouse.value.top = e.pageY
  isPop.value = true
  console.log(task)
  delTask.value = task
  document.addEventListener('click', () => {
    isPop.value = false
  })
}

// 删除任务
const toDeleteTask = async (taskId) => {
  console.log(taskId)
  await deleteTask(taskId)
  await getTasks() // 删除任务后刷新任务列表
  isPop.value = false // 关闭气泡
}

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
  <div class="myday-common-layout">
    <el-container>
      <el-header class="myday-header">
        <div class="myday-header-left">
          <h1>我的一天</h1>
          <h3>2月10日,星期一</h3>
        </div>
        <div class="myday-header-rigth">
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
      <el-scrollbar>
        <el-main class="myday-body">
          <div class="myday-taskList" v-if="tasks.length">
            <div
              v-for="(task, index) in tasks"
              :key="index"
              class="task-item"
              @contextmenu.prevent="pop($event, task)"
            >
              <!-- 显示任务是否完成 -->
              <div class="task-status">
                <el-icon v-if="task.isFinish" @click="toggleFinishStatus(task)">
                  <CircleCheck />
                </el-icon>
                <el-icon v-else @click="toggleFinishStatus(task)">
                  <CircleCheck :style="{ opacity: 0.2 }" />
                </el-icon>
              </div>
              <!-- 显示任务的文本 -->
              <h3 class="task-text" v-if="task.isFinish" @click="openDrawer(task)">
                <del>{{ task.text }}</del>
              </h3>
              <h3 class="task-text" v-else @click="openDrawer(task)">{{ task.text }}</h3>
              <!-- 显示任务是否重要 -->
              <div class="task-important">
                <el-icon v-if="task.isImportant" @click="toggleImportanceStatus(task)">
                  <StarFilled />
                </el-icon>
                <el-icon v-else @click="toggleImportanceStatus(task)">
                  <Star :style="{ opacity: 0.2 }" />
                </el-icon>
              </div>
            </div>
          </div>
          <div class="myday-noTaskList" v-else>
            <el-empty :image-size="200" description="没有任务"></el-empty>
          </div>
        </el-main>
      </el-scrollbar>
      <el-footer class="myday-foot">
        <el-input
          class="myday-foot-input"
          :style="transparency"
          v-model="inputValue"
          placeholder="添加任务"
          :prefix-icon="icon"
          @focus="inputFocus"
          @blur="inputBlur"
          @keydown.enter="inputEnter"
        />
      </el-footer>
    </el-container>
  </div>
  <el-drawer v-model="drawer" :with-header="false">
    <el-input
      class="drawer_input"
      @keydown.enter="toChangeTaskText(drawerShowTask)"
      v-model="drawerShowTask.text"
      style="width: 240px"
    />
  </el-drawer>
  <div v-if="isPop" class="pop" :style="{ left: mouse.left + 'px', top: mouse.top + 'px' }">
    <div @click="toDeleteTask(delTask.id)">删除任务</div>
  </div>
</template>
<style>
.pop {
  position: absolute;
  background: white;
  box-shadow: 1px 1px 4px #888;
  font-size: 12px;
  border-radius: 5px;
  z-index: 999;
}
.pop > div {
  margin: 15px 10px;
  cursor: pointer;
}

.myday-taskList {
  display: flex;
  flex-direction: column;
  gap: 5px;
  justify-content: center;
  align-items: center;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 20px;
  width: auto;
  margin: 0 auto;
}

.task-item {
  display: flex;
  align-items: center;
  width: 100%;
  background-color: #f9f9f9;
  border-radius: 8px;
  padding: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.task-status {
  margin-right: 15px;
}

.task-text {
  flex-grow: 1;
  font-size: 18px;
  font-weight: 600;
}

.task-important {
  margin-left: 15px;
}

.myday-noTaskList {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20%;
  height: 0px;
  flex-shrink: 1;
}

.myday-common-layout {
  display: flex;
  background-image: url('../../assets/【哲风壁纸】可爱女生-室内坐姿.png');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
  min-width: 400px;
}
.myday-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: auto;
}
.myday-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.myday-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}

.myday-taskLis {
  display: flex;
}

.myday-foot-input {
  display: flex;
}

.drawer_input .el-input__inner {
  color: #000000; /* 黑色 */
  font-weight: bold;
}

.scrollbar-demo-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 50px;
  margin: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
}
</style>

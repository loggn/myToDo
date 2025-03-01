<script setup>
import { Search } from '@element-plus/icons-vue'
import { useUserStore } from '@/stores/modules/user'
import { useRouter, useRoute } from 'vue-router'
import { ref, watch, onMounted, reactive } from 'vue'
import { changeUserName } from '@/api/user'

const userStore = useUserStore()
const router = useRouter()
const route = useRoute()
const activeMenu = ref(route.path) // 设置当前激活的菜单项
const menuVisible = ref(true) // 控制菜单是否显示
const dialogFormVisible = ref(false)
const formLabelWidth = '140px'

const form = reactive({
  name: '',
})

// 事件处理
const homeSearchBoxFocus = () => {
  console.log('输入框聚焦了')
}

const homeSearchBoxChangs = (value) => {
  console.log('输入框值改变了:', value)
}

const goTologin = () => {
  router.push('/login')
}

const logout = () => {
  goTologin()
  userStore.removeName()
  userStore.removeUserId()
  userStore.removeToken()
}

// 监听路由变化，动态更新激活的菜单项
watch(
  () => route.path,
  (newPath) => {
    activeMenu.value = newPath
  },
)

// 在页面加载时，监听窗口尺寸变化
onMounted(() => {
  window.addEventListener('resize', handleResize)
  handleResize() // 初始化一次
})

// 监听窗口大小变化，动态调整菜单显示状态
const handleResize = () => {
  if (window.innerWidth < 768) {
    menuVisible.value = false
  } else {
    menuVisible.value = true
  }
}

const toChangeUserName = async (name) => {
  await changeUserName(userStore.user_id, name)
  dialogFormVisible.value = false
  userStore.name = name
}
</script>

<template>
  <el-row class="homepage">
    <el-col :span="menuVisible ? 4 : 0" class="menu">
      <el-menu
        class="homepage-menu"
        active-text-color="#ffd04b"
        :default-active="activeMenu"
        router
      >
        <div class="menu-header">
          <el-popover placement="bottom" :width="250" trigger="click">
            <template #reference>
              <div class="header-user">
                <div>
                  <el-avatar class="title" Avatar="large" :size="80"></el-avatar>
                </div>
                <div class="header-user-text">
                  <h3>{{ userStore.name }}</h3>
                  <p>每天进步一点点</p>
                </div>
              </div>
            </template>
            <template #default>
              <ul>
                <li @click="logout">注销登录</li>
                <li @click="dialogFormVisible = true">更改昵称</li>
                <el-dialog v-model="dialogFormVisible" title="更改昵称" width="500">
                  <el-form :model="form">
                    <el-form-item label="新昵称" :label-width="formLabelWidth">
                      <el-input v-model="form.name" autocomplete="off" />
                    </el-form-item>
                  </el-form>
                  <template #footer>
                    <div class="dialog-footer">
                      <el-button @click="dialogFormVisible = false">取消</el-button>
                      <el-button type="primary" @click="toChangeUserName(form.name)">
                        确认
                      </el-button>
                    </div>
                  </template>
                </el-dialog>
              </ul>
            </template>
          </el-popover>
          <div>
            <el-input
              class="Home-el-input"
              placeholder="搜索"
              :suffix-icon="Search"
              @focus="homeSearchBoxFocus"
              @change="homeSearchBoxChangs"
            />
          </div>
        </div>
        <el-menu-item index="/article/myday">
          <el-icon><Sunny /></el-icon>
          <span>我的一天</span>
        </el-menu-item>
        <el-menu-item index="/article/ImportantToDo">
          <el-icon><Star /></el-icon>
          <span>重要</span>
        </el-menu-item>
        <el-menu-item index="/article/TaskInPlan">
          <el-icon><List /></el-icon>
          <span>计划内</span>
        </el-menu-item>
        <el-menu-item index="/article/AssignedToMe">
          <el-icon><User /></el-icon>
          <span>已分配给我</span>
        </el-menu-item>
        <el-menu-item index="/article/TaskList">
          <el-icon><Collection /></el-icon>
          <span>已完成</span>
        </el-menu-item>
      </el-menu>
    </el-col>
    <el-col :span="menuVisible ? 20 : 24">
      <router-view></router-view>
    </el-col>
  </el-row>
</template>
<style>
.menu {
  transition: width 0.3s ease;
}
.homepage-menu {
  height: 100vh;
  background-color: rgb(144, 142, 142);
  color: black;
}

.title {
  background-image: url(../../assets/83a52f189a9329b6fd03fcaedb6611a5.png);
  background-size: 100% 100%;
  background-repeat: no-repeat;
  margin: 10px;
}

.menu-header {
  display: flex;
  flex-direction: column;
  height: 135px;
  align-items: center;
  flex-shrink: 1;
}

.header-user {
  display: flex;
  height: 100px;
}

.header-user-text {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.Home-el-input {
  box-shadow: 0 1px 1px rgb(0, 0, 0);
}

.homepage {
  display: flex;
  overflow: auto;
  flex-wrap: nowrap;
  height: 100vh;
}

@media (max-width: 768px) {
  .menu {
    display: none;
  }
}
</style>

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
          <div>
            <el-icon><TopRight /></el-icon>
          </div>
          <div>...</div>
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
